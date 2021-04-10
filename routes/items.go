package routes

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4"
	"go-trok/models"
	"net/http"
)

func ItemsIndex(c *gin.Context) {
	db, _ := c.Get("db")
	conn := db.(pgx.Conn)
	items, err := models.GetAllItems(&conn)
	if err != nil {
		fmt.Println(err)
	}
	c.JSON(http.StatusOK, gin.H{"items": items})
}

func ItemsCreate(c *gin.Context) {
	userID := c.GetString("user_id")
	db, _ := c.Get("db")
	conn := db.(pgx.Conn)

	item := models.Item{}
	c.ShouldBindJSON(&item)
	err := item.Create(&conn, userID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, item)
}

func ItemsForSaleByCurrentUser(c *gin.Context) {
	userID := c.GetString("user_id")
	db, _ := c.Get("db")
	conn := db.(pgx.Conn)

	items, err := models.GetItemsBeingSoldByUser(userID, &conn)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"items": items})
}
