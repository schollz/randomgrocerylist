package main

import (
	"bufio"
	"log"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("templates/*")
	r.GET("/*num", func(c *gin.Context) {
		num := c.Param("num")
		numVal, err := strconv.Atoi(num[1:])
		var foods []string
		if err == nil {
			foods = RandomGroceryList(numVal)
		} else {
			foods = RandomGroceryList(10)
		}
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"foods": foods,
		})
	})
	r.Run(":8064")
}

// RandomGroceryList lists some random groceries
func RandomGroceryList(num int) []string {
	rand.Seed(time.Now().Unix())
	file, err := os.Open("sr28/FOOD_DES.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	foods := make([]string, 8789)
	i := 0
	for scanner.Scan() {
		items := strings.Split(scanner.Text(), "~")
		foods[i] = items[5]
		i++
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	randomFoods := make([]string, num)
	for i := 0; i < num; i++ {
		randomFoods[i] = foods[rand.Intn(len(foods))]
	}
	return randomFoods
}
