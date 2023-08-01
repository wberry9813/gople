// 编写一个rotate函数，通过一次循环完成旋转。

package main

import (
	"fmt"
)

// rotate函数用于将切片s中的元素向右旋转k个位置
func rotate(s []int, k int) {
	// 处理k超过切片s长度的情况
	k = k % len(s)

	// 创建一个新的切片temp，长度与s相同
	temp := make([]int, len(s))

	// 将切片s中的元素复制到temp中
	copy(temp, s)

	// 将切片s中的元素根据旋转要求重新排列
	for i, val := range temp {
		newIndex := (i + k) % len(s)
		s[newIndex] = val
	}
}

func main() {
	// 示例使用
	s := []int{1, 2, 3, 4, 5}
	k := 2
	rotate(s, k)
	fmt.Println(s) // 输出：[4 5 1 2 3]
}
