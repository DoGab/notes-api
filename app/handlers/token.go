package handlers

import (
	"time"

	"github.com/dogab/notes-api/app/model"
	"github.com/dogab/notes-api/database"
	"github.com/dogab/notes-api/pkg/config"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

type GetTokenRequest struct {
	Email    string
	Password string
}

func GetToken(c *fiber.Ctx) error {
	req := new(GetTokenRequest)
	if err := c.BodyParser(req); err != nil {
		return err
	}

	if req.Email == "" || req.Password == "" {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "invalid credentials", "data": nil})
	}

	db := database.DBconn
	var user model.User

	// // Read the param userId
	// id := c.Query("userid")

	// Find the user with the given Id
	err := db.Find(&user, "email = ?", req.Email).Error
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "invalid credentials", "data": nil})
	}

	// If no such user present return an error
	if user.ID == 0 {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "invalid input", "data": nil})
	}

	// Compare password with the provided one with post request
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		return c.Status(403).JSON(fiber.Map{"status": "error", "message": "invalid credentials", "data": nil})
	}

	// Create JWT Token
	token, exp, err := createJWTToken(user)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "invalid input", "data": nil})
	}

	// Return the token
	// return c.JSON(fiber.Map{"status": "success", "message": "User Found", "data": {"token": token, "exp": exp, "user": user}})
	return c.JSON(fiber.Map{"token": token, "exp": exp, "user": user})
}

func RefreshToken(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"status": "success", "message": "Notes Found", "data": nil})
}

func createJWTToken(user model.User) (string, int64, error) {
	exp := time.Now().Add(time.Minute * 30).Unix()
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["user_id"] = user.ID
	claims["exp"] = exp
	t, err := token.SignedString(config.SecretKey)
	if err != nil {
		return "", 0, err
	}
	return t, exp, nil
}
