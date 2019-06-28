 //Longest Substring Without Repeating Characters
// https://leetcode.com/problems/longest-substring-without-repeating-characters/
 func lengthOfLongestSubstring(s string) int {
    pos := map[rune]int{}
    ans := 0
    begin := 0
    for end, r := range s {
        if p := pos[r]; p > begin {
            begin = p
        }
        pos[r] = end + 1
        if cur := end - begin + 1; cur > ans {
            ans = cur
        }
    }
    return ans
}

// demo on https://play.golang.org/p/Jku3ezmr8WW