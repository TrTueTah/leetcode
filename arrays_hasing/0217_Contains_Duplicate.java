class Solution {
    public boolean hasDuplicate(int[] nums) {
        // Method 1
        // for(int i = 0; i < nums.length; i++) {
        //     for(int j = i + 1; j < nums.length; j++) {
        //         if(nums[i] == nums[j]) {
        //             return true;
        //         }
        //     }
        // }
        // return false;

        // Method 2: Using Hash
        Set<Integer> store = new HashSet<>();
        for(int i = 0; i < nums.length; i++) {
            if(store.contains(nums[i])) {
                return true;
            } else {
                store.add(nums[i]);
            }
        }
        return false;
    }
}
