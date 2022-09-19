package main

func fibRec(n int) int {
	if n == 1 {
		return 1
	}
	if n == 0 {
		return 0
	}
	return fibRec(n-2) + fibRec(n-1)
}

func fib(n int) int {
	l, r := 0, 1
	for ; n >= 1; n-- {
		l, r = r, l+r
	}
	return l
}
