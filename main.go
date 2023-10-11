package main

import "github.com/webview/webview_go"

func main() {
	w := webview.New(false)
	defer w.Destroy()
	w.SetTitle("BananoVault")
	w.SetSize(1200, 700, webview.HintNone)
  w.Navigate("https://app.element.io")
	w.Run()
}
