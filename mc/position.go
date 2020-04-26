package mc

// Position holds the position of players, mobs, vehicles etc., not bound to the
// block grid.
type Position struct {
	X float64
	Y float64
	Z float64

	Yaw   float32
	Pitch float32

	OnGround bool
}

// BlockPosition holds data for a block's position in the world. The position
// is absolute, relative to (0, 0, 0).
type BlockPosition struct {
	X int
	y int
	Z int
}

// ChunkPosition holds position of a chunk column on the chunk grid.
// ChunkPosition X and Z can be translated to world position by multiplying each
// by 16.
type ChunkPosition struct {
	X int
	Z int
}
