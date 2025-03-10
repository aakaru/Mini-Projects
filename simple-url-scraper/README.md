# Simple URL Scraper

A lightweight, concurrent web crawler written in Go that extracts and lists all unique URLs from specified web pages.

## Overview

Simple URL Scraper is a command-line tool that crawls web pages to find and collect all unique URLs. It uses Go's concurrency model with goroutines and channels to efficiently process multiple web pages simultaneously.

## Features

- Concurrent web page scraping using goroutines
- Extracts all unique URLs from HTML anchor tags
- Filters for only fully qualified URLs (those starting with "http")
- Simple command-line interface
- Displays a summary of unique URLs found

## Requirements

- Go 1.11 or higher
- `golang.org/x/net/html` package

## Installation

1. Clone the repository:
   ```
   git clone https://github.com/yourusername/simple-url-scraper.git
   cd simple-url-scraper
   ```

2. Install dependencies:
   ```
   go get golang.org/x/net/html
   ```

3. Build the executable:
   ```
   go build -o urlscraper main.go
   ```

## Usage

Run the tool by providing one or more URLs as command-line arguments:

```
./urlscraper https://example.com https://another-site.com
```

### Example Output

```
Found 15 Unique URLs:

-https://www.google.com
-https://github.com
-https://golang.org/pkg/net/http
-https://blog.example.com/posts
-https://example.com/about
````

## License

[MIT]
