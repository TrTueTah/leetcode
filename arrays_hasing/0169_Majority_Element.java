import java.util.Arrays;

class Solution {
    public int majorityElement(int[] nums) {
        // Method 1: Sorting
        // Arrays.sort(nums);
        // return nums[nums.length / 2];

        // Boyer-Mooer Voting Algorithm
        int count = 0;
        int res = 0;
        for(int num : nums) {
            if(count == 0) {
                res = num;
            }
            count += (num == res) ? 1 : -1;
        }
        return res;
    }
}
