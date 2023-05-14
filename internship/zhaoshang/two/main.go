package main

import (
	"fmt"
	"sort"
)

/*
	评估故障对历史数据造成的影响范围
	导出了某个时间段内故障集群设备,网络设备和上下游系统的日志.
	程序输出了N条数据,每套数据唯一对应一台设备,并且通过分析给出了这台设备关于这个同一批请求的推断..
	L ti 表示这台设备在本地某个时间ti之前(或恰好在ti这一时刻)收到
	G ti 表示这台设备在某本地时间ti之后(或恰好)收到该批次的健康检查
	求出接受到这条请求时,最少有多少台设备的时间是错误的.
*/

type log struct {
	t     int    // t
	state string // 大还是小
}

func main() {
	// 处理输入
	// 2 G 1 L 4
	N := 0 // N = 设备数量
	fmt.Scan(&N)
	logs := make([]log, N)

	for i := 0; i < N; i++ {
		x := ""
		y := 0
		fmt.Scan(&x, &y)
		logs[i] = log{
			t:     y,
			state: x,
		}
	}
	// 按照t进行排序
	sort.Slice(logs, func(i, j int) bool {
		return logs[i].t < logs[j].t
	})
	minT := 0
	maxT := len(logs)
	errNum := 0
	errNums := make([]int, 0)
	solution(logs, 0, errNum, minT, maxT, &errNums)

}

func solution(logs []log, start int, errNum int, minT int, maxT int, nums *[]int) int {
	//if logs[start].t < minT
	return 0
}
