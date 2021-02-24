package main

import (
	"fmt"
	"sort"
)

type people []string

// people has to implement all the methods (method set) of the sort interface
// called the Interface interface
func (p people) Len() int           { return len(p) }
func (p people) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func (p people) Less(i, j int) bool { return p[i] < p[j] }

func main() {
	studyGroup := people{"Zeno", "John", "Al", "Jenny"}

	fmt.Println(studyGroup)
	sort.Sort(studyGroup)
	fmt.Println(studyGroup)

	// Sorting a sclice
	s := []string{"Louis", "Rimo", "Brendon", "Dave", "Declan", "Tompy", "Wisdom"}
	fmt.Println(s)
	// StringSlice implement the Interface interface
	sort.Sort(sort.StringSlice(s))
	fmt.Println(s)
}
