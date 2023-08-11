package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	n := bufio.NewReader(os.Stdin)
	fmt.Print("Введите целое число: ")
	data, _ := n.ReadString('\n')
	fmt.Printf("Вы ввели число: %v\n", data)
}
