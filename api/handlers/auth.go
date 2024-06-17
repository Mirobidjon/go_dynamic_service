package handlers

import (
	"strings"
	"time"

	"github.com/mirobidjon/go_dynamic_service/api/http"
	"github.com/mirobidjon/go_dynamic_service/api/models"
	"github.com/mirobidjon/go_dynamic_service/config"
	"github.com/mirobidjon/go_dynamic_service/pkg/helper"

	fiber "github.com/gofiber/fiber/v2"
)

// Sign In godoc
// @ID sign_in
// @Router /client-api/auth/sign-in [POST]
// @Summary Sign In
// @Description Sign In
// @Accept json
// @Produce json
// @Tags Auth
// @Param location header string false "Location" default("Asia/Tashkent")
// @Param body body models.SignInRequest true "Sign In"
// @Success 200 {object} http.Response{data=models.SignInResponse} "Success"
// @Response 400 {object} http.Response{data=string} "Bad Request"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) SignIn(c *fiber.Ctx) error {
	var req models.SignInRequest
	if err := c.BodyParser(&req); err != nil {
		return h.handleResponse(c, http.BadRequest, nil, "invalid_request", err.Error())
	}

	if req.Username == "" {
		return h.handleResponse(c, http.BadRequest, nil, "invalid_username", "username is required")
	}

	if req.Password == "" {
		return h.handleResponse(c, http.BadRequest, nil, "invalid_password", "password is required")
	}

	if req.Username != h.cfg.Username || req.Password != h.cfg.Password {
		return h.handleResponse(c, http.BadRequest, nil, "invalid_credentials", "invalid credentials")
	}

	clientId := helper.GenerateUUID()

	h.cache.Set(h.cfg.SignInKey+clientId, clientId, time.Hour*8)

	resp := models.SignInResponse{
		AccessToken: clientId,
		ExpiredAt:   time.Now().Add(time.Hour * 24).Format(config.TimeStampLayout),
	}

	return h.handleResponse(c, http.OK, resp, "sign_in", "")
}

func (h *Handler) HasAccessMiddleware(c *fiber.Ctx) error {
	if h.cfg.Environment == "local" {
		return c.Next()
	}

	accessToken := c.Get("Authorization")

	if accessToken == "" {
		return h.handleResponse(c, http.Unauthorized, nil, "unauthorized", "unauthorized")
	}

	if !strings.HasPrefix(accessToken, "Bearer ") {
		return h.handleResponse(c, http.Unauthorized, nil, "unauthorized", "unauthorized")
	}

	accessToken = strings.TrimPrefix(accessToken, "Bearer ")

	if accessToken == "" {
		return h.handleResponse(c, http.Unauthorized, nil, "unauthorized", "unauthorized")
	}

	if h.cache.Get(h.cfg.SignInKey+accessToken) == nil {
		return h.handleResponse(c, http.Unauthorized, nil, "unauthorized", "unauthorized")
	}

	return c.Next()
}
