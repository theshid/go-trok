package routes

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4"
<<<<<<< HEAD:src/routes/items.go
	"github.com/theshid/go-trok/src/models"

=======
	"github.com/theshid/go-trok/models"
>>>>>>> 72c6936d1eed77ad5d46b9887202e9de69f48832:routes/items.go
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

func ItemsUpdate(c *gin.Context) {
	userID := c.GetString("user_id")
	db, _ := c.Get("db")
	conn := db.(pgx.Conn)

	itemSent := models.Item{}
	err := c.ShouldBindJSON(&itemSent)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid form sent"})
		return
	}

	itemBeingUpdated, err := models.FindItemById(itemSent.ID, &conn)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if itemBeingUpdated.SellerID.String() != userID {
		c.JSON(http.StatusForbidden, gin.H{"error": "You are not authorized to update this item"})
		return
	}

	itemSent.SellerID = itemBeingUpdated.SellerID
	err = itemSent.Update(&conn)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"item": itemSent})
}


