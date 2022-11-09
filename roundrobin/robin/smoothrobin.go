package roundrobin

type server struct {
	weight        int
	currentWeight int
	serverName    string
}

func smoothRoundRobin(servers []server) (index int) {
	index = -1
	total := getSum(servers)

	for i := 0; i < len(servers); i++ {
		servers[i].currentWeight += servers[i].weight

		if index == -1 || servers[index].currentWeight < servers[i].currentWeight {
			index = i
		}
	}

	servers[index].currentWeight -= total
	return index
}

func getSum(servers []server) (total int) {
	res := 0

	for i := 0; i < len(servers); i++ {
		res += servers[i].weight
	}
	return res
}
