package main

import (
	"Go-000/Week02/dao"
	"fmt"
)

func main() {
	user, err := dao.QueryList()
	if err != nil {
		fmt.Printf("err is %#v", err)
		return
	}
	fmt.Println(user)
}
