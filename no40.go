package main

import (
	"fmt"
	"runtime" // お、新しいの出てきた
)

/*

switch. デフォルトで break する挙動。 fallthrough キーワードをつけると break しなくなる。

*/

func main() {
	fmt.Print("Go runs on ")
	switch os := runtime.GOOS; os { // switch 書式注意ってこれは if のときもでてきた、条件式部分に変数代入式を書けるというやつか。既出だった。
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		// freebsd, openbsd,
		// plan9, windows...
		fmt.Printf("%s.", os)
	}
}