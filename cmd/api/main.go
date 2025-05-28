// // cmd/api/main.go
// package main

// import (
// 	"log"

// 	"github.com/gin-gonic/gin"

// 	"digital-market/config"
// 	"digital-market/internal/handlers"
// 	"digital-market/internal/middleware"
// 	"digital-market/internal/models"
// 	"digital-market/internal/repositories"
// 	"digital-market/internal/services"
// 	"digital-market/pkg/database"
// )

// func main() {
// 	// Загрузка конфигурации
// 	cfg := config.LoadConfig()

// 	// Подключение к базе данных
// 	db, err := database.NewPostgresDB(database.Config{
// 		Host:     cfg.DB.Host,
// 		Port:     cfg.DB.Port,
// 		User:     cfg.DB.User,
// 		Password: cfg.DB.Password,
// 		DBName:   cfg.DB.Name,
// 		SSLMode:  cfg.DB.SSLMode,
// 	})
// 	if err != nil {
// 		log.Fatalf("Failed to initialize database: %v", err)
// 	}

// 	// Автомиграция моделей
// 	if err := db.AutoMigrate(
// 		&models.User{},
// 		&models.Category{},
// 		&models.Product{},
// 		&models.Cart{},
// 		&models.CartItem{},
// 		&models.Order{},
// 		&models.OrderItem{},
// 		&models.Review{},
// 		&models.Address{},
// 		&models.Payment{},
// 		&models.Delivery{},
// 	); err != nil {
// 		log.Fatalf("Failed to migrate database: %v", err)
// 	}

// 	// Инициализация репозиториев
// 	userRepo := repositories.NewUserRepository(db)
// 	productRepo := repositories.NewProductRepository(db)
// 	// ... другие репозитории

// 	// Инициализация сервисов
// 	authService := services.NewAuthService(userRepo, cfg.JWT.Secret, cfg.JWT.ExpiresIn)
// 	productService := services.NewProductService(productRepo)
// 	// ... другие сервисы

// 	// Инициализация обработчиков
// 	authHandler := handlers.NewAuthHandler(authService)
// 	productHandler := handlers.NewProductHandler(productService)
// 	// ... другие обработчики

// 	// Инициализация маршрутизатора
// 	router := gin.Default()

// 	// Middleware
// 	authMiddleware := middleware.NewAuthMiddleware(cfg.JWT.Secret)

// 	// Группа API v1
// 	apiV1 := router.Group("/api/v1")

// 	// Маршруты аутентификации
// 	auth := apiV1.Group("/auth")
// 	{
// 		auth.POST("/register", authHandler.Register)
// 		auth.POST("/login", authHandler.Login)
// 		auth.POST("/refresh", authHandler.RefreshToken)
// 	}

// 	// Публичные маршруты для продуктов
// 	products := apiV1.Group("/products")
// 	{
// 		products.GET("", productHandler.List)
// 		products.GET("/:id", productHandler.GetByID)
// 		// Другие маршруты, требующие аутентификации
// 	}

// 	// Защищенные маршруты для пользователей
// 	user := apiV1.Group("/users")
// 	user.Use(authMiddleware.RequireAuth())
// 	{
// 		user.GET("/me", userHandler.GetProfile)
// 		// ... другие маршруты пользователя
// 	}

// 	// Административные маршруты
// 	admin := apiV1.Group("/admin")
// 	admin.Use(authMiddleware.RequireAdmin())
// 	{
// 		// Управление продуктами
// 		admin.POST("/products", productHandler.Create)
// 		admin.PUT("/products/:id", productHandler.Update)
// 		admin.DELETE("/products/:id", productHandler.Delete)
// 		// ... другие административные маршруты
// 	}

// 	// Запуск сервера
// 	if err := router.Run(":" + cfg.Server.Port); err != nil {
// 		log.Fatalf("Failed to start server: %v", err)
// 	}
// }

// Файл: cmd/api/main.go
package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	// "gorm.io/gorm"

	"digital-market/config"
	"digital-market/internal/models"
	"digital-market/pkg/database"
)

func main() {
	// Загрузка конфигурации
	cfg := config.LoadConfig()

	// Подключение к базе данных
	db, err := database.NewPostgresDB(database.Config{
		Host:     cfg.DB.Host,
		Port:     cfg.DB.Port,
		User:     cfg.DB.User,
		Password: cfg.DB.Password,
		DBName:   cfg.DB.Name,
		SSLMode:  cfg.DB.SSLMode,
	})
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}

	// Автомиграция моделей (только Product для теста)
	if err := db.AutoMigrate(&models.Product{}); err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}

	// Инициализация роутера
	router := gin.Default()

	// Тестовый эндпоинт для проверки работоспособности
	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "ok",
		})
	})

	// Тестовый эндпоинт для проверки базы данных
	router.GET("/products/test", func(c *gin.Context) {
		var products []models.Product
		
		// Создаем тестовый продукт
		testProduct := models.Product{
			Name:        "Test Product",
			Slug:        "test-product",
			Description: "This is a test product",
			Price:       99.99,
			Stock:       10,
		}
		
		// Сохраняем в БД
		result := db.Create(&testProduct)
		if result.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": result.Error.Error(),
			})
			return
		}
		
		// Получаем все продукты
		db.Find(&products)
		
		c.JSON(http.StatusOK, products)
	})

	// Запуск сервера
	if err := router.Run(":" + cfg.Server.Port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
