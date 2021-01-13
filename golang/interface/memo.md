# interface で遊ぶコーナー

- wrote : 2021/01/14

## 本日の主役

- `interface`

## interface

- [Goのinterfaceがわからない人へ - Qiita](https://qiita.com/rtok/items/46eadbf7b0b7a1b0eb08)
- [A Tour of Go](https://go-tour-jp.appspot.com/methods/11)
- interface で関数を定義しておくと、interface への代入はガワをかぶせる感じになる(変換される)
  - ex
    - `type I interface {M()}` を定義して `var i I` をすると `type I な i` が出来上がる
    - `type T struct { S string }` を定義して、`func (t *T) M() {fmt.Println(t.S)}` しておく。
    - `i = &T{"Hello"}` すると、type `T` をレシーバーとする func `M` が存在するため、interface `I` は成立できる
      - 雰囲気は分かったが、雑な typescript のイメージを思い出すと `i = &T{S:"Hello"}` ではないのはよく分かっていない
    - 同様に `type F float64` を定義して、`func (f F) M() { fmt.Println(f)}` を定義しておく。
    - そうすると、`i = F(math.Pi)` ができて、`i.M()` は 3.1415... になる

## インターフェースのレシーバ

- [【Go言語】関数にメソッドを定義する - Qiita](https://qiita.com/ksugimori/items/465f0a4c4fe315158df5) や [The Go Programming Language Specification - The Go Programming Language](https://golang.org/ref/spec#Method_declarations) を読んだところ、基本的な型などはレシーバになれない
