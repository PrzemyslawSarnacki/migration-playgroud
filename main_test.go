package main

import (
	// "github.com/PrzemyslawSarnacki/migration-playgroud/migrations"
	"github.com/PrzemyslawSarnacki/migration-playgroud/models"
	"testing"

	"gorm.io/datatypes"
)

func TestGORM(t *testing.T) {
	// user := User{
	// 	Name: "Ken",
	// 	Data: datatypes.JSON([]byte(`{"foo1": "bar1", "foo2": "bar3"}`)),
	// }

	// DB.Create(&user)

	t.Run("test records after migration", func(t *testing.T) {
		var results []models.User
		if err := DB.Find(&results).Error; err != nil {
			t.Errorf("Failed, got error: %v", err)
		}
		assertInt(len(results), 2, t)
		assertString(results[0].Name, "Ken", t)
	})

	// t.Run("test integrity", func(t *testing.T) {
	// 	var result User
	// 	if err := DB.First(&result, user.ID).Error; err != nil {
	// 		t.Errorf("Failed, got error: %v", err)
	// 	}
	// 	want := "Ken"
	// 	assertString(result.Name, want, t)
	// })

	t.Run("test one condition", func(t *testing.T) {
		var users []models.User

		query := datatypes.JSONQuery("data").Equals("bar3", "foo2")

		if err := DB.Where(query).Find(&users).Error; err != nil {
			t.Errorf("Failed, got error: %v", err)
		}
		want := "Ken"
		assertString(users[0].Name, want, t)
	})

	t.Run("test both conditions", func(t *testing.T) {
		var users []models.User
		DB.Create(&models.User{Name: "Jack",
			Data: datatypes.JSON([]byte(`{"foo1": "bar1", "foo2": "bar2"}`))})

		query1 := datatypes.JSONQuery("data").Equals("bar1", "foo1")
		query2 := datatypes.JSONQuery("data").Equals("bar2", "foo2")

		if err := DB.Where(query1, query2).Find(&users).Error; err != nil {
			t.Errorf("Got error %v", err)
		}
		want := "Jack"
		assertString(users[0].Name, want, t)
	})

}

func assertString(got, want string, t testing.TB) {
	t.Helper()
	if got != want {
		t.Errorf(`got "%s", want "%s"`, got, want)
	}
}
func assertInt(got, want int, t testing.TB) {
	t.Helper()
	if got != want {
		t.Errorf(`got "%d", want "%d"`, got, want)
	}
}
