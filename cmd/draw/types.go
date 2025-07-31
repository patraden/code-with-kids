package main

import "fmt"

type Bucket = [9]string

type TeamPair struct {
	A string
	B string
}

func (p TeamPair) String() string {
	return fmt.Sprintf("%s - %s", p.A, p.B)
}
