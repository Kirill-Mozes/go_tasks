package main

import (
	"fmt"
	"sort"
)

func removeDuplicateElement(addrs []string) []string {
	result := make([]string, 0, len(addrs))
	temp := map[string]struct{}{}
	for _, item := range addrs {
		if _, ok := temp[item]; !ok { //если такого элемента в карте еще нет
			temp[item] = struct{}{}
			result = append(result, item)
		}
	}
	return result
}

func main() {
	var arr = []string{"cat", "cat", "dog", "cat", "tree"}
	arr = removeDuplicateElement(arr)
	sort.Strings(arr)
	fmt.Println(arr)

}
