package main

import (
	"crypto/md5"
	"fmt"
)

func main() {

	// adapter := database.InitDB()

	// defer adapter.Close()

	// repo := rel.New(adapter)

	// repo.Ping(context.TODO())

	// var user model.User

	// if err := repo.Find(context.Background(), &user, where.Eq("id", 1)); err != nil {
	// 	panic(err)
	// }

	// fmt.Println(user)

	// fmt.Println(user.Username)

	// password := fmt.Sprintf("%x", md5.Sum([]byte(rawPassword)))

	passd := fmt.Sprintf("%x", md5.Sum([]byte("1234")))
	fmt.Println(passd)
}
