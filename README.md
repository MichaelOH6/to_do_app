---

# Todo App

## Overview

This is a simple Todo application written in Go. The app allows users to create, view, update and delete tasks. It uses a minimal web server and basic file-based storage to manage tasks.

## Getting Started

To get started with the Todo app, follow these instructions:

### Prerequisites

Make sure you have the following installed on your machine:

- [Go](https://golang.org/dl/) (version 1.18 or later recommended)

### Running the Application

1. **Clone the Repository**

   ```bash
   git clone <repository-url>
   cd <repository-directory>
   ```

2. **Run the Application**

   Start the Todo app with:

   ```bash
   go run main.go
   ```

   The application will be available at `http://localhost:8080`.

### Running Tests

To run the tests for the Todo app, use the following command:

```bash
go test -v
```

This will execute all the tests and provide verbose output.

## Project Structure

- `static/style.css`: Contains the CSS styles for the web application.
- `templates/index.html`: The HTML template for the main page of the application.
- `main.go`: The main file that sets up the web server and handles routing.
- `main_test.go`: Contains unit tests for the application logic.
- `savedTasks.json`: File-based storage for saving and loading tasks.

## Features

- **Add a Todo**: Create a new task with a name.
- **View Todos**: Display all existing tasks.
- **Update Todos**: Update the name of the todos.
- **Delete a Todo**: Remove a task from the list.

## Configuration

The application configuration is minimal. The `main.go` file contains all necessary configuration settings. If you need to adjust settings, you can modify `main.go`.

## How It Works

- The application serves an HTML page using `index.html` and applies styles from `style.css`.
- Tasks are stored in `savedTasks.json`, which is read and written by the application.
- `main.go` contains the logic for handling HTTP requests and managing tasks.
- `main_test.go` includes tests to verify the application's functionality.
---
