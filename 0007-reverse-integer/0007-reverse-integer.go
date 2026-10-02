func reverse(x int) int {
    rev := 0
    for x != 0 {
        remainder := x % 10
        rev = (rev * 10) + remainder
        x = x/10
    }

    if rev < math.MinInt32 || rev > math.MaxInt32 {
		return 0
	}
    return rev
}