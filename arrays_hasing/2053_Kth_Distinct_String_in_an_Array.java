import java.util.HashMap;
import java.util.Map;

class Solution {
    public String kthDistinct(String[] arr, int k) {
        Map<String, Integer> store = new HashMap<>();
        for(String s : arr) {
            if(!store.containsKey(s)) {
                store.put(s, 1);
            } else {
                store.put(s, store.get(s) + 1);
            }
        }
        for(String s : arr) {
            if(store.get(s) == 1) {
                k--;
                if(k == 0) return s;
            }
        }
        return "";
    }
}
