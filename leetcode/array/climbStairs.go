func climbStairs(n int) int {
	if n <0{
		return 0
	}
	if n>0 && n<=2{
		return n
	}

	a :=1
	b :=2
	res :=0
	for i := 3; i <=n ; i++ {
		res = a+b
		a = b
		b = res
	}
	
	return res

}
