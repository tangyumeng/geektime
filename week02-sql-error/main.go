package main

import (
	"log"

	"example.com/dao"
)

func main() {
	defer dao.CloseDB()

	u, err := dao.QueryUserByID(100)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("test: ", u.Id, u.Name)

}
