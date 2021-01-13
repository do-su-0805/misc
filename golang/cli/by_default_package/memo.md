# cli ツールを golang Standard library だけで作る

- wrote : 2021/01/14

## 参考文献

- [Golangの標準パッケージで簡単なCLIツールを作ってみる - Qiita](https://qiita.com/nightswinger/items/8abc3ee7db41a3365784)

## Standard library

- `os` とか `fmt` とかああいうやつ
- なんていうんだろと調べたら標準ライブラリって出てきてなるほど〜ってなった

## 本日の主役

- Package `flag` 

## Package flag

[flag - The Go Programming Language](https://golang.org/pkg/flag/)

- flag を受け付けて値を取り出して返すような奴らは `フラグ`,`初期値`,`ヘルプメッセージ` の形で受け付けられそう
  - ex : `flag.Int`
- 変数はポインタ渡しで使う場合は `変数ポインタ`,`フラグ`,`初期値`,`ヘルプメッセージ` の形で受け付けられそう
  - ex : `flag.IntVar`
- 自分て定義した type に引き渡すこともできそう。interface が絡んでくるのでちょっと難しいけど。`Value interface に引き渡す内容`,`フラグ`,`ヘルプメッセージ` の形で受け付けられそう
  - [flag - The Go Programming Language](https://golang.org/pkg/flag/#Var)
  - Value Interface は [flag - The Go Programming Language](https://golang.org/pkg/flag/#Value) で、`String() string` と `Set(string) error` を要求する。
  - [【Go言語】flagパッケージで独自の型のオプションを定義する - フリエン生活](https://free-engineer.life/golang-flag-package-custom-type/) がわかりやすかった。
- まぁそんな感じで受け付けたいやつを用意できたら、`flag.Parse()` で全部整えてくれる
  - 上記の例だと、 `flag.IntVar` とか `flag.Var` はイメージつくんだけど、`flag.Int` ができる理由がイメージつかない。ポインタ渡してその先を書き換えた？
  - あと、`flag.IntVar` を二つ用意した時とかも不思議。引数が別なものは別なものとして認識されるのだろうか。ちゃんと実装を見たらわかるんだけど今は混乱しそうなので御茶を濁す
- `hoge -a a -b b -cde efg hij k l` みたいにした時、`flag.Parce` した後に `flag.Args` でスライスとして取得できたり、`flag.Arg(i)` で取得できたりする。
  - また、`flag.NArg()` で引数の数を表示してくれるらしい
