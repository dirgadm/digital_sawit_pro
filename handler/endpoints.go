package handler

import (
	"context"
	"fmt"
	"net/http"

	"github.com/SawitProRecruitment/UserService/generated"
	"github.com/SawitProRecruitment/UserService/utils"
	"github.com/labstack/echo/v4"

	"golang.org/x/crypto/bcrypt"
)

// This is just a test endpoint to get you started. Please delete this endpoint.
// (GET /hello)
func (s *Server) Hello(ctx echo.Context, params generated.HelloParams) error {

	var resp generated.HelloResponse
	resp.Message = fmt.Sprintf("Hello User %d", params.Id)
	return ctx.JSON(http.StatusOK, resp)
}

// (POST /registration)
func (s *Server) PostRegistration(ctx echo.Context) error {

	var request generated.PostRegistrationJSONBody

	if err := ctx.Bind(&request); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request data"})
	}

	// Perform validation checks on request data
	validationErrors := make(map[string]string)

	// Validate phone number
	// ...

	// Validate full name
	// ...

	// Validate password
	// ...

	if len(validationErrors) > 0 {
		return ctx.JSON(http.StatusBadRequest, validationErrors)
	}

	// Hash the password before storing it in the database
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": "Internal server error"})
	}

	fmt.Println(request.PhoneNumber, "------------", request.FullName, "---------------", string(hashedPassword), "----------------", s)
	userID, err := s.Repository.RegisterUser(context.Background(), request.PhoneNumber, request.FullName, string(hashedPassword))
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to create user"})
	}

	// Return the user ID as a response
	return ctx.JSON(http.StatusOK, map[string]int{"user_id": userID})
}

// GetProfile retrieves the user's profile
func (s *Server) GetProfile(c echo.Context) error {
	// Get the user ID from the JWT token
	var (
		userID int
		err    error
	)
	if userID, err = utils.GetUserIDFromToken(c); err != nil {
		return err
	}

	// Get the user's profile information
	fullName, phoneNumber, err := s.Repository.GetMyProfile(context.Background(), userID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to fetch profile"})
	}

	// Return the user's profile information
	return c.JSON(http.StatusOK, map[string]interface{}{"fullName": fullName, "phoneNumber": phoneNumber})
}

// PutProfile updates the user's profile
func (s *Server) PutProfile(c echo.Context) error {

	var (
		userID int
		err    error
		req    generated.PutProfileJSONRequestBody
	)
	// Get the user ID from the JWT token
	if userID, err = utils.GetUserIDFromToken(c); err != nil {
		return err
	}

	// Parse the request body into an UpdateProfileRequest struct
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request"})
	}

	// Validate the update profile request
	if err := c.Validate(req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Validation failed"})
	}

	// Check if the provided phone number exists for another user
	if _, _, err := s.Repository.GetMyProfile(context.Background(), userID, req.PhoneNumber); err != nil {
		return c.JSON(http.StatusConflict, map[string]string{"error": "Phone number already exists for another user"})
	}

	if err := s.Repository.UpdateMyProfile(context.Background(), userID, req.PhoneNumber, req.FullName); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to update profile"})
	}

	// Return a success response
	return c.JSON(http.StatusOK, map[string]string{"message": "Profile updated successfully"})
}

// (POST /login)
func (s *Server) PostLogin(ctx echo.Context) error {

	var req generated.PostLoginJSONRequestBody

	// Bind the request with body to the req variable
	if err := ctx.Bind(&req); err != nil {
		return err
	}
	// Validate the login request
	if err := ctx.Validate(req); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "Validation failed"})
	}

	// Check if the provided phone number and password match a user in the database
	userID, err := s.Repository.LoginUser(context.Background(), req.PhoneNumber, req.Password)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "Login failed"})
	}

	// Generate a JWT token for the user
	jwtToken, err := utils.GenerateJWTToken(userID)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to generate JWT token"})
	}

	// Return the user ID and JWT token
	return ctx.JSON(http.StatusOK, map[string]interface{}{"userId": userID, "jwtToken": jwtToken})
}
