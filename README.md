# Country-Search

Country-Search is a RESTful API service that fetches and provides information about countries using the REST Countries API. This service implements custom caching logic to optimize performance and includes comprehensive unit tests to ensure reliability.

## Features

- **Country Information**: Retrieve detailed information about countries, including name, capital, currency, and population.If the country details present in cache it will retrieve from cache other wise it call the external third party api and retrieve from there.
- **Caching**: Implements custom caching to reduce the number of requests to the third-party API and improve response times.
- **Error Handling**: Robust error handling for invalid requests and server errors.
- **Unit Testing**: Comprehensive test coverage to ensure the service functions correctly.

## Installation

1. Clone the repository:
   ```bash
   git clone https://github.com/Akshayvij07/country-search.git
   ```

2. Navigate to the project directory:
   ```bash
   cd country-search
   ```

3. Install dependencies:
   ```bash
   go mod tidy
   ```

## Usage

1. Start the server:
   ```bash
   go run cmd/country_search/main.go
   ```

2. Access the API:
   - Use an HTTP client like curl or Postman to send requests to the API.
   curl http://localhost:8000/api/countries/search?name=?
   - Example GET request to fetch country information:
     ```bash
     curl http://localhost:8000/api/countries/search?name=India
     ```

## Testing

Run unit tests to verify the functionality:
```bash
go test ./...
```

