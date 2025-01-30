# Go User API

A simple Go web server that manages user data in an in-memory cache with thread safety.

## Features
- Thread-Safe In-Memory Storage: Uses sync.RWMutex to handle concurrent access.
- Abstraction of storage logic for decoupled database
- RESTful API: Supports creating, retrieving, and deleting users via HTTP endpoints.
- Minimal & Efficient: Simple implementation with a focus on performance.

## Installation
1. Clone this repository:
   ```sh
   git clone https://github.com/travboz/go-user-web-server.git
   cd go-user-web-server
   ```
2. Run server:
   ```sh
   make run
   ```
3. Navigate to `http://localhost<SERVER_PORT>` and call an endpoint

### `.env` file
This server uses a `.env` file for basic configuration.
Here is an example of the `.env`:
   ```sh
   SERVER_PORT=":8080"
   ENV="development"
   ```
   
## API Endpoints

| Method   | Endpoint        | Description          |
|----------|----------------|----------------------|
| `GET`    | `/`            | Welcome message/health check     |
| `POST`   | `/users`       | Create a new user   |
| `GET`    | `/users`       | Get all users       |
| `GET`    | `/users/{id}`  | Get user by ID      |
| `PUT`    | `/users/{id}`  | Update a user       |
| `DELETE` | `/users/{id}`  | Delete a user       |

## Example Usage

#### User Payload

```json
{
   "name": "example"
}
```

#### Create a User
```sh
curl -X POST -H "Content-Type: application/json" -d '{"name": "Alice"}' http://localhost:4545/users
```

#### Get All Users
```sh
curl http://localhost:4545/users
```

## Contributing
Feel free to fork and submit PRs!

## License:
`MIT`


This should work for GitHub! Let me know if you need any tweaks. 
