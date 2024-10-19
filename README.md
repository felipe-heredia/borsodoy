# Radovid

Radovid is an Auction House API, written in Go, using SQLite and Docker. The API
has some tests to ensure the application works well.

The Auction House stores on Database the bids, items and users. The API uses JWT
token to authenticate users.

## Getting Started

All you need to run the application is having Docker installed.

```bash
docker compose up -d
```

This command will build and start Radovid API.

## Routes

1. GET /ping
- **Description**: A simple health check route to verify that the server is running.
- **Response**:
```json
{
  "message": "pong"
}
```
2. GET /users
- **Description**: Retrieves a list of all users.
- **Response**: Array of user objects.
3. POST /user
- **Description**: Creates a new user.
- **Request Body**:
```json
{
  "name": "string",
  "email": "string",
  "password": "string"
}
```
- **Response**: The created user object or error message.
4. GET /user/:id
- **Description**: Retrieves a specific user by their ID.
- **Response**: The user object or error message if not found.
5. GET /items
- **Description**: Retrieves a list of all available items.
- **Response**: Array of item objects.
6. GET /item/:id
- **Description**: Retrieves details of a specific item by its ID.
- **Response**: The item object or error message if not found.
7. POST /login
- **Description**: Logs in a user and returns a JWT token.
- **Request Body**:
```json
{
  "email": "string",
  "password": "string"
}
```
- **Response**:
```json
{
  "access_token": "string",
  "expire_at": "Date"
}
```
8. Protected Routes (Authentication Required)

The following routes require a valid JWT token for access. The token must be
included in the Authorization header as Bearer <token>.

8.1 POST /item
- **Description**: Creates a new item (protected route).
- **Request Body**:
```json
{
  "name": "string",
  "description": "string",
  "price": "integer",
  "image_url": "string",
  "user_id": "string"
}
```
- **Response**: The created item object or error message.

8.2 POST /bid
- **Description**: Places a bid on an item (protected route).
- **Request Body**:
```json
{
  "amount": "integer",
  "withdrawn_in": "integer",
  "item_id": "string",
  "user_id": "string"
}
```
- **Response**: The created bid object or error message.

8.3 DELETE /bid/:id
- **Description**: Withdraws a bid by its ID (protected route).
- **Responsei**: Success message or error message if the bid cannot be withdrawn.

### Authentication

To access protected routes, you must authenticate by logging in through the
/login route and include the JWT token in subsequent requests using the
Authorization: Bearer <token> header.

### Error Handling

All routes return appropriate error messages if something goes wrong, including
validation errors, resource not found errors, or unauthorized access attempts.

This documentation provides a detailed overview of how to use the API, what
each route does, and what is expected when making POST requests.

---

Made with :purple_heart: by Felipe Heredia!
