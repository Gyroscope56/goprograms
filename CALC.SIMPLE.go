package main

import (
	"bufio"
	"fmt"
	"os"
)

func readFromKeyboard() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		str := scanner.Text()
		fmt.Println(str)
	}

	// intgr := 0
	// fmt.Scanf("%d", &intgr)
	// op := ' '
	// fmt.Scanf("%c", &op)
	// fmt.Scanf("\n")
	// fmt.Println(intgr, op)
}
