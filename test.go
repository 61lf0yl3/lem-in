package main

import "fmt"

func main() {
	sliice1 := []string{"1", "3", "4", "0"}
	sliice2 := []string{"1", "3", "4", "2", "5", "6", "0"}
	sliice3 := []string{"2", "5", "6", "7"}

	var slicegrop1 = [][]string{}
	slicegrop1 = append(slicegrop1, sliice1)
	slicegrop1 = append(slicegrop1, sliice2)
	slicegrop1 = append(slicegrop1, sliice3)

	var group = [][][]string{}

	for i1, path1 := range slicegrop1 {
		var slicegrop2 = [][]string{}
		slicegrop2 = append(slicegrop2, path1)
		for i2 := i1 + 1; i2 < len(slicegrop1); i2++ {
			path2 := slicegrop1[i2]
			flag := true
			for _, ps := range slicegrop2 {
				for _, room1 := range ps {
					for _, room2 := range path2 {
						if room1 == room2 {
							flag = false
						}
					}
				}
			}
			if flag == true {
				slicegrop2 = append(slicegrop2, path2)
			}
		}
		group = append(group, slicegrop2)
	}

	fmt.Println("this is slicegroup2:", group)
}
