package planner

import (
	"net/http"
	"time"

	"log"

	"github.com/gin-gonic/gin"
	"github.com/himanshu/daily-planner/internal/models"
	"github.com/himanshu/daily-planner/internal/repository"
)

type PlannerHandler struct {
	db *repository.Database
}

func NewPlannerHandler(db *repository.Database) *PlannerHandler {
	return &PlannerHandler{db: db}
}

// ShowDashboard renders the dashboard page
// Add logging to debug the database query
func (h *PlannerHandler) ShowDashboard(c *gin.Context) {
	userID, _ := c.Get("user_id")

	// Log userID and today's date
	log.Printf("ShowDashboard: userID=%v", userID)

	// Get today's data
	var todos []models.TodoItem
	var priorities []models.Priority
	var contacts []models.Contact
	var waterIntake models.WaterIntake
	var thought models.Thought

	// Get today's date
	today := time.Now().Truncate(24 * time.Hour)
	log.Printf("ShowDashboard: today=%v", today.Format("2006-01-02"))

	// Fetch all data for today
	if err := h.db.DB.Where("user_id = ? AND DATE(due_date) = ?", userID, today.Format("2006-01-02")).Find(&todos).Error; err != nil {
		log.Printf("Error fetching todos: %v", err)
	}
	log.Printf("Fetched todos: %v", todos)

	if err := h.db.DB.Where("user_id = ? AND date = ?", userID, today).Find(&priorities).Error; err != nil {
		log.Printf("Error fetching priorities: %v", err)
	}
	log.Printf("Fetched priorities: %v", priorities)

	if err := h.db.DB.Where("user_id = ? AND date = ?", userID, today).Find(&contacts).Error; err != nil {
		log.Printf("Error fetching contacts: %v", err)
	}
	log.Printf("Fetched contacts: %v", contacts)

	// Fetch water intake, create default if not found
	if err := h.db.DB.Where("user_id = ? AND date = ?", userID, today).First(&waterIntake).Error; err != nil {
		log.Printf("Error fetching water intake: %v", err)
		waterIntake = models.WaterIntake{
			UserID:  userID.(uint),
			Date:    today,
			Glasses: 0,
			Target:  10,
		}
	}
	log.Printf("Fetched water intake: %v", waterIntake)

	// Fetch thought, create default if not found
	if err := h.db.DB.Where("user_id = ? AND date = ?", userID, today).First(&thought).Error; err != nil {
		log.Printf("Error fetching thought: %v", err)
		thought = models.Thought{
			UserID:  userID.(uint),
			Date:    today,
			Content: "",
		}
	}
	log.Printf("Fetched thought: %v", thought)

	// Create water glasses array for the template
	waterGlasses := make([]int, waterIntake.Target)
	for i := 0; i < waterIntake.Target; i++ {
		waterGlasses[i] = i
	}

	// Prepare data for the template
	data := gin.H{
		"Title":        "Daily Planner",
		"Todos":        todos,
		"Priorities":   priorities,
		"Contacts":     contacts,
		"WaterIntake":  waterIntake,
		"WaterGlasses": waterGlasses,
		"Thought":      thought,
		"ShowForms":    false,
	}

	// Check if any data is missing
	if len(todos) == 0 || len(priorities) == 0 || len(contacts) == 0 || thought.Content == "" {
		data["ShowForms"] = true
	}

	c.HTML(http.StatusOK, "dashboard.html", data)
}

// CreateTodo handles creating a new todo item
func (h *PlannerHandler) CreateTodo(c *gin.Context) {
	userID, _ := c.Get("user_id")

	var todo struct {
		Title       string `json:"title"`
		Description string `json:"description"`
		DueDate     string `json:"dueDate"`
	}
	if err := c.ShouldBindJSON(&todo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	dueDate, err := time.Parse("2006-01-02", todo.DueDate)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid date format. Use YYYY-MM-DD"})
		return
	}

	newTodo := models.TodoItem{
		UserID:      userID.(uint),
		Title:       todo.Title,
		Description: todo.Description,
		DueDate:     dueDate,
	}

	if err := h.db.DB.Create(&newTodo).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create todo"})
		return
	}

	c.JSON(http.StatusCreated, newTodo)
}

// GetTodos handles retrieving all todo items
func (h *PlannerHandler) GetTodos(c *gin.Context) {
	userID, _ := c.Get("user_id")

	var todos []models.TodoItem
	if err := h.db.DB.Where("user_id = ?", userID).Find(&todos).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch todos"})
		return
	}

	c.JSON(http.StatusOK, todos)
}

// UpdateTodo handles updating a todo item
func (h *PlannerHandler) UpdateTodo(c *gin.Context) {
	userID, _ := c.Get("user_id")
	todoID := c.Param("id")

	var todo models.TodoItem
	if err := h.db.DB.Where("id = ? AND user_id = ?", todoID, userID).First(&todo).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Todo not found"})
		return
	}

	var updateData struct {
		Completed bool `json:"completed"`
	}
	if err := c.ShouldBindJSON(&updateData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	todo.Completed = updateData.Completed
	if err := h.db.DB.Save(&todo).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update todo"})
		return
	}

	c.JSON(http.StatusOK, todo)
}

// DeleteTodo handles deleting a todo item
func (h *PlannerHandler) DeleteTodo(c *gin.Context) {
	userID, _ := c.Get("user_id")
	todoID := c.Param("id")

	if err := h.db.DB.Where("id = ? AND user_id = ?", todoID, userID).Delete(&models.TodoItem{}).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete todo"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Todo deleted successfully"})
}

// CreatePriority handles creating a new priority
func (h *PlannerHandler) CreatePriority(c *gin.Context) {
	userID, _ := c.Get("user_id")

	var priorityData struct {
		Title       string `json:"title"`
		Description string `json:"description"`
	}
	if err := c.ShouldBindJSON(&priorityData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	priority := models.Priority{
		UserID:      userID.(uint),
		Title:       priorityData.Title,
		Description: priorityData.Description,
		Date:        time.Now().Truncate(24 * time.Hour),
	}

	if err := h.db.DB.Create(&priority).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create priority"})
		return
	}

	c.JSON(http.StatusCreated, priority)
}

// GetPriorities handles retrieving all priorities
func (h *PlannerHandler) GetPriorities(c *gin.Context) {
	userID, _ := c.Get("user_id")

	var priorities []models.Priority
	if err := h.db.DB.Where("user_id = ?", userID).Find(&priorities).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch priorities"})
		return
	}

	c.JSON(http.StatusOK, priorities)
}

// UpdatePriority handles updating a priority
func (h *PlannerHandler) UpdatePriority(c *gin.Context) {
	userID, _ := c.Get("user_id")
	priorityID := c.Param("id")

	var priority models.Priority
	if err := h.db.DB.Where("id = ? AND user_id = ?", priorityID, userID).First(&priority).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Priority not found"})
		return
	}

	if err := c.ShouldBindJSON(&priority); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.db.DB.Save(&priority).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update priority"})
		return
	}

	c.JSON(http.StatusOK, priority)
}

// DeletePriority handles deleting a priority
func (h *PlannerHandler) DeletePriority(c *gin.Context) {
	userID, _ := c.Get("user_id")
	priorityID := c.Param("id")

	if err := h.db.DB.Where("id = ? AND user_id = ?", priorityID, userID).Delete(&models.Priority{}).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete priority"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Priority deleted successfully"})
}

// CreateContact handles creating a new contact
func (h *PlannerHandler) CreateContact(c *gin.Context) {
	userID, _ := c.Get("user_id")

	var contactData struct {
		Name        string `json:"name"`
		Type        string `json:"type"`
		Description string `json:"description"`
	}
	if err := c.ShouldBindJSON(&contactData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	contact := models.Contact{
		UserID:      userID.(uint),
		Name:        contactData.Name,
		Type:        contactData.Type,
		Description: contactData.Description,
		Date:        time.Now().Truncate(24 * time.Hour),
	}

	if err := h.db.DB.Create(&contact).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create contact"})
		return
	}

	c.JSON(http.StatusCreated, contact)
}

// GetContacts handles retrieving all contacts
func (h *PlannerHandler) GetContacts(c *gin.Context) {
	userID, _ := c.Get("user_id")

	var contacts []models.Contact
	if err := h.db.DB.Where("user_id = ?", userID).Find(&contacts).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch contacts"})
		return
	}

	c.JSON(http.StatusOK, contacts)
}

// UpdateContact handles updating a contact
func (h *PlannerHandler) UpdateContact(c *gin.Context) {
	userID, _ := c.Get("user_id")
	contactID := c.Param("id")

	var contact models.Contact
	if err := h.db.DB.Where("id = ? AND user_id = ?", contactID, userID).First(&contact).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Contact not found"})
		return
	}

	if err := c.ShouldBindJSON(&contact); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.db.DB.Save(&contact).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update contact"})
		return
	}

	c.JSON(http.StatusOK, contact)
}

// DeleteContact handles deleting a contact
func (h *PlannerHandler) DeleteContact(c *gin.Context) {
	userID, _ := c.Get("user_id")
	contactID := c.Param("id")

	if err := h.db.DB.Where("id = ? AND user_id = ?", contactID, userID).Delete(&models.Contact{}).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete contact"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Contact deleted successfully"})
}

// UpdateWaterIntake handles updating water intake
func (h *PlannerHandler) UpdateWaterIntake(c *gin.Context) {
	userID, _ := c.Get("user_id")
	today := time.Now().Truncate(24 * time.Hour)

	var intakeData struct {
		Glasses int `json:"glasses"`
	}
	if err := c.ShouldBindJSON(&intakeData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Try to find existing record for today
	var waterIntake models.WaterIntake
	if err := h.db.DB.Where("user_id = ? AND date = ?", userID, today).First(&waterIntake).Error; err != nil {
		// Create new record if not found
		waterIntake = models.WaterIntake{
			UserID:  userID.(uint),
			Date:    today,
			Glasses: intakeData.Glasses,
			Target:  10,
		}
		if err := h.db.DB.Create(&waterIntake).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create water intake record"})
			return
		}
	} else {
		// Update existing record
		waterIntake.Glasses = intakeData.Glasses
		if err := h.db.DB.Save(&waterIntake).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update water intake"})
			return
		}
	}

	c.JSON(http.StatusOK, waterIntake)
}

// GetWaterIntake handles retrieving water intake
func (h *PlannerHandler) GetWaterIntake(c *gin.Context) {
	userID, _ := c.Get("user_id")
	today := time.Now().Truncate(24 * time.Hour)

	var intake models.WaterIntake
	if err := h.db.DB.Where("user_id = ? AND date = ?", userID, today).First(&intake).Error; err != nil {
		// If no record exists, return default values
		intake = models.WaterIntake{
			UserID:  userID.(uint),
			Date:    today,
			Glasses: 0,
			Target:  10,
		}
	}

	c.JSON(http.StatusOK, intake)
}

// CreateThought handles creating a new thought
func (h *PlannerHandler) CreateThought(c *gin.Context) {
	userID, _ := c.Get("user_id")

	var thoughtData struct {
		Content string `json:"content"`
	}
	if err := c.ShouldBindJSON(&thoughtData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	thought := models.Thought{
		UserID:  userID.(uint),
		Content: thoughtData.Content,
		Date:    time.Now().Truncate(24 * time.Hour),
	}

	if err := h.db.DB.Create(&thought).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create thought"})
		return
	}

	c.JSON(http.StatusCreated, thought)
}

// GetTodayThought handles retrieving today's thought
func (h *PlannerHandler) GetTodayThought(c *gin.Context) {
	userID, _ := c.Get("user_id")
	today := time.Now().Truncate(24 * time.Hour)

	var thought models.Thought
	if err := h.db.DB.Where("user_id = ? AND date = ?", userID, today).First(&thought).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "No thought found for today"})
		return
	}

	c.JSON(http.StatusOK, thought)
}

// GenerateThought handles generating a new thought
func (h *PlannerHandler) GenerateThought(c *gin.Context) {
	// For now, return a simple placeholder thought
	// In a real application, you might want to integrate with an AI service
	thought := models.Thought{
		Content: "Today is a new opportunity to make a difference. Focus on what matters most.",
		Date:    time.Now().Truncate(24 * time.Hour),
	}

	c.JSON(http.StatusOK, thought)
}
