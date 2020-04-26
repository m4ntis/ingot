package mc

// Block represends a single block in a ChunkSection.
type Block struct {
	ID BlockID

	// TODO: Textual ID
}

func NewBlock(id BlockID) Block {
	return Block{id}
}
