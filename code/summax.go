package code

import (
	"fmt"
	"math"
)

/**
 * 例题1.3 总和最大区间问题 Page35
 * 给定一个实数序列，设计一个最有效的算法，找到一个总和最大的区间。
 */

func SumMax() {
	list := []float32{1.5, -12.3, 3.2, -5.5, 23.2, 3.2, -1.4, -12.2, 34.2, 5.4, -7.8, 1.1, -4.9}

	if len(list) <= 1 {
		return
	}

	// original data
	fmt.Println(list)

	// one
	fmt.Println("1:")
	fmt.Println(one(list))

	// two
	fmt.Println("2:")
	fmt.Println(two(list))

	// three
	fmt.Println("3:")
	fmt.Println(three(list))

	// four
	fmt.Println("4:")
	fmt.Println(four(list))
}

// one 方法1 做一次三重循环
func one(inputs []float32) (outputs []float32, p, q int) {
	k := len(inputs)

	pq := make([]int, 2)
	sums := make(map[int]map[int]float32)

	pq[0] = 0
	pq[1] = 1

	for indexP := 0; indexP < k; indexP++ {
		if _, exist := sums[indexP]; !exist {
			sums[indexP] = make(map[int]float32)
		}
		for indexQ := indexP + 1; indexQ < k; indexQ++ {
			for idx := indexP; idx < indexQ; idx++ {
				sums[indexP][indexQ] = sums[indexP][indexQ] + inputs[idx]
			}

			if sums[indexP][indexQ] > sums[pq[0]][pq[1]] {
				pq[0] = indexP
				pq[1] = indexQ
			}
		}
	}

	return inputs[pq[0]:pq[1]], pq[0], pq[1]
}

// two 方法2 做两重循环
func two(inputs []float32) (outputs []float32) {
	// todo
	return inputs
}

// three 方法3 利用分治算法
func three(inputs []float32) (outputs []float32, p, q int) {

	iLen := len(inputs)

	k := int(math.Floor(float64(iLen) / 2))
	leftPart := inputs[0 : k+1]
	rightPart := inputs[k+1 : int(math.Max(float64(k+2), float64(iLen)))]

	leftOutputs, leftP, leftQ := one(leftPart)     // ?
	rightOutputs, rightP, rightQ := one(rightPart) // ?

	leftSum := sumList(leftOutputs)
	rightSum := sumList(rightOutputs)

	if leftQ == k+rightP {
		if leftSum > 0 && rightSum > 0 {
			return leftOutputs, leftP, k + rightQ + 1
		} else if leftSum > rightSum {
			return leftOutputs, leftP, leftQ
		} else {
			return rightOutputs, k + rightP, k + rightQ + 1
		}
	}

	midOutputs := inputs[leftP:int(math.Min(float64(k+rightQ+1), float64(iLen)))]
	midSum := sumList(midOutputs)

	if rightSum > leftSum && rightSum > midSum {
		return rightOutputs, k + rightP, k + rightQ + 1
	} else if midSum > rightSum && midSum > leftSum {
		return midOutputs, leftP, k + rightQ + 1
	}

	return leftOutputs, leftP, leftQ
}

// four 方法4 正反两遍扫描的方法
func four(inputs []float32) (outputs []float32) {
	// todo
	return inputs
}

// sumList 计算 list 的总和
func sumList(inputs []float32) float32 {
	var sum float32
	for i := 0; i < len(inputs); i++ {
		sum += inputs[i]
	}
	return sum
}
