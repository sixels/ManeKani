package hooks

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"net/http"
	"time"

	"github.com/sixels/manekani/core/domain/user"
	"github.com/sixels/manekani/core/ports"
)

type RegisterUserRequest struct {
	UserID    string    `json:"user_id"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
}

func RegisterUser(userDB ports.UserRepository) echo.HandlerFunc {
	return func(c echo.Context) error {
		log.Info("registering user")
		var data RegisterUserRequest
		if err := c.Bind(&data); err != nil {
			c.Error(err)
			return c.NoContent(http.StatusBadRequest)
		}
		log.Infof("%##v\n", data)

		ctx := c.Request().Context()
		exists, err := userDB.Exists(ctx, data.UserID)
		if err != nil {
			c.Error(err)
			return c.NoContent(http.StatusInternalServerError)
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
				return c.NoContent(http.StatusInternalServerError)
			}
			log.Infof("user created. id: %v\n", u.ID)
		} else {
			log.Info("user already exists")
		}

		return c.NoContent(http.StatusOK)
	}
}
