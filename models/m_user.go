package models

import "a04-go-mvc-web-v1/config"

type User struct {
	ID       int    `json:"id"`
	NAME     string `json:"name"`
	USERNAME string `json:"username"`
}

func GetUsers() ([]User, error) {
	rows, err := config.DB.Query("SELECT ID, NAME, USERNAME FROM USERADISMONLAP.VER_USERS")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []User
	for rows.Next() {
		var user User
		if err := rows.Scan(&user.ID, &user.NAME, &user.USERNAME); err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}
