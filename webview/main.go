package main

import (
	"log"

	"github.com/webview/webview"
)

func main() {
	w := webview.New(true)
	defer w.Destroy()
	w.SetTitle("Hello")
	w.Bind("noop", func() string {
		log.Println("hello")
		return "hello"
	})
	w.Bind("add", func(a, b int) int {
		return a + b
	})
	w.Bind("quit", func() {
		w.Terminate()
	})
	w.Navigate(`data:text/html,
		<!doctype html>
		<html>
			<body>hello</body>
			<button id="myBtn">QUIT</button>
			<button id="addBtn">1 + 2</button>
			<script>
				doquit = function() {
					quit();
				};
				doadd = function() {
					add(1,2).then(function(res){
						alert(res);
					});
				}
				window.onload = function() {
					document.getElementById("myBtn").addEventListener("click", doquit); 
					document.getElementById("addBtn").addEventListener("click", doadd);
				};
			</script>
		</html>
	`)
	w.Run()
}
