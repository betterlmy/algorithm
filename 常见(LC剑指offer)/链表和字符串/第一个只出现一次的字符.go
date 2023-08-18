package main

/*
在字符串 s 中找出第一个只出现一次的字符。如果没有，返回一个单空格。 s 只包含小写字母。
输入：s = "abaccdeff"
输出：'b'
*/
type unique struct {
	uniqueString []byte       // 记录唯一的字符串(有序)
	stringIndex  map[byte]int // 判断是否重复
	length       int
}

func firstUniqChar(s string) byte {
	u := unique{
		uniqueString: make([]byte, 0),
		stringIndex:  make(map[byte]int),
		length:       0,
	}
	for _, v := range s {
		if index, ok := u.stringIndex[byte(v)]; ok {
			// 已经重复了 则从[]byte中标记为1删除
			u.uniqueString[index] = '1'
		} else {
			// 不存在则添加
			u.length++
			u.uniqueString = append(u.uniqueString, byte(v))
			u.stringIndex[byte(v)] = u.length - 1
		}
	}

	// 遍历uniqueString
	for _, b := range u.uniqueString {
		if b == '1' {
			continue
		} else {
			return b
		}
	}
	return ' '
}

func main() {

}
