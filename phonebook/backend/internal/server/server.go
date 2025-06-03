package server

import (
	"github.com/backend-d/phonebook/internal/pgsql"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

type Server struct {
	App *fiber.App

	cfg  *Config
	deps *Deps
}

type Config struct{}

type Deps struct {
	PG *pgsql.Client
}

func New(cfg *Config, deps *Deps) *Server {
	s := &Server{
		App: fiber.New(fiber.Config{}),

		cfg:  cfg,
		deps: deps,
	}

	s.App.Use(cors.New())

	// panic recovery
	s.App.Use(recover.New(recover.Config{
		Next:             nil,
		EnableStackTrace: true,
	}))

	api := s.App.Group("/api")

	contacts := api.Group("/contacts")
	contacts.Post("/list", s.contactsListHandler)

	departments := api.Group("/departments")
	departments.Post("/list", s.departmentsListHandler)

	return s
}

type response struct {
	Message string `json:"message"`
}

func (s *Server) responseOK(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(response{"ok"})
}

func (s *Server) internalServerError(c *fiber.Ctx, err error) error {
	log.Error(err)
	return c.Status(fiber.StatusInternalServerError).
		JSON(response{fiber.ErrInternalServerError.Error()})
}

func (s *Server) badRequest(c *fiber.Ctx, err error, errResp error) error {
	log.Warn(err)
	return c.Status(fiber.StatusBadRequest).
		JSON(response{errResp.Error()})
}
