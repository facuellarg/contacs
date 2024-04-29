# Contact Management API

This is a RESTful API for managing contacts, built using the Go programming language and the Echo framework. The API follows the hexagonal architecture, with separate folders for external services, interfaces, models, and use cases.

## Folders

- `external-service`: Contains code for integrating with external services, such as the webserver, databases or third-party APIs on case of be need it.
- `interface`: Contains code for handling HTTP requests and responses, implementing the API endpoints.
- `models`: Contains data structures representing contacts and other entities.
- `use-case`: Contains business logic for managing contacts.

## Functional Requirements

### 1. Create Contact Endpoint

- **Method**: `POST`
- **Path**: `/contacts`
- **Request Body**: JSON object representing the contact
- **Response**: JSON object of the created contact with an assigned ID

### 2. Get Contact by ID Endpoint

- **Method**: `GET`
- **Path**: `/contacts/{id}`
- **Path Parameter**: `id` (contact ID)
- **Response**: JSON object of the contact with the provided ID

### 3. Delete Contact Endpoint

- **Method**: `DELETE`
- **Path**: `/contacts/{id}`
- **Path Parameter**: `id` (contact ID)
- **Response**: No content (204 status code) upon successful deletion

### 4. Update Contact Endpoint

- **Method**: `PUT`
- **Path**: `/contacts/{id}`
- **Path Parameter**: `id` (contact ID)
- **Request Body**: JSON object representing the updated contact
- **Response**: JSON object of the updated contact

## Dependencies

- Go programming language
- Echo framework for HTTP routing and handling
- Any additional dependencies required for integrating with external services (e.g., database drivers)

## Architecture

The project follows the hexagonal architecture pattern, which separates concerns and promotes loose coupling between components. The main folders are:

- **external-service**: Contains adapters for external services, such as databases or third-party APIs. This layer handles the communication with these services and abstracts their implementation details from the rest of the application.

- **interface**: Contains the HTTP handlers that implement the API endpoints. This layer translates incoming HTTP requests into method calls on the use case layer and transforms the responses into appropriate HTTP responses.

- **models**: Contains data structures that represent the entities in the application, such as contacts. These models are used throughout the application layers.

- **use-case**: Contains the business logic of the application. This layer orchestrates the interactions between the external services and the domain models, implementing the core functionality of the application.

## Usage

1. Clone the repository.
2. Run `go mod tidy` to Install any required dependencies.
4. Run `go run main.go` to run the application.
5. Use a tool like Postman or cURL to interact with the API endpoints. To facilitate testing, a Postman collection is provided in the [postman](./postman) folder.
6. The API will be available at `http://localhost:8080`.


