package mapred

import (
	"fmt"
	"regexp"
	"strings"
)

type MapReduce struct {
}

// todo implement mapreduce

func (m MapReduce) Run(input []string) map[string]int {
	//TODO using go routines to parallize
	mappings := m.wordCountMapper(strings.Join(input, " "))
	fmt.Println(mappings)
	panic("implement me")
}

func (m MapReduce) wordCountMapper(text string) []KeyValue {
	//TODO implement me
	pairs := make([]KeyValue, 0)
	// filter out special characters and numericals
	reg := regexp.MustCompile(`[^A-Za-z]`)
	text = reg.ReplaceAllString(text, " ")

	for _, value := range strings.Fields(text) {
		pair := KeyValue{Key: value, Value: 1}
		pairs = append(pairs, pair)
	}

	return pairs
}

func (m MapReduce) wordCountReducer(key string, values []int) KeyValue {
	//TODO implement me
	panic("implement me")
}
