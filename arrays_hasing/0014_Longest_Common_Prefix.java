class Solution {
    public String longestCommonPrefix(String[] strs) {
        String res = strs[0];
        for(int i = 1; i < strs.length; i++) {
            int j = 0;
            int len = Math.min(res.length(), strs[i].length());
            while(j < len) {
                if(res.charAt(j) != strs[i].charAt(j)) {
                    break;
                }
                j++;
            }
            res = res.substring(0, j);
        }
        return res;
    }
}
