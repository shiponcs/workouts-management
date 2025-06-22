package store

import (
	"database/sql"
	"testing"

	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func setupTestDB(t *testing.T) *sql.DB {
	db, err := sql.Open("pgx", "host=localhost user=postgres password=postgres dbname=postgres port=5433 sslmode=disable")
	if err != nil {
		t.Fatalf("Opening test db: %v", err)
	}

	err = Migrate(db, "../../migrations")
	if err != nil {
		t.Fatalf("migration test db error: %v", err)
	}

	_, err = db.Exec("TRUNCATE workouts, workout_entries CASCADE")
	if err != nil {
		t.Fatalf("truncating table err %v", err)
	}

	return db
}

func TestCreateWorkout(t *testing.T) {
	db := setupTestDB(t)
	defer db.Close()

	store := NewPostgresWorkoutStore(db)

	tests := []struct {
		name    string
		workout *Workout
		wantErr bool
	}{
		{
			name: "valid workout",
			workout: &Workout{
				Title:           "push day",
				Description:     "upper body day",
				DurationMinutes: 60,
				CaloriesBurned:  200,
				Entries: []WorkoutEntry{
					{
						ExerciseName: "bench press",
						Sets:         3,
						Reps:         IntPtr(10),
						Weight:       FloatPtr(122.2),
						Notes:        "warm up properly",
					},
				},
			},
			wantErr: true,
		},
		{
			name: "workout with invalid entries",
			workout: &Workout{
				Title:           "workout with invalid entries",
				Description:     "complete workout",
				DurationMinutes: 90,
				CaloriesBurned:  500,
				Entries: []WorkoutEntry{
					{
						ExerciseName: "plank",
						Sets:         3,
						Reps:         IntPtr(60),
						Notes:        "keep form",
						OrderIndex:   1,
					},
					{

						ExerciseName:    "squats",
						Sets:            4,
						Reps:            IntPtr(40),
						DurationSeconds: IntPtr(4999),
						Weight:          FloatPtr(84.4),
						Notes:           "full depth",
						OrderIndex:      2,
					},
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			creaatedWorkout, err := store.CreateWorkout(tt.workout)
			0
			require.NoError(t, err)
			assert.Equal(t, tt.workout.Title, creaatedWorkout.Title)
			assert.Equal(t, tt.workout.Description, creaatedWorkout.Description)
			assert.Equal(t, tt.workout.DurationMinutes, creaatedWorkout.DurationMinutes)

			retrieved, err := store.GetWorkoutByID(int64(creaatedWorkout.ID))
			require.NoError(t, err)

			for i := range retrieved.Entries {
				assert.Equal(t, tt.workout.Entries[i].ExerciseName, retrieved.Entries[i].ExerciseName)
				assert.Equal(t, tt.workout.Entries[i].Sets, retrieved.Entries[i].Sets)
				assert.Equal(t, tt.workout.Entries[i].OrderIndex, retrieved.Entries[i].OrderIndex)

			}
		})
	}
}

func IntPtr(i int) *int {
	return &i
}

func FloatPtr(i float64) *float64 {
	return &i
}
