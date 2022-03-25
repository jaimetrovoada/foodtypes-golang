package main

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func connectDb() *sql.DB {
	db, err := sql.Open("mysql", "jaime:jaimetdl@/foodtypes")
	if err != nil {
		fmt.Println(err)
	}
	return db
}

type Food struct {
	FoodName       string `json:"foodName"`
	ScientificName string `json:"scientificName"`
	Group          string `json:"group"`
	SubGroup       string `json:"subGroup"`
}

func getAllFoods(c *gin.Context) {
	db := connectDb()
	rows, err := db.Query("SELECT * FROM `mockdata-food`")
	if err != nil {
		fmt.Println(err)
	}
	var allFoodsData []Food
	for rows.Next() {
		var foodData Food
		err = rows.Scan(&foodData.FoodName, &foodData.ScientificName, &foodData.Group, &foodData.SubGroup)
		if err != nil {
			fmt.Println(err)
		}
		allFoodsData = append(allFoodsData, foodData)
		fmt.Println(allFoodsData)
	}

	c.IndentedJSON(http.StatusOK, allFoodsData)

}

func getFood(c *gin.Context) {
	name := c.Param("foodName")
	fmt.Printf("param is, %s", name)
	db := connectDb()
	rows, err := db.Query("SELECT * FROM `mockdata-food` WHERE `FOOD NAME` LIKE CONCAT('%', ?, '%')", name)
	if err != nil {
		fmt.Println(err)
	}
	var data []Food
	for rows.Next() {
		var foodData Food
		err = rows.Scan(&foodData.FoodName, &foodData.ScientificName, &foodData.Group, &foodData.SubGroup)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(foodData)
		data = append(data, foodData)
	}

	if len(data) > 0 {
		c.IndentedJSON(http.StatusOK, data)
	} else {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "food not found"})
	}

	return
}

func main() {
	router := gin.Default()
	router.GET("/foodTypes", getAllFoods)
	router.GET("/foodTypes/:foodName", getFood)
	router.Run("localhost:8080")
}
