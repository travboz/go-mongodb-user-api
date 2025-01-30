# Go MongoDB User API

TODO:

## Features

TODO:

## Installation

1. Clone this repository:
   ```sh
   git clone TODO:
   cd TODO:
   ```
2. Run docker container:
    ```sh
    make up
    ```
3. Seed MongoDB instance:
   ```sh
   make seed-db
   ```
4. Run server:
    ```sh
    make run
    ```
5. Navigate to `http://localhost<SERVER_PORT>` and call an endpoint

### `.env` file
This server uses a `.env` file for basic configuration.
Here is an example of the `.env`:
   ```sh
    DB_CONTAINER_NAME=MONGO-USER-CRUD
    SERVER_PORT=":8080"
    MONGO_DB_NAME=mongo_user_crud
    MONGO_DB_USERNAME=user
    MONGO_DB_PASSWORD=secret
    MONGODB_URI=mongodb://user:secret@localhost:27017/mongo_user_crud?authSource=admin&readPreference=primary&appname=MongDB%20Compass&directConnection=true&ssl=false
    COMPASS_USER_MONGODB_URI=mongodb://user:secret@localhost:27017/mongo_user_crud?authSource=admin&readPreference=primary&appname=MongDB%20Compass&directConnection=true&ssl=false
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
  "name": "bob jones",
  "email": "bob@jones.com",
  "favourite_number": 25,
  "active": false
}
```

#### Create a User
```sh
curl -X POST "http://localhost:8080/users" \
     -H "Content-Type: application/json" \
     -d '{
       "name": "bob jones",
       "email": "bob@jones.com",
       "favourite_number": 25,
       "active": false
     }'
```

#### Get All Users
```sh
curl http://localhost:8080/users
```

## Contributing
Feel free to fork and submit PRs!

## License:
`MIT`


This should work for GitHub! Let me know if you need any tweaks. 
