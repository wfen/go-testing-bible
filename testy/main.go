package testy

import "fmt"

// Developer - models our developer
type Developer struct {
	Name string
	Age  int
}

// FilterUnique - filters out unique developers
func FilterUnique(developers []Developer) []string {
	var uniques []string
	check := make(map[string]int)
	for _, developer := range developers {
		check[developer.Name] = 1
	}

	for name := range check {
		uniques = append(uniques, name)
	}
	return uniques
}

func main() {
	fmt.Println("Getting Started with Testify")
}
