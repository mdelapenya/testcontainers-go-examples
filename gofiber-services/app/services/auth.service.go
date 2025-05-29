package services

import (
	"errors"

	"github.com/gofiber/fiber/v3"
	"gorm.io/gorm"

	"github.com/mdelapenya/testcontainers-go-examples/gofiber-services/app/dal"
	"github.com/mdelapenya/testcontainers-go-examples/gofiber-services/app/types"
	"github.com/mdelapenya/testcontainers-go-examples/gofiber-services/config/database"
	"github.com/mdelapenya/testcontainers-go-examples/gofiber-services/utils"
	"github.com/mdelapenya/testcontainers-go-examples/gofiber-services/utils/jwt"
	"github.com/mdelapenya/testcontainers-go-examples/gofiber-services/utils/password"
)

// Login service logs in a user
func Login(ctx fiber.Ctx) error {
	b := new(types.LoginDTO)

	if err := utils.ParseBodyAndValidate(ctx, b); err != nil {
		return err
	}

	u := &types.UserResponse{}

	err := dal.FindUserByEmail(database.DB, u, b.Email).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return fiber.NewError(fiber.StatusUnauthorized, "Invalid email or password")
	}

	if err := password.Verify(u.Password, b.Password); err != nil {
		return fiber.NewError(fiber.StatusUnauthorized, "Invalid email or password")
	}

	t := jwt.Generate(&jwt.TokenPayload{
		ID: u.ID,
	})

	return ctx.JSON(&types.AuthResponse{
		User: u,
		Auth: &types.AccessResponse{
			Token: t,
		},
	})
}

// Signup service creates a user
func Signup(ctx fiber.Ctx) error {
	b := new(types.SignupDTO)

	if err := utils.ParseBodyAndValidate(ctx, b); err != nil {
		return err
	}

	err := dal.FindUserByEmail(database.DB, &struct{ ID string }{}, b.Email).Error

	// If email already exists, return
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return fiber.NewError(fiber.StatusConflict, "Email already exists")
	}

	user := &dal.User{
		Name:     b.Name,
		Password: password.Generate(b.Password),
		Email:    b.Email,
	}

	uid := uint64(user.ID)

	// Create a user, if error return
	if err := dal.CreateUser(database.DB, user); err.Error != nil {
		return fiber.NewError(fiber.StatusConflict, err.Error.Error())
	}

	// generate access token
	t := jwt.Generate(&jwt.TokenPayload{
		ID: uid,
	})

	return ctx.JSON(&types.AuthResponse{
		User: &types.UserResponse{
			ID:    uid,
			Name:  user.Name,
			Email: user.Email,
		},
		Auth: &types.AccessResponse{
			Token: t,
		},
	})
}
