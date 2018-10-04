//Ejemplo basico de un server http en go usando la libreria standard http.
package main

import (
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/li", handler2)
	http.ListenAndServe(":8000", nil)
}

func handler(rsp http.ResponseWriter, rqt *http.Request) {
	io.WriteString(rsp, "<br>Hola servidor Go!</br>")
}

func handler2(rsp http.ResponseWriter, rqt *http.Request) {
	io.WriteString(rsp, "<br>Esta es otra pagina!</br>")
}
