package parser

import (
	"io"
	"sync"
)

type Reader struct {
	mu sync.Mutex

	reader io.ReaderAt
	offset int64
}

// func (sel *Reader) ReadAt(buf []byte, offset int64) (n int, err error) {
// 	if offset < 0 {
// 		return 0, io.EOF
// 	}

// 	sel.mu.Lock()
// 	defer self.mu.Unlock()

// }

func NewReader(reader io.ReaderAt, offset int64) *Reader {
	return &Reader{
		reader: reader,
		offset: offset,
	}
}
