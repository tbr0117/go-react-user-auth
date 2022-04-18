package api

import (
	"database/sql"

	"github.com/gofiber/fiber/v2"
)

func linkRoutesToServer(apiRouter fiber.Router, dbConn *sql.DB) {
	publicRoutes(apiRouter, dbConn)

	authedRoutes(apiRouter, dbConn)
}

func publicRoutes(apiRouter fiber.Router, dbConn *sql.DB) {
	apiRouter.Get("/ping", Ping)

	apiRouter.Post("/login", WithDB(Login, dbConn))
	apiRouter.Post("/refister", WithDB(CreateUser, dbConn))
	apiRouter.Get("/logot", Logout)

}

func authedRoutes(apiRouter fiber.Router, dbConn *sql.DB) {

}

func WithDB(apiFunction func(c *fiber.Ctx, db *sql.DB) error, db *sql.DB) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		return apiFunction(c, db)
	}
}
