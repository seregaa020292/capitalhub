package repository

import "fmt"

const (
	queryGetTotalByWhere = `SELECT m.market_id, m.title, m.ticker, r.identify, m.image_url as image_url,
							SUM(a.quantity) as total_quantity, COUNT(*) as total_count,
							(
							   SELECT SUM(amount * quantity)
							   FROM assets at WHERE at.market_id = m.market_id
						   	) as total_amount,
						   	(
							   SELECT ROUND(SUM(amount * quantity) / SUM(quantity))
							   FROM assets at WHERE at.market_id = m.market_id
						   	) as average_total_amount,
       						MIN(a.notation_at) as first_notation_at
							FROM assets a
							LEFT JOIN portfolios p on a.portfolio_id = p.portfolio_id
							LEFT JOIN markets m on a.market_id = m.market_id
							LEFT JOIN registers r on a.market_id = r.market_id
							WHERE %s AND p.active GROUP BY m.market_id, r.register_id`

	createAsset = `INSERT INTO assets (portfolio_id, market_id, amount, quantity, notation_at) VALUES ($1, $2, $3, $4, $5) RETURNING *`

	updateAsset = `UPDATE assets SET amount = $1, quantity = $2, updated_at = CURRENT_TIMESTAMP WHERE asset_id = $3 RETURNING *`

	deleteAsset = `DELETE FROM assets WHERE asset_id = $1`

	getAssetByUserID = `SELECT m.title, m.ticker, m.image_url as image_url, a.asset_id, a.user_id, a.amount, a.quantity, a.updated_at	
						FROM assets a
        				LEFT JOIN markets m on a.market_id = m.market_id
						WHERE a.user_id = $1`

	getAssetByID = `SELECT u.name as user, m.image_url as image_url, a.amount, a.quantity, a.updated_at, a.user_id, a.asset_id	
						FROM assets a
        				LEFT JOIN portfolios p on a.portfolio_id = p.portfolio_id
						LEFT JOIN users u on p.user_id = u.user_id
        				LEFT JOIN markets m on a.market_id = m.market_id
						WHERE a.asset_id = $1`

	getTotalCountByAssetID = `SELECT COUNT(asset_id) FROM assets WHERE asset_id = $1`

	getAssetsByAssetID = `SELECT u.name as user, m.image_url as image_url, a.amount, a.quantity, a.updated_at, a.user_id, a.asset_id
							FROM assets a
        					LEFT JOIN users u on a.user_id = u.user_id
        					LEFT JOIN markets m on a.market_id = m.market_id
        					WHERE a.asset_id = $1 
							ORDER BY updated_at OFFSET $2 LIMIT $3`
)

var (
	getTotalAssetByUserID = fmt.Sprintf(queryGetTotalByWhere, "p.user_id = $1")
	getTotalAssetByMarketID = fmt.Sprintf(queryGetTotalByWhere, "a.market_id = $1 AND a.portfolio_id = $2")
)
