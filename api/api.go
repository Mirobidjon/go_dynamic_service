package api

import (
	"encoding/json"
	"time"

	_ "github.com/mirobidjon/go_dynamic_service/api/docs"
	"github.com/mirobidjon/go_dynamic_service/api/handlers"
	"github.com/mirobidjon/go_dynamic_service/config"
	"github.com/mirobidjon/go_dynamic_service/grpc/client"

	swagger "github.com/arsmn/fiber-swagger/v2"
	fiber "github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/basicauth"
	"github.com/gofiber/fiber/v2/middleware/cors"
	fiberLogger "github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/saidamir98/udevs_pkg/logger"
)

func StartHTTPServer(cfg config.Config, log logger.LoggerI) error {
	svcs, err := client.NewGrpcClients(cfg)
	if err != nil {
		log.Error("client.NewGrpcClients", logger.Error(err))
		return err
	}

	r := SetUpRouter(handlers.NewHandler(cfg, log, svcs), cfg)

	log.Info("HTTP: Server being started...", logger.String("port", cfg.HTTPPort))
	if err = r.Listen(cfg.HTTPPort); err != nil {
		log.Error("HTTP server error", logger.Error(err))
		return err
	}

	return nil
}

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
// @securityDefinitions.basic BasicAuth
// @in header
// @name Authorization
// SetUpRouter godoc
// @description This is a api gateway
// @termsOfService https://mirobidjon.uz
func SetUpRouter(h handlers.Handler, cfg config.Config) *fiber.App {
	r := fiber.New(fiber.Config{JSONEncoder: json.Marshal, BodyLimit: 100 * 1024 * 1024})
	r.Use(fiberLogger.New(), cors.New())
	r.Use(h.HandlerLanguage())
	r.Server().MaxConnsPerIP = 100
	r.Server().DisableKeepalive = true
	r.Server().IdleTimeout = time.Minute * 2
	r.Server().LogAllErrors = true

	open := r.Group("/client-api/service", h.Limiter(10, 5))
	{
		open.Get("/:slug/get-all", h.GetAllServices)
		open.Get("/:slug/get/:id", h.GetEntityService)
	}

	r.Use(h.Limiter(300, 10))

	r.Use("/client-api/swagger/", basicauth.New(basicauth.Config{
		Users: map[string]string{
			cfg.Username: cfg.Password,
		},
		Realm: "Restricted",
	}))

	r.Get("/client-api/swagger/*", swagger.HandlerDefault)

	r.Post("/client-api/auth/sign-in", h.SignIn)

	dynamic := r.Group("/client-api/dynamic", h.HasAccessMiddleware)
	{
		dynamic.Post("/group", h.CreateGroup)
		dynamic.Get("/group", h.GetAllGroup)
		dynamic.Get("/group/:id", h.GetGroup)
		dynamic.Put("/group/:id", h.UpdateGroup)
		dynamic.Delete("/group/:id", h.DeleteGroup)
		dynamic.Post("/field", h.CreateField)
		dynamic.Get("/field", h.GetAllField)
		dynamic.Get("/field/:id", h.GetFieldById)
		dynamic.Put("/field/:id", h.UpdateField)
		dynamic.Delete("/field/:id", h.DeleteField)
		dynamic.Get("/group/:slug/full", h.GetFullGroup)
	}

	entity := r.Group("/client-api/entity", h.HasAccessMiddleware)
	{
		entity.Post("/:slug/create", h.CreateEntity)
		entity.Put("/:slug/update/:id", h.UpdateEntity)
		entity.Patch("/:slug/update/:id", h.PatchUpdateEntity)
		entity.Delete("/:slug/delete/:id", h.DeleteEntity)
		entity.Get("/:slug/get/:id", h.GetEntity)
		entity.Post("/:slug/get-all", h.GetAllEntityPost)
		entity.Get("/:slug/get-all", h.GetAllEntityGet)
		entity.Post("/:slug/get-join", h.GetJoinEntity)
	}

	configuration := r.Group("/client-api/configuration", h.HasAccessMiddleware)
	{
		configuration.Get("/field_types", h.FieldTypeConfiguration)
		configuration.Get("/default_values", h.DefaultValuesConfiguration)
		configuration.Get("/group_types", h.GroupTypeConfiguration)
		configuration.Get("/validation_functions", h.ValidationFunctionConfiguration)
		configuration.Get("/regex", h.RegexConfiguration)
	}

	r.Post("/client-api/upload", h.HasAccessMiddleware, h.Upload)
	r.Get("/client-api/download-file", h.DownloadFile)

	return r
}
