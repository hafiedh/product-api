
## Getting Started

### Prerequisites

- Docker
- Docker Compose
- Go

### Installation

1. Clone the repository:
    ```sh
    git clone <repository-url>
    cd <repository-directory>
    ```

2. Build and run the containers:
    ```sh
    docker-compose up --build
    ```

3. The Product API will be available at `http://localhost:8090`.

### Configuration

Configuration is managed using environment variables. You can set these variables in the [`.env`](command:_github.copilot.openRelativePath?%5B%7B%22scheme%22%3A%22file%22%2C%22authority%22%3A%22%22%2C%22path%22%3A%22%2FUsers%2Fhafiedh%2Fgo%2Fsrc%2Fside_projects%2Fapi-product%2F.env%22%2C%22query%22%3A%22%22%2C%22fragment%22%3A%22%22%7D%2C%22cb46b11c-ca19-40e9-ba04-d48960866c0c%22%5D "/Users/hafiedh/go/src/side_projects/api-product/.env") file.

### Database

The project uses PostgreSQL as the database. The database schema can be found in the [`product.sql`](command:_github.copilot.openRelativePath?%5B%7B%22scheme%22%3A%22file%22%2C%22authority%22%3A%22%22%2C%22path%22%3A%22%2FUsers%2Fhafiedh%2Fgo%2Fsrc%2Fside_projects%2Fapi-product%2Fproduct.sql%22%2C%22query%22%3A%22%22%2C%22fragment%22%3A%22%22%7D%2C%22cb46b11c-ca19-40e9-ba04-d48960866c0c%22%5D "/Users/hafiedh/go/src/side_projects/api-product/product.sql") file.

### Running the Application

To run the application locally without Docker, use the following commands:

1. Install dependencies:
    ```sh
    go mod download
    ```

2. Run the application:
    ```sh
    go run main.go
    ```

### Project Structure

- [`cmd/`](command:_github.copilot.openRelativePath?%5B%7B%22scheme%22%3A%22file%22%2C%22authority%22%3A%22%22%2C%22path%22%3A%22%2FUsers%2Fhafiedh%2Fgo%2Fsrc%2Fside_projects%2Fapi-product%2Fcmd%2F%22%2C%22query%22%3A%22%22%2C%22fragment%22%3A%22%22%7D%2C%22cb46b11c-ca19-40e9-ba04-d48960866c0c%22%5D "/Users/hafiedh/go/src/side_projects/api-product/cmd/"): Entry point of the application.
- [`internal/config/`](command:_github.copilot.openRelativePath?%5B%7B%22scheme%22%3A%22file%22%2C%22authority%22%3A%22%22%2C%22path%22%3A%22%2FUsers%2Fhafiedh%2Fgo%2Fsrc%2Fside_projects%2Fapi-product%2Finternal%2Fconfig%2F%22%2C%22query%22%3A%22%22%2C%22fragment%22%3A%22%22%7D%2C%22cb46b11c-ca19-40e9-ba04-d48960866c0c%22%5D "/Users/hafiedh/go/src/side_projects/api-product/internal/config/"): Configuration management.
- [`internal/domain/entities/`](command:_github.copilot.openRelativePath?%5B%7B%22scheme%22%3A%22file%22%2C%22authority%22%3A%22%22%2C%22path%22%3A%22%2FUsers%2Fhafiedh%2Fgo%2Fsrc%2Fside_projects%2Fapi-product%2Finternal%2Fdomain%2Fentities%2F%22%2C%22query%22%3A%22%22%2C%22fragment%22%3A%22%22%7D%2C%22cb46b11c-ca19-40e9-ba04-d48960866c0c%22%5D "/Users/hafiedh/go/src/side_projects/api-product/internal/domain/entities/"): Domain entities.
- [`internal/domain/repositories/`](command:_github.copilot.openRelativePath?%5B%7B%22scheme%22%3A%22file%22%2C%22authority%22%3A%22%22%2C%22path%22%3A%22%2FUsers%2Fhafiedh%2Fgo%2Fsrc%2Fside_projects%2Fapi-product%2Finternal%2Fdomain%2Frepositories%2F%22%2C%22query%22%3A%22%22%2C%22fragment%22%3A%22%22%7D%2C%22cb46b11c-ca19-40e9-ba04-d48960866c0c%22%5D "/Users/hafiedh/go/src/side_projects/api-product/internal/domain/repositories/"): Repository interfaces and implementations.
- [`internal/infrastructure/container/`](command:_github.copilot.openRelativePath?%5B%7B%22scheme%22%3A%22file%22%2C%22authority%22%3A%22%22%2C%22path%22%3A%22%2FUsers%2Fhafiedh%2Fgo%2Fsrc%2Fside_projects%2Fapi-product%2Finternal%2Finfrastructure%2Fcontainer%2F%22%2C%22query%22%3A%22%22%2C%22fragment%22%3A%22%22%7D%2C%22cb46b11c-ca19-40e9-ba04-d48960866c0c%22%5D "/Users/hafiedh/go/src/side_projects/api-product/internal/infrastructure/container/"): Dependency injection container.
- [`internal/infrastructure/postgres/`](command:_github.copilot.openRelativePath?%5B%7B%22scheme%22%3A%22file%22%2C%22authority%22%3A%22%22%2C%22path%22%3A%22%2FUsers%2Fhafiedh%2Fgo%2Fsrc%2Fside_projects%2Fapi-product%2Finternal%2Finfrastructure%2Fpostgres%2F%22%2C%22query%22%3A%22%22%2C%22fragment%22%3A%22%22%7D%2C%22cb46b11c-ca19-40e9-ba04-d48960866c0c%22%5D "/Users/hafiedh/go/src/side_projects/api-product/internal/infrastructure/postgres/"): PostgreSQL database connection.
- `pkg/`: Shared constants and utilities.
- `server/`: HTTP server and handlers.
- `usecase/`: Use case implementations.