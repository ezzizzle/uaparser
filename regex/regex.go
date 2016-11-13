package regex

// MapRegexNames takes an array of regexp matches and
// an array of named matches and combines them into a
// map
func MapRegexNames(m, n []string) map[string]string {
    m, n = m[1:], n[1:]
    r := make(map[string]string, len(m))
    for i := range n {
        r[n[i]] = m[i]
    }
    return r
}
