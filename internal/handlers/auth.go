	package handlers
	import (
		"github.com/gofiber/fiber/v2"
		"gorm.io/gorm"
		"golang.org/x/crypto/bcrypt"
		"auth-go/internal/database"
		"auth-go/internal/utils"

	)


func Register(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var input struct {
			Name     string `json:"name"`
			Email    string `json:"email"`
			Password string `json:"password"`
		}
		if err := c.BodyParser(&input); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"Error": err.Error()})
		}

		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"Error": err.Error()})
		}

		user := database.User{
			Name:     input.Name,
			Email:    input.Email,
			Password: string(hashedPassword),
		}
		if err := db.Create(&user).Error; err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"Error": err.Error()})
		}

		return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "User registered successfully"})
	}
}

func Login(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var input struct {
			Email    string `json:"email"`
			Password string `json:"password"`
		}
		if err := c.BodyParser(&input); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"Error": err.Error()})
		}

		var user database.User
		if err := db.Where("email = ?", input.Email).First(&user).Error; err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"Error": "User not found"})
		}

		if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"Error": "Invalid password"})
		}

		token, err := utils.GenerateJWT(user.Email)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"Error": err.Error()})
		}

		return c.Status(fiber.StatusOK).JSON(fiber.Map{"token": token})
	}
}

func Logout() fiber.Handler {
	return func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"message": "Logout successful"})
	}
}