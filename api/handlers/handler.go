package handlers

import (
	"kassa360/kassa360_go_dynamic_service/api/http"
	"kassa360/kassa360_go_dynamic_service/config"
	"kassa360/kassa360_go_dynamic_service/grpc/client"
	"strconv"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/jellydator/ttlcache/v3"

	"github.com/saidamir98/udevs_pkg/logger"

	fiber "github.com/gofiber/fiber/v2"
)

type Handler struct {
	cfg      config.Config
	log      logger.LoggerI
	services client.ServiceManagerI
	cache    *ttlcache.Cache[string, string]
}

type Map map[string]interface{}

var (
	ErrAlreadyExists       = "ALREADY_EXISTS"
	ErrNotFound            = "NOT_FOUND"
	ErrInternalServerError = "INTERNAL_SERVER_ERROR"
	ErrServiceUnavailable  = "SERVICE_UNAVAILABLE"
	SigningKey             = []byte("FfLbN7pIEYe8@!EqrttOLiwa(H8)7Ddo")
	SuperAdminUserType     = "superadmin"
	SystemUserType         = "admin"
)

func NewHandler(cfg config.Config, log logger.LoggerI, svcs client.ServiceManagerI) Handler {
	return Handler{
		cfg:      cfg,
		log:      log,
		services: svcs,
		cache: ttlcache.New[string, string](
			ttlcache.WithTTL[string, string](time.Duration(cfg.CacheTTL) * time.Minute),
		),
	}
}

func (h *Handler) handleResponse(c *fiber.Ctx, status http.Status, data interface{}, key string, err string) error {
	customMessage := getCustomMessage(status, c.Get("lang", "uz"), key)

	switch code := status.Code; {
	case code < 300:
		h.log.Info(
			"response",
			logger.Int("code", status.Code),
			logger.String("status", status.Status),
			logger.Any("description", status.Description),
			logger.Any("data", data),
			logger.Any("key", key),
			logger.Any("custom-request-id", c.Get("custom-request-id", "default")),
		)
	case code < 400:
		h.log.Warn(
			"response",
			logger.Int("code", status.Code),
			logger.String("status", status.Status),
			logger.Any("description", status.Description),
			logger.Any("data", data),
			logger.Any("key", key),
			logger.Any("error", err),
			logger.Any("custom-request-id", c.Get("custom-request-id", "default")),
		)
	default:
		h.log.Error(
			"response",
			logger.Int("code", status.Code),
			logger.String("status", status.Status),
			logger.Any("description", status.Description),
			logger.Any("data", data),
			logger.Any("error", err),
			logger.Any("key", key),
			logger.Any("customMessage", customMessage),
			logger.Any("custom-request-id", c.Get("custom-request-id", "default")),
		)
	}

	return c.Status(status.Code).JSON(http.Response{
		Status:      status.Status,
		Description: status.Description,
		Data:        data,
		RequestId:   c.Get("custom-request-id"),
		Error:       err,
		Message:     customMessage,
	})
}

// func (h *Handler) handleIntegrationResponse(c *fiber.Ctx, status http.Status, data interface{}) error {
// 	switch code := status.Code; {
// 	case code < 300:
// 		h.log.Info(
// 			"response",
// 			logger.Int("code", status.Code),
// 			logger.String("status", status.Status),
// 			logger.Any("description", status.Description),
// 			logger.Any("data", data),
// 			logger.Any("custom-request-id", c.Get("custom-request-id", "default")),
// 		)
// 	case code < 400:
// 		h.log.Warn(
// 			"response",
// 			logger.Int("code", status.Code),
// 			logger.String("status", status.Status),
// 			logger.Any("description", status.Description),
// 			logger.Any("data", data),
// 			logger.Any("custom-request-id", c.Get("custom-request-id", "default")),
// 		)
// 	default:
// 		h.log.Error(
// 			"response",
// 			logger.Int("code", status.Code),
// 			logger.String("status", status.Status),
// 			logger.Any("description", status.Description),
// 			logger.Any("data", data),
// 			logger.Any("custom-request-id", c.Get("custom-request-id", "default")),
// 		)
// 	}

// 	return c.Status(status.Code).JSON(data)
// }

func (h *Handler) getOffsetParam(c *fiber.Ctx) (offset int, err error) {
	offsetStr := c.Query("offset", h.cfg.DefaultOffset)
	return strconv.Atoi(offsetStr)
}

func (h *Handler) getLimitParam(c *fiber.Ctx) (offset int, err error) {
	offsetStr := c.Query("limit", h.cfg.DefaultLimit)
	return strconv.Atoi(offsetStr)
}

func (h *Handler) HandlerLanguage() fiber.Handler {
	return func(c *fiber.Ctx) error {
		lang := c.Get("lang", "uz")
		if lang != "ru" && lang != "uz" && lang != "oz" {
			c.Context().Request.Header.Set("lang", "uz")
		}

		id := uuid.NewString()
		c.Request().Header.Set("custom-request-id", id)
		c.Context().SetUserValue("request_id", id)

		return c.Next()
	}
}

func MakePrettySlice(s string) []string {
	var empty []string
	str := strings.Split(s, ",")
	for _, item := range str {
		if item != "" {
			empty = append(empty, item)
		}
	}

	return empty
}

func getCustomMessage(status http.Status, userLang, key string) string {
	var customMessage string

	if status.Code < 300 {
		if _, ok := http.CustomSuccessMessages[key][userLang]; ok {
			customMessage = http.CustomSuccessMessages[key][userLang]
		} else {
			customMessage = http.CustomSuccessMessages["default_success_message"][userLang]
		}
		return customMessage
	}

	if _, ok := http.CustomErrorMessages[key][userLang]; ok {
		customMessage = http.CustomErrorMessages[key][userLang]
	} else {
		customMessage = http.CustomErrorMessages["default_error_message"][userLang]
	}

	return customMessage
}
