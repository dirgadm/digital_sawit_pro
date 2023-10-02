package handler

import (
	"context"
	"fmt"
	"net/http"

	"github.com/SawitProRecruitment/UserService/generated"
	"github.com/SawitProRecruitment/UserService/repository"
	"github.com/SawitProRecruitment/UserService/utils"
	"github.com/go-playground/validator/v10"
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
	s.Repository = repository.NewRepository(repository.NewRepositoryOptions{})

	var request generated.PostRegistrationJSONBody

	if err := ctx.Bind(&request); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request data"})
	}

	// Perform validation checks on request data
	validationErrors := make(map[string]string)

	// Validate phone number
	if len(request.PhoneNumber) < 10 || len(request.PhoneNumber) > 13 {
		validationErrors["phoneNumber"] = utils.LengthPhoneNumber
	}
	// Check if the phone number starts with "+62"
	if request.PhoneNumber[:3] != "+62" {
		validationErrors["phoneNumber"] += utils.StartingCode
	}

	// Validate full name
	if len(request.FullName) < 3 || len(request.FullName) > 60 {
		validationErrors["fullName"] = utils.LengthFullName
	}

	// Validate password
	if !utils.IsStrongPassword(request.Password) {
		validationErrors["password"] = utils.NotStrongPassword
	}

	if len(validationErrors) > 0 {
		return ctx.JSON(http.StatusBadRequest, validationErrors)
	}

	// Hash the password before storing it in the database
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": "Internal server error"})
	}

	userID, err := s.Repository.RegisterUser(context.Background(), request.PhoneNumber, request.FullName, string(hashedPassword))
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to create user"})
	}

	// Return the user ID as a response
	return ctx.JSON(http.StatusOK, map[string]int{"user_id": userID})
}

// GetProfile retrieves the user's profile
func (s *Server) GetProfile(c echo.Context) error {
	s.Repository = repository.NewRepository(repository.NewRepositoryOptions{})
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
		return c.JSON(http.StatusForbidden, map[string]string{"error": "Failed to fetch profile"})
	}

	// Return the user's profile information
	return c.JSON(http.StatusOK, map[string]interface{}{"fullName": fullName, "phoneNumber": phoneNumber})
}

// PutProfile updates the user's profile
func (s *Server) PutProfile(c echo.Context) error {
	s.Repository = repository.NewRepository(repository.NewRepositoryOptions{})

	var (
		userID int
		err    error
	)
	validator := validator.New()

	req := &generated.PutProfileJSONRequestBody{}
	// Get the user ID from the JWT token
	if userID, err = utils.GetUserIDFromToken(c); err != nil {
		return err
	}

	// Parse the request body into an UpdateProfileRequest struct
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request"})
	}

	// Perform validation checks on request data
	validationErrors := make(map[string]string)

	// Validate phone number
	if len(req.PhoneNumber) > 0 && len(req.PhoneNumber) < 10 || len(req.PhoneNumber) > 13 {
		validationErrors["phoneNumber"] = utils.LengthPhoneNumber
	}
	// Check if the phone number starts with "+62"
	if len(req.PhoneNumber) > 0 && req.PhoneNumber[:3] != "+62" {
		validationErrors["phoneNumber"] += utils.StartingCode
	}

	// Validate full name
	if len(req.FullName) > 0 && len(req.FullName) < 3 || len(req.FullName) > 60 {
		validationErrors["fullName"] = utils.LengthFullName
	}

	if len(validationErrors) > 0 {
		return c.JSON(http.StatusBadRequest, validationErrors)
	}

	// Validate the update profile request
	if err := validator.Struct(req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Validation failed"})
	}

	if err := s.Repository.UpdateMyProfile(context.Background(), userID, req.PhoneNumber, req.FullName); err != nil {
		return c.JSON(http.StatusConflict, map[string]string{"error": err.Error()})
	}

	// Return a success response
	return c.JSON(http.StatusOK, map[string]string{"message": "Profile updated successfully"})
}

// (POST /login)
func (s *Server) PostLogin(ctx echo.Context) error {
	s.Repository = repository.NewRepository(repository.NewRepositoryOptions{})

	var req generated.PostLoginJSONRequestBody

	validator := validator.New()

	// Bind the request with body to the req variable
	if err := ctx.Bind(&req); err != nil {
		return err
	}
	// Validate the login request
	if err := validator.Struct(req); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "Validation failed"})
	}

	// Hash the password before storing it in the database
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "Internal server error"})
	}

	// Check if the provided phone number and password match a user in the database
	user, err := s.Repository.LoginUser(context.Background(), req.PhoneNumber, string(hashedPassword))
	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(req.Password))

	if err != nil {

		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "Login failed"})
	}

	// Generate a JWT token for the user
	jwtToken, err := utils.GenerateJWTToken(int(user.ID))
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "Failed to generate JWT token"})
	}

	// Return the user ID and JWT token
	return ctx.JSON(http.StatusOK, map[string]interface{}{"userId": user.ID, "jwtToken": jwtToken})
}
