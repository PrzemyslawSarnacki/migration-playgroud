package main

import (
	"fmt"

	"gorm.io/datatypes"
)

func main() {
	user := User{Name: "Ken",
		Data: datatypes.JSON([]byte(`{"foo1": "bar1", "foo2": "bar3"}`))}
	var result User
	DB.First(&result, user.ID)

	DB.Create(&user)
	fmt.Println(result.Name)

}
