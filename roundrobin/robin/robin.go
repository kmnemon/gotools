package roundrobin

import "fmt"

func robin() {
	serverNum := 4
	selectID := serverNum - 1

	j := selectID

	for c := 0; c < 100; c++ {
		j = (j + 1) % serverNum
		selectID = j
		fmt.Println(selectID)
	}
}
