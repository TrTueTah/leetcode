import java.util.HashMap;
import java.util.Map;

class Solution {
    public int findLucky(int[] arr) {
        Map<Integer, Integer> store = new HashMap<>();
        int res = -1;
        for(int num : arr) {
            if(store.containsKey(num)) {
                store.put(num, store.get(num) + 1);
            } else {
                store.put(num, 1);
            }
        }
        for(int num : arr) {
            if(num == store.get(num)) {
                if(num > res) {
                    res = num;
                }
            }
        }
        return res;
    }
}
