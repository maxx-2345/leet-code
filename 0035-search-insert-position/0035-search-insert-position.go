func searchInsert(nums []int, target int) int {
    for k, v := range nums{
        // Check if number exist
        if v == target {
            return k
        }
        // If not exist then return potential place 
        if v > target {
            return k
        }
    }
    return len(nums)
}