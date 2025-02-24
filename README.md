## Project Structure

```bash
/root-dir
│── main.go              # Entry point
│── config/
│   ├── config.go        # Load .env config
│── db/
│   ├── db.go            # Database connection setup
│── handlers/
│   ├── user_handler.go  # User API handlers
│── models/
│   ├── user.go          # User model
│── repositories/
│   ├── user_repo.go     # Data access layer
│── services/
│   ├── user_service.go  # Business logic layer
│── middleware/
│   ├── auth_middleware.go  # JWT Middleware
│── routes/
│   ├── routes.go        # Route definitions
│── .env                 # Environment variables
│── go.mod
│── go.sum
```
