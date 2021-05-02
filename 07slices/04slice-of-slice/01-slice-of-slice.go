package main

import "fmt"

func main() {
	sliceOfSlice := [][]string{
		{"Go", "Node"},
		{"Python", ".Net"},
	}

	fmt.Println(sliceOfSlice)

	sliceWtf()
}

func sliceWtf() {
	sliceWtf := [][][][]string{
		{
			{
				{
					"Hello",
				},
			},
		},
		{
			{
				{
					"World !",
				},
			},
		},
	}

	fmt.Println(sliceWtf[0][0][0][0], sliceWtf[1][0][0][0])
}
