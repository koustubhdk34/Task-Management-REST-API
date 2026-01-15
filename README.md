# Task Management REST API

A simple backend project built using **Golang and the Gin framework** that demonstrates core backend development concepts.

## Features

- RESTful APIs using GET and POST methods
- JSON format for data exchange
- In-memory data storage (no database required)
- Proper HTTP status codes
- Basic error handling

## Prerequisites

- Go 1.21 or higher installed on your system

## Installation

1. Clone or download this project
2. Navigate to the project directory
3. Install dependencies:
   ```bash
   go mod download
   ```

## Running the Application

Start the server:
```bash
go run main.go
```

The server will start on `http://localhost:8080`

## API Endpoints

### GET /tasks
Fetches all tasks.

**Response:**
```json
{
  "message": "Tasks retrieved successfully",
  "tasks": [
    {
      "id": 1,
      "title": "Sample Task",
      "description": "Task description",
      "status": "pending",
      "created_at": "2024-01-01T12:00:00Z"
    }
  ]
}
```

### POST /tasks
Creates a new task.

**Request Body:**
```json
{
  "title": "New Task",
  "description": "Task description",
  "status": "pending"
}
```

**Response (201 Created):**
```json
{
  "message": "Task created successfully",
  "task": {
    "id": 1,
    "title": "New Task",
    "description": "Task description",
    "status": "pending",
    "created_at": "2024-01-01T12:00:00Z"
  }
}
```

**Error Response (400 Bad Request):**
```json
{
  "error": "Title is required"
}
```

## Testing with Postman

1. **GET Request:**
   - Method: GET
   - URL: `http://localhost:8080/tasks`
   - Headers: None required

2. **POST Request:**
   - Method: POST
   - URL: `http://localhost:8080/tasks`
   - Headers: `Content-Type: application/json`
   - Body (raw JSON):
     ```json
     {
       "title": "Complete project",
       "description": "Finish the task management API",
       "status": "in-progress"
     }
     ```

## Project Structure

```
Task_Manager_GO/
├── main.go          # Main application file with routes and handlers
├── go.mod           # Go module dependencies
└── README.md        # This file
```

## Notes

- Tasks are stored in-memory and will be lost when the server restarts
- The `title` field is required when creating a task
- If `status` is not provided, it defaults to "pending"
- Task IDs are auto-incremented starting from 1

