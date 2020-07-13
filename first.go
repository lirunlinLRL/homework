package main

import (
	"fmt"
)

func main() {
	mymap := make(map[string]int)
	mymap["zhanshan"] = 23
	mymap["lisi"] = 19
	fmt.Println(mymap)

	for k, v := range mymap {
		fmt.Println(k, v)
	}
}
