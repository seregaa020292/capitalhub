package repository

import "fmt"

const (
	createQuery = `INSERT INTO portfolios (title, active, user_id, currency_id)
						VALUES ($1, $2, $3, $4) 
						RETURNING *`

	clearActiveQuery = `UPDATE portfolios SET active = FALSE
						WHERE user_id = $1`

	editQuery = `UPDATE portfolios SET title = $1, currency_id = $2
						WHERE portfolio_id = $3
						RETURNING *`

	deleteQuery = `DELETE FROM portfolios WHERE portfolio_id = $1 and user_id = $2`

	setActiveQuery = `UPDATE portfolios SET active = TRUE
						WHERE portfolio_id = $1 and user_id = $2`

	getActiveQuery = `SELECT *
					 FROM portfolios
					 WHERE active and user_id = $1`

	hasPortfolioUserQuery = `SELECT TRUE as exist FROM portfolios WHERE portfolio_id = $1 and user_id = $2`

	getAllQuery = `SELECT *
					FROM portfolios
					WHERE user_id = $1`

	queryStatsWhere = `SELECT p.title, p.active, p.currency_id, p.portfolio_id,
					c.title as currency_title, c.description as currency_desc,
					COALESCE(SUM(a.quantity), 0) as asset_quantity, COALESCE(SUM(a.amount), 0) as asset_amount
					FROM portfolios p
					LEFT JOIN assets a on p.portfolio_id = a.portfolio_id
					LEFT JOIN currencies c on p.currency_id = c.currency_id
					WHERE %s GROUP BY p.portfolio_id, c.currency_id ORDER BY p.created_at ASC`
)

var (
	getStatsQuery = fmt.Sprintf(queryStatsWhere, "p.portfolio_id = $1")
	getAllTotalQuery = fmt.Sprintf(queryStatsWhere, "p.user_id = $1")
)
