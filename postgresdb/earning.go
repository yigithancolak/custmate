package postgresdb

import "github.com/yigithancolak/custmate/graph/model"

func (s *Store) ListEarningsByOrganization(orgID string, offset *int, limit *int, startDate string, endDate string) ([]*model.Earning, int, error) {
	query := `
SELECT
    g.id,
    g.name,
    COALESCE(SUM(CASE WHEN p.currency = 'try' THEN p.amount ELSE 0 END), 0) AS total_amount_try,
    COALESCE(SUM(CASE WHEN p.currency = 'usd' THEN p.amount ELSE 0 END), 0) AS total_amount_usd,
    COALESCE(SUM(CASE WHEN p.currency = 'eur' THEN p.amount ELSE 0 END), 0) AS total_amount_eur,
    COUNT(*) OVER() as total_count
FROM
    org_groups g
JOIN
    payments p ON g.id = p.org_group_id
WHERE
    g.organization_id = $1
AND
    p.date BETWEEN $2 AND $3
GROUP BY
    g.id
ORDER BY 
    g.name`

	var args []interface{}
	args = append(args, orgID, startDate, endDate)

	if limit != nil && offset != nil {
		query += ` LIMIT $4 OFFSET $5`
		args = append(args, limit, offset)
	}

	rows, err := s.DB.Query(query, args...)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	var earnings []*model.Earning
	var totalCount int

	for rows.Next() {
		//TODO: MAYBE OTHER FIELDS WILL BE INCLUDES IN GROUP IN FUTURE
		var earning model.Earning
		var group model.Group
		if err := rows.Scan(&group.ID, &group.Name, &earning.Try, &earning.Usd, &earning.Eur, &totalCount); err != nil {
			return nil, 0, err
		}

		earning.Group = &group
		earnings = append(earnings, &earning)
	}
	if err := rows.Err(); err != nil {
		return nil, 0, err
	}

	return earnings, totalCount, nil
}
