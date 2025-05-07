package auth

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/himanshu/daily-planner/internal/models"
	"github.com/himanshu/daily-planner/internal/repository"
	"golang.org/x/crypto/bcrypt"
)

type AuthHandler struct {
	db *repository.Database
}

func NewAuthHandler(db *repository.Database) *AuthHandler {
	return &AuthHandler{db: db}
}

// ShowLoginPage renders the login page
func (h *AuthHandler) ShowLoginPage(c *gin.Context) {
	fmt.Println("Showing login page")
	c.HTML(http.StatusOK, "login.html", gin.H{
		"Title": "Login",
	})
}

// ShowRegisterPage renders the registration page
func (h *AuthHandler) ShowRegisterPage(c *gin.Context) {
	c.HTML(http.StatusOK, "register.html", gin.H{
		"Title": "Register",
	})
}

func (h *AuthHandler) RegisterHandler(c *gin.Context) {
	var registerData struct {
		Username        string `form:"username" binding:"required,min=3,max=50"`
		Email           string `form:"email" binding:"required,email"`
		Password        string `form:"password" binding:"required,min=6"`
		ConfirmPassword string `form:"confirm_password" binding:"required"`
	}

	if err := c.ShouldBind(&registerData); err != nil {
		c.HTML(http.StatusBadRequest, "register.html", gin.H{
			"Title": "Register",
			"Error": "Please fill in all fields correctly",
		})
		return
	}

	// Validate password match
	if registerData.Password != registerData.ConfirmPassword {
		c.HTML(http.StatusBadRequest, "register.html", gin.H{
			"Title": "Register",
			"Error": "Passwords do not match",
		})
		return
	}

	// Check if username or email already exists
	var existingUser models.User
	if err := h.db.DB.Where("username = ? OR email = ?", registerData.Username, registerData.Email).First(&existingUser).Error; err == nil {
		c.HTML(http.StatusBadRequest, "register.html", gin.H{
			"Title": "Register",
			"Error": "Username or email already exists",
		})
		return
	}

	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(registerData.Password), bcrypt.DefaultCost)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "register.html", gin.H{
			"Title": "Register",
			"Error": "Failed to process registration",
		})
		return
	}

	// Create user with GoogleID as nil
	user := models.User{
		Username: registerData.Username,
		Email:    registerData.Email,
		Password: string(hashedPassword),
		GoogleID: nil, // This will be stored as NULL in the database
	}

	if err := h.db.DB.Create(&user).Error; err != nil {
		c.HTML(http.StatusInternalServerError, "register.html", gin.H{
			"Title": "Register",
			"Error": "Failed to create account",
		})
		return
	}

	// Redirect to login page
	c.Redirect(http.StatusSeeOther, "/auth/login")
}

func (h *AuthHandler) LoginHandler(c *gin.Context) {
	var loginData struct {
		Username string `form:"username" binding:"required"`
		Password string `form:"password" binding:"required"`
	}

	if err := c.ShouldBind(&loginData); err != nil {
		c.HTML(http.StatusBadRequest, "login.html", gin.H{
			"Title": "Login",
			"Error": "Invalid input data",
		})
		return
	}

	var user models.User
	if err := h.db.DB.Where("username = ?", loginData.Username).First(&user).Error; err != nil {
		c.HTML(http.StatusUnauthorized, "login.html", gin.H{
			"Title": "Login",
			"Error": "Invalid credentials",
		})
		return
	}

	// Check password
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginData.Password)); err != nil {
		c.HTML(http.StatusUnauthorized, "login.html", gin.H{
			"Title": "Login",
			"Error": "Invalid credentials",
		})
		return
	}

	// Generate JWT token
	token, err := generateJWTToken(user.ID)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "login.html", gin.H{
			"Title": "Login",
			"Error": "Failed to process login",
		})
		return
	}

	c.SetCookie("auth_token", token, 3600*24, "/", "", false, true)
	c.Redirect(http.StatusFound, "/planner")
}

func (h *AuthHandler) LogoutHandler(c *gin.Context) {
	c.SetCookie("auth_token", "", -1, "/", "", false, true)
	c.Redirect(http.StatusSeeOther, "/auth/login")
}

func (h *AuthHandler) ForgotPasswordHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "login.html", gin.H{
		"Title": "Forgot Password",
		"Error": "Password reset functionality not implemented yet",
	})
}

func (h *AuthHandler) ResetPasswordHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "login.html", gin.H{
		"Title": "Reset Password",
		"Error": "Password reset functionality not implemented yet",
	})
}

func (h *AuthHandler) GoogleLoginHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "login.html", gin.H{
		"Title": "Login",
		"Error": "Google login not implemented yet",
	})
}

func (h *AuthHandler) GoogleCallbackHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "login.html", gin.H{
		"Title": "Login",
		"Error": "Google login not implemented yet",
	})
}
