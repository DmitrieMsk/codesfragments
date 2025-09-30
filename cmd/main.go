package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type user struct {
	id   int
	name string
	age  int
	city string
}

func main() {
	connStr := "user=dmitriemoskvin password=0000 dbname=codesfragments sslmode=disable"

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		fmt.Println("ошибка подключения: ", err)
	}
	defer db.Close()
	fmt.Println("check ping")
	err = db.Ping()
	if err != nil {
		fmt.Println("ошибка ping: ", err)
	}

	// sqlRes, err := db.Exec("insert into test_table2 (name, age, city) values ('DIMA', $1, $2)", "22", "moskow_city")
	// if err != nil {
	// 	fmt.Println("ошибка exec: ", err)
	// }

	rows, err := db.Query("select * from test_table2")
	if err != nil {
		fmt.Println("ошибка exec: ", err)
	}
	defer rows.Close()

	users := []user{}

	for rows.Next() {
		u := user{}
		err := rows.Scan(&u.id, &u.name, &u.age, &u.city)
		if err != nil {
			fmt.Println(err)
			continue
		}
		users = append(users, u)
	}
	for _, u := range users {
		fmt.Println(u.id, u.name, u.age, u.city)
	}

	//id, _ := sqlRes.LastInsertId()
	fmt.Println("end")

}
