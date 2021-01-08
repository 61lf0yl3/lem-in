package tools

// SortPathsGroups1 sort according to number of rooms in ascending order
func SortPathsGroups1() {
	for i1 := range Paths {
		for i2 := range Paths {
			if len(Paths[i1]) < len(Paths[i2]) {
				Paths[i1], Paths[i2] = Paths[i2], Paths[i1]
			}
		}
	}
}
