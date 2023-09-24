package string

func strStr(haystack string, needle string) int {
  var j = 0
  var i = 0

  if len(needle) > len(haystack) {
    return -1
  }

  for i < len(haystack) {
    for j < len(needle) && i + j < len(haystack) {
      if haystack[i+j] != needle[j] {
        j = 0
        break
      }
      if j == len(needle) - 1 {
        return i
      }
      j++
    }
    i++
  }

  return -1
}
