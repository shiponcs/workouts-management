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