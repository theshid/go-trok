package main

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4"
	"github.com/theshid/go-trok/models"
	"go-trok/routes"
	"net/http"
	"strings"
)

func main() {
	conn, err := connectDb()
	if err != nil {
		return
	}
	router := gin.Default()
	router.Use(dbMiddleware(*conn))
	usersGroup := router.Group("users")
	{
		usersGroup.POST("register", routes.UserRegister)
		usersGroup.POST("login", routes.UsersLogin)
	}

	itemsGroup := router.Group("items")
	{
		itemsGroup.GET("index", routes.ItemsIndex)
		itemsGroup.POST("create",authMiddleware(),routes.ItemsCreate)
		itemsGroup.GET("sold_by_user",authMiddleware(),routes.ItemsForSaleByCurrentUser)
	}
	router.Run(":3000")
}

func connectDb() (c *pgx.Conn, err error) {
	conn, err := pgx.Connect(context.Background(), "postgresql://postgres:@localhost:5434/trok")
	if err != nil {
		fmt.Println("Error connecting to DB")
		fmt.Println(err.Error())
	}
	_ = conn.Ping(context.Background())
	return conn, err
}

func dbMiddleware(conn pgx.Conn)gin.HandlerFunc{
	return func (c *gin.Context){
		c.Set("db",conn)
		c.Next()
	}
}

func authMiddleware() gin.HandlerFunc{
	return func(c *gin.Context){
		bearer := c.Request.Header.Get("Authorization")
		split := strings.Split(bearer,"Bearer ")
		if len(split)<2 {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "not authenticated."})
			c.Abort()
			return
		}
		token := split[1]
		//fmt.Printf("Bearer (%v) \n", token)
		isValid, userID := models.IsTokenValid(token)
		if isValid == false {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Not authenticated."})
			c.Abort()
		} else {
			c.Set("user_id", userID)
			c.Next()
		}
	}
}
