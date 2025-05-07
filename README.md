# Daily Planner

A comprehensive daily planner web application built with Go, Gin, PostgreSQL, and Bootstrap. The application helps users manage their daily tasks, priorities, contacts, water intake, and random thoughts.

## Features

- **Authentication**
  - Username/password login
  - Google SSO integration
  - Secure password hashing
  - JWT-based session management

- **Planner Features**
  - To-Do List management
  - Daily Priorities tracking
  - Contact reminders (Call/Email/Text)
  - Water intake tracker (10 glasses)
  - Random Thought of the Day

## Tech Stack

- Backend:
  - Go 1.22+
  - Gin web framework
  - GORM ORM
  - PostgreSQL database
  - JWT authentication

- Frontend:
  - HTML/CSS
  - Bootstrap 5
  - Vanilla JavaScript
  - Font Awesome icons

## Prerequisites

- Go 1.22 or higher
- PostgreSQL 12 or higher
- Google OAuth credentials (for SSO)

## Installation

1. Clone the repository:
   ```bash
   git clone https://github.com/yourusername/daily-planner.git
   cd daily-planner
   ```

2. Install dependencies:
   ```bash
   go mod download
   ```

3. Create a PostgreSQL database:
   ```sql
   CREATE DATABASE daily_planner;
   ```

4. Create a `.env` file in the project root:
   ```env
   SERVER_ADDRESS=:8080
   DB_HOST=localhost
   DB_PORT=5432
   DB_USER=postgres
   DB_PASSWORD=your_password
   DB_NAME=daily_planner
   JWT_SECRET=your-secret-key-change-this-in-production
   
   # Google OAuth credentials
   GOOGLE_CLIENT_ID=your-google-client-id
   GOOGLE_CLIENT_SECRET=your-google-client-secret
   GOOGLE_REDIRECT_URL=http://localhost:8080/auth/google/callback
   ```

5. Run the application:
   ```bash
   go run cmd/api/main.go
   ```

## Project Structure

```
daily-planner/
├── cmd/
│   └── api/
│       └── main.go
├── internal/
│   ├── auth/
│   │   ├── handlers.go
│   │   └── jwt.go
│   ├── config/
│   │   └── config.go
│   ├── models/
│   │   └── models.go
│   ├── planner/
│   │   └── handlers.go
│   └── repository/
│       └── db.go
├── pkg/
│   └── middleware/
│       └── auth.go
├── static/
│   ├── css/
│   │   └── style.css
│   └── js/
│       └── main.js
├── templates/
│   ├── auth/
│   │   ├── login.html
│   │   └── register.html
│   ├── planner/
│   │   ├── dashboard.html
│   │   └── modals.html
│   └── partials/
│       └── base.html
├── .env
├── .gitignore
├── go.mod
├── go.sum
└── README.md
```

## API Endpoints

### Authentication
- `POST /auth/register` - Register new user
- `POST /auth/login` - Login user
- `GET /auth/logout` - Logout user
- `POST /auth/forgot-password` - Request password reset
- `POST /auth/reset-password` - Reset password
- `GET /auth/google/login` - Google SSO login
- `GET /auth/google/callback` - Google SSO callback

### Planner
- `GET /planner` - Dashboard
- `GET /planner/todos` - Get todos
- `POST /planner/todos` - Create todo
- `PUT /planner/todos/:id` - Update todo
- `DELETE /planner/todos/:id` - Delete todo
- `GET /planner/priorities` - Get priorities
- `POST /planner/priorities` - Create priority
- `PUT /planner/priorities/:id` - Update priority
- `DELETE /planner/priorities/:id` - Delete priority
- `GET /planner/contacts` - Get contacts
- `POST /planner/contacts` - Create contact
- `PUT /planner/contacts/:id` - Update contact
- `DELETE /planner/contacts/:id` - Delete contact
- `GET /planner/water-intake` - Get water intake
- `POST /planner/water-intake` - Update water intake
- `GET /planner/thought` - Get today's thought
- `POST /planner/thought/generate` - Generate new thought

## Contributing

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## License

This project is licensed under the MIT License - see the LICENSE file for details. 