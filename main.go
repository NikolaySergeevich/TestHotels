package main

import (
	"fmt"
	"strconv"
	"strings"
	"sync"
)

func main() {
	printMultiplicationTable(5)
}

func printMultiplicationTable(n int) {
	fmt.Printf("Таблица умножения до %d:\n", n)

	var lineNum strings.Builder
	var overLine strings.Builder
	var res strings.Builder

	wg := sync.WaitGroup{}
	wg.Add(1)
	go func(n int){
		defer wg.Done()
		lineNum.WriteString("   ")
		for i := 1; i <= n; i++ {
			lineNum.WriteString(strconv.Itoa(i))
			lineNum.WriteString("  \t")
		}
	}(n)
	wg.Add(1)
	go func(n int){
		defer wg.Done()
		overLine.WriteString("   ")
		for i := 1; i <= n; i++ {
			overLine.WriteString("_")
			overLine.WriteString("   \t")
		}
	}(n)

	wg.Add(1)
	go func(n int){
		defer wg.Done()
		for i := 1; i <= n; i++ {
			res.WriteString(fmt.Sprintf("%d |", i))
			for j := 1; j <= n; j++ {
				res.WriteString(fmt.Sprintf("%d\t", i*j))
			}
			res.WriteString("\n")
		}
	}(n)

	wg.Wait()
	fmt.Println(lineNum.String())
	fmt.Println(overLine.String())
	fmt.Println(res.String())
}