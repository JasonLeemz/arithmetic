package main

import (
	"fmt"
	"reflect"
	"strings"
	"strconv"
)

func main() {

	a := "10"
	b := 1
	c := "2"
	d := "1.0"
	e := 1.0
	var f int64
	f = 1.0
	var h float32
	h = 3.1
	x := int64(9)

	res := Plus(a, b, c, d, e, f, h, x)
	fmt.Println(res)

	res = Minus(a, b, c, d, e, f, h, x)
	fmt.Println(res)

	res = Multip(a, b, c, d, e, f, h, x)
	fmt.Println(res)

	res = Division(a, b, c, d, e, f, h, x)
	fmt.Println(res)
}

//减法
func Minus(args ...interface{}) (float64) {
	integer, decimal := extract(args...)

	//乘以-1后转化为加法
	for k1, _ := range integer {
		if k1 != 0 {
			integer[k1] = integer[k1] * -1
		}
	}
	for k2, _ := range decimal {
		if k2 != 0 {
			decimal[k2] = decimal[k2] * -1
		}
	}

	res := plusMapValues(integer, decimal)
	return res

}

//加法
func Plus(args ...interface{}) (float64) {
	integer, decimal := extract(args...)
	res := plusMapValues(integer, decimal)
	return res
}

//加法运算的具体过程
func plusMapValues(integer map[int]int64, decimal map[int]float64) (float64) {
	var i int64
	i = 0
	var d float64
	d = 0.0

	//整数部分相加
	for _, v1 := range integer {
		i += v1
	}

	//小数部分相加
	for _, v2 := range decimal {
		d += v2
	}
	return float64(i) + d
}

//乘法
func Multip(args ...interface{}) (float64) {

	numMap := cleanArgs(args...)
	result := float64(1)

	//将每一项相乘
	for _, num := range numMap {
		result = result * num
	}
	return result
}

//除法
func Division(args ...interface{}) (float64) {
	numMap := cleanArgs(args...)
	result := numMap[0] //被除数
	for i, num := range numMap {
		if i == 0 {
			continue
		}

		result = result / num
	}
	return result
}

//将参数分隔为整数部分和小数部分
func extract(args ...interface{}) (map[int]int64, map[int]float64) {
	integer := map[int]int64{}
	decimal := map[int]float64{}

	for i, arg := range args {
		ty := reflect.TypeOf(arg).String()
		switch ty {
		case "string":
			nums := strings.Split(arg.(string), ".")
			if len(nums) > 2 || len(nums) < 1 {
				err := fmt.Sprintf("args value is unvalid of Minus th. %v :type is %v , value is %v", i, ty, arg)
				panic(err)
			}

			if len(nums) == 2 {
				d, err := strconv.ParseFloat("0."+nums[1], 64)
				if err != nil {
					panic(err)
				}

				decimal[i] = d
			}

			inte, err := strconv.ParseInt(nums[0], 10, 64)
			if err != nil {
				panic(err)
			}
			//fmt.Println(nums)
			integer[i] = inte

		case "int", "int64":
			if ty == "int" {
				integer[i] = int64(arg.(int))
			} else {
				integer[i] = arg.(int64)
			}

		case "float32", "float64":
			//不保证精度,使用的时候需要注意
			if ty == "float32" {
				integer[i] = int64(arg.(float32))
				decimal[i] = float64(arg.(float32)) - float64(integer[i])
			} else {
				integer[i] = int64(arg.(float64))
				decimal[i] = arg.(float64) - float64(integer[i])
			}
		default:
			err := fmt.Sprintf("args value is unvalid of Minus th. %v :type is %v , value is %v", i, ty, arg)
			panic(err)
		}
	}

	return integer, decimal
}

//将参数中为字符串类型的数据转换成int或float
func cleanArgs(args ...interface{}) (map[int]float64) {
	num := map[int]float64{}

	for i, arg := range args {
		ty := reflect.TypeOf(arg).String()
		switch ty {
		case "string":
			nums := strings.Split(arg.(string), ".")
			if len(nums) > 2 || len(nums) < 1 {
				err := fmt.Sprintf("args value is unvalid of th. %v :type is %v , value is %v", i, ty, arg)
				panic(err)
			}

			if len(nums) == 2 {
				res, err := strconv.ParseFloat(nums[0]+"."+nums[1], 64)
				if err != nil {
					panic(err)
				}

				num[i] = res
			} else {
				res, err := strconv.ParseFloat(nums[0], 64)
				if err != nil {
					panic(err)
				}
				num[i] = res
			}

		case "int", "int64":
			if ty == "int" {
				num[i] = float64(arg.(int))
			} else {
				num[i] = float64(arg.(int64))
			}

		case "float32", "float64":
			if ty == "float32" {
				num[i] = float64(arg.(float32))
			} else {
				num[i] = arg.(float64)
			}
		default:
			err := fmt.Sprintf("args value is unvalid of th. %v :type is %v , value is %v", i, ty, arg)
			panic(err)
		}
	}

	return num
}
