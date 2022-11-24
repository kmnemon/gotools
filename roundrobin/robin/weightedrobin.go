package roundrobin

func weightedrobin() {
	// maxS := 100
	// selectID := -1
	// currentWeight := maxS

}

func gcd(a int, b int) int {
	var c int
	for b != 0 {
		c = b
		b = a % b
		a = c
	}
	return a
}
