package tools

import "fmt"

var visited = make(map[string]bool)

// Path contain room in one path
var Path = []string{}

// Paths slice of slice of string, contain all paths
var Paths = [][]string{}

// PathsGroups slie of slice of slice of string, contain all disjoins path
var PathsGroups = [][][]string{}

// CleanedPathsGroups is disjointed paths with uniq flow, hops
var CleanedPathsGroups = [][][]string{}

//why *Lemin not just Lemin in method of func

// DFS implements depth-first traversal algorithm
func (g *Lemin) DFS(room string) {
	if visited[room] {
		return
	}
	visited[room] = true
	fmt.Print(room, " --> ")
	for _, v := range g.Edges[room] {
		g.DFS(v)
	}
}

// GetAllPaths implements DFS algorithm to find paths from start to end rooms
func (g *Lemin) GetAllPaths(room string) {
	if visited[room] {
		return
	}
	visited[room] = true
	Path = append(Path, room)
	if room != g.End {
		for _, neighbor := range g.Edges[room] {
			g.GetAllPaths(neighbor)
		}
	} else {
		var copy = make([]string, len(Path)) //need to copy because can't change origin path
		for index, room := range Path {
			copy[index] = room
		}
		Paths = append(Paths, copy)
	}
	Path = Path[:len(Path)-1] // delete end room for finding others path
	visited[room] = false     // make end room invisited for including end room for other path
}

// GetDisjointPaths get all disjointed (not intersenct) path from Paths
func (g *Lemin) GetDisjointPaths() {
	RmStart()
	for i1, path1 := range Paths {
		var group = [][]string{} // group is combination of disjoined paths
		group = append(group, path1)
		PathsGroups = append(PathsGroups, group)
		for i2 := i1 + 1; i2 < len(Paths); i2++ {
			path2 := Paths[i2]
			flag := true
			for _, temp2 := range group { // temp2 is path in the paths of disjoined paths
				for _, room1 := range temp2 {
					for _, room2 := range path2 {
						if room1 == room2 && room2 != g.End {
							flag = false
						}
					}
				}
			}
			if flag == true {
				group = append(group, path2)
			}
		}
		PathsGroups = append(PathsGroups, group)

	}
}

//ShortenPathsGroups if we have the same path with same flow and same hops we leave  uniq
// func ShortenPathsGroups() {
// 	for i1, flow1 := range PathsGroups {
// 		flag1 := true
// 		for i2 := i1 + 1; i2 < len(PathsGroups); i2++ {
// 			flow2 := PathsGroups[i2]
// 			if len(flow1) == len(flow2) { //get sorted PathsGroups according to number of rooms
// 				for i3 := 0; i3 < len(flow1); i3++ {
// 					if len(flow1) != 1 && len(flow1[i3]) == len(flow2[i3]) {
// 						flag1 = false
// 					}
// 					if len(flow1) == 1 && len(flow1[i3]) >= len(flow2[i3]) {
// 						flag1 = false
// 					}
// 				}
// 			}
// 		}
// 		if flag1 == true {
// 			CleanedPathsGroups = append(CleanedPathsGroups, flow1)
// 		}
// 	}
// }

// OptimalPath choose wich is optimal path from CleanedPathGruops
func (g Lemin) OptimalPath() [][]string {
	var heights = make(map[int][][]string)
	for _, flow := range PathsGroups {
		if len(flow) == 1 && len(flow[0]) == 2 { // if start and end connected
			//path := flow[0]
			// if path[0] == g.Start && path[1] == g.End {
			// 	fmt.Println(flow)
			return flow
			//}
		}
		maxhops := MaxHops(flow) // the number of hops in the combination of path
		fillants := FillAnts(flow)
		if fillants == NbAnts {
			continue
		}
		nbpaths := len(flow)
		height := maxhops + (NbAnts-fillants)/nbpaths
		if (NbAnts-fillants)%nbpaths != 0 {
			height++
		}
		heights[height] = flow
	}
	//fmt.Println("heights:", heights, "\nlen of heights:", len(heights), "\n")
	optimalflow := GetMinHeightValue(heights)
	return optimalflow
}

// SendAnts2 prints ants according to their paths
func (g *Lemin) SendAnts2(group [][]string, n int) {
	levels := n / len(group)
	if n%len(group) != 0 {
		levels++
	}

	var ants = make([]Ant, n+1)
	// ignore zero because ants start from 1
	ants[0].Ignore = true
	id := 0

	for j := 0; j < levels; j++ {
		for _, path := range group {
			id++
			ants[id].Path = path
			ants[id].RoomID = 0
			ants[id].Ignore = false
			fmt.Println("ant", ants[id])
			if id == n {
				break
			}
		}
	}

	fmt.Println("ants:", ants)
	fmt.Println()

	exit := false
	var taken = make(map[string]bool)
	for !exit {
		for id, ant := range ants {
			if ant.Ignore {
				continue
			}
			room := ant.Path[ant.RoomID]
			if taken[room] {
				fmt.Println()
				break
			}
			fmt.Print("L", id, "-", room, " ")
			if id == n {
				fmt.Println()
				if room == g.End {
					exit = true
				}
			}
			ants[id].RoomID++
			taken[ants[id].Previous] = false
			if room != g.End {
				taken[room] = true
				ants[id].Previous = room
			}
			if room == g.End {
				ants[id].Ignore = true
			}
		}
	}
}

// SendAnts prints ants according to their paths
func (g *Lemin) SendAnts(optimal [][]string) {
	// level := NbAnts / len(optimal)
	// if NbAnts%len(optimal) != 0 {
	// 	level++
	// }
	//fmt.Println("level:", len(optimal))
	for i1 := 0; i1 < len(optimal); i1++ {
		for i2 := 1; i2 <= NbAnts; i2++ {
			for _, path := range optimal {

			}
		}
	}
}
