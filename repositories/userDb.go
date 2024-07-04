package repositories

import (
	"time"

	"alexpy.com/julia/db"

	"alexpy.com/julia/models"
)

func CreateUser(user models.User) (models.User, error) {
	db := db.GetDB()
	sqlStatement := `INSERT INTO users (name, email, password, created_at, updated_at) VALUES ($1, $2, $3, $4, $5) RETURNING id`
	err := db.QueryRow(sqlStatement, user.Name, user.Email, user.Password, time.Now(), time.Now()).Scan(&user.Id)

	if err != nil {
		return user, err
	}
	return user, nil
}

func UpdateUser(user models.User, id int) (models.User, error) {
	db := db.GetDB()
	sqlStatement := `
		UPDATE users
		SET name = $2, email = $3, password = $4, updated_at = $5
		WHERE id = $1
		RETURNING id`
	err := db.QueryRow(sqlStatement, id, user.Name, user.Email, user.Password, time.Now()).Scan(&id)
	if err != nil {
		return models.User{}, err
	}
	user.Id = id
	return user, nil
}

func GetAllUsers() ([]models.User, error) {
	db := db.GetDB()
	sqlStatement := `
	SELECT id, name, email, created_at, updated_at
	FROM users;`
	rows, err := db.Query(sqlStatement)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var users []models.User
	for rows.Next() {
		var user models.User
		err := rows.Scan(&user.Id, &user.Name, &user.Email, &user.CreatedAt, &user.UpdatedAt)
		if err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	// Check for errors from iterating over rows.
	return users, nil

}
