package main

import (
	"fmt"

	"github.com/alecthomas/kingpin/v2"
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
	kingpin.FatalIfError(err, "Error opening file: %s", err)

	buf := make([]byte, parser.FS3_Descriptor_Size)
	_, err = reader.ReadAt(buf, 0)
	kingpin.FatalIfError(err, "Error reading file: %s", err)

	descriptor, err := parser.NewFS3Descriptor(buf)
	kingpin.FatalIfError(err, "Error parsing descriptor: %s", err)

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
