class Solution {
    public String mergeAlternately(String word1, String word2) {
        int i = 0, j = 0;
        StringBuilder builder = new StringBuilder();
        while(i < word1.length() && j < word2.length()) {
            builder.append(word1.charAt(i));
            builder.append(word2.charAt(j));
            i++;
            j++;
            if(i == word1.length()) {
                while(j < word2.length()) {
                    builder.append(word2.charAt(j));
                    j++;
                }
            } else if(j == word2.length()) {
                while(i < word1.length()) {
                    builder.append(word1.charAt(i));
                    i++;
                }
            }
        }
        return builder.toString();
    }
}
