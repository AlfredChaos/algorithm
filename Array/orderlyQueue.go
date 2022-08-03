func orderlyQueue(s string, k int) string {
    if k > 1 {
        //return bubbleSort([]rune(s))
        return quick([]rune(s))
    }
    min_string := s
    for i:=0; i<len(s); i++ {
        s = s[1:] + s[:1]
        if s < min_string {
            min_string = s
        }
    }
    return min_string
}