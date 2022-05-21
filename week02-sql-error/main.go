package main

import (
	"log"

	"example.com/dao"
)

func main() {
	//自己实现。思路不对
	// defer dao.CloseDB()

	// u, err := dao.QueryUserByID(100)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// log.Println("test: ", u.Id, u.Name)

	defer dao.CloseDB()

	u, err := dao.QueryUserByID(100)
	if dao.IsNoRow(err) { // 这里是sql NoRow error
		log.Fatal(err)
	} else if err != nil {

	}

	log.Println("test: ", u.Id, u.Name)

}
