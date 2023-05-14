package main

import (
	"fmt"
	"sort"
)

/*
	N个设备 处理业务的能力分别是a1,a2,...aN.流量分发策略保证这部分设备始终满负荷运行
	需要依次下线这些设备,并且下线持续一分钟,
	下线操作期间,这部分设备处理的业务量最小值是T T是如何随着各个设备的吞吐量变化而变化的呢?
	假设通过一些指令来盖面一些设备的吞吐量 Q个质量,每个指令由(i,j)构成,使第i个设备的吞吐量变为j
	请问在该指令影响下 T会变成多少?
	指令的影响可以看成临时且独立的 = 指令生效前 之前指令的影响会被重置?
*/

func main() {
	// 处理输入
	// 5 1 10 3 4 6 3 2 1 2 4 4 6
	N := 0 // N = 设备数量
	fmt.Scan(&N)
	devices := make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Scan(&devices[i])
	}
	Q := 0 // Q = 指令数量
	fmt.Scan(&Q)
	for k := 0; k < Q; k++ {
		i, j := 0, 0
		fmt.Scan(&i, &j)
		i -= 1 // 控制改变的设备
		newDevices := make([]int, N)
		copy(newDevices, devices)
		newDevices[i] = j
		fmt.Println(getT(newDevices))
	}

}

func getT(devices []int) int {
	//lens := len(devices)
	T := 0
	// 将devices按照从大到小排序
	sort.Slice(devices, func(i, j int) bool {
		return devices[i] > devices[j]
	})
	for i, device := range devices {
		T += (i + 1) * device
	}
	return T
}
