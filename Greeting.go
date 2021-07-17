package main
import ("fmt")

func main () {
	fmt.Println("Good morning")

	fmt.Println("Good afternoon")

	fmt.Println("Good evening")
}

// ●Goは何らかのパッケージに属している必要がある。
// ●Goは必ずmainパッケージに属する必要がある。
// ●fmtはフォーマットの略。
// ●fmtパッケージをインポートすることでパッケージ内の関数を使うことができるようになる。
// ●fmtには文字列を参照させるパッケージが入っている。
// ●よく使う関数をまとめてパッケージと呼んでいる。