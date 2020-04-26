package mc

import (
	"github.com/google/uuid"
)

// Player holds an online player's state.
type Player struct {
	UUID uuid.UUID

	Position Position
}
