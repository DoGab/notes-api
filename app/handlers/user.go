package handlers

import (
	"github.com/dogab/notes-api/app/model"
	"github.com/dogab/notes-api/database"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

func GetUsers(c *fiber.Ctx) error {
	db := database.DBconn
	var users []model.User

	// find all users in the database
	db.Find(&users)

	// If no User is present return an error
	if len(users) == 0 {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "No users present", "data": nil})
	}

	// Else return users
	return c.JSON(fiber.Map{"status": "success", "message": "Users Found", "data": users})
}

func CreateUser(c *fiber.Ctx) error {
	db := database.DBconn
	user := new(model.User)

	// Store the body in the user and return error if encountered
	err := c.BodyParser(user)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Invalid request, review your input", "data": err})
	}
	// Add a uuid to the user
	// user.ID = uuid.New()

	if user.Name == "" || user.Email == "" || user.Password == "" {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Invalid request, review your input", "data": err})
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hash)

	// Create the User and return error if encountered
	err = db.Create(&user).Error
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Could not create user", "data": err})
	}

	// Return the created user
	return c.JSON(fiber.Map{"status": "success", "message": "Created User", "data": user})
}

func GetUser(c *fiber.Ctx) error {
	db := database.DBconn
	var user model.User

	id := c.Params("userId")
	if err := db.Find(&user, "id = ?", id).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "No user present", "data": nil})
	}

	if user.ID == 0 {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "No user present", "data": nil})
	}

	return c.JSON(fiber.Map{"status": "success", "message": "Notes Found", "data": user})
}

func UpdateUser(c *fiber.Ctx) error {
	type updateUserRequest struct {
		Name      string `json:"name"`
		Email     string `json:"email"`
		Password1 string `json:"password1"`
		Password2 string `json:"password2"`
	}

	db := database.DBconn
	var user model.User

	id := c.Params("userId")
	db.Find(&user, "id = ?", id)

	if user.ID == 0 {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "No user present", "data": nil})
	}

	req := new(updateUserRequest)
	if err := c.BodyParser(req); err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Review your input", "data": err})
	}

	if req.Password1 != req.Password2 {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Passwords do not match", "data": nil})
	}

	if err := db.Find(&user, "id = ?", id).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "No user present", "data": nil})
	}

	user.Email = req.Email
	user.Name = req.Name
	hash, err := bcrypt.GenerateFromPassword([]byte(req.Password1), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hash)

	db.Save(&user)

	return c.JSON(fiber.Map{"status": "success", "message": "User updated", "data": user})
}
