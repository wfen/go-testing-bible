package linting

import "fmt"

func checkFlag(flag bool) bool { //nolint:deadcode,unused
	// if flag == true {
	//     return true
	// } else {
	//     return false
	return flag
}

func errChecking() (int, error) { //nolint:unused
	// a := 1
	// a = 2
	a := 2
	fmt.Println(a)
	return 0, nil
}

func callsErrChecking() int { //nolint:deadcode,unused
	val, _ := errChecking()
	return val
}
