package main

var exp = 0   // 当前经验
var level = 1 // 当前等级

// 升级模块
func UpLevel(e int) {
	exp += e

	level = (exp / 10) + 1
}
