package array

// 最长公共前缀 https://leetcode-cn.com/problems/longest-common-prefix/description/
func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	res := make([]byte, 0)
	minL := len(strs[0])
	for _, str := range strs {
		l := len(str)
		if l < minL {
			minL = l
		}
	}

	for i := 0; i < minL; i++ {
		tstr := strs[0][i]
		for _, str := range strs {
			if str[i] != tstr {
				goto END // 一定不再相同，得到结果
			}
		}
		res = append(res, tstr)
	}

END:
	return string(res)
}
