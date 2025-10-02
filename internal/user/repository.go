package user

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

type UserRepo struct {
	db *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) *UserRepo {
	return &UserRepo{db: db}
}

func (u *UserRepo) GetAll() *[]User {
	rows, err := u.db.Query("select * from test_table3")
	if err != nil {
		fmt.Println("ошибка select: ", err)
	}
	defer rows.Close()

	users := []User{}

	for rows.Next() {
		u := User{}
		err := rows.Scan(&u.Id, &u.Name, &u.Age, &u.City)
		if err != nil {
			fmt.Println(err)
			continue
		}
		users = append(users, u)
	}
	for _, u := range users {
		fmt.Println(u.Id, u.Name, u.Age, u.City)
	}

	return &users
}
