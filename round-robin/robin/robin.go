package roundrobin

import (
	"fmt"
)

func robin() {
	sn := 2
	i := sn - 1

	j := i

	fmt.Println((j + 1) / sn)
}


func main(){
	robin()
}