package main

import (
	"bufio"
	"fmt"
	"math/big"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	// 	a, _ := new(big.Int).SetString("14236789613472891364972084768919", 10)
	line1, _ := reader.ReadString('\n')
	line1 = strings.TrimSpace(line1)
	line2, _ := reader.ReadString('\n')
	line2 = strings.TrimSpace(line2)
	a, ok1 := new(big.Int).SetString(line1, 10)
	b, ok2 := new(big.Int).SetString(line2, 10)
	if !ok1 || !ok2 {
		fmt.Println("Incorrect input")
		return
	}
	sum := new(big.Int).Add(a, b)
	diff := new(big.Int).Sub(a, b)
	prod := new(big.Int).Mul(a, b)
	quot := new(big.Int)
	rem := new(big.Int)
	if b.Sign() != 0 {
		quot.Div(a, b)
		rem.Mod(a, b)
	}
	fmt.Println("Sum:      ", sum)
	fmt.Println("diff:   ", diff)
	fmt.Println("prod:", prod)
	if b.Sign() != 0 {
		fmt.Println("quot:    ", quot)
		fmt.Println("rem:    ", rem)
	} else {
		fmt.Println("division by zero")
	}
}
