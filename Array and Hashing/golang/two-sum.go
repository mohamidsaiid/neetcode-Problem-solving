// https://leetcode.com/problems/two-sum/
func twoSum(nums []int, target int) []int {
    res := make([]int, 0, 2)
    n := len(nums)
    mp := make(map[int]int, n)
    for idx, val := range nums{
        v, ok := mp[val]
        if ok{
            res = append(res, v, idx)
            break 
        }else{
            mp[target-val] = idx
        }
    }
    return res
}
