package main

import (
	"fmt"
	"io"
	"log"
)

type ReaderWriter struct {
	buffer []byte
	index  int
}

func (rw *ReaderWriter) Clear(b []byte) {
	if len(rw.buffer) > 0 {
		rw.buffer = rw.buffer[:0]
	}
	if len(b) > 0 {
		b = b[:0]
	}
}

func (rw *ReaderWriter) Read(b []byte) (int, error) {
	if rw.index >= len(rw.buffer) {
		rw.Clear(b)
		return 0, io.EOF
	}
	count := copy(b, rw.buffer[rw.index:])
	rw.index += count
	return count, nil
}

func (rw *ReaderWriter) Write(b []byte) (int, error) {
	rw.buffer = append(rw.buffer, b...)
	return len(b), nil
}

func main() {
	var rw = &ReaderWriter{}

	l, err := io.WriteString(rw, "Lorem ipsum dolor sit amet, officia excepteur ex fugiat reprehenderit enim labore culpa sint ad nisi Lorem pariatur mollit ex esse exercitation amet. Nisi anim cupidatat excepteur officia. Reprehenderit nostrud nostrud ipsum Lorem est aliquip amet voluptate voluptate dolor minim nulla est proident. Nostrud officia pariatur ut officia. Sit irure elit esse ea nulla sunt ex occaecat reprehenderit commodo officia dolor Lorem duis laboris cupidatat officia voluptate. Culpa proident adipisicing id nulla nisi laboris ex in Lorem sunt duis officia eiusmod. Aliqua reprehenderit commodo ex non excepteur duis sunt velit enim. Voluptate laboris sint cupidatat ullamco ut ea consectetur et est culpa et culpa duis.")
	if err != nil {
		log.Fatalf("Ошибка записи: %v", err)
	}
	fmt.Printf("< Запись в буфер: %d байт >\n%s\n\n", l, rw.buffer)

	buf, err := io.ReadAll(rw)
	if err != nil {
		log.Fatalf("Ошибка чтения: %v", err)
	}
	fmt.Printf("< Чтение из буфера: %d байт >\n%s\n", rw.index, string(buf))
}
