package main

import (
	"fmt"
	"strings"
)

func isRobotBounded(instructions string) bool {

	x := 0
	y := 0
	direct := 0 // 0,1,2,3 分别代表北东南西
	instructions = strings.ReplaceAll(instructions, "LR", "")
	instructions = strings.ReplaceAll(instructions, "RL", "")
	for _, v := range instructions {
		if v == 'L' {
			direct -= 1
			if direct < 0 {
				direct = 3
				continue
			}

			direct %= 4
		} else if v == 'R' {
			direct += 1
			direct %= 4
		} else {
			x, y = next(x, y, direct)
		}
	}

	if direct == 0 {
		// 朝向不变
		if x == 0 && y == 0 {
			return true
		}
		return false
	}
	if direct == 2 {
		// 朝向相反
		if x == 0 || y == 0 || x == y {
			return true
		}
		return false
	}

	if direct == 1 || direct == 3 {
		for i := 0; i < 3; i++ {
			for _, v := range instructions {
				if v == 'L' {
					direct -= 1
					if direct < 0 {
						direct = 3
					}

				} else if v == 'R' {
					direct += 1
					if direct > 3 {
						direct = 0
					}
				} else {
					x, y = next(x, y, direct)
				}
			}
		}
		if x == 0 && y == 0 {
			return true
		}
		return false

	}
	return false

}

func next(x, y, direct int) (x1, y1 int) {
	switch direct {
	case 0:
		// 北
		y += 1
	case 1:
		// 东
		x += 1
	case 2:
		// 南
		y -= 1
	case 3:
		// 西
		x -= 1
	}
	return x, y
}
func main() {
	fmt.Println(isRobotBounded("GGRGGR"))
}
