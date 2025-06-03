package server

import (
	"github.com/backend-d/phonebook/internal/types"
	"github.com/gofiber/fiber/v2"
)

type contactsListRequest struct {
	Page       int    `json:"page"`
	Limit      int    `json:"limit"`
	Search     string `json:"search"`
	Department string `json:"department"`
}

type contactsListResponse struct {
	Contacts []*types.Contact `json:"contacts"`
	Total    int              `json:"total"`
	Page     int              `json:"page"`
	Limit    int              `json:"limit"`
}

func (req *contactsListRequest) validate() {
	if req.Page < 1 {
		req.Page = 1
	}

	if req.Limit < 1 || req.Limit > 100 {
		req.Limit = 40
	}
}

func (s *Server) contactsListHandler(c *fiber.Ctx) error {
	req := &contactsListRequest{}
	if err := c.BodyParser(&req); err != nil {
		return s.badRequest(c, err, errInvalidJSON)
	}

	req.validate()

	offset := (req.Page - 1) * req.Limit
	contacts, total, err := s.deps.PG.SelectContactsWithFilters(req.Limit, offset, req.Search, req.Department)
	if err != nil {
		return s.internalServerError(c, err)
	}

	res := &contactsListResponse{
		Contacts: contacts,
		Total:    total,
		Page:     req.Page,
		Limit:    req.Limit,
	}

	return c.JSON(res)
}
