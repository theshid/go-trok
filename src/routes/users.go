package routes

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4"
<<<<<<< HEAD:src/routes/users.go
	"github.com/theshid/go-trok/src/models"

=======
	"github.com/theshid/go-trok/models"
>>>>>>> 72c6936d1eed77ad5d46b9887202e9de69f48832:routes/users.go
	"net/http"
)

func UsersLogin(c *gin.Context) {
	user := models.User{}
	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	db, _ := c.Get("db")
	conn := db.(pgx.Conn)
	err = user.IsAuthenticated(&conn)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	token, err := user.GetAuthToken()
	if err == nil {
		c.JSON(http.StatusOK, gin.H{
			"token": token,
		})
		return
	}

	c.JSON(http.StatusBadRequest, gin.H{
		"error": "There was an error authenticating.",
	})
}

func UserRegister(c *gin.Context) {
	user := models.User{}
	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	db, _ := c.Get("db")
	conn := db.(pgx.Conn)
	err = user.Register(&conn)
	if err != nil {
		fmt.Println("Error in user.Register()")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	token, err := user.GetAuthToken()
	if err == nil {
		c.JSON(http.StatusOK, gin.H{
			"token": token,
		})
		return
	}


	c.JSON(http.StatusOK, gin.H{
		"user_id": user.ID,
	})
}
