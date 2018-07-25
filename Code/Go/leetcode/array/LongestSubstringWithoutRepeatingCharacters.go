package array

// 无重复字符的最长子串 https://leetcode-cn.com/problems/longest-substring-without-repeating-characters/description/
func lengthOfLongestSubstring(s string) int {
	mp := make(map[byte]int)
	maxL := 0
	nowL := 0
	for i := 0; i < len(s); i++ {
		if _, ok := mp[s[i]]; !ok {
			//不重复，则++
			nowL++
			mp[s[i]] = i
			if nowL > maxL {
				maxL = nowL
			}
		} else {
			if nowL > maxL {
				maxL = nowL
			}
			i = mp[s[i]] + 1
			nowL = 1
			mp = make(map[byte]int)
			mp[s[i]] = i
		}
	}
	return maxL
}
