class Solution {
    public int[] replaceElements(int[] arr) {
        // Method 1: Brute Force
        // if(arr.length == 1) return new int[]{-1};
        // for(int i = 0; i < arr.length - 1; i++) {
        //     int max = arr[i + 1];
        //     for(int j = i + 1; j < arr.length; j++) {
        //         if(max < arr[j]) {
        //             max = arr[j];
        //         }
        //     }
        //     arr[i] = max;
        // }
        // arr[arr.length - 1] = -1;
        // return arr;

        // Method 2: Right to left travel
        int max = -1;
        for(int i = arr.length - 1; i >= 0; i--) {
            if(arr[i] > max) {
                int temp = max;
                max = arr[i];
                arr[i] = temp;
            } else {
                arr[i] = max;
            }
        }
        return arr;
    }
}
