package main

import ("github.com/gin-gonic/gin"
	"math/rand"
    "time"
	"fmt"
)

func main() {

	gin.SetMode(gin.ReleaseMode)

	router := gin.Default()


	router.GET("/ping", func(c *gin.Context) {
		pause_time := randInt(100,1200)
		duration := time.Duration(pause_time)*time.Millisecond
		time.Sleep(duration)
		fmt.Print(duration)

		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	router.Run(":80")
	fmt.Scanln()
}

func randInt(min int, max int) int {
    return min + rand.Intn(max-min)
}