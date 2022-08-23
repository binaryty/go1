package main

import (
	"fmt"
	"io"
	"log"
)

type Customizer struct {
	buffer []byte
	index  int
}

func (c *Customizer) Read(b []byte) (int, error) {
	if c.index >= len(c.buffer) {
		return 0, io.EOF
	}

	if cap(b) >= cap(c.buffer) {
		for i, v := range c.buffer {
			b[i] = v
			c.index += i
		}
		return c.index, nil
	} else {
		//Todo: fix slice overflow!
		return c.index, fmt.Errorf("Slice overflow in Read func\n")
	}
}

func (c *Customizer) Write(b []byte) (int, error) {
	if len(c.buffer) == 0 {
		c.buffer = b
		return len(b), nil
	}
	c.buffer = append(c.buffer, b...)
	return len(b), nil
}

func main() {
	c := &Customizer{}

	l, err := io.WriteString(c, "Lorem ipsum dolor sit amet, officia excepteur ex fugiat reprehenderit enim labore culpa sint ad nisi Lorem pariatur mollit ex esse exercitation amet. Nisi anim cupidatat excepteur officia. Reprehenderit nostrud nostrud ipsum Lorem est aliquip amet voluptate voluptate dolor minim nulla est proident. Nostrud officia pariatur ut officia. Sit irure elit esse ea nulla sunt ex occaecat reprehenderit commodo officia dolor Lorem duis laboris cupidatat officia voluptate. Culpa proident adipisicing id nulla nisi laboris ex in Lorem sunt duis officia eiusmod. Aliqua reprehenderit commodo ex non excepteur duis sunt velit enim. Voluptate laboris sint cupidatat ullamco ut ea consectetur et est culpa et culpa duis.")
	if err != nil {
		log.Fatalf("Ошибка записи: %v", err)
	}
	fmt.Printf("Записано в буфер: %s | %d байт\n\n", c.buffer, l)

	buf, err := io.ReadAll(c)
	if err != nil {
		log.Fatalf("Ошибка чтения: %v", err)
	}
	fmt.Printf("Чтение из буфера: %s | %d байт", string(buf), len(c.buffer))
}
