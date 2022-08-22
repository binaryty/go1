package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
)

type Node struct {
	value int
	next  *Node
}

type List struct {
	head *Node
	len  int
}

func (L *List) InsertFront(value int) {
	newNode := &Node{
		value: value,
	}

	if L.head == nil {
		L.head = newNode
	} else {
		newNode.next = L.head
		L.head = newNode
	}
	L.len++
}

func (L *List) InsertBack(value int) {
	newNode := &Node{
		value: value,
	}
	if L.head == nil {
		L.head = newNode
	} else {
		p := L.head
		for p.next != nil {
			p = p.next
		}
		p.next = newNode
	}
	L.len++
}

func (L *List) RemoveFront() bool {
	if L.head == nil {
		return false
	}
	L.head = L.head.next
	L.len--
	return true
}

func (L *List) RemoveBack() bool {
	if L.head == nil {
		return false
	}
	var prev *Node
	p := L.head
	for p.next != nil {
		prev = p
		p = p.next
	}
	if prev != nil {
		prev.next = nil
	} else {
		L.head = nil
	}
	L.len--
	return true
}

func (L *List) Print() {
	if L.head == nil {
		fmt.Println("Элементы списка: Спсок пуст")
		return
	}
	p := L.head
	fmt.Print("Элементы списка: ")
	for p != nil {
		fmt.Printf("%d|", p.value)
		p = p.next
	}
	fmt.Printf("%v\n", nil)
}

func clear() {
	fmt.Print("\033[H\033[2J")
}

func main() {

	list := List{}
	scanner := bufio.NewScanner(os.Stdin)
	clear()

	for {
		fmt.Println("Выберите команду для работы с односвязным списком из предложенных ниже.")
		fmt.Print("1-Добавить в начало\n2-Добавить в конец\n3-Удалить из начала\n4-Удалить с конца\n")
		list.Print()
		scanner.Scan()
		input := scanner.Text()
		if input == "q" {
			break
		}
		operation, err := strconv.ParseInt(input, 10, 64)
		if err != nil || operation < 0 || operation > 4 {
			clear()
			fmt.Println("Неизвестная команда. Повторите попытку")
			continue
		}
		switch operation {
		case 1:
			list.InsertFront(rand.Intn(100))
		case 2:
			list.InsertBack(rand.Intn(100))
		case 3:
			list.RemoveFront()
		case 4:
			list.RemoveBack()
		}
		clear()
	}
}
