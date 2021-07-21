package repository

const (
	createUserQuery = `INSERT INTO users (name, email, password, role, avatar, created_at, updated_at)
						VALUES ($1, $2, $3, COALESCE(NULLIF($4, ''), 'user'), $5, now(), now()) 
						RETURNING *`

	updateUserQuery = `UPDATE users 
						SET first_name = COALESCE(NULLIF($1, ''), name),
						    email = COALESCE(NULLIF($2, ''), email),
						    role = COALESCE(NULLIF($3, ''), role),
						    avatar = COALESCE(NULLIF($4, ''), avatar),
						    updated_at = now()
						WHERE user_id = $5
						RETURNING *
						`
	confirmedUserQuery = `UPDATE users
						SET confirmed = NULL
						WHERE confirmed = $1
						RETURNING *
						`

	deleteUserQuery = `DELETE FROM users WHERE user_id = $1`

	getUserQuery = `SELECT user_id, name, email, role, avatar, confirmed, created_at, updated_at
					 FROM users 
					 WHERE user_id = $1`

	getTotalCount = `SELECT COUNT(user_id) FROM users 
						WHERE name ILIKE '%' || $1 || '%' or email ILIKE '%' || $1 || '%'`

	findUsers = `SELECT user_id, name, email, role, avatar, created_at, updated_at 
				  FROM users 
				  WHERE name ILIKE '%' || $1 || '%' or email ILIKE '%' || $1 || '%'
				  ORDER BY name, email
				  OFFSET $2 LIMIT $3
				  `

	getTotal = `SELECT COUNT(user_id) FROM users`

	getUsers = `SELECT user_id, name, email, role, avatar, created_at, updated_at
				 FROM users 
				 ORDER BY COALESCE(NULLIF($1, ''), name) OFFSET $2 LIMIT $3`

	findUserByEmail = `SELECT user_id, name, email, role, avatar, password, confirmed, created_at, updated_at
				 		FROM users 
				 		WHERE email = $1`
)
