package foo

import "fmt"

// before
func makeSlice(n int) []string {
	var r []string
	for i := 0; i < n; i++ {
		r = append(r, fmt.Sprintf("%03d..", i))
	}
	return r
}

// after
//func makeSlice(n int) []string {
//	r := make([]string, n)
//	for i := 0; i < n; i++ {
//		r[i] = fmt.Sprintf("%03d だよーん", i)
//	}
//	return r
//}
