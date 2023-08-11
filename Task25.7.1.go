package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	n := bufio.NewReader(os.Stdin)
	fmt.Print("Введите данные: ")
	data, _ := n.ReadString('\n')
	fmt.Printf("Вы ввели следующие данные: %v\n", data)
}
