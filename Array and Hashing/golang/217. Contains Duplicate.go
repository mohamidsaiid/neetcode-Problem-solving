// https://leetcode.com/problems/contains-duplicate/description/
func containsDuplicate(nums []int) bool {
    mp := make(map[int]bool, len(nums))
    for _, val := range nums{
        if mp[val] {
            return true
        }else {
            mp[val] = true
        }
    }
    return false
}
