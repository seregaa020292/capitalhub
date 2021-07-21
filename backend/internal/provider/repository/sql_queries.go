package repository

const (
	getByTitle = `SELECT *	
						FROM providers
						WHERE title = $1`
)
