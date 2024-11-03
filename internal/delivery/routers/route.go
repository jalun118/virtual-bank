package routers

import (
	"virtual-bank/internal/delivery/controllers"
	"virtual-bank/internal/delivery/middleware"

	"github.com/gin-gonic/gin"
)

type Routers struct {
	App                   *gin.RouterGroup
	AuthController        *controllers.AuthController
	UserController        *controllers.UserController
	AccountTypeController *controllers.AccountTypeController
	AccountController     *controllers.AccountController
	TransactionController *controllers.TransactionController
	Middleware            *middleware.Middlewares
}

func (r *Routers) SetUp() {
	r.SetupAuthRoute()
	r.SetupUserRoute()
	r.SetupAccountRoute()
	r.SetupAccountTypeRoute()
	r.SetupTransactionRoute()
}

func (r *Routers) SetupAuthRoute() {
	g := r.App.Group("/auth")
	g.POST("/register", r.Middleware.NotLogin(), r.AuthController.RegisterUser)
	g.POST("/login", r.Middleware.NotLogin(), r.AuthController.LoginUser)
	g.GET("/refresh", r.Middleware.Auth(), r.AuthController.SetRefreshToken)
	g.DELETE("/logout", r.AuthController.LogOut)
}

func (r *Routers) SetupUserRoute() {
	g := r.App.Group("/user")
	g.Use(r.Middleware.TokenGuard())
	g.GET("", r.UserController.GetMe)
	g.PUT("", r.UserController.Update)
	g.PUT("/password", r.UserController.ChangePassword)
}

func (r *Routers) SetupAccountTypeRoute() {
	g := r.App.Group("/account-type")
	g.Use(r.Middleware.TokenGuard())
	g.GET("", r.AccountTypeController.GetAllType)
	g.GET("/:id", r.AccountTypeController.GetOne)
}

func (r *Routers) SetupAccountRoute() {
	g := r.App.Group("/account")
	g.POST("/top-up", r.AccountController.TopAccount)

	t := g.Use(r.Middleware.TokenGuard())
	t.POST("", r.AccountController.CreateAccount)
	t.GET("", r.AccountController.GetMyAccount)
	t.POST("/pin", r.AccountController.MatchPinAccount)
	t.GET("/:id", r.AccountController.GetAccount)
}

func (r *Routers) SetupTransactionRoute() {
	g := r.App.Group("/transaction")
	g.Use(r.Middleware.TokenGuard())
	g.POST("", r.TransactionController.CreateTransaction)
	g.GET("", r.TransactionController.GetAllTransaction)
	g.GET("/submition", r.TransactionController.GetAllSubmition)
	g.GET("/:id", r.TransactionController.GetOne)
}
