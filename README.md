# go-academy-prep

Go Academyに向けて、Goを基礎から勉強するために作ったリポジトリです。

もともとHTML / CSS / JavaScript / PHPなどは触っていましたが、Goはほぼ初めてなので、
書きながら分からなかったところを確認しつつ進めています。

## 今やっていること

現在はこのあたりまで学習しています。

- 変数、型
- if / for
- 関数
- slice / map
- struct
- method
- pointer
- error処理
- package
- Go Modules

最初はmain.goに全部書いていましたが、
途中から役割ごとにファイルやpackageを分けるようにしています。

```text
.
├── main.go
├── user.go
├── practice.go
├── calc/
│   └── calc.go
└── go.mod

calc は計算処理を別packageに分ける練習として作っています。

func Divide(a int, b int) (int, error) {
	err := checkZero(b)

	if err != nil {
		return 0, err
	}

	return a / b, nil
}

func checkZero(b int) error {
	if b == 0 {
		return fmt.Errorf("0では割れません")
	}

	return nil
}

```
Divide は外から使えるようにして、
checkZero はcalcの中だけで使う処理にしています。


##次にやること
次はテストを書いて、その後JSONやHTTPを触る予定です。
