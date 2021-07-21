package repository

const (
	create = `INSERT INTO markets (title, ticker, content, image_url, currency_id, instrument_id) 
					VALUES ($1, $2, $3, NULLIF($4, ''), $5, $6) 
					RETURNING *`

	update = `UPDATE markets 
					SET title = COALESCE(NULLIF($1, ''), title),
						ticker = COALESCE(NULLIF($2, ''), ticker),
						content = COALESCE(NULLIF($3, ''), content),
					    image_url = COALESCE(NULLIF($4, ''), image_url),
					    instrument_id = COALESCE(NULLIF($5, ''), instrument_id),
					    updated_at = now() 
					WHERE market_id = $6
					RETURNING *`

	getByID = `SELECT n.market_id,
       n.title,
       n.content,
       n.updated_at,
       n.image_url,
       n.category,
       CONCAT(u.first_name, ' ', u.last_name) as author,
       u.user_id as author_id
FROM markets n
         LEFT JOIN users u on u.user_id = n.author_id
WHERE market_id = $1`

	getByUserID = `SELECT r.identify, m.ticker
FROM markets m
         LEFT JOIN assets a on a.market_id = m.market_id
         LEFT JOIN portfolios p on p.portfolio_id = a.portfolio_id
         LEFT JOIN registers r on r.market_id = m.market_id
WHERE p.user_id = $1 GROUP BY r.identify, m.ticker`

	deleteById = `DELETE FROM markets WHERE market_id = $1`

	getTotalCount = `SELECT COUNT(market_id) FROM markets`

	getAll = `SELECT market_id, title, content, image_url, category, updated_at, created_at 
				FROM markets 
				ORDER BY created_at, updated_at OFFSET $1 LIMIT $2`

	findByTitleCount = `SELECT COUNT(*)
					FROM markets
					WHERE ticker ILIKE $1 || '%' OR title ILIKE $1 || '%'`

	findByTitle = `SELECT m.market_id, m.title, i.title as title_instrument, i.description as desc_instrument, m.ticker, m.content, m.image_url, m.updated_at
					FROM markets m
					LEFT JOIN instruments i ON i.instrument_id = m.instrument_id
					WHERE m.ticker ILIKE $1 || '%' OR m.title ILIKE $1 || '%'
					ORDER BY m.ticker, m.title
					OFFSET $2 LIMIT $3`
)
