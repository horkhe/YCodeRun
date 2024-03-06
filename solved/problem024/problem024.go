// #24. Deadends
// https://coderun.yandex.ru/problem/dead-ends?currentPage=3&pageSize=10&rowNumber=24
package main

import (
	"bufio"
	"container/heap"
	"os"
	"strconv"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)
	sc.Scan()
	deadEndCount, _ := strconv.Atoi(sc.Text())
	sc.Scan()
	trainCount, _ := strconv.Atoi(sc.Text())

	schedule := make([]train, trainCount)
	for i := 0; i < trainCount; i++ {
		train := &schedule[i]
		sc.Scan()
		train.arrivalTime, _ = strconv.Atoi(sc.Text())
		sc.Scan()
		train.departureTime, _ = strconv.Atoi(sc.Text())
	}

	arrangements, failedOn := arrangeTrains(deadEndCount, schedule)

	bw := bufio.NewWriter(os.Stdout)
	if failedOn > 0 {
		_, _ = bw.WriteString("0 ")
		_, _ = bw.WriteString(strconv.Itoa(failedOn))
		_ = bw.Flush()
		return
	}
	for _, deadEnd := range arrangements {
		_, _ = bw.WriteString(strconv.Itoa(deadEnd))
		_ = bw.WriteByte(' ')
	}
	_ = bw.Flush()
}

type train struct {
	arrivalTime   int
	departureTime int
}

func arrangeTrains(deadEndCount int, schedule []train) ([]int, int) {
	arrangements := make([]int, 0, len(schedule))
	// Initialize the list of dead ends available for arriving trains.
	var availableDeadEnds deadEnds
	for i := 1; i <= deadEndCount; i++ {
		heap.Push(&availableDeadEnds, i)
	}

	// Iterate over arriving trains.
	var scheduledDepartures departures
	for i, currTrain := range schedule {
		// Release all dead ends that have been vacated by the time the
		// current train arrived.
		nextDeparture := scheduledDepartures.peek()
		for nextDeparture != nil && nextDeparture.time < currTrain.arrivalTime {
			availableDeadEnds.push(nextDeparture.deadEnd)
			scheduledDepartures.pop()
			nextDeparture = scheduledDepartures.peek()
		}

		// Select the dead end with the smallest number, but return the current
		// train number if none is available.
		selectedDeadEnd, ok := availableDeadEnds.pop()
		if !ok {
			return nil, i + 1
		}

		// Schedule the selected dead end release up on the current train
		// departure from the station.
		scheduledDepartures.push(currTrain.departureTime, selectedDeadEnd)
		// Add the selected dead end to the arrangement list.
		arrangements = append(arrangements, selectedDeadEnd)
	}
	return arrangements, 0
}

type deadEnds []int

func (d *deadEnds) push(deadEnd int) {
	heap.Push(d, deadEnd)
}

func (d *deadEnds) pop() (int, bool) {
	if len(*d) == 0 {
		return 0, false
	}
	return heap.Pop(d).(int), true
}

func (d *deadEnds) Len() int           { return len(*d) }
func (d *deadEnds) Less(i, j int) bool { return (*d)[i] < (*d)[j] }
func (d *deadEnds) Swap(i, j int)      { (*d)[i], (*d)[j] = (*d)[j], (*d)[i] }
func (d *deadEnds) Push(v any)         { *d = append(*d, v.(int)) }
func (d *deadEnds) Pop() any {
	lastIdx := len(*d) - 1
	v := (*d)[lastIdx]
	*d = (*d)[:lastIdx]
	return v
}

type departures []departure

type departure struct {
	time    int
	deadEnd int
}

func (d *departures) peek() *departure {
	if len(*d) == 0 {
		return nil
	}
	return &(*d)[0]
}

func (d *departures) push(departureTime, deadEnd int) {
	heap.Push(d, departure{departureTime, deadEnd})
}

func (d *departures) pop() (departure, bool) {
	if len(*d) == 0 {
		return departure{}, false
	}
	return heap.Pop(d).(departure), true
}

func (d *departures) Len() int           { return len(*d) }
func (d *departures) Less(i, j int) bool { return (*d)[i].time < (*d)[j].time }
func (d *departures) Swap(i, j int)      { (*d)[i], (*d)[j] = (*d)[j], (*d)[i] }
func (d *departures) Push(v any)         { *d = append(*d, v.(departure)) }
func (d *departures) Pop() any {
	lastIdx := len(*d) - 1
	v := (*d)[lastIdx]
	*d = (*d)[:lastIdx]
	return v
}
