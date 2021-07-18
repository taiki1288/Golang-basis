package main

import "fmt"

func main() {
	a := [2][2]string{{"yamashita", "yamamoto"}, {{"suzuki", "tanaka"}}

	fmt.Println(a[0][0])
	fmt.Println(a[0][1])
	fmt.Println(a[1][0])
	fmt.Println(a[1][1])
	
}

// ●二次元配列は角括弧を2回入れることで出力される。
// "yamashita"は一つ目の配列の1番目の文字列だから[0][0]とすると出力される。
// "yamamoto"は一つ目の配列の2番目の文字列だから[0][1]とすると出力される。
// "suzuki"は二つ目の配列の一番目だから[1][0]とすると出力される。
// "tanaka"は二つ目の配列の二番目だから[1][1]とすると出力される。