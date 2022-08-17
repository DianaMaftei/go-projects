package main

import (
	"fmt"
	"sort"
)

type byLength []string

func main() {

	names := []string{"Dorian", "Alexander", "Jayce", "Ian", "Zoey"}
	sort.Strings(names) //natural order
	fmt.Println(names)

	namesByLength := byLength(names)

	sort.Sort(namesByLength) //custom sort
	fmt.Println(names)

	ages := []int{58, 24, 15, 42, 8}
	sort.Ints(ages)
	fmt.Println(ages)

	sorted := sort.IntsAreSorted(ages)
	fmt.Println("Are ages sorted?", sorted)

}

func (s byLength) Len() int {
	return len(s)
}

func (s byLength) Less(i, j int) bool {
	return len(s[i]) < len(s[j])
}

func (s byLength) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
