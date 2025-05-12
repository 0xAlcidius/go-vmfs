package parser

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

const (
	VMFS_FS3_MAGIC  = 0x2fabf15e
	VMFSL_FS3_MAGIC = 0x2fabf15f
)

type FDS_VolInfo struct {
	ID [32]byte
}
type FS3_Checksum struct {
	Value       uint64
	ChecksumGen uint64
}

type FS3FileDescriptor struct {
	Magic                  uint32
	MajorVersion           uint32
	MinorVersion           uint8
	UUID                   [16]byte
	Config                 uint32
	FSLabel                [128]byte
	DiskBlockSize          uint32
	FileBlockSize          uint64
	CreationTime           uint32
	SnapID                 uint32
	VolInfo                FDS_VolInfo
	FDCClusterGroupOffset  uint32
	FDCClustersPerGroup    uint32
	SubBlockSize           uint32
	MaxJournalSlotsPerTxn  uint32
	PB2VolAddr             uint64
	PB2FDAddr              uint32
	HostUUID               [16]byte
	GBLGeneration          uint64
	SDDVolAddr             uint64
	SDDFDAddr              uint32
	ChecksumType           uint8
	UnmapPriority          uint16
	Pad1                   [4]byte
	ChecksumGen            uint64
	Checksum               FS3_Checksum
	PhysDiskBlockSize      uint32
	MDAlignment            uint32
	SFBToLFBShift          uint16
	Reserved16_1           uint16
	Reserved16_2           uint16
	PtrBlockShift          uint16
	SFBAddrBits            uint16
	Reserved16_3           uint16
	TBZGranularity         uint32
	JournalBlockSize       uint32
	LeaseIntervalMs        uint32
	ReclaimWindowMs        uint32
	LocalStampUS           uint64
	LocalMountOwnerMacAddr [6]byte
}

func (self *FS3FileDescriptor) IsVMFS5() bool {
	return self.MajorVersion <= 0x17
}

func (self *FS3FileDescriptor) IsVMFS6() bool {
	return self.MajorVersion >= 0x18
}

func (self *FS3FileDescriptor) DebugString() string {
	return fmt.Sprintf("FS3FileDescriptor{\nMagic: %x,\n\tMajorVersion: %d,\n\tMinorVersion: %d,\n\tUUID: %x,\n\tConfig: %d,\n\tFSLabel: %s,\n\tDiskBlockSize: %d,\n\tFileBlockSize: %d,\n\tCreationTime: %d,\n\tSnapID: %d,\n\tVolInfo: %x,\n\tFDCClusterGroupOffset: %d,\n\tFDCClustersPerGroup: %d,\n\tSubBlockSize: %d,\n\tMaxJournalSlotsPerTxn: %d,\n\tPB2VolAddr: %x,\n\tPB2FDAddr: %d,\n\tHostUUID: %x,\n\tGBLGeneration: %d,\n\tSDDVolAddr: %x,\n\tSDDFDAddr: %d,\n\tChecksumType: %d,\n\tUnmapPriority: %d,\n\tPad1: %x,\n\tChecksumGen: %d,\n\tChecksum: {Value:%x, ChecksumGen:%x},\n\tPhysDiskBlockSize: %d,\n\tMDAlignment: %d,\n\tSFBToLFBShift: %d,\n\tReserved16_1: %d,\n\tReserved16_2: %d,\n\tPtrBlockShift: %d,\n\tSFBAddrBits: %d,\n\tReserved16_3: %d,\n}\n",
		self.Magic, self.MajorVersion, self.MinorVersion, self.UUID, self.Config, self.FSLabel[:], self.DiskBlockSize, self.FileBlockSize, self.CreationTime, self.SnapID, self.VolInfo.ID[:], self.FDCClusterGroupOffset, self.FDCClustersPerGroup, self.SubBlockSize, self.MaxJournalSlotsPerTxn, self.PB2VolAddr, self.PB2FDAddr, self.HostUUID[:], self.GBLGeneration, self.SDDVolAddr, self.SDDFDAddr, self.ChecksumType, self.UnmapPriority, self.Pad1[:], self.ChecksumGen, self.Checksum.Value, self.Checksum.ChecksumGen, self.PhysDiskBlockSize, self.MDAlignment, self.SFBToLFBShift, self.Reserved16_1, self.Reserved16_2, self.PtrBlockShift, self.SFBAddrBits, self.Reserved16_3)
}

func NewFS3FileDescriptor(buf []byte) (*FS3FileDescriptor, error) {
	if len(buf) < 512 {
		return nil, fmt.Errorf("buffer too small")
	}

	fd := &FS3FileDescriptor{}
	err := binary.Read(bytes.NewReader(buf), binary.LittleEndian, fd)
	if err != nil {
		return nil, err
	}

	if fd.Magic != VMFS_FS3_MAGIC && fd.Magic != VMFSL_FS3_MAGIC {
		return nil, fmt.Errorf("invalid magic number")
	}

	return fd, nil
}
