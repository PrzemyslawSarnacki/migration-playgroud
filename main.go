package main

import (
	"fmt"

	"github.com/PrzemyslawSarnacki/migration-playgroud/models"
	"gorm.io/datatypes"
)

func main() {
	user := models.User{Name: "Ken",
		Data: datatypes.JSON([]byte(`{"foo1": "bar1", "foo2": "bar3"}`))}
	var result models.User
	DB.First(&result, user.ID)

	DB.Create(&user)
	fmt.Println(result.Name)

}
