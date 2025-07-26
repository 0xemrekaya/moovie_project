# Moovie API

This is a RESTful API for managing a collection of movies. It's built with Go (httprouter) and uses PostgreSQL as its database.

## Features

*   **HTTP Router:** Uses `github.com/julienschmidt/httprouter` for high-performance routing.
*   **CRUD Operations:** Create, Read, Update, and Delete movies.
*   **Filtering:** Filter movies by title and genres.
*   **Pagination:** Paginate through the list of movies.
*   **Sorting:** Sort movies by ID, title, year, or runtime.
*   **Health Check:** An endpoint to check the status of the API.
*   **Middleware:** Includes IP based rate limiting and panic recovery.

## Prerequisites

*   [Go](https://golang.org/doc/install) (version 1.24 or newer)
*   [PostgreSQL](https://www.postgresql.org/download/)
*   [Docker](https://www.docker.com/products/docker-desktop) (for running PostgreSQL in a container)
*   [make](https://www.gnu.org/software/make/)

## Getting Started

### 1. Clone the Repository

```bash
git clone https://github.com/your-username/moovie_project.git
cd moovie_project
```

### 2. Set Up Environment Variables

There is a `.env` file in the root of the project and add the following environment variables. These will be used by `docker-compose`.

### 3. Run the Application with Docker

The easiest way to get the application running is by using Docker. The provided `docker-compose.yml` file will set up both the API service and the PostgreSQL database.

Run the following command to build and start the containers:

```bash
make docker-up
```

This will start the API server on `http://localhost:8000`.

To stop the application and the database container, you can run:

```bash
make docker-down
```

### 4. Run Database Migrations

Database migrations are handled automatically by the application on startup. The migration files are located in the `/migrations` directory.

## Makefile Commands

*   `make build`: Builds the application binary.
*   `make run`: Runs the application locally (without Docker).
*   `make test`: Runs the tests.
*   `make clean`: Removes the build artifacts.
*   `make docker-up`: Starts the PostgreSQL container.
*   `make docker-down`: Stops the PostgreSQL container.

## API Endpoints

All endpoints are prefixed with `/v1`.

| Method  | Endpoint                | Description                                         |
| :------ | :---------------------- | :-------------------------------------------------- |
| `GET`   | `/healthcheck`          | Checks the health of the application.               |
| `GET`   | `/movies`               | Returns a list of movies (with filtering/sorting). |
| `POST`  | `/movies`               | Creates a new movie.                                |
| `GET`   | `/movies/:id`           | Retrieves a specific movie by its ID.               |
| `PATCH` | `/movies/:id`           | Updates a specific movie.                           |
| `DELETE`| `/movies/:id`           | Deletes a specific movie.                           |

### Filtering, Sorting, and Pagination

You can query the `/v1/movies` endpoint with the following parameters:

*   **`title`**: Filter by movie title (e.g., `?title=Inception`).
*   **`genres`**: Filter by genres, comma-separated (e.g., `?genres=action,sci-fi`).
*   **`page`**: The page number for pagination (e.g., `?page=2`).
*   **`page_size`**: The number of results per page (e.g., `?page_size=10`).
*   **`sort`**: Sort order. Use a field name (e.g., `?sort=year`). Prepend with `-` for descending order (e.g., `?sort=-year`).
    *   **Allowed sort fields**: `id`, `title`, `year`, `runtime`.

### Example cURL Requests

**Create a movie:**

```bash
curl -X POST -H "Content-Type: application/json" -d '{
  "title": "Inception",
  "year": 2010,
  "runtime": 148,
  "genres": ["action", "sci-fi", "thriller"]
}' http://localhost:8000/v1/movies
```

**Get a list of movies:**

```bash
curl http://localhost:8000/v1/movies
```

**Get a list of movies with filtering and sorting:**

```bash
curl "http://localhost:8000/v1/movies?genres=sci-fi,action&sort=-year"
```

**Get a specific movie:**

```bash
curl http://localhost:8000/v1/movies/1
```
