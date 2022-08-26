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

func (c *ReaderWriter) Read(b []byte) (int, error) {
	if c.index >= len(c.buffer) {
		return 0, io.EOF
	}
  count := copy(b, c.buffer[c.index:])
  c.index += count
  return count, nil
}

func (c *ReaderWriter) Write(b []byte) (int, error) {
	c.buffer = append(c.buffer, b...)
	return len(b), nil
}

func main() {
	c := &ReaderWriter{}

	l, err := io.WriteString(c, "Lorem ipsum dolor sit amet, officia excepteur ex fugiat reprehenderit enim labore culpa sint ad nisi Lorem pariatur mollit ex esse exercitation amet. Nisi anim cupidatat excepteur officia. Reprehenderit nostrud nostrud ipsum Lorem est aliquip amet voluptate voluptate dolor minim nulla est proident. Nostrud officia pariatur ut officia. Sit irure elit esse ea nulla sunt ex occaecat reprehenderit commodo officia dolor Lorem duis laboris cupidatat officia voluptate. Culpa proident adipisicing id nulla nisi laboris ex in Lorem sunt duis officia eiusmod. Aliqua reprehenderit commodo ex non excepteur duis sunt velit enim. Voluptate laboris sint cupidatat ullamco ut ea consectetur et est culpa et culpa duis.")
	if err != nil {
		log.Fatalf("Ошибка записи: %v", err)
	}
	fmt.Printf("< Запись в буфер: %d байт >\n%s\n\n", l, c.buffer)

	buf, err := io.ReadAll(c)
	if err != nil {
		log.Fatalf("Ошибка чтения: %v", err)
	}
	fmt.Printf("< Чтение из буфера: %d байт >\n%s\n", len(c.buffer), string(buf))
}
