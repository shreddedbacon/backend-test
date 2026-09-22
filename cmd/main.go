package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	lagoonProject = os.Getenv("LAGOON_PROJECT")
	lagoonEnvironment = os.Getenv("LAGOON_ENVIRONMENT")
	http.HandleFunc("/", handler)
	http.HandleFunc("/cms", cmsHandler)
	http.HandleFunc("/api", apiHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

var lagoonProject, lagoonEnvironment string

var page = `<html>
<body bgcolor="#4c7ab0">
<h1 style="color:#000000">Hello world Backend!</h1>
</body>
</html>\n`

var cms = `<html>
<body bgcolor="#56b54e">
<h1 style="color:#000000">CMS</h1>
<h1 style="color:#000000">Project: %s</h1>
<h1 style="color:#000000">Environment: %s</h1>
</body>
</html>\n`

var api = `<html>
<body bgcolor="#fcba03">
<h1 style="color:#000000">API</h1>
<h1 style="color:#000000">Project: %s</h1>
<h1 style="color:#000000">Environment: %s</h1>
</body>
</html>\n`

func handler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(fmt.Sprintf("%v\n", page)))
}
func cmsHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(fmt.Sprintf(cms, lagoonProject, lagoonEnvironment)))
}
func apiHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(fmt.Sprintf(api, lagoonProject, lagoonEnvironment)))
}
