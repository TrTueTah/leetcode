class Solution {
    public int maxAscendingSum(int[] nums) {
        int res = nums[0];
        int max = nums[0];
        for(int i = 1; i < nums.length; i++) {
            if(nums[i] > nums[i - 1]) {
                max += nums[i];
            } else {
                max = nums[i];
            }
            if(max > res) res = max;
        }
        return res;
    }
}
