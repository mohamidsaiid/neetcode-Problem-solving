// https://leetcode.com/problems/two-sum/
class Solution {
public:
    vector<int> twoSum(vector<int>& nums, int target) {
        unordered_map<int, int> x;
        int size = nums.size();
        for (int i = 0; i < size; i++)
        {
            auto t = x.find(nums[i]);
            if(t != x.end())
            {
                return {i, t->second};
            }
            else x[target - nums[i]] = i;
        }
        return {0, 0};
    }
};
