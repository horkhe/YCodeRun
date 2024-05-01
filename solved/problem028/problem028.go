// #28. Longest ascending sequence with response recovery
// https://coderun.yandex.ru/problem/nvp-with-response-recovery/description?currentPage=3&pageSize=10&rowNumber=28
// The idea is to traverse the sequence and build a list of possible ascending
// subsequences. Each number goes to just one subsequence that would be the
// longest ending in that number.
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	scanner.Scan()
	seqLen, _ := strconv.Atoi(scanner.Text())
	seq := make([]int, seqLen)
	for i := range seq {
		scanner.Scan()
		seq[i], _ = strconv.Atoi(scanner.Text())
	}
	las := findLongestAscendingSubseq(seq)
	for i := range las {
		fmt.Print(las[i], " ")
	}
	fmt.Print("\n")
}

func findLongestAscendingSubseq(seq []int) []int {
	if len(seq) < 1 {
		return nil
	}
	subseqs := [][]int{seq[:1]}
	seq = seq[1:]
	for _, n := range seq {
		chosenSeqIdx := 0
		chosenSeq := subseqs[chosenSeqIdx]
		chosenSeqPos := findInsertPos(chosenSeq, n)
		for i := 1; i < len(subseqs); i++ {
			ss := subseqs[i]
			pos := findInsertPos(ss, n)
			if pos > chosenSeqPos {
				chosenSeqIdx = i
				chosenSeq = ss
				chosenSeqPos = pos
			}
		}
		if chosenSeqPos == len(chosenSeq) {
			subseqs[chosenSeqIdx] = append(chosenSeq, n)
			continue
		}
		newSubSeq := make([]int, chosenSeqPos+1)
		copy(newSubSeq, chosenSeq[:chosenSeqPos])
		newSubSeq[chosenSeqPos] = n
		subseqs = append(subseqs, newSubSeq)
	}
	longestAscendingSubseqIdx := 0
	for i := 1; i < len(subseqs); i++ {
		if len(subseqs[longestAscendingSubseqIdx]) < len(subseqs[i]) {
			longestAscendingSubseqIdx = i
		}
	}
	return subseqs[longestAscendingSubseqIdx]
}

func findInsertPos(s []int, n int) int {
	begin := 0
	end := len(s)
	for {
		if end == begin {
			return begin
		}
		pivot := begin + (end-begin)/2
		if s[pivot] < n {
			begin = pivot + 1
			continue
		}
		end = pivot
	}
}
