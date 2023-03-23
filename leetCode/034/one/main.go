package main

import (
	"sort"
	"strconv"
	"time"
)

func groupAnagrams(strs []string) [][]string {
	var res [][]string
	var m = make(map[string]int) //string是单个字符串排序后的intStr,如果存在的话则找到他在res数组中的下标插入
	var intStr []string
	for _, str := range strs {
		intStr = append(intStr, convertStr2intStr(str))
	}
	for i, str := range intStr {
		if index, ok := m[str]; ok {
			// 存在的话则插入原str到对应的数组中
			res[index] = append(res[index], strs[i])
		} else {
			// 不存在则新建一个
			m[str] = len(res)
			res = append(res, []string{strs[i]})
		}

	}
	time.Sleep(1)
	return res
}

func convertStr2intStr(str string) string {
	var tmp []int
	var res string
	for _, v := range str {
		tmp = append(tmp, int(v))
	}
	sort.Slice(tmp, func(i, j int) bool {
		return tmp[i] < tmp[j]
	})

	for _, v := range tmp {
		res += strconv.Itoa(v)
	}
	return res
}

func main() {
	groupAnagrams([]string{"zxc", "zxd", "zcx"})

}
