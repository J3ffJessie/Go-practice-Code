package main

//The power of slices
import (
	"fmt"
	"strings"
)
//hyperspace removes the space surrounding worlds
func hyperspace(worlds []string) {
	for i := range worlds {
		worlds[i] = strings.TrimSpace(worlds[i])
	}
}
//passing our planets array into the hyperspace function will remove all the spaces surrounding our planet strings
func main() {
	planets := []string{" Venus   ", "Earth  ", " Mars"}
	hyperspace(planets)

	fmt.Println(strings.Join(planets, ""))
}
