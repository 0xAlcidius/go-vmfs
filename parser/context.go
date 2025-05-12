package parser

import (
	"fmt"
	"io"
)

type VMFSContext struct {
	Volume     io.ReaderAt
	BlockSize  uint64
	Descriptor *FS3Descriptor
}

func GetVMFSContext(reader io.ReaderAt) (*VMFSContext, error) {
	buf := make([]byte, FS3_Descriptor_Size)
	_, err := reader.ReadAt(buf, 0)
	if err != nil {
		return nil, err
	}

	desc, err := NewFS3Descriptor(buf)
	if err != nil {
		return nil, err
	}

	fmt.Println(desc.DebugString())

	ctx := &VMFSContext{
		Volume:     reader,
		BlockSize:  uint64(desc.FileBlockSize),
		Descriptor: desc,
	}

	return ctx, nil
}
