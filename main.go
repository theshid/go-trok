package main

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4"
	"go-trok/routes"
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
