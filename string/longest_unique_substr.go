package string

func lengthOfLongestSubstring(s string) int {
    var m = make(map[rune]int)
    var longestLength int
    var head int
    var v rune
    var tail int = -1

    for head, v = range s {
        if idx, ok := m[v]; ok {
            if idx > tail {
                tail = idx
            } else {
                idx = tail
            }

            if head - idx > longestLength {
                longestLength = head - idx
            }
            m[v] = head
        } else {
            if head - tail > longestLength {
                longestLength = head - tail
            }
        }
        m[v] = head
    }

    return longestLength
}
