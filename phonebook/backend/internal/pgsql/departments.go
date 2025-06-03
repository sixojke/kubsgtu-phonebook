package pgsql

import (
	"github.com/backend-d/phonebook/internal/types"
)

func (c *Client) SelectDepartments() ([]*types.Department, error) {
	sess := c.GetSession()

	departments := make([]*types.Department, 0)
	if _, err := sess.Select("*").From("departments").Load(&departments); err != nil {
		return nil, err
	}

	return departments, nil
}
