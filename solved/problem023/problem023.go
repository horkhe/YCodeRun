// #23. Goblins and Chess
// https://coderun.yandex.ru/problem/goblins-and-chess?currentPage=3&pageSize=10&rowNumber=23
// The idea is to implement a dequeue that maintains pointers to its head,
// middle and tail, so that pushTail, pushMiddle and popHead operations take O(1).
package main

import (
	"bufio"
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {
	process(os.Stdin, os.Stdout)
}

func process(r io.Reader, w io.Writer) {
	br := bufio.NewReader(r)
	bw := bufio.NewWriter(w)
	line, _ := br.ReadString('\n')
	count, _ := strconv.Atoi(strings.TrimSpace(line))

	// The problem condition states that there could be no more than 100000
	// commands. So we allocate memory to fit all of them.
	q := newQueue(100000)
	for i := 0; i < count; i++ {
		line, _ = br.ReadString('\n')
		cmd := line[0]
		switch cmd {
		case '+':
			num, _ := strconv.Atoi(strings.TrimSpace(line[1:]))
			q.pushTail(num)
		case '*':
			num, _ := strconv.Atoi(strings.TrimSpace(line[1:]))
			q.pushMiddle(num)
		case '-':
			num := q.popHead()
			_, _ = bw.WriteString(strconv.Itoa(num))
			_ = bw.WriteByte('\n')
		}
	}
	_ = bw.Flush()
}

func newQueue(len int) *queue {
	q := queue{buf: make([]queueElem, len)}
	return &q
}

type queue struct {
	buf       []queueElem
	head      int // index of the head element
	tail      int // index of the tail element
	middle    int // index of the middle element
	insertPos int // index of the next free slot to add a new element
	count     int // the number of elements in the queue.
}

func (q *queue) pushTail(v int) {
	insertElem := queueElem{val: v, prev: -1, next: -1}
	if q.count == 0 {
		q.head = q.insertPos
		q.middle = q.insertPos
	} else {
		insertElem.prev = q.tail
		q.buf[q.tail].next = q.insertPos
		if q.count&1 == 0 {
			q.middle = q.buf[q.middle].next
		}
	}
	q.buf[q.insertPos] = insertElem
	q.tail = q.insertPos
	q.count++
	q.insertPos++
}

func (q *queue) pushMiddle(v int) {
	insertElem := queueElem{val: v, prev: -1, next: -1}
	if q.count == 0 {
		q.head = q.insertPos
		q.middle = q.insertPos
		q.tail = q.insertPos
	} else {
		insertElem.prev = q.middle
		insertElem.next = q.buf[q.middle].next
		q.buf[q.middle].next = q.insertPos
		if q.count&1 == 0 {
			q.middle = q.insertPos
		}
		if q.count == 1 {
			q.tail = q.insertPos
		}
	}
	q.buf[q.insertPos] = insertElem
	q.count++
	q.insertPos++
}

func (q *queue) popHead() int {
	headElem := q.buf[q.head]
	q.count--
	if q.count > 0 {
		q.head = headElem.next
		if q.count&1 == 1 {
			q.middle = q.buf[q.middle].next
		}
	}
	return headElem.val
}

type queueElem struct {
	val  int
	next int
	prev int
}
