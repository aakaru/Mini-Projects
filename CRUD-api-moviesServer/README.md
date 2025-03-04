# Go Movies CRUD API Server 

## Project Overview 🚀
This is a simple CRUD (Create, Read, Update, Delete) API server for managing a movie database, built using Go and the Gorilla Mux router. The application provides endpoints to interact with a collection of movies.

## Features ✨
- 📋 List all movies
- 🔍 Get a specific movie by ID
- ➕ Create a new movie
- ✏️ Update an existing movie
- 🗑️ Delete a movie

## Prerequisites 🛠️
- Go (1.16+)
- Gorilla Mux Router

## Installation 💻

1. Clone the repository:
```bash
git clone https://your-repository-url.git
cd your-project-directory
```

2. Install dependencies:
```bash
go mod init movies-server
go get -u github.com/gorilla/mux
```

3. Run the server:
```bash
go run main.go
```

## API Endpoints 🌐

| Method | Endpoint | Description |
|--------|----------|-------------|
| GET | `/movies` | Retrieve all movies |
| GET | `/movies/{id}` | Retrieve a specific movie by ID |
| POST | `/movies/` | Create a new movie |
| PUT | `/movies/{id}` | Update an existing movie |
| DELETE | `/movies/{id}` | Delete a movie |

## Example Payload 📝
```json
{
  "isbn": "123456",
  "title": "Inception",
  "director": {
    "firstname": "Christopher",
    "lastname": "Nolan"
  }
}
```

## Running Tests 🧪
*Coming soon* - Test suite to be added

## Contributing 🤝
Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

## License 📄
[MIT](https://choosealicense.com/licenses/mit/)
