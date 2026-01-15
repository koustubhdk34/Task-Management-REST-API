package main

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// Task represents a task structure
type Task struct {
	ID          int       `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Status      string    `json:"status"`
	CreatedAt   time.Time `json:"created_at"`
}

// In-memory storage for tasks
var tasks []Task
var nextID int = 1

func main() {
	// Initialize Gin router
	router := gin.Default()

	// GET endpoint to fetch all tasks
	router.GET("/tasks", getTasks)

	// POST endpoint to create a new task
	router.POST("/tasks", createTask)

	// Start server on port 8080
	router.Run(":8080")
}

// getTasks handles GET /tasks - returns all tasks
func getTasks(c *gin.Context) {
	if len(tasks) == 0 {
		c.JSON(http.StatusOK, gin.H{
			"message": "No tasks found",
			"tasks":   []Task{},
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Tasks retrieved successfully",
		"tasks":   tasks,
	})
}

// createTask handles POST /tasks - creates a new task
func createTask(c *gin.Context) {
	var newTask Task

	// Bind JSON request body to Task struct
	if err := c.ShouldBindJSON(&newTask); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid request body",
			"details": err.Error(),
		})
		return
	}

	// Validate required fields
	if newTask.Title == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Title is required",
		})
		return
	}

	// Set task ID and timestamp
	newTask.ID = nextID
	nextID++
	newTask.CreatedAt = time.Now()

	// Set default status if not provided
	if newTask.Status == "" {
		newTask.Status = "pending"
	}

	// Add task to in-memory storage
	tasks = append(tasks, newTask)

	// Return created task with 201 status code
	c.JSON(http.StatusCreated, gin.H{
		"message": "Task created successfully",
		"task":    newTask,
	})
}

