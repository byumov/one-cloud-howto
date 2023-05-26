package main

import (
	"fmt"
	"net/http"
)

const content = `
<style>
.box {
	height: 400px;
	width: 1800px;
	background: #FFFFFF;
	font-size: 50px;
	color: #000000;
	text-align: center;
	padding: 0 20px;
	margin: 20px;
	display: flex;
	justify-content: center;
	/* align horizontal */
	align-items: center;
	text: center;
	/* align vertical */
  }
</style>
<div class=box>
    Welcome to One-cloud, Dzen and VK!
</div>
`

func hello(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, content)
}

func main() {
	http.HandleFunc("/", hello)
	listenPort := 443
	err := http.ListenAndServeTLS(fmt.Sprintf(":%d", listenPort), "/one/conf/ssl/tls.crt", "/one/conf/ssl/tls.key", nil)
	if err != nil {
		fmt.Printf("can't start on port %d: %s", listenPort, err.Error())
	}
}
