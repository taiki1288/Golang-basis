//変数名の定義
//var 変数名 型
var n int

//構造体型の変数
var p struct {
	name string
	age int
}

p := struct {
	name string
	age int
}{ name: "Gopher", age: 10 }
//フィールドにアクセスする例
p.age++ //p.age = p.age + 1と同じ
Println(p.name, p.age)

//配列
//■同じ方のデータを集めて並べたデータ構造
  //●要素の方は全て同じ
  //●要素数が違えば別の型
  //●要素数は変更できない
  //●型はリテラルで記述することが多い

  // 型と要素数がセット
  var ns [5]int

  //配列の初期化
  //■ 配列の初期化のいろいろ

  //ゼロ値で初期化
  var ns1 [5]int
  //配列リテラルで初期化
  var ns2 = [5]int{10, 20, 30, 40, 50}
  //要素数を値から推論
  ns3 := [...]int{10, 20, 30, 40, 50}
  //5番目が50、10番目が100で他が0の要素数11の配列
  ns4 := [...]int{5: 50, 10: 100}

  //配列の操作
  ns := [...]int{10, 20, 30, 40, 50}

//要素にアクセス
Println(len(ns))
//スライス演算
fmt.Println(ns[1:2])


//スライスの初期化

//ゼロ値はnil
var ns1 []int

//長さと容量を指定して初期化
//各要素はゼロ値で初期化される
ns1 = make([]int, 3, 10)

//スライスリテラルで初期化
//要素数は指定しなくてよい
//自動で配列は作られる
var ns2 = []int{10, 20, 30, 40, 50}

//5番目が50, 10版目が100で他が0の要素数11のスライス
ns3 := []int{5: 50, 10: 100}

//ユーザー定義型
//typeで名前をつけて新しい型を定義する
type 型名 基底型

//組み込み型を元にする
type MyInt int

//他のパッケージを元にする
type MyWriter io.MyWriter

//型リテラルを基にする
type Person struct {
	Name string
}

//ユーザー定義型の特徴
//基底型とユーザー型の相互キャストが可能
type MyInt int
var n int = 100
m := MyInt(n)
n := int(m)

//型なし定義から明示的なキャストは不要
//●デフォルトの型からユーザー定義型へキャストができる場合
//10秒を戻す
d := 10 * time.Second
        //↓
type Duration int64

//メソッド
//■レシーバと紐つけられた関数
//データとそれに対する操作を紐つけるために用いる
//ドッドでメソッドにアクセスする

type Hex int
func (h Hex) String() string{
	return fmt.Sprintf("%x", int(h))
}

//100をHex型として代入
var hex Hex = 100
//stringメソッドを呼び出す
fmt.Println(hex.String())

//☆レシーバの値を変えたい場合はポインタに変更する

//メソッド値
//■メソッドも値として扱える
//レシーバは束縛された状態
type Hex int
func (h Hex) String() string {
	return fmt.Sprintf("%x", int(n))
}

func main() {
	var hex Hex = 100
	f := hex.String
	fmrt.Println(f())
}

//レシーバは引数と同じイメージ
