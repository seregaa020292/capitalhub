package repository

const (
	create = `INSERT INTO registers (identify, provider_id, market_id) 
					VALUES ($1, $2, $3) 
					RETURNING *`
)
