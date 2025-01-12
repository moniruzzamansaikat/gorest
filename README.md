# Fiber Test Application

A simple web application built using the [Fiber](https://github.com/gofiber/fiber) web framework in Go.

## Prerequisites

Before you begin, ensure you have the following installed:

- [Go](https://golang.org/doc/install) (v1.18+ recommended)

## Installation

1. **Clone the repository**:

   ```bash
   git clone https://github.com/your-username/fiber-app.git
   cd fiber-app
   ```

2. **Install dependencies**:

   If you are using Go modules (recommended), run:

   ```bash
   go mod tidy
   ```

3. **Run the application**:

   Start the app with:

   ```bash
   go run main.go
   ```

4. **Access the application**:

   Open your browser and go to `http://localhost:3000`. You should see the "Hello, Fiber!" message.

## Endpoints

- **GET `/`** - Returns "Hello, Fiber!"
- **GET `/about`** - Returns a message about the application
- **GET `/user`** - Returns JSON data for a user
- **GET `/users`** - Returns a list of users in JSON format
- **GET `/user/:id`** - Dynamic route returning the user ID
- **GET `/search?q=some-query`** - Handles search query parameters