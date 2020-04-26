package mc

import "fmt"

// BlockID is the numerical block ID generated from minecraft's blocks.json
// file. To generate the file, download the latest minecraft-server.jar, and
// run following command:
//
// java -cp minecraft-server.jar net.minecraft.data.Main --reports
//
// For further information, please refer to https://wiki.vg/Data_Generators
type BlockID int

const (
	Air = iota
	Stone
)

const blockCount = 2

func GetBlockID(id int) (BlockID, error) {
	if id < 0 || id >= blockCount {
		return 0, ErrInvalidBlockID{id}
	}

	return BlockID(id), nil
}

// ErrInvalidBlockID returnes when attempting to create a BlockID using an out
// of range ID.
type ErrInvalidBlockID struct {
	id int
}

func (e ErrInvalidBlockID) Error() string {
	return fmt.Sprintf("invalid block id. id: %d, expected -1<id<%d", blockCount)
}
