# Golang Clean Architecture Template API

A starter template for building scalable, maintainable, and testable RESTful APIs in Go using the principles of Clean Architecture.

## Features

- Layered architecture (domain, usecase, handle, repository)
- Dependency injection
- RESTful API structure
- Swagger/OpenAPI documentation
- Example modules and handlers
- Easy to extend and customize

## Getting Started

### Prerequisites

- Go 1.18+
- Docker (optional, for running with containers)

### Installation

Clone the repository:

```bash
git clone https://github.com/yourusername/golang-clean-arch-template.git
cd golang-clean-arch-template
```

Install dependencies:

```bash
go mod tidy
```

### Running the API

```bash
go run cmd/api/main.go
```

Or with Docker:

```bash
docker-compose up --build
```

### API Documentation

Access the Swagger UI at:

```
http://localhost:5000/swagger/index.html
```

## Project Structure

```
.
├── cmd/            # Application entrypoints
├── internal/       # Application code (domain, usecase, handle, repository)
├── pkg/            # Shared packages
├── docs/           # API documentation
├── Dockerfile
└── README.md
```

## Contributing

Contributions are welcome! Please open issues or submit pull requests.

## License

This project is licensed under the MIT License.
