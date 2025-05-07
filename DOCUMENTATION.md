# Daily Planner - Technical Documentation

## Overview

The Daily Planner is a comprehensive web application built using Go and modern web technologies. It provides users with a suite of planning and productivity tools while implementing secure authentication and clean architecture principles.

## Architecture

### Backend Architecture

1. **Clean Architecture**
   - Routes (URL routing and middleware)
   - Handlers (Presentation Layer)
   - Database (Data Access Layer)
   - Models (Domain Layer)

2. **Key Components**
   - Authentication System (JWT + Google SSO)
   - GORM for database operations
   - Gin for routing and middleware
   - PostgreSQL for data persistence

### Frontend Architecture

1. **Template-Based Rendering**
   - Base layout with common elements
   - Modular templates for features
   - AJAX for dynamic updates

2. **Static Assets**
   - Custom CSS for styling
   - Vanilla JavaScript for interactivity
   - Bootstrap for responsive design

## Implementation Details

### Project Structure

1. **Routes (`internal/routes/`)**
   - Centralized route definitions
   - Grouped by feature (auth, planner)
   - Middleware integration

2. **Database (`internal/repository/`)**
   - Single Database struct for all operations
   - CRUD operations for all models
   - Connection management
   - Migration handling

3. **Handlers**
   - Auth Handler (`internal/auth/`)
   - Planner Handler (`internal/planner/`)
   - Clear separation of concerns

### Authentication System

1. **Local Authentication**
   - Secure password hashing with bcrypt
   - JWT-based session management
   - Password reset functionality

2. **Google SSO**
   - OAuth2 integration
   - Secure user profile handling
   - Unified session management

### Planner Features

1. **To-Do List**
   - CRUD operations
   - User-specific items
   - Completion status

2. **Priorities**
   - Daily priority setting
   - Progress tracking
   - Automatic date management

3. **Contact Reminders**
   - Contact management
   - Notes and details
   - Date tracking

4. **Water Intake**
   - Visual glass counter
   - Daily goal tracking
   - Progress persistence

5. **Random Thoughts**
   - Daily thought generation
   - Persistent storage
   - User association

## Database Schema

```sql
-- Users table
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    username VARCHAR(255) NOT NULL UNIQUE,
    email VARCHAR(255) NOT NULL UNIQUE,
    password_hash VARCHAR(255) NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- Todo items table
CREATE TABLE todo_items (
    id SERIAL PRIMARY KEY,
    user_id INTEGER NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    title VARCHAR(255) NOT NULL,
    description TEXT,
    completed BOOLEAN DEFAULT FALSE,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- Priorities table
CREATE TABLE priorities (
    id SERIAL PRIMARY KEY,
    user_id INTEGER NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    title VARCHAR(255) NOT NULL,
    description TEXT,
    date DATE NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- Contacts table
CREATE TABLE contacts (
    id SERIAL PRIMARY KEY,
    user_id INTEGER NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    name VARCHAR(255) NOT NULL,
    email VARCHAR(255),
    phone VARCHAR(255),
    notes TEXT,
    date DATE NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- Water intake table
CREATE TABLE water_intake (
    id SERIAL PRIMARY KEY,
    user_id INTEGER NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    glasses INTEGER NOT NULL DEFAULT 0,
    target INTEGER NOT NULL DEFAULT 10,
    date DATE NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    UNIQUE(user_id, date)
);

-- Thoughts table
CREATE TABLE thoughts (
    id SERIAL PRIMARY KEY,
    user_id INTEGER NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    content TEXT NOT NULL,
    date DATE NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    UNIQUE(user_id, date)
);
```

## Security Considerations

1. **Authentication**
   - Secure password hashing
   - JWT token encryption
   - CSRF protection
   - Session management

2. **Data Protection**
   - Input validation
   - SQL injection prevention (GORM)
   - XSS protection
   - CORS configuration

## Development Process

1. **Initial Setup**
   - Project structure creation
   - Dependency management
   - Environment configuration

2. **Core Development**
   - Database setup and migrations
   - Route organization
   - Handler implementation
   - Frontend templates

3. **Feature Implementation**
   - Planner features
   - AJAX functionality
   - UI/UX improvements

## Challenges and Solutions

1. **Database Organization**
   - Challenge: Managing multiple repository types
   - Solution: Consolidated into single Database struct with clear operations

2. **Route Management**
   - Challenge: Scattered route definitions
   - Solution: Centralized route package with feature grouping

3. **Migration System**
   - Challenge: Database schema versioning
   - Solution: SQL-based migrations with version tracking

## Future Improvements

1. **Features**
   - Calendar integration
   - Mobile app development
   - Email notifications
   - Data export/import

2. **Technical**
   - Redis caching
   - WebSocket real-time updates
   - API documentation
   - Test coverage improvement

## Deployment Considerations

1. **Requirements**
   - Go 1.22+
   - PostgreSQL 12+
   - Environment configuration
   - SSL certificate

2. **Process**
   - Database migration
   - Environment setup
   - Binary compilation
   - Service configuration

## Maintenance

1. **Regular Tasks**
   - Database backups
   - Log rotation
   - Security updates
   - Performance monitoring

2. **Monitoring**
   - Error tracking
   - Performance metrics
   - User analytics
   - Security audits 