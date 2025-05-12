package parser

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
)

const (
	FS3_FileDescriptor_Size = 0x16c
)

type FS3_FileDescriptor struct {
	Address             uint32
	Generation          uint32
	NumLinks            uint32
	Type                uint32
	Flags               uint32
	Length              uint64
	BlockSize           uint64
	NumBlocks           uint64
	ModificationTime    uint32
	CreationTime        uint32
	AccessTime          uint32
	UID                 uint32
	GID                 uint32
	Mode                uint32
	ZLA                 uint32
	TBZLo               uint32
	COWLo               uint32
	NewSinceEpochLo     uint32
	TBZHi               uint32
	COWHi               uint32
	NumPointerBlocks    uint32
	NewSinceEpochHi     uint32
	Unk1                uint32
	AffinityFD          uint32
	TBZGranularityShift uint32
	ParentFD            uint32
	LastSFBClusterNum   uint32
	Unk4                uint32
	Unk5                uint32
	Unk6                uint32
	NumPreAllocBlocks   uint8
	Unk7                uint8
	Unk8                uint8
	Unk9                uint8
	Unk10               uint8
	BlockOffsetShift    uint8
	NumTracked          uint8
	Unk12               uint8
	NumLFB              uint32
	Unk13               [216]byte
	LastFreeSFBC        uint32
}

func NewFS3FileDescriptor(r io.ReaderAt, addr uint64, blockSize uint64) (*FS3_FileDescriptor, error) {
	buf := make([]byte, blockSize)
	offset := int64(addr) * int64(blockSize)
	_, err := r.ReadAt(buf, offset)
	if err != nil {
		return nil, fmt.Errorf("failed to read descriptor at %x: %v", offset, err)
	}

	var fd FS3_FileDescriptor
	err = binary.Read(bytes.NewReader(buf), binary.LittleEndian, &fd)
	if err != nil {
		return nil, fmt.Errorf("binary.Read failed: %v", err)
	}

	return &fd, nil
}
