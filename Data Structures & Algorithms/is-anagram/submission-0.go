func isAnagram(s string, t string) bool {
    if len(s) != len(t) {
        return false
    }

    alpha := [26]int{}
    for i := 0; i < len(s); i++ {
        alpha[s[i] - 'a']++
        alpha[t[i] - 'a']--
    }

    for _, sym := range alpha {
        if sym != 0 {
            return false 
        }
    }

    return true 
}
