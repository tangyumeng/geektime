package dao

//
import (
	"database/sql"
	"errors"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
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

var NotFoundCode = 400001
var OtherErrorCode = 500001

var NotFoundError = errors.New("sql row not found")
var OtherError = errors.New("other error")

var db *sql.DB
var err error

func init() {

	db, err = sql.Open("mysql",
		"root@tcp(127.0.0.1:3306)/week2")

	if err != nil {
		log.Fatal(err)
	}
}

func QueryUserByID(id int) (*User, error) {
	var u User
	sql := "select id, name from users where id = ?"
	stmt, err := db.Prepare(sql)
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	err = stmt.QueryRow(id).Scan(&u.Id, &u.Name)
	if err != nil {
		if err == sql.ErrNoRows {
			// return u, errors.Wrap(err, "QueryUserByID failed")
			// return nil, fmt.Errorf("%d not found", NotFoundCode)
			// return errors.Wrapf(code.NotFound, fmt.Sprintf("sql: %s error: %v", sql, err))
			return errors.Wrapf(NotFoundError, fmt.Sprintf("sql: %s error: %v", sql, err))
		} else {
			// return nil, fmt.Errorf("%d not found", OtherErrorCode)
			return errors.Wrapf(OtherError, fmt.Sprintf("sql: %s error: %v", sql, err))
		}
	}
	return &u, nil
}

// func IsNoRow(err error) bool {
// 	return strings.HasPrefix(err.Error(), fmt.Sprintf("%d", NotFoundCode))
// }

func CloseDB() {
	db.Close()
}
