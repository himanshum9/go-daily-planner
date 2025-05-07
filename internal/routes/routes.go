package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/himanshu/daily-planner/internal/auth"
	"github.com/himanshu/daily-planner/internal/planner"
	"github.com/himanshu/daily-planner/internal/repository"
)

func SetupRoutes(r *gin.Engine, db *repository.Database) {
	// Initialize handlers
	authHandler := auth.NewAuthHandler(db)
	plannerHandler := planner.NewPlannerHandler(db)

	// Auth routes
	authGroup := r.Group("/auth")
	{
		authGroup.GET("/login", authHandler.ShowLoginPage)
		authGroup.POST("/login", authHandler.LoginHandler)
		authGroup.GET("/register", authHandler.ShowRegisterPage)
		authGroup.POST("/register", authHandler.RegisterHandler)
		authGroup.GET("/logout", authHandler.LogoutHandler)
		authGroup.POST("/forgot-password", authHandler.ForgotPasswordHandler)
		authGroup.POST("/reset-password", authHandler.ResetPasswordHandler)
		authGroup.GET("/google/login", authHandler.GoogleLoginHandler)
		authGroup.GET("/google/callback", authHandler.GoogleCallbackHandler)
	}

	// Planner routes
	plannerGroup := r.Group("/planner")
	{
		plannerGroup.GET("/", plannerHandler.ShowDashboard)
		plannerGroup.POST("/todos", plannerHandler.CreateTodo)
		plannerGroup.GET("/todos", plannerHandler.GetTodos)
		plannerGroup.PUT("/todos/:id", plannerHandler.UpdateTodo)
		plannerGroup.DELETE("/todos/:id", plannerHandler.DeleteTodo)

		plannerGroup.POST("/priorities", plannerHandler.CreatePriority)
		plannerGroup.GET("/priorities", plannerHandler.GetPriorities)
		plannerGroup.PUT("/priorities/:id", plannerHandler.UpdatePriority)
		plannerGroup.DELETE("/priorities/:id", plannerHandler.DeletePriority)

		plannerGroup.POST("/contacts", plannerHandler.CreateContact)
		plannerGroup.GET("/contacts", plannerHandler.GetContacts)
		plannerGroup.PUT("/contacts/:id", plannerHandler.UpdateContact)
		plannerGroup.DELETE("/contacts/:id", plannerHandler.DeleteContact)

		plannerGroup.POST("/water-intake", plannerHandler.UpdateWaterIntake)
		plannerGroup.GET("/water-intake", plannerHandler.GetWaterIntake)

		plannerGroup.POST("/thought", plannerHandler.CreateThought)
		plannerGroup.GET("/thought", plannerHandler.GetTodayThought)
		plannerGroup.POST("/thought/generate", plannerHandler.GenerateThought)
	}

	// Root route redirects to login if not authenticated
	r.GET("/", func(c *gin.Context) {
		if _, exists := c.Get("user_id"); !exists {
			c.Redirect(302, "/auth/login")
			return
		}
		c.Redirect(302, "/planner")
	})
}
