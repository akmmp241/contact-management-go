//go:build wireinject
// +build wireinject

package main

import (
	"contact-management-restful/app"
	controllerContracts "contact-management-restful/controller/contracts"
	"contact-management-restful/controller/impl"
	"contact-management-restful/middleware"
	"contact-management-restful/repositories/contracts"
	repostoryimpl "contact-management-restful/repositories/impl"
	serviceContracts "contact-management-restful/services/contracts"
	serviceimpl "contact-management-restful/services/impl"
	"github.com/google/wire"
	"net/http"
)

var userControllerSet = wire.NewSet(
	repostoryimpl.NewUserRepositoryImpl,
	wire.Bind(new(contracts.UserRepository), new(*repostoryimpl.UserRepositoryImpl)),
	serviceimpl.NewUserServiceImpl,
	wire.Bind(new(serviceContracts.UserService), new(*serviceimpl.UserServiceImpl)),
	impl.NewUserController,
	wire.Bind(new(controllerContracts.UserController), new(*impl.UserControllerImpl)),
)

var contactControllerSet = wire.NewSet(
	repostoryimpl.NewContactRepositoryImpl,
	wire.Bind(new(contracts.ContactRepository), new(*repostoryimpl.ContactRepositoryImpl)),
	serviceimpl.NewContactServiceImpl,
	wire.Bind(new(serviceContracts.ContactService), new(*serviceimpl.ContactServiceImpl)),
	impl.NewContactControllerImpl,
	wire.Bind(new(controllerContracts.ContactController), new(*impl.ContactControllerImpl)),
)

var addressControllerSet = wire.NewSet(
	repostoryimpl.NewAddressRepositoryImpl,
	wire.Bind(new(contracts.AddressRepository), new(*repostoryimpl.AddressRepositoryImpl)),
	serviceimpl.NewAddressServiceImpl,
	wire.Bind(new(serviceContracts.AddressService), new(*serviceimpl.AddressServiceImpl)),
	impl.NewAddressControllerImpl,
	wire.Bind(new(controllerContracts.AddressController), new(*impl.AddressControllerImpl)),
)

func InitializedServer() *http.Server {
	wire.Build(
		app.NewDB,
		NewValidator,
		NewServer,
		userControllerSet,
		contactControllerSet,
		addressControllerSet,
		app.NewControllers,
		middleware.NewAuthMiddleware,
		app.NewRouter,
		app.NewConfig,
	)
	return nil
}
