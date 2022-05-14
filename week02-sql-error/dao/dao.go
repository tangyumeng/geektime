package dao

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/pkg/errors"
)

// DROP TABLE IF EXISTS `users`;
// CREATE TABLE `users` (
//   `id` int NOT NULL AUTO_INCREMENT,
//   `name` char(100) DEFAULT NULL,
//   PRIMARY KEY (`id`)
// ) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3;

type User struct {
	Id   int
	Name string
}

var db *sql.DB
var err error

func init() {

	db, err = sql.Open("mysql",
		"root@tcp(127.0.0.1:3306)/week2")

	if err != nil {
		log.Fatal(err)
	}
}

func QueryUserByID(id int) (User, error) {
	var u User
	stmt, err := db.Prepare("select id, name from users where id = ?")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	err = stmt.QueryRow(id).Scan(&u.Id, &u.Name)
	if err != nil {
		if err == sql.ErrNoRows {
			return u, errors.Wrap(err, "QueryUserByID failed")
		}
	}
	return u, nil
}

func CloseDB() {
	db.Close()
}
