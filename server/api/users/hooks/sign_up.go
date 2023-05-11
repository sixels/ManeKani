package hooks

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sixels/manekani/core/domain/user"
	"github.com/sixels/manekani/core/ports"
)

type RegisterUserRequest struct {
	UserID    string    `json:"user_id"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
}

func RegisterUser(userDB ports.UserRepository) gin.HandlerFunc {
	return func(c *gin.Context) {
		log.Println("registering user")
		var data RegisterUserRequest
		if err := c.BindJSON(&data); err != nil {
			c.Error(err)
			c.Status(http.StatusBadRequest)
			return
		}
		log.Printf("%##v\n", data)

		ctx := c.Request.Context()
		exists, err := userDB.Exists(ctx, data.UserID)
		if err != nil {
			c.Error(err)
			c.Status(http.StatusInternalServerError)
			return
		}
		if !exists {
			createUserReq := user.CreateUserRequest{
				ID:        data.UserID,
				Email:     data.Email,
				Username:  "banana",
				CreatedAt: &data.CreatedAt,
			}
			u, err := userDB.CreateUser(ctx, createUserReq)
			if err != nil {
				c.Error(err)
				c.Status(http.StatusInternalServerError)
				return
			}
			log.Printf("user created. id: %v\n", u.ID)
		} else {
			log.Println("user already exists")
		}

		c.Status(http.StatusOK)
	}
}
