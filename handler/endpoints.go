package handler

import (
	"context"
	"fmt"
	"net/http"

	"github.com/SawitProRecruitment/UserService/generated"
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
	var context context.Context

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
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte("request.Password"), bcrypt.DefaultCost)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": "Internal server error"})
	}

	// Store the user in the database (models.CreateUser should be implemented)
	userID, err := s.Repository.RegisterUser(context, request.PhoneNumber, request.FullName, string(hashedPassword))
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to create user"})
	}

	// Return the user ID as a response
	return ctx.JSON(http.StatusOK, map[string]int{"user_id": userID})
}

// (POST /login)
func (s *Server) PostLogin(ctx echo.Context) error {

	var req generated.PostLoginJSONRequestBody

	// Bind the request with body to the req variable
	if err := ctx.Bind(&req); err != nil {
		return err
	}
	fmt.Println(req.Password, "User Login", req.PhoneNumber)
	return ctx.JSON(http.StatusOK, req)
}
