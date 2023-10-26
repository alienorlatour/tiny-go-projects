package habit

import (
	"context"

	"github.com/google/uuid"
)

type habitCreator interface {
	Add(ctx context.Context, habit Habit) error
}

// CreateHabit adds a habit into the DB.
func CreateHabit(ctx context.Context, db habitCreator, h Habit) error {
	if h.ID == "" {
		h.ID = ID(uuid.NewString())
	}

	return db.Add(ctx, h)
}
