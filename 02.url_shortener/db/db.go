package db

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"go-projects/02.url_shortener/io"
	"log"
	"strings"
)

func GetPathsFromDb() map[string]string {
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/pathsDB")
	defer db.Close()
	handleError(err)

	runScript(db)

	return getAllPaths(db)
}

func getAllPaths(db *sql.DB) map[string]string {
	res, err2 := db.Query("SELECT * FROM paths")
	handleError(err2)

	paths := make(map[string]string)

	for res.Next() {

		var path string
		var url string
		err3 := res.Scan(&path, &url)
		handleError(err3)

		paths[path] = url
	}

	return paths
}

func runScript(db *sql.DB) {
	sqlFile := io.ReadFile("paths.sql")
	commands := strings.Split(string(sqlFile), "\n")

	for _, command := range commands {
		_, err := db.Exec(command)
		handleError(err)
	}
}

func handleError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
