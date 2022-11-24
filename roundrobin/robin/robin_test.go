package roundrobin

import (
	"fmt"
	"testing"
)

func TestRoundRobin(t *testing.T) {

	robin()

}

func TestWeightedRoundRobin(t *testing.T) {

	weightedrobin()

}

func TestSmoothRoundRobin(t *testing.T) {
	serverA := server{7, 0, "A"}
	serverB := server{3, 0, "B"}
	// serverC := server{1, 0, "C"}

	servers := []server{serverA, serverB}

	indexs := make([]int, 0, 50)
	for i := 0; i < 50; i++ {
		index := smoothRoundRobin(servers)
		indexs = append(indexs, index)
	}
	fmt.Println(indexs)

}

func SmoothRoundRobinStatus() {

}
