import java.util.Arrays;

class Solution {
    public int minimumDifference(int[] nums, int k) {
        Arrays.sort(nums);
        int res = 100001;
        for(int i = 0; i <= nums.length - k; i++) {
            res = Math.min(nums[i + k - 1] - nums[i], res);
        }
        return res;
    }
}
