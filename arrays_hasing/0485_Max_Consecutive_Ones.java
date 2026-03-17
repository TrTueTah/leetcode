class Solution {
    public int findMaxConsecutiveOnes(int[] nums) {
        int res = 0, curr = 0;
        int i = 0;
        while(i < nums.length) {
            if(nums[i] == 1) {
                curr++;
            } else {
                if(curr > res) res = curr;
                curr = 0;
            }
            i++;
        }
        return curr > res ? curr : res;
    }
}
