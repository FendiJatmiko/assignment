package main

import (
	"fmt"
	"io/ioutil"
	"sort"
)

type uint64arr []byte

func (a uint64arr) Len() int           { return len(a) }
func (a uint64arr) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a uint64arr) Less(i, j int) bool { return a[i] < a[j] }
func (a uint64arr) String() (s string) {
	sep := ""
	for _, el := range a {
		s += sep
		sep = "\n"
		s += fmt.Sprintf("%d", el)
	}
	return
}

func main() {
	file, err := ioutil.ReadFile("age.txt")
	if err != nil {
		return
	}

	arr := uint64arr(file)
	sort.Sort(arr)
	fmt.Printf("%s\n", arr)
}
