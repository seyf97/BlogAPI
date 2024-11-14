# BlogAPI
This project is a simple Go-based RESTful API built with the Gin framework, serving as the backend for a blog. It offers endpoints for managing blog posts and includes authentication, enabling creation, reading, updating, and deletion (CRUD) of entries.

## Installation

1. Clone the repository:
   ```bash
   git clone https://github.com/seyf97/BlogAPI.git
   cd BlogAPI
   ```

2. Create a `.env` file in the project root with the following fields:
   ```
   PORT=your_port_number
   JWT_SECRET=your_jwt_secret
   ```
   - `PORT`: Specifies the port the server will run on (e.g. 8080)
   - `JWT_SECRET`: Secret key used for signing JWT tokens

3. Install dependencies:
   ```bash
   go mod download
   ```

4. Run the application:
   ```bash
   go run main.go
   ```


## Features

- **JWT Authentication**  
  Secure user authentication using JSON Web Tokens (JWT).

- **CRUD Operations for Articles**  
  Supports Creating, Reading, Updating, and Deleting articles.

- **User Management**  
  Allows user signup and login functionality.

- **SQLite Database**  
  Utilizes SQLite as the database backend for lightweight and easy data storage.


## API Endpoints

### Articles

| Method | Route                  | Auth Required | Description                 |
|--------|------------------------|---------------|-----------------------------|
| GET    | `/articles/:id`        | No            | Get a single article by ID  |
| GET    | `/articles`            | No            | Get all articles            |
| POST   | `/articles`            | Yes           | Create a new article        |
| DELETE | `/articles/:id`        | Yes           | Delete an article by ID     |
| PUT    | `/articles/:id`        | Yes           | Update an article by ID     |

### Users

| Method | Route     | Description                  |
|--------|-----------|------------------------------|
| POST   | `/signup` | Register a new user          |
| POST   | `/login`  | Authenticate user and get token |


## Structure

```plaintext
BlogAPI/
├── db/
│   └── db.go                  // DB connection and operations
├── middlewares/
│   └── auth.go                // JWT authentication middleware
├── models/
│   ├── article.go             // Article model and data handling
│   └── user.go                // User model and data handling
├── routes/
│   ├── articles.go            // Routes for article-related endpoints
│   ├── routes.go              // Contains all routes
│   └── users.go               // Routes for user-related endpoints
├── utils/
│   ├── hash.go                // Utility for password hashing
│   └── jwt.go                 // Utility for JWT generation and validation
├── .gitignore                 
├── go.mod                     
├── go.sum                     
├── LICENSE                    
├── main.go                    // Main entry point of the application
└── README.md                 
```
