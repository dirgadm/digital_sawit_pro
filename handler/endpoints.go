package handler

import (
	"fmt"
	"net/http"

	"github.com/SawitProRecruitment/UserService/generated"
	"github.com/labstack/echo/v4"
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

	var resp generated.HelloResponse
	resp.Message = fmt.Sprintf("User Registered")
	return ctx.JSON(http.StatusOK, resp)
}

// (POST /login)
func (s *Server) PostLogin(ctx echo.Context) error {

	var req generated.PostLoginJSONRequestBody
	fmt.Println("User Login", req)
	return ctx.JSON(http.StatusOK, req)
}
