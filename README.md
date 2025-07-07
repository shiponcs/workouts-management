# Workouts Management Service API
This is a service that helps managing Workouts

### What is it based upon?
- [jackc/pgx](https://github.com/jackc/pgx) as PostgreSQL driver
- [pressly/goose](https://github.com/pressly/goose) as Database Migration Tool
- [go-chi/chi](https://github.com/go-chi/chi) as HTTP routing
- Stateful Token for authentication


### Sample curl commands
#### Create a new user
```bash
curl -X POST "http://localhost:8080/users" \
     -H "Content-Type: application/json" \
     -d '{
          "username": "melkey",
          "email": "melkey@example.com",
          "password": "SecureP@ssword123",
          "bio": "Fitness enthusiast and software developer"
        }'
```
#### Get a token (aka login)
```bash
curl -X POST "http://localhost:8080/tokens/authentication" \
     -H "Content-Type: application/json" \
     -d '{
          "username": "johndoe",
          "password": "SecureP@ssword123"
        }'
```
#### create a workout
```bash
curl -X POST "http://localhost:8080/workouts" \
     -H "Authorization: Bearer 3VGSV7LBT4IVASURIKUOUWOF5WGHOWRXUZA6OLKK7U4ABNCLPVSA" \
     -H "Content-Type: application/json" \
     -d '{
          "title": "Morning Cardio",
          "description": "A light 30-minute jog to start the day.",
          "duration_minutes": 30,
          "calories_burned": 300,
          "entries": [
              {
                  "exercise_name": "Jogging",
                  "sets": 1,
                  "duration_seconds": 1800,
                  "weight": 0,
                  "notes": "Maintain a steady pace",
                  "order_index": 1
              }
          ]
        }'
```
#### Update a workout that you own

```bash
curl -X PUT "http://localhost:8080/workouts/6" \
     -H "Authorization: Bearer M6BOKTWSQJL74GJJO5OIOMOG3V63MYRIKLUBZ6ILEUQUPCN7472Q" \
     -H "Content-Type: application/json" \
     -d '{
          "title": "Updated Cardio",
          "description": "A relaxed 45-minute walk after dinner.",
          "duration_minutes": 45,
          "calories_burned": 250,
          "entries": [
              {
                  "exercise_name": "Walking",
                  "sets": 1,
                  "duration_seconds": 2700,
                  "weight": 0,
                  "notes": "Keep a steady pace",
                  "order_index": 1
              }
          ]
        }'
```

# Project Architecture Documentation

## Overview
This project follows several well-established architectural principles and patterns for building a robust and maintainable Go web service.

## Primary Architecture: **Layered/Clean Architecture**

The project is organized into distinct layers with clear separation of concerns:

1. **Presentation Layer** - `internal/api` handlers
2. **Business/Application Layer** - `internal/app` application setup
3. **Data Access Layer** - `internal/store` data persistence
4. **Infrastructure Layer** - `internal/middleware`, `internal/routes`

## Key Architectural Principles Applied

### 1. **Dependency Injection Pattern**
- The `app.NewApplication()` function creates and wires dependencies
- Handlers receive store interfaces through constructors (e.g., `NewWorkoutHandler`)

### 2. **Repository Pattern**
- Abstract interfaces like `WorkoutStore`, `UserStore`, and `TokenStore`
- Concrete implementations like `PostgresWorkoutStore`

### 3. **Interface Segregation Principle (ISP)**
- Small, focused interfaces for each store type
- Handlers depend on interfaces, not concrete implementations

### 4. **Single Responsibility Principle (SRP)**
- Each handler has a single responsibility (e.g., `WorkoutHandler` only handles workout operations)
- Separate concerns: authentication (`middleware`), routing (`routes`), data access (`store`)

### 5. **Middleware Pattern**
- Authentication and authorization implemented as middleware in `internal/middleware/middleware.go`
- Applied via `routes.SetupRoutes()`

### 6. **Domain-Driven Design (DDD) Elements**
- Domain entities like `Workout`, `User`
- Business logic encapsulated in methods (e.g., password hashing in `user_store.go`)

### 7. **Database Migration Pattern**
- Structured migrations in `migrations` directory
- Embedded filesystem pattern with `migrations/fs.go`

### 8. **Optimistic Concurrency Control**
- Version field in workouts for handling concurrent updates
- Implemented in `UpdateWorkout` method
- Test scripts demonstrate this pattern: `test_scripts/test_optimistic_concurrency_control.go`

### 9. **RESTful API Design**
- Standard HTTP methods and status codes
- Resource-based URLs (e.g., `/workouts/{id}`)
- JSON request/response format using `utils.WriteJSON`

## Benefits of This Architecture

This architecture promotes:
- **Testability** - Clear interfaces make unit testing easier
- **Maintainability** - Separation of concerns allows for easier modifications
- **Loose Coupling** - Components depend on abstractions, not concrete implementations
- **Scalability** - Layered approach allows for easy horizontal scaling
- **Code Reusability** - Interface-based design promotes code reuse

