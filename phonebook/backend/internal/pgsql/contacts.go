package pgsql

import (
	"fmt"
	"strings"

	"github.com/backend-d/phonebook/internal/types"
	"github.com/gocraft/dbr/v2"
)

func (c *Client) SelectContactsWithFilters(limit, offset int, search, department string) ([]*types.Contact, int, error) {
	sess := c.GetSession()

	query := sess.Select(
		"c.id",
		"c.name",
		"c.position",
		"c.phone",
		"c.email",
		"c.room",
		"departments.name AS department",
	).From("contacts c").
		Join("departments", "c.department_id = departments.id")

	countQuery := sess.Select("COUNT(*)").
		From("contacts c").
		Join("departments", "c.department_id = departments.id")

	if search != "" {
		searchTerms := strings.Fields(search)
		for _, term := range searchTerms {
			lowerTerm := strings.ToLower(term)
			cond := dbr.Or(
				dbr.Expr("LOWER(c.name) LIKE ?", "%"+lowerTerm+"%"),
				dbr.Expr("LOWER(c.position) LIKE ?", "%"+lowerTerm+"%"),
			)
			query.Where(cond)
			countQuery.Where(cond)
		}
	}

	if department != "" {
		query.Where("departments.name = ?", department)
		countQuery.Where("departments.name = ?", department)
	}

	var total int
	if err := countQuery.LoadOne(&total); err != nil {
		return nil, 0, fmt.Errorf("failed to get the count of contacts: %v", err)
	}

	contacts := make([]*types.Contact, 0)
	if _, err := query.
		OrderBy("c.name").
		Limit(uint64(limit)).
		Offset(uint64(offset)).
		Load(&contacts); err != nil {
		return nil, 0, fmt.Errorf("failed to select contacts: %v", err)
	}

	return contacts, 0, nil
}
