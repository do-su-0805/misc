package main

import "fmt"

// I : Interface だから
type I interface {
	Print()
}

// String : 別名定義しないと interface のレシーバになれない
type String string

// Print : String を受け付けるプリント関数。後述の StructC 構造体でも利用できるかと思ったらできなかったけどなるほどってなった。
func (str String) Print() {
	fmt.Println(str)
}

// StructA : str だと文字列と衝突してしまう。これだけ独自実装の type の練習で試す。こっちはレシーバ書く時に string を受け付けるので問題なし。
type StructA struct {
	S string
}

// Print : StructA 構造体を受け付けるプリント関数
func (s *StructA) Print() {
	fmt.Println(s.S)
}

// StructB : 上記 Struct A に対して int を受け付けてみた場合
type StructB struct {
	S int
}

// Print : StructB 向け
func (s *StructB) Print() {
	fmt.Println(s.S)
}

// StructC : 独自定義の String 型を持つ Struct
type StructC struct {
	S String
}

// Print : Struct C が持つ String 型向けプリント関数
func (s *StructC) Print() {
	fmt.Println(s.S)
}

// Int : String 同様に別名定義しないと interface のレシーバになれない
type Int int

// Print : Int 型向け
func (i Int) Print() {
	fmt.Println(i)
}

func main() {
	var i I

	// String & Int
	i = String("Hello")
	i.Print()
	i = Int(20)
	i.Print()

	// Struct A & B
	i = &StructA{"Hello2"}
	i.Print()
	i = &StructB{21}
	i.Print()

	// Struct C
	i = &StructC{"Hello3"}
	i.Print()
}
