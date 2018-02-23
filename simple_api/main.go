package main

import (
	"database/sql"
	"fmt"

	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/lib/pq"
)

const (
	host   = "localhost"
	port   = 5432
	user   = "Katie"
	dbname = "go_practice"
)

func main() {
	// db setup
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"dbname=%s sslmode=disable",
		host, port, user, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println("Successfully connected!")

	// gin
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello World",
		})
	})
	r.Run()

}

// func () Query() {

// 	id := 3
// 	rows, err := db.Query("SELECT * FROM users WHERE id=?", id)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer rows.Close()
// 	for rows.Next() {
// 		var name string
// 		if err := rows.Scan(&name); err != nil {
// 			log.Fatal(err)
// 		}
// 		fmt.Println("%s is the id # for %d\n", name, id)
// 	}
// 	if err := rows.Err(); err != nil {
// 		log.Fatal(err)
// 	}
// }
