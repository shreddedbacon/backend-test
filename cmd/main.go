package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/cms", cmsHandler)
	http.HandleFunc("/api", apiHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

var page = `<html>
<body bgcolor="#4c7ab0">
<h1 style="color:#000000">Hello world Backend!</h1>
</body>
</html>`

var cms = `<html>
<body bgcolor="#56b54e">
<h1 style="color:#000000">CMS</h1>
</body>
</html>`

var api = `<html>
<body bgcolor="#fcba03">
<h1 style="color:#000000">API</h1>
</body>
</html>`

func handler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(fmt.Sprintf("%v\n", page)))
}
func cmsHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(fmt.Sprintf("%v\n", cms)))
}
func apiHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(fmt.Sprintf("%v\n", api)))
}
