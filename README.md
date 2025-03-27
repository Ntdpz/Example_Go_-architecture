# Project Structure (Go)

## Introduction
This project is structured following best practices for a Go application. The organization ensures maintainability, scalability, and clarity.

## Folder Structure
```
/project-root
│── /database
│    └── db.go
│── /cmd
│    └── main.go
│── /internal
│    ├── /repository
│    │    └── user_repository.go
│    ├── /handlers
│    │    └── user_handlers.go
│    ├── /service
│    │    └── user_service.go
│── /config
│    ├── .env
│    └── config.go
│── /test
│    └── test.go
│── docker-compose.yaml
│── Dockerfile
```

### **Folder Descriptions**

#### **`/database`**
- `db.go` - Contains database connection logic and setup.

#### **`/cmd`**
- `main.go` - The main entry point of the application where the server is initialized and started.

#### **`/internal/repository`**
- `user_repository.go` - Handles all interactions with the database related to user data, including queries and CRUD operations.

#### **`/internal/handlers`**
- `user_handlers.go` - Defines HTTP handlers that process incoming API requests and forward them to the appropriate services.

#### **`/internal/service`**
- `user_service.go` - Implements business logic and acts as a bridge between handlers and repositories.

#### **`/config`**
- `.env` - Stores environment variables such as database credentials and configuration settings.
- `config.go` - Loads and manages application configurations from environment variables.

#### **`/test`**
- `test.go` - Contains unit and integration tests for different components of the application.

#### **`docker-compose.yaml`**
- Defines services for running the application with Docker, including the database and any dependencies.

#### **`Dockerfile`**
- Specifies instructions to build a Docker image for deploying the application in a containerized environment.

## Getting Started
1. Clone the repository:
   ```sh
   git clone <repository-url>
   ```
2. Navigate to the project directory:
   ```sh
   cd project-root
   ```
3. Set up the environment variables in `.env` file.
4. Build and run the project:
   ```sh
   go run cmd/main.go
   ```
5. Run using Docker:
   ```sh
   docker-compose up --build
   ```

## Contributing
Feel free to open issues and submit pull requests to improve the project!

## License
This project is licensed under the MIT License.

