# Task Management API
### GDSC UIT club work
This is a Task Management API built with Go and MongoDB. It allows users to create, read, update, and delete projects and tasks, including support for subtasks.

## Features

- **Project Management**: Create, read, update, and delete projects. Get tasks by project.
- **Task Management**: Create, read, update, and delete tasks.
- **Subtasks**: Support for creating subtasks under main tasks.
- **Search and Filter**: Search tasks by name, status, or priority.

## Getting Started

### Prerequisites

- Go (latest version)
- MongoDB

### Installation

1. Clone the repository:
   ```sh
   git clone https://github.com/iknizzz1807/task-management-api.git
   cd task-management-api
   ```

2. Install dependencies:
   ```sh
   go mod tidy
   ```

3. Set up MongoDB connection in `database.go`
   ```go
   client, err = mongo.Connect(context.Background(), options.Client().ApplyURI("mongodb://localhost:27017"))
   ```

4. Run the application:
   ```sh
   go run main.go
   ```

## API Documentation

### Projects

#### Get All Projects

- **Endpoint**: `GET /projects`
- **Response**:
  ```json
  {
    "status": "success",
    "count": 2,
    "projects": [
      {
        "id": "60d5ec49f1e4c2d2f8f0e4b1",
        "name": "Study",
        "description": "Study related tasks",
        "created_at": "2023-01-01T00:00:00Z",
        "updated_at": "2023-01-01T00:00:00Z"
      }
    ]
  }
  ```

#### Get Project by ID

- **Endpoint**: `GET /project/{id}`
- **Response**:
  ```json
  {
    "status": "success",
    "project": {
      "id": "60d5ec49f1e4c2d2f8f0e4b1",
      "name": "Study",
      "description": "Study related tasks",
      "created_at": "2023-01-01T00:00:00Z",
      "updated_at": "2023-01-01T00:00:00Z"
    },
    "tasks": [
      {
        "id": "60d5ec49f1e4c2d2f8f0e4b2",
        "name": "Complete homework",
        "description": "Finish math homework",
        "deadline": "2023-12-31T23:59:59Z",
        "priority": "High",
        "status": "Incomplete",
        "parent_id": null,
        "project_id": "60d5ec49f1e4c2d2f8f0e4b1",
        "created_at": "2023-01-01T00:00:00Z",
        "updated_at": "2023-01-01T00:00:00Z"
      }
    ],
    "count": 1
  }
  ```

#### Create Project

- **Endpoint**: `POST /project`
- **Request Body**:
  ```json
  {
    "name": "Study",
    "description": "Study related tasks"
  }
  ```
- **Response**:
  ```json
  {
    "status": "success",
    "project": {
      "id": "60d5ec49f1e4c2d2f8f0e4b1",
      "name": "Study",
      "description": "Study related tasks",
      "created_at": "2023-01-01T00:00:00Z",
      "updated_at": "2023-01-01T00:00:00Z"
    }
  }
  ```

#### Update Project

- **Endpoint**: `PUT /project/{id}`
- **Request Body**:
  ```json
  {
    "name": "Study Updated",
    "description": "Updated study related tasks"
  }
  ```
- **Response**:
  ```json
  {
    "status": "success",
    "project": {
      "id": "60d5ec49f1e4c2d2f8f0e4b1",
      "name": "Study Updated",
      "description": "Updated study related tasks",
      "created_at": "2023-01-01T00:00:00Z",
      "updated_at": "2023-01-01T00:00:00Z"
    }
  }
  ```

#### Delete Project

- **Endpoint**: `DELETE /project/{id}`
- **Response**:
  ```json
  {
    "status": "success"
  }
  ```

### Tasks

#### Get All Tasks

- **Endpoint**: `GET /tasks`
- **Query Parameters**:
  - `name` (optional)
  - `status` (optional)
  - `priority` (optional)
- **Response**:
  ```json
  {
    "status": "success",
    "count": 2,
    "tasks": [
      {
        "id": "60d5ec49f1e4c2d2f8f0e4b2",
        "name": "Complete homework",
        "description": "Finish math homework",
        "deadline": "2023-12-31T23:59:59Z",
        "priority": "High",
        "status": "Incomplete",
        "parent_id": null,
        "project_id": "60d5ec49f1e4c2d2f8f0e4b1",
        "created_at": "2023-01-01T00:00:00Z",
        "updated_at": "2023-01-01T00:00:00Z"
      }
    ]
  }
  ```

#### Create Task

- **Endpoint**: `POST /task`
- **Request Body**:
  ```json
  {
    "name": "Complete homework",
    "description": "Finish math homework",
    "deadline": "2023-12-31T23:59:59Z",
    "priority": "High",
    "status": "Incomplete",
    "parent_id": "60d5ec49f1e4c2d2f8f0e4b3",  // Optional
    "project_id": "60d5ec49f1e4c2d2f8f0e4b1"  // Required
  }
  ```
- **Response**:
  ```json
  {
    "status": "success",
    "task": {
      "id": "60d5ec49f1e4c2d2f8f0e4b2",
      "name": "Complete homework",
      "description": "Finish math homework",
      "deadline": "2023-12-31T23:59:59Z",
      "priority": "High",
      "status": "Incomplete",
      "parent_id": "60d5ec49f1e4c2d2f8f0e4b3",
      "project_id": "60d5ec49f1e4c2d2f8f0e4b1",
      "created_at": "2023-01-01T00:00:00Z",
      "updated_at": "2023-01-01T00:00:00Z"
    }
  }
  ```

#### Update Task

- **Endpoint**: `PUT /task/{id}`
- **Request Body**:
  ```json
  {
    "name": "Complete homework updated",
    "description": "Finish math homework updated",
    "deadline": "2023-12-31T23:59:59Z",
    "priority": "Medium",
    "status": "Ongoing",
    "parent_id": "60d5ec49f1e4c2d2f8f0e4b3",  // Optional
    "project_id": "60d5ec49f1e4c2d2f8f0e4b1"  // Required
  }
  ```
- **Response**:
  ```json
  {
    "status": "success",
    "task": {
      "id": "60d5ec49f1e4c2d2f8f0e4b2",
      "name": "Complete homework updated",
      "description": "Finish math homework updated",
      "deadline": "2023-12-31T23:59:59Z",
      "priority": "Medium",
      "status": "Ongoing",
      "parent_id": "60d5ec49f1e4c2d2f8f0e4b3",
      "project_id": "60d5ec49f1e4c2d2f8f0e4b1",
      "created_at": "2023-01-01T00:00:00Z",
      "updated_at": "2023-01-01T00:00:00Z"
    }
  }
  ```

#### Delete Task

- **Endpoint**: `DELETE /task/{id}`
- **Response**:
  ```json
  {
    "status": "success"
  }
  ```

## Contributing

Contributions are welcome! Please open an issue or submit a pull request.

## License

This project is licensed under the MIT License.

---

This README provides an overview of the project, installation instructions, and detailed API documentation.
