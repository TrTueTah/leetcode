class Solution {
    public int maxProfit(int[] prices) {
        int max = 0;
        int minPrice = prices[0];
        for(int price : prices) {
            max = Math.max(price - minPrice, max);
            minPrice = Math.min(price, minPrice);
        }
        return max;
    }
}
