package main

import "fmt"

type data struct {
	a int
	b int
}

func nilnil(a, b int) (*data, error) {
	if a == 0 && b == 0 {
		return nil, nil
	}
	return &data{a: a, b: b}, nil
}

func main() {
	const a = 0
	const b = 0
	res, err := nilnil(a, b)
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}
	fmt.Printf("res: %v\n", res)
}
