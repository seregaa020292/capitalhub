package repository

const (
	createQuery = `INSERT INTO portfolios (title, active, user_id, currency_id)
						VALUES ($1, $2, $3, $4) 
						RETURNING *`

	getActiveQuery = `SELECT *
					 FROM portfolios
					 WHERE active and user_id = $1`

	hasPortfolioUser = `SELECT TRUE as exist FROM portfolios WHERE portfolio_id = $1 and user_id = $2`
)
