package tools

// MaxHops the number of hops in the combination of path
func MaxHops(flow [][]string) int {
	maxhops := 0
	for _, paths := range flow {
		if maxhops <= len(paths) {
			maxhops = len(paths)
		}
	}
	return maxhops
}

// FillAnts number of ants use for fill the other paths for equato to each other
func FillAnts(flow [][]string) int {
	fillants := 0
	for _, paths := range flow {
		for i := len(paths); i < MaxHops(flow); i++ {
			fillants++
		}
	}
	return fillants
}

// GetMinHeightValue get the value of min height
func GetMinHeightValue(heights map[int][][]string) [][]string {
	var optimalflow = [][]string{}
	var min int
	for key := range heights {
		min = key
		break
	}

	for key, value := range heights {
		if min >= key {
			min = key
			optimalflow = value
		}
	}
	return optimalflow
}

//RmStart remove rooms start from all paths
func RmStart() {
	for i, rooms := range Paths {
		Paths[i] = rooms[1:]
	}
}
