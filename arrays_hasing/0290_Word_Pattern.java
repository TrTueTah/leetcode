import java.util.HashMap;
import java.util.Map;

class Solution {
    public boolean wordPattern(String pattern, String s) {
        String[] words = s.split(" ");
        if(words.length != pattern.length()) return false;
        Map<Character, String> store1 = new HashMap<>();
        Map<String, Character> store2 = new HashMap<>();
        for(int i = 0; i < pattern.length(); i++) {
            if(store1.containsKey(pattern.charAt(i)) && !store1.get(pattern.charAt(i)).equals(words[i]) ||
        store2.containsKey(words[i]) && !store2.get(words[i]).equals(pattern.charAt(i)) ) {
                return false;
            } else {
                store1.put(pattern.charAt(i), words[i]);
                store2.put(words[i], pattern.charAt(i));
            }
        }
        return true;
    }
}
