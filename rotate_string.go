package main
import (
	"fmt"
)
func rotateLeft(s string, k int) string {
	n := len(s)
	if n == 0 {
		return s
	}
	k = k % n
	if k == 0 {
		return s
	}
	result := s[k:] + s[:k]
	return result
}

func main() {
	var s string
	var a string
	var k int
	fmt.Print("Enter a string : ")
	fmt.Scanln(&s)
	fmt.Print("Enter the number of positions to rotate: ")
	fmt.Scanln(&k)
	a = s
	rotatedstring := rotateLeft(s, k)
	fmt.Printf("String before rotation: %s\n", a)
	fmt.Printf("String after rotation: %s\n", rotatedstring)
}
