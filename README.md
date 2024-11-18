# Quiz App Project

## 📖 Overview

This project is a **Quiz App** composed of two main components:

1. **Server**: A REST API backend for handling quiz questions, answers, and leaderboards.
2. **CLI**: A command-line interface to start a quiz and view the leaderboard.

The application is designed to be modular, testable, and cross-platform. It supports in-memory storage for simplicity but can be extended for database integration.

---

## 🗂️ Project Structure

Here’s the folder structure of the project:

```plaintext
quiz-app/
├── cli/                     # CLI application
│   ├── leaderboard.go       # CLI command to display leaderboard
│   ├── root.go              # CLI root command
│   ├── start.go             # CLI command to start the quiz
│   └── main.go              # CLI entry point
├── server/                  # Server application
│   ├── handler/             # HTTP handlers for API endpoints
│   │   ├── answers.go
│   │   ├── leaderboard.go
│   │   └── questions.go
│   ├── logger/              # Logging utilities
│   │   └── logger.go
│   ├── middleware/          # Middleware for HTTP requests
│   │   ├── compression.go
│   │   ├── logging.go
│   │   ├── middleware.go
│   │   └── security.go
│   ├── model/               # Core domain models
│   │   ├── answer.go
│   │   └── question.go
│   ├── repository/          # In-memory repositories
│   │   ├── in_memory_answer_repository.go
│   │   ├── in_memory_leaderboard_repository.go
│   │   ├── in_memory_question_repository.go
│   │   └── leaderboards_repository.go
│   ├── router/              # API routing
│   │   └── router.go
│   ├── tests/handler/       # Integration tests for handlers
│   │   ├── answers_test.go
│   │   ├── leaderboard_test.go
│   │   └── questions_test.go
│   └── main.go              # Server entry point
├── shared/                  # Shared utilities and DTOs
│   ├── client/              # API client for interacting with the server
│   │   └── api_client.go
│   └── dto/                 # Data Transfer Objects (DTOs)
│       ├── answer.go
│       ├── leaderboard.go
│       ├── question.go
│       ├── submit_answer_response.go
│       └── submit_answer_request.go
├── .gitignore               # Ignored files for version control
├── go.mod                   # Go module definition
├── go.sum                   # Go module checksums
├── Makefile                 # Build and run automation
└── README.md                # Documentation (this file)
```

## 🚀 Features

### Server

- **API Endpoints**:
  - Serve quiz questions.
  - Evaluate submitted answers.
  - Maintain and retrieve leaderboard rankings.
- **In-Memory Repositories**:
  - Store questions, answers, and scores without requiring a database.

### CLI

- **Interactive Quiz**:
  - Start a quiz and submit answers via the terminal.
- **Leaderboard Viewer**:
  - Fetch and display rankings from the server.

---

## 🔧 Prerequisites

Ensure the following tools are installed:

- **Go** (version 1.18 or higher)
- A terminal or command prompt.
- **Make** (version 4.0 or higher)

---

## 🛠️ How to Build and Run

### Using the Makefile

The `Makefile` provides commands to simplify builds and execution.

### Build Applications

Build binaries for your current platform:

Before running the `make build` command, ensure that the **Go environment variables** `GOOS` (for the operating system) and `GOARCH` (for the architecture) are set properly for your target platform. This step is crucial for cross-compiling the binaries.

```bash
make build
```

The `Makefile` provides specific build commands for each platform (Linux, macOS, and Windows). You **don't need to manually set the `GOOS` and `GOARCH` variables**—the `Makefile` already handles this for you. Simply choose the appropriate build command for your platform:

#### Build for Linux (amd64 architecture):

```bash
make build-linux
```

#### Build for macOS (amd64 architecture):

```bash
make build-macos
```

#### Build for Windows (amd64 architecture):

```bash
make build-windows
```

### Run Applications

#### Run Server

Start the server locally:

```bash
make run-server
```

This runs the server binary for your platform. The server will be accessible at `http://localhost:3000`.

#### Run CLI

Start the CLI application with the `start` command:

```bash
make run-cli ARGS=start
```

You can also view the leaderboard:

```bash
make run-cli ARGS=leaderboard
```

---

## 📜 API Endpoints (Server)

- GET /questions: Fetch quiz questions.
- POST /answers: Submit quiz answers and get scores.
- GET /leaderboard: View leaderboard rankings.

---

## 🛡️ Running Tests

Run integration tests in the server/tests/handler directory:

```bash
go test ./server/tests/handler
```
