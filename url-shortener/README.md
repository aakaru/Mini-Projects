# URL Shortener

![Gopher in Space](gopher-space.png)

This is a simple URL shortener service built using Go. It allows you to create shortened URLs for long web addresses, making them easier to share and manage.

## Features

* Generates unique short URLs using MD5 hashing.
* Stores URL mappings in an in-memory database (for demonstration purposes).
* Provides API endpoints for shortening URLs and redirecting to the original URL.
* Simple and easy to use.

## Getting Started

1. **Clone the repository:**

   ```bash
   git clone https://github.com/your-username/go-url-shortener.git
   ```
1. Navigate to the project directory:
   ```bash
   cd go-url-shortener
   ```
2. Run the application:
   ```bash
   go run main.go
   ```
   The server will start running on http://localhost:3000.

## Usage

* Shortening a URL:

Send a POST request to /shortener with the URL you want to shorten in JSON format:

```bash
curl -X POST -H "Content-Type: application/json" -d '{"url": "https://www.example.com"}' http://localhost:3000/shortener
```

The response will contain the shortened URL:
```bash
{"short_url": "e5b7e999"}
```
* Redirecting to the Original URL:

Access the shortened URL by appending it to /redirect/:

```bash
http://localhost:3000/redirect/e5b7e999
```
This will redirect you to the original URL (https://www.example.com in this example).

## API Endpoints

* POST /shortener: Creates a short URL. Request body: {"url": "long_url"}. Response: {"short_url": "short_url"}
* GET /redirect/{short_url}: Redirects to the original URL.

## Future Improvements

* Persistent Storage: Replace the in-memory database with a persistent storage solution (e.g., PostgreSQL, MySQL, Redis) for production use.
* Custom Short URLs: Allow users to customize their short URLs.
* Analytics: Track click statistics for shortened URLs.
*  Rate Limiting: Implement rate limiting to prevent abuse.
* Improved Error Handling: Provide more informative error messages.
* URL Validation: Add validation to ensure that the provided URLs are valid.

 ## Contributing

 Contributions are welcome! Feel free to open issues and submit pull requests.

## License

This project is licensed under the MIT License - see the LICENSE file for details.
