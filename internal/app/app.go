package app

import (
	"virtual-bank/internal/delivery/controllers"
	"virtual-bank/internal/delivery/middleware"
	"virtual-bank/internal/delivery/routers"
	"virtual-bank/internal/helpers"
	"virtual-bank/internal/repository"
	"virtual-bank/internal/usecase"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type BootstrapConfig struct {
	DB  *gorm.DB
	App *gin.Engine
	Log *helpers.Logger
}

func AppBootstrap(conf *BootstrapConfig) {

	// validator
	val := validator.New()
	validation := helpers.NewValidation(val)

	// repositorys
	userRepo := repository.NewUserRepository()
	accoutTypeRepo := repository.NewAccountTypeRepository()
	accoutRepo := repository.NewAccountRepository()
	transactionRepo := repository.NewTransactionRepository()
	sessionRepo := repository.NewSessionRepository()

	// usecases
	userUsecase := usecase.NewUserUsecase(conf.DB, userRepo, conf.Log)
	accountTypeUsecase := usecase.NewAccountTypeUsecase(conf.DB, accoutTypeRepo, conf.Log)
	accountUsecase := usecase.NewAccountUsecase(conf.DB, accoutRepo, accoutTypeRepo, userRepo, conf.Log)
	transactionUsecase := usecase.NewTransactionUsecase(conf.DB, transactionRepo, accoutRepo, conf.Log)
	sessionUsecase := usecase.NewSessionUsecase(conf.DB, sessionRepo, userRepo, conf.Log)

	// controllers
	authContoller := controllers.NewAuthController(userUsecase, sessionUsecase, validation)
	userController := controllers.NewUserController(userUsecase, validation)
	accountTypeController := controllers.NewAccountTypeController(accountTypeUsecase)
	accountController := controllers.NewAccountController(accountUsecase, validation)
	transactionController := controllers.NewTransactionController(transactionUsecase, accountUsecase, validation)

	// middleware
	middleware := middleware.NewMiddlewares(sessionUsecase, userUsecase)

	routerConfig := routers.Routers{
		App:                   conf.App.Group("/api/v1"),
		AuthController:        authContoller,
		UserController:        userController,
		AccountController:     accountController,
		AccountTypeController: accountTypeController,
		TransactionController: transactionController,
		Middleware:            middleware,
	}

	routerConfig.SetUp()
}
