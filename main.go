package main

import (
	"fmt"
	"os"
	"sort"
)

// create a struct to store kv for calculating the mode
type occur struct {
	Key   int
	Count int
}

func main() {
	// Using a slice because in the future the array input
	// could be more dynamic - also makes this a bit more challenging
	list := []int{1, 3, 4, 15, 18, 3, 2, 12, 18, 22, 18}

	if len(list) < 1 {
		fmt.Println("Input slice is empty")
		os.Exit(1)
	}

	// sort the array
	sort.Ints(list)

	// get the median index
	med_ind := getMedian(list)

	// assure the index is not out of range before using it as the result
	if med_ind <= len(list) {
		fmt.Printf("The median is %d\n", list[med_ind])
	} else {
		fmt.Println("Something went wrong! index is out of range!")
	}

	// Get the mode as a struct
	modestruct := getMode(list)

	// let the user know the mode
	fmt.Printf("The mode is: %d and the occurences are: %d\n", modestruct.Key, modestruct.Count)
}

func getMedian(list []int) int {
	// if the list has a length of 1
	// return an index of 0 since that is the only value
	if len(list) == 1 {
		return 0
	} else {
		return len(list) / 2
	}
}

func getMode(list []int) occur {

	// create a new map to calculate occurrences
	count := make(map[int]int)

	// populate the map and count occurences
	for _, val := range list {
		count[val]++
	}

	// create a slice for occur values
	var modeslice []occur
	// populate vector with
	for k, v := range count {
		modeslice = append(modeslice, occur{k, v})
	}

	// Sort the slice using an anonymous function
	// in descending order
	sort.Slice(modeslice, func(i, j int) bool {
		return modeslice[i].Count > modeslice[j].Count
	})

	// return the first index
	return modeslice[0]
}
