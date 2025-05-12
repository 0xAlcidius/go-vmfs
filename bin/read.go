package main

import (
	"fmt"

	"github.com/velocidex/go-vmfs/parser"
	ntfs_parser "www.velocidex.com/golang/go-ntfs/parser"
)

var (
	info_command = app.Command(
		"read", "Stat a vmdk file.")

	info_command_file_arg = info_command.Arg(
		"file", "The image file to inspect",
	).Required().File()
)

func doRead() {
	reader, err := ntfs_parser.NewPagedReader(&ntfs_parser.OffsetReader{
		Offset: PartitionOffset,
		Reader: *info_command_file_arg,
	}, 1024, 10000)

	if err != nil {
		fmt.Printf("Error creating reader: %s\n", err)
		return
	}

	buf := make([]byte, 512)
	_, err = reader.ReadAt(buf, 0)
	if err != nil {
		fmt.Printf("Error reading from file: %s\n", err)
		return
	}

	descriptor, err := parser.NewFS3FileDescriptor(buf)
	if err != nil {
		fmt.Printf("Error parsing descriptor: %s\n", err)
		return
	}

	fmt.Printf("Descriptor: %s\n", descriptor.DebugString())
}

func init() {
	command_handlers = append(command_handlers, func(command string) bool {
		switch command {
		case info_command.FullCommand():
			doRead()
		default:
			return false
		}
		return true
	})
}
