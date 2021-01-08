package lib

// SortPaths sorts paths by their length
func (g *Graph) SortPaths() {
	for i1 := 0; i1 < len(Paths); i1++ {
		for i2 := 0; i2 < len(Paths); i2++ {
			if len(Paths[i1]) < len(Paths[i2]) {
				Paths[i1], Paths[i2] = Paths[i2], Paths[i1]
			}
		}
	}
	// fmt.Println("sorted paths:", paths)
}
