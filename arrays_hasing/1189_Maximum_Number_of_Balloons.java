import java.util.HashMap;
import java.util.Map;

class Solution {
    public int maxNumberOfBalloons(String text) {
        Map<Character, Integer> store = new HashMap<>();
        for(char c : text.toCharArray()) {
            if(store.containsKey(c)) {
                store.put(c, store.get(c) + 1);
            } else {
                store.put(c, 1);
            }
        }
        if(store.containsKey('b') && store.containsKey('a') && store.containsKey('l') && store.containsKey('o') && store.containsKey('n')) {
            int b = store.get('b');
            int a = store.get('a');
            int l = store.get('l') / 2;
            int o = store.get('o') / 2;
            int n = store.get('n');
            return Math.min(b, Math.min(a, Math.min(l, Math.min(o, n))));
        } else return 0;
    }
}
