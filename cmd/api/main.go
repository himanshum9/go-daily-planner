package main

import (
	"flag"
	"html/template"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/himanshu/daily-planner/internal/config"
	"github.com/himanshu/daily-planner/internal/repository"
	"github.com/himanshu/daily-planner/internal/routes"
	"github.com/himanshu/daily-planner/pkg/middleware"
	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	// Parse command line flags
	migrate := flag.Bool("migrate", false, "Run database migrations")
	flag.Parse()

	// Initialize configuration
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	// Initialize database
	db, err := repository.NewDatabase()
	if err != nil {
		log.Fatal("Failed to initialize database:", err)
	}

	// Run migrations if flag is set
	if *migrate {
		if err := db.Migrate(); err != nil {
			log.Fatal("Failed to run migrations:", err)
		}
		log.Println("Migrations completed successfully")
		os.Exit(0)
	}

	// Create Gin router
	r := gin.Default()

	// Set up middleware
	r.Use(middleware.CORS())
	r.Use(middleware.SessionAuth())

	// Add template functions
	r.SetFuncMap(template.FuncMap{
		"add": func(a, b int) int {
			return a + b
		},
	})

	// Load templates
	r.LoadHTMLGlob("templates/**/*.html")
	log.Printf("Loading templates from: templates/**/*.html")
	r.Static("/static", "./static")

	// Set up routes
	routes.SetupRoutes(r, db)

	// Start server
	log.Printf("Server starting on %s", cfg.ServerAddress)
	if err := r.Run(cfg.ServerAddress); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
