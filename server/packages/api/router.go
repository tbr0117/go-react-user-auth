package api

import "github.com/gofiber/fiber/v2"

func linkRoutesToServer(apiRouter fiber.Router) {
	publicRoutes(apiRouter)

	authedRoutes(apiRouter)
}

func publicRoutes(apiRouter fiber.Router) {
	apiRouter.Get("/ping", Ping)
}

func authedRoutes(apiRouter fiber.Router) {

}