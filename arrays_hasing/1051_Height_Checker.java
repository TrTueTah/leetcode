import java.util.Arrays;

class Solution {
    public int heightChecker(int[] heights) {
        int[] sorted = heights.clone();
        int count = 0;
        Arrays.sort(sorted);
        for(int i = 0; i < heights.length; i++) {
            if(sorted[i] != heights[i]) {
                count++;
            }
        }
        return count;
    }
}
