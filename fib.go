package main

//func main() {
//	var store = map[int]int{}
//	t := time.Now()
//	r := fib(500, store)
//	println(r, time.Now().Sub(t).String())
//}

func fib(n int, store map[int]int) int {
	if n == 0 { return 0 }
	if n == 1 { return 1 }

	n1, ok := store[n-1]
	if !ok {
		n1 = fib(n-1, store)
		store[n-1] = n1
	}

	n2, ok := store[n-2]
	if !ok {
		n2 = fib(n-2, store)
		store[n-2] = n2
	}

	return n1 + n2
}
