// display

package app

import (
	// 必要なpackageのみimportすること
	"fmt"
	"net/http"
)



func init() {

	// 入力が入ったとき, handlePataで受け取ります
	// 第一引数が、URLになる感じや
	http.HandleFunc("/", handlePata)
}

func handlePata(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprintf(w, "パタトカクシーー\n")
	fmt.Fprintf(w, `<form method=get>`)
	fmt.Fprintf(w, `<div><input type=text name="a"></div>`)
	fmt.Fprintf(w, `<div><input type=text name="b"></div>`) 
	fmt.Fprintf(w, `<div><input type="submit" value="送信"></div>`)
	fmt.Fprintf(w, `</form>`)


}
