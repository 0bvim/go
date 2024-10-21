package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"sort"
	"strings"
)

type Pair struct {
	Key   string
	Value int
}

type PairList []Pair

func main() {

	file, err := os.Open("text.txt")
	if err != nil {
		log.Fatalf("error: %s", err)
	}
	defer file.Close()

	words, err := word_counter(file)
	if err != nil {
		log.Fatalf("error from freq in main: %s", err)
	}

	sorted := sortWordFrequency(words)

	for _, pair := range sorted {
		fmt.Printf("%v\t %v\n", pair.Key, pair.Value)
	}
}

func word_counter(reader io.Reader) (map[string]int, error) {
	// creating a map to store word frequency
	wordFreq := make(map[string]int)

	// Create a scanner to read the file
	s := bufio.NewScanner(reader)
	s.Split((bufio.ScanWords))

	for s.Scan() {
		word := strings.ToLower(s.Text())
		word = strings.Trim(word, ".")
		word = strings.Trim(word, ",")
		word = strings.Trim(word, "\"")
		wordFreq[word]++
	}
	if err := s.Err(); err != nil {
		return nil, err
	}

	return wordFreq, nil
}

// implement methods to PairList
func (pair PairList) Len() int {
	return len(pair)
}

// sort in descending order
func (pair PairList) Less(i, j int) bool {
	return pair[i].Value > pair[j].Value
}

func (pair PairList) Swap(i, j int) {
	pair[i], pair[j] = pair[j], pair[i]
}

func sortWordFrequency(word_map map[string]int) PairList {
	// convert the map to a pair list
	pairs := make(PairList, len(word_map))

	i := 0
	for key, value := range word_map {
		pairs[i] = Pair{key, value}
		i++
	}

	// std module
	sort.Sort(pairs)

	return pairs
}
