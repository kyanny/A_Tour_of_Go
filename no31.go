package main

import "fmt"

func main() {
	m := make(map[string]int) // 復習。 map の定義だけ書くことができて初期化して map インスタンスを作るには make を使う。関数の中であれば var 宣言なしの := で変数代入できる。この map はキーが文字列で値が整数。

	m["Answer"] = 42 // map には添字のキーで参照と代入ができる
	fmt.Println("The value:", m["Answer"])

	m["Answer"] = 48 // map は mutable なので値を変更できる
	fmt.Println("The value:", m["Answer"])

	delete(m, "Answer") // delete で map からキーと値のペアを削除できる
	fmt.Println("The value:", m["Answer"])

	v, ok := m["Answer"] // map に添字アクセスして値を参照するとき、多値を受け取ると二つ目にそのキーバリューが存在するかの bool 値を返してくれるのか
	fmt.Println("The value:", v, "Present?", ok) // この場合だと ok は false になる。指定したキーが map に存在しない場合、値 v としては型ごとに決まっている zero 値が入る。
}

// 最後の例に類似して、 map に存在しないキーで参照した場合返り値の値もやはり型ごとの zero 値になる。