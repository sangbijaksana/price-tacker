# Crypto Price Tracker

Crypto Price Tracker is a simple REST API built with Golang to track cryptocurrency prices using the CoinCap API. The API allows authenticated users to track the prices of various cryptocurrencies in Indonesian Rupiah (IDR).

## Features

- User Registration and Authentication
- Add Coins to Track
- List Tracked Coins
- Remove Coins from Tracking

## Prerequisites

- Go 1.18 or later
- SQLite3
- Git

## Installation

1. Clone the repository:

    ```bash
    git clone https://github.com/sangbijaksana/price-tacker.git
    cd price-tacker
    ```

2. Install the dependencies:

    ```bash
    go mod tidy
    ```

3. Run the application:

    ```bash
    go build -o price-tracker && ./price-tracker
    ```

## API Endpoints

### Register

- **URL:** `/register`
- **Method:** `POST`
- **Description:** Registers a new user.
- **Request Body:**
    ```json
    {
        "email": "your_username",
        "password": "your_password",
        "password_conf": "your_password",
    }
    ```
- **Response:**
    - `201 Created` on success
    - `400 Bad Request` if the username is already taken
    - `500 Internal Server Error` on failure

### Login

- **URL:** `/login`
- **Method:** `POST`
- **Description:** Authenticates a user and returns a JWT token.
- **Request Body:**
    ```json
    {
        "email": "your_username",
        "password": "your_password"
    }
    ```
- **Response:**
    ```json
    {
        "jwt_token": "your_jwt_token"
    }
    ```
    - `200 OK` on success
    - `401 Unauthorized` if the credentials are invalid
    - `500 Internal Server Error` on failure

### Logout

- **URL:** `/logout`
- **Method:** `POST`
- **Description:** Logout an already sign in user.
- **Headers:**
    - `Authorization: your_jwt_token`
- **Response:**
    - `200 OK` on success
    - `500 Internal Server Error` on failure

### Get Tracked Coins

- **URL:** `/api/coins`
- **Method:** `GET`
- **Description:** Returns a list of tracked coins for the authenticated user.
- **Headers:**
    - `Authorization: your_jwt_token`
- **Response:**
    ```json
    {
        "data": [
            {
                "id": 4,
                "name": "bitcoin",
                "price": 960085007.7774923
            },
            {
                "id": 5,
                "name": "thorchain",
                "price": 88740.50776832688
            }
        ]
    }
    ```
    - `200 OK` on success
    - `401 Unauthorized` if the token is missing or invalid
    - `500 Internal Server Error` on failure

### Add Coin

- **URL:** `/api/coins`
- **Method:** `POST`
- **Description:** Adds a coin to the user's tracked coins.
- **Headers:**
    - `Authorization: your_jwt_token`
- **Request Body:**
    ```json
    {
        "name_id": "bitcoin"
    }
    ```
- **Response:**
    - `201 Created` on success
    - `400 Bad Request` if the coin does not exist in CoinCap API
    - `401 Unauthorized` if the token is missing or invalid
    - `500 Internal Server Error` on failure

### Remove Coin

- **URL:** `/api/coins/{id}`
- **Method:** `DELETE`
- **Description:** Removes a coin from the user's tracked coins by id (not token's name).
- **Headers:**
    - `Authorization: your_jwt_token`
- **Response:**
    - `200 OK` on success
    - `401 Unauthorized` if the token is missing or invalid
    - `500 Internal Server Error` on failure

## Database Schema

### Users Table

| Column    | Type    | Description          |
|-----------|---------|----------------------|
| id        | INTEGER | Primary key          |
| username  | TEXT    | Unique, not null     |
| password  | TEXT    | Hashed password      |

### Coins Table

| Column  | Type    | Description          |
|---------|---------|----------------------|
| id      | INTEGER | Primary key          |
| name_id | TEXT    | Coin name            |
| user_id | INTEGER | Foreign key to users |

## License

This project is licensed under the MIT License.
