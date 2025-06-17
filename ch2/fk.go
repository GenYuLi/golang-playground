package main

import "fmt"

const imconst = 100 // 常數宣告，imconst 是常數名稱，100 是值

func main() {
	println(imconst) // 輸出常數值
	// imconst = 200 // 這裡會報錯，因為常數不能被重新賦值
	var x = 10
	var b byte = 100
	var sum3 = x + int(b)
	// var sum4 int = x + int8(b) // 這裡會報錯，因為 int8 不能直接轉換成 int
	var sum4 = byte(x) + b
	fmt.Printf("sum3: %d, sum4: %d\n", sum3, sum4)
	var str = string(b) // byte 轉換成 string
	fmt.Printf("str: %s\n", str)

	// in golang literal dose not have type
	var _ float64 = 10
	var _ = 200.3 * 5
	yoman := 2e3 + 3
	fmt.Printf("yoman: %f\n", yoman)
}
