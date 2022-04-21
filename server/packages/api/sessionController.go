package api

import (
	"authapp/packages/config"
	userFunctions "authapp/packages/coponents/user/functions"
	userModel "authapp/packages/coponents/user/models"
	"database/sql"
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
)

type Claims struct {
	userModel.User
	jwt.StandardClaims
}

func Ping(c *fiber.Ctx) error {
	return c.SendString("Ping Connected")
}

func CreateUser(c *fiber.Ctx, dbConn *sql.DB) error {
	user := &userModel.User{}

	if err := c.BodyParser(user); err != nil {
		return err
	}

	if errs := userFunctions.ValidateUser(*user); len(errs) > 0 {
		return c.Status(http.StatusUnprocessableEntity).JSON(&fiber.Map{"success": false, "errors": errs})
	}

	if user.HasUserExists(dbConn) {
		return c.Status(400).JSON(&fiber.Map{"success": false, "errors": []string{"email already exists"}})
	}

	user.UpdateHashPassword()
	_, err := dbConn.Query(userModel.CreateUserQuery, user.Name, user.Password, user.Email)
	if err != nil {
		return nil
	}
	return c.JSON(&fiber.Map{"success": true})
	// if errs := 
}

func Login(c *fiber.Ctx, dbConn *sql.DB) error {
	loginUser := &userModel.User{}

	if err := c.BodyParser(loginUser); err != nil {
		return err
	}

	user := &userModel.User{}
	if err := dbConn.QueryRow(userModel.GetUserByEmailQuery, loginUser.Email).
	Scan(&user.ID, &user.Name, &user.Password, &user.Email, &user.CreatedAt, &user.UpdatedAt); err != nil {
		if err == sql.ErrNoRows {
			return c.Status(http.StatusUnauthorized).JSON(fiber.Map{"success": false, "errors": []string{"Incorrect credentials"}})
		}
	}

	match := userFunctions.ComparePassword(user.Password, loginUser.Password)
	if !match {
		return c.Status(http.StatusUnauthorized).JSON(fiber.Map{"success": false, "errors": []string{"Incorrect credentials"}})
	}

	//expiration time of the token ->30 mins
	expirationTime := time.Now().Add(30 * time.Minute)

	user.Password = ""
	claims := &Claims{
		User: *user,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	var jwtKey = []byte(config.ConfigSettings[config.JWT_KEY])
	tokenValue, err := token.SignedString(jwtKey)

	if err != nil {
		return err
	}

	c.Cookie(&fiber.Cookie{
		Name:     "token",
		Value:    tokenValue,
		Expires:  expirationTime,
		Domain:   config.ConfigSettings[config.CLIENT_URL],
		HTTPOnly: true,
	})

	return c.JSON(&fiber.Map{"success": true, "user": claims.User, "token": tokenValue})
}

func Logout(c *fiber.Ctx) error {
	c.ClearCookie()
	return c.SendStatus(http.StatusOK)
}