package helper

// This function needs to start with a
// Capital letter to be visible to `client.go`
func Add(a, b int) int {
	return a + b
}

// This function cannot be accessed from outside the helper package
func add(a, b int) int {
	return a + b
}

func Colatz(a int) int {
	if a%2 == 0 {
		return a / 2
	}
	return 3*a + 1
}
