/**
*  @Author: xu756 (https://github.com/xu756)
*  @Email: 756334744@qq.com
*  @Date: 2022/9/12 19:36
*  @To: 桥面函数计算
 */

package models

import "math"

func (data *BA) MDP() {
	data.MdpH = ToMDP(data.BA1, data.BA2, data.BA3, data.BA4, data.BA5, data.BA6, data.BA7)

}
func (data *BB) MDP() {
	data.MdpH = ToMDP(data.BB1, data.BB2)

}
func (data *BC) MDP() {
	data.MdpH = ToMDP(data.BC1, data.BC2, data.BC3, data.BC4, data.BC5, data.BC6, data.BC7, data.BC8, data.BC9)

}
func (data *BD) MDP() {
	data.MdpH = ToMDP(data.BD1, data.BD2, data.BD3, data.BD4)

}

func (data *BE) MDP() {
	data.MdpH = ToMDP(data.BE1, data.BE2, data.BE3)

}
func (data *BF) MDP() {
	data.MdpH = ToMDP(data.BF1, data.BF2, data.BF3)

}

// ToSum 求和
func ToSum(nums []float64) float64 {
	var sumNum float64 = 0
	for _, num := range nums {
		sumNum += num
	}
	return sumNum
}

// ToSqrt 平方
func ToSqrt(num1 float64, num2 float64) float64 {
	num2 = num1 * (3*(num1/num2)*(num1/num2)*(num1/num2) - 5.5*(num1/num2)*(num1/num2) + 3.5*(num1/num2))
	return num2

}

func ToMDP(nums ...float64) float64 {
	sum := ToSum(nums)
	n := len(nums)
	var Nums = make([]float64, n)
	maxn := 0.0

	for i, num := range nums {
		Nums[i] = ToSqrt(num, sum)
		if maxn < Nums[i] {
			maxn = Nums[i]
		}
	}
	return math.Min(math.Max(maxn, ToSum(Nums)), 100)

}
