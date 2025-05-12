package main

import (
	"fmt"

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
	fmt.Println("Read data:")
	for i, b := range buf {
		if i%16 == 0 {
			fmt.Printf("\n%08x: ", i)
		}
		fmt.Printf("%02x ", b)
	}

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
