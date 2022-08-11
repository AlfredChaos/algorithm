func reformat(s string) string {
    if len(s) <= 1 {
        return s
    }
    even := 0
    odd := 1
    lenLetter := 0
    lenNumber := 0
    result := make([]byte, len(s))
    for i:=0; i<len(s); i++ {
        if s[i] >= 'a' && s[i] <= 'z' {
            lenLetter++
        }
        if s[i] >= '0' && s[i] <= '9' {
            lenNumber++
        }
    }
    if lenLetter == len(s) || lenNumber == len(s) || lenLetter >= lenNumber + 2 || lenNumber >= lenLetter + 2 {
        return ""
    }
    for i:=0; i<len(s); i++ {
        if lenLetter > lenNumber {
            if s[i] >= 'a' && s[i] <= 'z' {
                result[even] = s[i]
                even += 2
            }
            if s[i] >= '0' && s[i] <= '9' {
                result[odd] = s[i]
                odd += 2
            }
        } else {
            if s[i] >= 'a' && s[i] <= 'z' {
                result[odd] = s[i]
                odd += 2
            }
            if s[i] >= '0' && s[i] <= '9' {
                result[even] = s[i]
                even += 2
            }
        }
    }
    return string(result)
}