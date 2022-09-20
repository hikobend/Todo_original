package main

import (
	"fmt"
	"todo/app/models"
)

func main() {
	fmt.Println(models.Db)

	u := models.GetUser(1)
	fmt.Println(u)
}
