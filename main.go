package main

import (
	"github.com/gin-gonic/gin"

	auth "github.com/hafizdkn/toko-belanja/apps/domain/auth/jwt"
	"github.com/hafizdkn/toko-belanja/apps/domain/auth/middleware"
	"github.com/hafizdkn/toko-belanja/apps/domain/category"
	"github.com/hafizdkn/toko-belanja/apps/domain/db"
	"github.com/hafizdkn/toko-belanja/apps/domain/handler"
	"github.com/hafizdkn/toko-belanja/apps/domain/product"
	transactionhistory "github.com/hafizdkn/toko-belanja/apps/domain/transactionHistory"
	"github.com/hafizdkn/toko-belanja/apps/domain/user"
)

func main() {
	db, err := db.ConnectionDB()
	if err != nil {
		panic(err)
	}

	authService := auth.NewJwtService

	userRepo := user.NewUserRepository(db)
	userSvc := user.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userSvc)

	categoryRepo := category.NewCategoryRepository(db)
	categorySvc := category.NewCategoryService(categoryRepo)
	categoryHandler := handler.NewCategoryHandler(categorySvc)

	productRepo := product.NewProductRepository(db)
	productSvc := product.NewServiceProduct(productRepo, categorySvc)
	productHandler := handler.NewProductHandler(productSvc)

	trasactionRepo := transactionhistory.NewTransactionRepository(db)
	trasactionSvc := transactionhistory.NewTransactionService(productSvc, userSvc, trasactionRepo)
	trasactionHandler := handler.NewTransactionHandler(trasactionSvc)

	app := gin.Default()

	users := app.Group("/users")
	{
		users.GET("/", userHandler.GetUsers)
		users.POST("/register", userHandler.UserRegister)
		users.POST("/login", userHandler.UserLogin)
		users.Use(middleware.AuthenticationMiddleware(authService))
		users.PATCH("/topup", userHandler.UserUpdateBalance)
	}

	categorys := app.Group("/categories")
	{
		categorys.GET("/", categoryHandler.GetCategorys)
		categorys.Use(middleware.AuthenticationMiddleware(authService), middleware.AuthorizationMiddleware(userSvc))
		categorys.POST("/", categoryHandler.CreateCategory)
		categorys.PATCH("/:categoryId", categoryHandler.UpdateCategory)
		categorys.DELETE("/:categoryId", categoryHandler.DeleteCategory)
	}

	products := app.Group("/products")
	{
		products.GET("/", productHandler.GetProduct)
		products.Use(middleware.AuthenticationMiddleware(authService), middleware.AuthorizationMiddleware(userSvc))
		products.POST("/", productHandler.CreateProduct)
		products.PUT("/:productId", productHandler.UpdateProduct)
		products.DELETE("/:productId", productHandler.DeleteProduct)
	}

	transactions := app.Group("/transactions")
	{
		transactions.Use(middleware.AuthenticationMiddleware(authService))
		transactions.POST("/", trasactionHandler.Transaction)
		transactions.GET("/user-transaction", trasactionHandler.GetUserTransactions)
		transactions.GET("/my-transaction", trasactionHandler.GetMyTransactions)

	}
	app.Run(":8080")
}
