// 写一个函数在原地完成消除[]string中相邻重复的字符串的操作

// 具体步骤如下：

// 使用两个指针，一个用于遍历原始字符串数组，另一个用于指示当前有效字符串的位置。
// 从第二个字符串开始遍历原始字符串数组。
// 每当遍历到的当前字符串与前一个字符串不相同时，将当前字符串复制到当前有效位置，并将有效位置指针向前移动一步。
// 最后，将有效位置指针之后的字符串设置为空字符串""，以标记原数组的结束。

package main

import "fmt"

func removeDuplicates(strings []string) []string {
	if len(strings) == 0 {
		return []string{}
	}

	validIndex := 1

	for i := 1; i < len(strings); i++ {
		if strings[i] != strings[i-1] {
			strings[validIndex] = strings[i]
			validIndex++
		}
	}

	// for i := validIndex; i < len(strings); i++ {
	// 	strings[i] = ""
	// }

	return strings[:validIndex]
}

func main() {
	strings := []string{"aa", "bb", "bb", "cc", "cc", "cc", "dd", "dd", "ee"}
	result := removeDuplicates(strings)
	fmt.Println(result) // Output: [aa bb cc dd ee]
}
