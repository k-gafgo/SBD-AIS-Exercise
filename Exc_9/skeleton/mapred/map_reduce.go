package mapred

import (
	"regexp"
	"strings"
	"sync"
)

type MapReduce struct {
}

// todo implement mapreduce

func (m MapReduce) Run(input []string) map[string]int {
	// creating channels for data exchange
	mapOut := make(chan KeyValue)
	reduceOut := make(chan KeyValue)

	// MAP
	var mapWaitGroup sync.WaitGroup

	for _, text := range input {
		mapWaitGroup.Add(1)
		// start a new go routine for every line
		go func(t string) {
			defer mapWaitGroup.Done()
			// write the pairs into the channel
			for _, pairs := range m.wordCountMapper(t) {
				mapOut <- pairs
			}
		}(text)
	}
	// close mapOut once mapping is done
	go func() {
		mapWaitGroup.Wait()
		close(mapOut)
	}()

	// SHUFFLE
	shuffled := make(map[string][]int)

	for pair := range mapOut {
		shuffled[pair.Key] = append(shuffled[pair.Key], pair.Value)
	}

	// REDUCE
	var reduceWaitGroup sync.WaitGroup

	for key, values := range shuffled {
		reduceWaitGroup.Add(1)
		// start new go routine for every key
		go func(k string, v []int) {
			defer reduceWaitGroup.Done()
			// write the result to the channel
			reduceOut <- m.wordCountReducer(k, v)
		}(key, values)
	}

	// Close reduceOut when done
	go func() {
		reduceWaitGroup.Wait()
		close(reduceOut)
	}()

	result := make(map[string]int)
	for pair := range reduceOut {
		result[pair.Key] = pair.Value
	}

	return result
}

func (m MapReduce) wordCountMapper(text string) []KeyValue {
	pairs := make([]KeyValue, 0)
	text = strings.ToLower(text)
	// filter out special characters and numericals
	reg := regexp.MustCompile(`[^A-Za-z]`)
	text = reg.ReplaceAllString(text, " ")

	for _, word := range strings.Fields(text) {
		pair := KeyValue{Key: word, Value: 1}
		pairs = append(pairs, pair)
	}

	return pairs
}

func (m MapReduce) wordCountReducer(key string, values []int) KeyValue {
	sum := 0
	for _, value := range values {
		sum += value
	}
	return KeyValue{Key: key, Value: sum}
}
