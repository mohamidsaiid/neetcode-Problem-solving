// https://leetcode.com/problems/valid-anagram/description/

func isAnagram(s string, t string) bool {
    ss := strings.Split(s, "")
    tt := strings.Split(t, "")
    sort.Strings(ss)
    sort.Strings(tt)
    return strings.Join(ss, "") == strings.Join(tt, "")
}


func isAnagram(s string, t string) bool {
    // count sort
    n := len(s)
    if n != len(t){
        return false
    }


    count := make([]int, 26)
    for i := 0; i < n; i++{
        count[s[i] - 'a']++
        count[t[i] - 'a' ]--
    }
    for _, val := range count{
        if val != 0{
            return false
        }
    }
    return true
}
