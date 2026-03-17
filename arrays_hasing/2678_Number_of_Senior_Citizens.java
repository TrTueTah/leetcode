class Solution {
    public int countSeniors(String[] details) {
        int res = 0;
        for(String citizen : details) {
            if(Integer.parseInt(citizen.substring(11, 13)) > 60) res++;
        }
        return res;
    }
}
