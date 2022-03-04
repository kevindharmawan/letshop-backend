# LetShop Backend

## Getting Started

### Prerequisites

1. Make sure you have Docker, git, and golang installed in your computer
2. Make sure you already have a copy of this project

### Setting Up Environment

1. First, you need to run the PostgreSQL database by using `bash start_db.sh`
2. Next, make a copy of `.env.example` and rename it to `.env`
3. If you make any changes in the database settings from `start_db.sh`, edit your `.env` file accordingly, if not you can ignore this step

### Running the Program

1. Clean the golang modules in `go.mod` by running `go mod tidy`
2. To run the program, you can use `go run main.go`
3. By default, the server will run on `localhost:8080`
4. To stop the server, just terminate the program

## API Documentation

TODO

## Project Structure

    .
    ├── ...
    ├── config/           # Configurations and environment file
    ├── constants/        # Constants
    ├── controller/       # Request handlers
    │   ├── api/          # API request handlers
    │   └── middleware/   # Middleware
    ├── models/           # Database and data transfer models
    ├── router/           # Router
    └── services/         # External services
