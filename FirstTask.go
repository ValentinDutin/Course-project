package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

type Pair struct{
	Key string
	Value int
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	str ,_ := reader.ReadString('\n')
	str = strings.Trim(str, "\n \t")
	dictMap := make(map[string]int)
	for _, symbol := range str {
		_, exist := dictMap[string(symbol)]
		if !exist {
			dictMap[string(symbol)] = 1;
		} else {
			dictMap[string(symbol)] += 1;
		}
	}
	fmt.Println(dictMap)
	sortedArray := make([]Pair, len(dictMap))
	for key, value := range dictMap {
		sortedArray = append(sortedArray, Pair{key, value})
	}
	sort.Slice(sortedArray, func(i int, j int) bool{
		return sortedArray[i].Value > sortedArray[j].Value
	})
	for _, item := range sortedArray {
		for i := 0; i < item.Value; i++ {
			fmt.Print(item.Key)
		}
	}
}