package main

import "fmt"

//实现 strStr() 函数。
//
//给你两个字符串 haystack 和 needle ，请你在 haystack 字符串中找出 needle 字符串出现的第一个位置（下标从 0 开始）。如果不存在，则返回  -1 。
//
//
//
//说明：
//
//当 needle 是空字符串时，我们应当返回什么值呢？这是一个在面试中很好的问题。
//
//对于本题而言，当 needle 是空字符串时我们应当返回 0 。这与 C 语言的 strstr() 以及 Java 的 indexOf() 定义相符。
//
//
//
//示例 1：
//
//输入：haystack = "hello", needle = "ll"
//输出：2
//示例 2：
//
//输入：haystack = "aaaaa", needle = "bba"
//输出：-1
//示例 3：
//
//输入：haystack = "", needle = ""
//输出：0
//
//
//提示：
//
//0 <= haystack.length, needle.length <= 5 * 104
//haystack 和 needle 仅由小写英文字符组成
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/implement-strstr
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

func main() {
	data := []struct {
		Haystack string
		Needle   string
		Index    int
	}{
		{Haystack: "hello", Needle: "ll", Index: 2},
		{Haystack: "aaaaa", Needle: "bba", Index: -1},
		{Haystack: "aaaaa", Needle: "abba", Index: -1},
		{Haystack: "", Needle: "", Index: 0},
		{Haystack: "a", Needle: "a", Index: 0},
		{Haystack: "", Needle: "s", Index: -1},
	}

	for _, v := range data {
		i := strStr(v.Haystack, v.Needle)
		fmt.Println(i, i == v.Index)
	}
}

func strStr(haystack string, needle string) int {
	if needle == "" {
		return 0
	}
	hl := len(haystack)
	nl := len(needle)
	for i := 0; i < hl; i++ {
		if haystack[i] == needle[0] && hl >= i+nl && haystack[i:nl+i] == needle {
			return i
		}
	}

	return -1
}
