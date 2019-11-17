package main

import (
	"fmt"
	"math"
	"math/rand"
	"sort"
	"time"
)

type selected struct {
	name  string
	times int
}

func main() {

	source := []string{}

	for i := 1; i <= 10; i++ {
		source = append(source, fmt.Sprintf("Participant %d", i))
	}

	startTime := time.Now()
	resultsByMap(source)
	fmt.Println("Executed in ", time.Now().Sub(startTime).Nanoseconds(), "ns")
	fmt.Println()
	startTime = time.Now()
	resultsByRand(source)
	fmt.Println("Executed in ", time.Now().Sub(startTime).Nanoseconds(), "ns")
}

func resultsByMap(source []string) {
	m := createMap(source)

	vv := map[string]int{}
	for i := 1; i <= 5000; i++ {
		v := getFirst(m)
		vv[v]++
	}

	ov := sortSelected(vv)

	fmt.Println(ov[:int(math.Min(float64(len(ov)), 20))])
	fmt.Println("Results produced by map randomness")
	fmt.Println("Participants: ", len(source))
	fmt.Println("Selected: ", len(ov))
	fmt.Println("Highest: ", ov[0])
	fmt.Println("lowest: ", ov[len(ov)-1])
}

func resultsByRand(source []string) {

	vv := map[string]int{}
	for i := 1; i <= 5000; i++ {
		rand.Seed(time.Now().UnixNano())
		v := source[rand.Intn(len(source))]
		vv[v]++
	}

	ov := sortSelected(vv)

	fmt.Println(ov[:int(math.Min(float64(len(ov)), 20))])
	fmt.Println("Results produced by math/rand randomness")
	fmt.Println("Participants: ", len(source))
	fmt.Println("Selected: ", len(ov))
	fmt.Println("Highest: ", ov[0])
	fmt.Println("lowest: ", ov[len(ov)-1])
}

func createMap(source []string) map[string]bool {
	m := map[string]bool{}
	for _, v := range source {
		m[v] = true
	}
	return m
}

func getFirst(m map[string]bool) string {
	for k := range m {
		return k
	}
	return ""
}

func sortSelected(vv map[string]int) []selected {
	ov := []selected{}
	for k, v := range vv {
		ov = append(ov, selected{
			name:  k,
			times: v,
		})
	}

	sort.SliceStable(ov, func(i, j int) bool {
		return ov[i].times > ov[j].times
	})

	return ov
}
