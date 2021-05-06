package solution

import "fmt"

// import (
// 	"reflect"
// 	"unsafe"
// )

// func str2bytes(s string) []byte {
// 	b := (*reflect.SliceHeader)(unsafe.Pointer(&s))
// 	b.Cap = b.Len
// 	return *(*[]byte)(unsafe.Pointer(b))
// }

// func bytes2str(b []byte) string {
// 	s := (*reflect.StringHeader)(unsafe.Pointer(&b))
// 	return *(*string)(unsafe.Pointer(s))
// }

func reverse(b []byte, begin int, end int) []byte {
	for i, j := begin, end; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}
	return b
}

func trimSpace(b []byte) []byte {
	i, j := 0, 0
	space := byte(' ')
	// 处理单词中间的空格, 将连续多个空格合并为1个
	for j < len(b) {
		if b[j] != space {
			b[i] = b[j]
			i++
			j++
		} else {
			for j < len(b) && b[j] == space {
				j++
			}
			// 确保新的slice开头和结尾不会是空格
			if i != 0 && j < len(b) {
				b[i] = space
				i++
			}
		}
	}
	return b[:i]
}

func reverseEachWord(b []byte) []byte {

	space := byte(' ')
	for i, j := 0, 0; j <= len(b); j++ {
		if j == len(b) || b[j] == space {
			b = reverse(b, i, j-1)
			i = j + 1
		}
	}
	return b
}

func reverseWords(s string) string {
	// arr := str2bytes(s)
	arr := []byte(s)
	fmt.Printf("origial string: %q\n", string(arr))
	// 1. 去掉多余的空格
	arr = trimSpace(arr)
	fmt.Printf("after trim: %q\n", string(arr))

	// 2. 翻转所有字母
	arr = reverse(arr, 0, len(arr)-1)
	fmt.Printf("after reverse whole letter: %q\n", string(arr))

	// 3. 翻转每个单词内的字母
	arr = reverseEachWord(arr)
	fmt.Printf("after reverse letter for each word: %q\n", string(arr))

	// return bytes2str(arr)
	return string(arr)
}
