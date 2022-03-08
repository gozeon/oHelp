package lib

import (
	"fmt"
	"testing"
)

/**
 * bash: go test oHelp/lib -v -run TestDB
 */
func TestDB(t *testing.T) {

	InitDb()

	// DBClient.Create(&model.Order{
	// 	Title:       "test",
	// 	Description: "test",
	// 	Tags:        "test,test",
	// 	Status:      true,
	// 	CreateUser:  "wzy",
	// })

	// var data []model.Order
	// result := DBClient.Find(&data)
	// fmt.Println(result.RowsAffected)

	fmt.Println(DBClient)
}
