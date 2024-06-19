package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
)

const (
	plus  = "+"
	minus = "-"
	mul   = "*"
	div   = "/"
)

func parseExpression(expression string) (op string, x, y float64, err error) {
	_, err = fmt.Sscanln(expression, &x, &op, &y)
	if err != nil {
		err = errors.New("invalid expression")
	}
	return
}

func computeExpression(op string, x, y float64) (res float64, err error) {
	switch op {
	case plus:
		res = x + y
	case minus:
		res = x - y
	case mul:
		res = x * y
	case div:
		res = x / y
	default:
		err = errors.New("unsupported operation")
	}
	return
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	input := ""
	for {
		fmt.Print("Enter your expression: ")
		input, _ = reader.ReadString('\n')

		if input == "q" {
			break
		}

		op, x, y, err := parseExpression(input)

		if err != nil {
			fmt.Println(err)
			continue
		}

		res, err := computeExpression(op, x, y)

		if err != nil {
			fmt.Println(err)
			continue
		}

		fmt.Println(res)
	}

	fmt.Println("Goodbye!")
}
