class Solution {
    public int maxDifference(String s) {
        int[] count = new int[26];
        for(char c : s.toCharArray()) {
            count[c - 'a']++;
        }
        int oddMax = 0;
        int evenMin = s.length();
        for(int c : count) {
            if(c % 2 == 1) {
                oddMax = Math.max(oddMax, c);
            } else {
                evenMin = Math.min(evenMin, c);
            }
        }
        return oddMax - evenMin;
    }
}
