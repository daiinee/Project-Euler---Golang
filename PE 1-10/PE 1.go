package main

func description() {
	print("If we list all the natural numbers below 10 that are multiples of 3 or 5, we get 3, 5, 6 and 9. The sum of these multiples is 23. \nFind the sum of all the multiples of 3 or 5 below 1000.\n\n")
}

func analysis() {
	print("This problem can be solved using an arithmetic progression with a common diffetence to find the sum of all the numbers multiples of 5 and 3, \nbut there's a catch, we must substract those numbers multiples of 15 as they're factors of both 3 and 5, and the will be counted twice.")
}

func main() {
	const number int = 100
	const a int = 3
	const b int = 5
	const c int = 15
	ans := sumSeries(a, 999) + sumSeries(b, 999) - sumSeries(c, 999)
	print(ans)
}

func sumSeries(x, y int) int {
	return x
	/*I need to do the series here.*/
}
