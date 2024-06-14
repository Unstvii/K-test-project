package main

type Request struct {
	A *int `json:"a"`
	B *int `json:"b"`
}

type Response struct {
	FactorialA string `json:"factorial_a"`
	FactorialB string `json:"factorial_b"`
}
