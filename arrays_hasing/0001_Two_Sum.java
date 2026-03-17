import java.util.HashMap;
import java.util.Map;

class Solution {
    public int[] twoSum(int[] nums, int target) {
        Map<Integer, Integer> store = new HashMap<>();
        for(int i = 0; i < nums.length; i++) {
            if(store.containsKey(target - nums[i]) && store.get(target - nums[i]) != i) {
                return new int[]{i, store.get(target - nums[i])};
            }
            store.put(nums[i], i);
        }
        return new int[]{};
    }
}
