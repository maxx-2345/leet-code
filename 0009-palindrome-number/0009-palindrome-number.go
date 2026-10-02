func isPalindrome(x int) bool {
reversed := 0
original :=x
if x < 0 || (x % 10 == 0 && x != 0){
    return false
} 
for x > 0{
    reversed = (reversed * 10) + (x % 10)
    x /= 10
}
if original == reversed {
    return true
}
return false
}