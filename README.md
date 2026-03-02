# Go Authentication API

## Overview
This project is a **simple Authentication API** built using **Go (Golang)** with the **Fiber framework** and **GORM ORM** for database management.  
The purpose of this project is to provide practical training for building a user authentication system including **registration, login, and logout** using JWT tokens.

---

## Features

### User Authentication
- **Register**: Create a new user account with hashed password (bcrypt).  
- **Login**: Authenticate user and return a JWT token.  
- **Logout**: Stateless logout returning a success message.  

### Technology Stack
- **Programming Language:** Go 1.25.5  
- **Web Framework:** Fiber v2  
- **ORM:** GORM v1.31  
- **Database:** SQLite (can be replaced with MySQL/PostgreSQL)  
- **Password Hashing:** bcrypt  
- **Authentication:** JWT  

---

## 📂 Project Structure

```text
auth-go/
├── internal/
│   ├── database.go    # Database connection and User model
│   ├── handlers.go    # Functions: Register, Login, Logout
│   └── routes.go      # Define API routes
├── utils/
│   └── utils.go       # JWT generation
├── main.go            # Application entry point
├── go.mod             # Go module file with dependencies
└── README.md          # Project documentation


### Database
ID: Primary key

Name: User's full name

Email: Unique email

Password: Hashed password (bcrypt)
**User Model:**

```go
type User struct {
	ID       uint   `gorm:"primaryKey" json:"id"`
	Name     string `gorm:"not null" json:"name"`
	Email    string `gorm:"unique;not null" json:"email"`
	Password string `gorm:"not null" json:"password"`
}

---


---

## 3️⃣ Installation & Run

```markdown
### Installation & Run

1. Clone the repository or create a folder `auth-go`
2. Copy all project files as shown
3. Open terminal in project folder

Install dependencies and tidy modules:

```bash
go mod tidy
