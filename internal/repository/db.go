package repository

import (
	"fmt"
	"log"
	"os"

	"github.com/himanshu/daily-planner/internal/models"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Database struct that handles all database operations
type Database struct {
	DB *gorm.DB
}

// NewDatabase creates a new database instance
func NewDatabase() (*Database, error) {
	// Load environment variables
	if err := godotenv.Load(); err != nil {
		return nil, fmt.Errorf("error loading .env file: %v", err)
	}

	// Get database configuration from environment variables
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	// Create database connection string
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		dbHost, dbPort, dbUser, dbPassword, dbName)

	// Open database connection
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %v", err)
	}

	// Test the connection
	sqlDB, err := db.DB()
	if err != nil {
		return nil, fmt.Errorf("failed to get database instance: %v", err)
	}

	if err := sqlDB.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping database: %v", err)
	}

	log.Println("Successfully connected to database")
	return &Database{DB: db}, nil
}

// Migrate runs database migrations
func (db *Database) Migrate() error {
	// Read migration file
	migrationSQL, err := os.ReadFile("migrations/001_initial_schema.sql")
	if err != nil {
		return fmt.Errorf("failed to read migration file: %v", err)
	}

	// Execute migration
	if err := db.DB.Exec(string(migrationSQL)).Error; err != nil {
		return fmt.Errorf("failed to execute migration: %v", err)
	}

	log.Println("Database migrations completed successfully")
	return nil
}

// User operations
func (db *Database) CreateUser(user *models.User) error {
	return db.DB.Create(user).Error
}

func (db *Database) FindUserByID(id uint) (*models.User, error) {
	var user models.User
	err := db.DB.First(&user, id).Error
	return &user, err
}

func (db *Database) FindUserByUsername(username string) (*models.User, error) {
	var user models.User
	err := db.DB.Where("username = ?", username).First(&user).Error
	return &user, err
}

func (db *Database) FindUserByEmail(email string) (*models.User, error) {
	var user models.User
	err := db.DB.Where("email = ?", email).First(&user).Error
	return &user, err
}

func (db *Database) UpdateUser(user *models.User) error {
	return db.DB.Save(user).Error
}

func (db *Database) DeleteUser(id uint) error {
	return db.DB.Delete(&models.User{}, id).Error
}

// Todo operations
func (db *Database) CreateTodo(todo *models.TodoItem) error {
	return db.DB.Create(todo).Error
}

func (db *Database) FindTodoByID(id uint) (*models.TodoItem, error) {
	var todo models.TodoItem
	err := db.DB.First(&todo, id).Error
	return &todo, err
}

func (db *Database) FindTodosByUserID(userID uint) ([]models.TodoItem, error) {
	var todos []models.TodoItem
	err := db.DB.Where("user_id = ?", userID).Find(&todos).Error
	return todos, err
}

func (db *Database) UpdateTodo(todo *models.TodoItem) error {
	return db.DB.Save(todo).Error
}

func (db *Database) DeleteTodo(id uint) error {
	return db.DB.Delete(&models.TodoItem{}, id).Error
}

// Priority operations
func (db *Database) CreatePriority(priority *models.Priority) error {
	return db.DB.Create(priority).Error
}

func (db *Database) FindPriorityByID(id uint) (*models.Priority, error) {
	var priority models.Priority
	err := db.DB.First(&priority, id).Error
	return &priority, err
}

func (db *Database) FindPrioritiesByUserID(userID uint) ([]models.Priority, error) {
	var priorities []models.Priority
	err := db.DB.Where("user_id = ?", userID).Find(&priorities).Error
	return priorities, err
}

func (db *Database) UpdatePriority(priority *models.Priority) error {
	return db.DB.Save(priority).Error
}

func (db *Database) DeletePriority(id uint) error {
	return db.DB.Delete(&models.Priority{}, id).Error
}

// Contact operations
func (db *Database) CreateContact(contact *models.Contact) error {
	return db.DB.Create(contact).Error
}

func (db *Database) FindContactByID(id uint) (*models.Contact, error) {
	var contact models.Contact
	err := db.DB.First(&contact, id).Error
	return &contact, err
}

func (db *Database) FindContactsByUserID(userID uint) ([]models.Contact, error) {
	var contacts []models.Contact
	err := db.DB.Where("user_id = ?", userID).Find(&contacts).Error
	return contacts, err
}

func (db *Database) UpdateContact(contact *models.Contact) error {
	return db.DB.Save(contact).Error
}

func (db *Database) DeleteContact(id uint) error {
	return db.DB.Delete(&models.Contact{}, id).Error
}

// WaterIntake operations
func (db *Database) CreateWaterIntake(waterIntake *models.WaterIntake) error {
	return db.DB.Create(waterIntake).Error
}

func (db *Database) FindWaterIntakeByID(id uint) (*models.WaterIntake, error) {
	var waterIntake models.WaterIntake
	err := db.DB.First(&waterIntake, id).Error
	return &waterIntake, err
}

func (db *Database) FindWaterIntakeByUserIDAndDate(userID uint, date string) (*models.WaterIntake, error) {
	var waterIntake models.WaterIntake
	err := db.DB.Where("user_id = ? AND date = ?", userID, date).First(&waterIntake).Error
	return &waterIntake, err
}

func (db *Database) UpdateWaterIntake(waterIntake *models.WaterIntake) error {
	return db.DB.Save(waterIntake).Error
}

func (db *Database) DeleteWaterIntake(id uint) error {
	return db.DB.Delete(&models.WaterIntake{}, id).Error
}

// Thought operations
func (db *Database) CreateThought(thought *models.Thought) error {
	return db.DB.Create(thought).Error
}

func (db *Database) FindThoughtByID(id uint) (*models.Thought, error) {
	var thought models.Thought
	err := db.DB.First(&thought, id).Error
	return &thought, err
}

func (db *Database) FindThoughtByUserIDAndDate(userID uint, date string) (*models.Thought, error) {
	var thought models.Thought
	err := db.DB.Where("user_id = ? AND date = ?", userID, date).First(&thought).Error
	return &thought, err
}

func (db *Database) UpdateThought(thought *models.Thought) error {
	return db.DB.Save(thought).Error
}

func (db *Database) DeleteThought(id uint) error {
	return db.DB.Delete(&models.Thought{}, id).Error
}
