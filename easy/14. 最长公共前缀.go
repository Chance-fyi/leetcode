package main

import "fmt"

//编写一个函数来查找字符串数组中的最长公共前缀。
//
//如果不存在公共前缀，返回空字符串 ""。
//
//
//
//示例 1：
//
//输入：strs = ["flower","flow","flight"]
//输出："fl"
//示例 2：
//
//输入：strs = ["dog","racecar","car"]
//输出：""
//解释：输入不存在公共前缀。
//
//
//提示：
//
//1 <= strs.length <= 200
//0 <= strs[i].length <= 200
//strs[i] 仅由小写英文字母组成
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/longest-common-prefix
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

func main() {
	s := ""
	s = longestCommonPrefix([]string{""})
	fmt.Println(s, s == "")
	s = longestCommonPrefix([]string{"a"})
	fmt.Println(s, s == "a")
	s = longestCommonPrefix([]string{"flower", "flow", "flight"})
	fmt.Println(s, s == "fl")
	s = longestCommonPrefix([]string{"dog", "racecar", "car"})
	fmt.Println(s, s == "")
	s = longestCommonPrefix([]string{"flower", "flow", "flight", ""})
	fmt.Println(s, s == "")
	s = longestCommonPrefix([]string{"a", "ab"})
	fmt.Println(s, s == "a")
}

func longestCommonPrefix(strs []string) string {
	if len(strs) == 1 {
		return strs[0]
	}
	s, i := "", 1
	for {
		for _, str := range strs {
			if len(str) < i {
				return s[0 : i-1]
			}
			if len(s) != i {
				s = str[0:i]
			}
			if s != str[0:i] {
				return s[0 : i-1]
			}
		}
		i++
	}
}
