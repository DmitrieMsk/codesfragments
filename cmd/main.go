package main

import (
	"codesfragments/internal/user"
	"codesfragments/storage"
	"fmt"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func main() {
	//router
	router := gin.Default()
	//router

	connStr := "user=dmitriemoskvin password=0000 dbname=codesfragments sslmode=disable"
	db, err := storage.NewDBConnect(connStr)
	if err != nil {
		panic("ошибка подключения к базе данных")
	}
	userRepos := user.NewUserRepository(db)
	user.NewUserHadnler(router, userRepos)

	// sqlRes, err := db.Exec("insert into test_table2 (name, age, city) values ("DIMA", $1, $2)", "22", "moskow_city")
	// if err != nil {
	// 	fmt.Println("ошибка exec: ", err)
	// }

	//id, _ := sqlRes.LastInsertId()
	fmt.Println(db)

	router.Run() // по умолчанию слушает 0.0.0.0:8080

}
