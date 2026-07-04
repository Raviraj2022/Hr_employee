package auth

import (
	

	"example.com/hr-emp-mgmt/internal/user"
	"example.com/hr-emp-mgmt/pkg/errors"
	"example.com/hr-emp-mgmt/pkg/response"
	"example.com/hr-emp-mgmt/pkg/validator"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	service user.Service
}

func NewHandler(service user.Service) *Handler {
	return &Handler{
		service: service,
	}
}

func (h *Handler) Register(c *gin.Context) {

	var req RegisterRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		// c.JSON(http.StatusBadRequest, gin.H{
		// 	"message": err.Error(),
		// })
		response.Error(c, 400, "Invalid request body", nil)
		return
	}
validationErrors := validator.ValidateStruct(req)


if validationErrors != nil {
	response.Error(c, 422, "Validation failed", validationErrors)
	return
}
	newUser := user.User{
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Email:     req.Email,
		Password:  req.Password,
	}

	err := h.service.Register(&newUser)
	if err != nil {

	switch err {

	case errors.ErrEmailExists:
		response.Error(c, 409, err.Error(), nil)

	default:
		response.Error(c, 500, "Internal server error", nil)
	}

	return
}

	response.Success(c, 201, "User registered successfully", gin.H{
	"data": newUser,
})
}


func (h *Handler) Login(c *gin.Context) {

	var req LoginRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, 400, "Invalid request body", err.Error())
		return
	}

	validationErrors := validator.ValidateStruct(req)
	if validationErrors != nil {
		response.Error(c, 422, "Validation failed", validationErrors)
		return
	}

	loginResponse, err := h.service.Login(req)

	if err != nil {

		switch err {

		case errors.ErrInvalidCredentials:
			response.Error(c, 401, err.Error(), nil)

		default:
			response.Error(c, 500, "Internal server error", nil)
		}

		return
	}

	response.Success(c, 200, "Login successful", loginResponse)
}