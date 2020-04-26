package mc

const (
	numSections = 16
)

// Chunk holds data for a chunk column, a 16x256x16 area cosisting of 16
// vertical chunk sections. For more information please refer to
// https://wiki.vg/Chunk_Format
type Chunk struct {
	Position ChunkPosition

	Sections [numSections]ChunkSection
}

// PrimaryBitMask returns a mask describing which sections of a chunk aren't
// empty and should be sent to the client.
func (c *Chunk) PrimaryBitMask() (mask int16) {
	for i, sect := range c.Sections {
		if !sect.Empty {
			mask |= 1 << i
		}
	}

	return mask
}

// ChunkSection is a 16x16x16 collection of blocks. A chunk section holds the
// actual block data. See https://wiki.vg/Chunk_Format for further reading.
type ChunkSection struct {
	// Empty chunk sections consist only of air and no lighting data.
	Empty bool

	// TODO: Blocks data, palette and methods
}
