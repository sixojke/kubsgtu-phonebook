package server

import (
	"github.com/backend-d/phonebook/internal/types"
	"github.com/gofiber/fiber/v2"
)

type departmentsListResponse struct {
	Departments []*types.Department `json:"departments"`
}

func (s *Server) departmentsListHandler(c *fiber.Ctx) error {
	departments, err := s.deps.PG.SelectDepartments()
	if err != nil {
		return s.internalServerError(c, err)
	}

	res := &departmentsListResponse{
		Departments: departments,
	}

	return c.JSON(res)
}
