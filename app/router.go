package app

import (
	"contact-management-restful/controller/contracts"
	"contact-management-restful/exception"
	"contact-management-restful/middleware"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type Controllers struct {
	contracts.UserController
	contracts.ContactController
	contracts.AddressController
}

func NewControllers(userController contracts.UserController, contactController contracts.ContactController, addressController contracts.AddressController) *Controllers {
	return &Controllers{UserController: userController, ContactController: contactController, AddressController: addressController}
}

func NewRouter(authMiddleware *middleware.AuthMiddleware, controllers *Controllers) *httprouter.Router {
	router := httprouter.New()

	router.GET("/ping", func(writer http.ResponseWriter, request *http.Request, _ httprouter.Params) {
		writer.Write([]byte("pong"))
	})

	registerUserRouter(router, controllers.UserController, authMiddleware)
	registerContactRouter(router, controllers.ContactController, authMiddleware)
	registerAddressRouter(router, controllers.AddressController, authMiddleware)

	router.PanicHandler = exception.ErrorHandler

	return router
}

func registerUserRouter(router *httprouter.Router, userController contracts.UserController, authMiddleware *middleware.AuthMiddleware) {
	router.POST("/api/users", userController.Register)
	router.POST("/api/users/login", userController.Login)
	router.GET("/api/users/", authMiddleware.ApiAuthMiddleware(userController.Current))
	router.PATCH("/api/users/", authMiddleware.ApiAuthMiddleware(userController.Update))
	router.DELETE("/api/users/", authMiddleware.ApiAuthMiddleware(userController.Logout))
}

func registerContactRouter(router *httprouter.Router, controller contracts.ContactController, authMiddleware *middleware.AuthMiddleware) {
	router.POST("/api/contact", authMiddleware.ApiAuthMiddleware(controller.Create))
	router.GET("/api/contact", authMiddleware.ApiAuthMiddleware(controller.List))
	router.PUT("/api/contact/:id", authMiddleware.ApiAuthMiddleware(controller.Update))
	router.GET("/api/contact/:id", authMiddleware.ApiAuthMiddleware(controller.Get))
	router.DELETE("/api/contact/:id", authMiddleware.ApiAuthMiddleware(controller.Delete))
}

func registerAddressRouter(router *httprouter.Router, controller contracts.AddressController, authMiddleware *middleware.AuthMiddleware) {
	router.POST("/api/contacts/:contactId/addresses", authMiddleware.ApiAuthMiddleware(controller.Create))
	router.GET("/api/contacts/:contactId/addresses", authMiddleware.ApiAuthMiddleware(controller.List))
	router.GET("/api/contacts/:contactId/addresses/:addressId", authMiddleware.ApiAuthMiddleware(controller.Get))
	router.PUT("/api/contacts/:contactId/addresses/:addressId", authMiddleware.ApiAuthMiddleware(controller.Update))
	router.DELETE("/api/contacts/:contactId/addresses/:addressId", authMiddleware.ApiAuthMiddleware(controller.Delete))
}
