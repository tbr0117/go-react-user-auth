package api

import (
	"authapp/packages/db"

	"github.com/apex/log"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/requestid"
)

// import

func StartServer() {
	// fiber

	appServer := startHttpServer()
	port := ":8084"
	serverErr := appServer.Listen(port)

	if serverErr != nil {
		log.WithField("reason", serverErr.Error()).Fatal("Server error")
	}
}

func startHttpServer() *fiber.App {
	dbConn, err := db.ConnectDB()
	if err != nil {
		log.WithField("reason", err.Error()).Fatal("Db connection error occurred")
	}
	defer dbConn.Close()

	app := fiber.New()
	app.Use(logger.New())
	app.Use(requestid.New())

	apiRouter := app.Group("/api")

	apiRouter.Use(cors.New(cors.Config{
		AllowOrigins:     "http://localhost:3000",
		AllowCredentials: true,
		AllowHeaders:     "Content-Type, Content-Length, Accept-Encoding, Authorization, accept, origin",
		AllowMethods:     "GET",
		ExposeHeaders:    "Set-Cookie",
	}))

	linkRoutesToServer(apiRouter, dbConn)

	return app
}
