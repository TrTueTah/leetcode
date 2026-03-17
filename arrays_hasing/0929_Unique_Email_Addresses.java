import java.util.HashSet;
import java.util.Set;

class Solution {
    public int numUniqueEmails(String[] emails) {
        Set<String> store = new HashSet<>();
        for(String email : emails) {
            String parts[] = email.split("@");
            String local = parts[0];
            String domain = parts[1];

            local = local.split("\\+")[0];
            local = local.replace(".", "");
            store.add(local + "@" + domain);
        }
        return store.size();
    }
}
