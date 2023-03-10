package main

import (
	"flag"
	"log"
	"net/http"
	"text/template"
)

var addr = flag.String("addr", ":1718", "https://chart.apis.google.com") // Q=17, R=18
var templ = template.Must(template.New("qr").Parse(templateStr))

func main() {
	flag.Parse()
	http.Handle("/", http.HandlerFunc(QR))
	err := http.ListenAndServe(*addr, nil)
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
func QR(w http.ResponseWriter, req *http.Request) {
	templ.Execute(w, req.FormValue("s"))
}

const templateStr = `
<html>
<head>
<title>QR Link Generator</title>
</head>
<body>
{{if .}}
<img src=
"http://chart.apis.google.com/chart?chs=300x300&cht=qr&choe=UTF-8&chl={{urlquery .}}"/>
<br>
{{html .}}
<br>
<br>
37
{{end}}
<form action="/" name=f method="GET"><input maxLength=1024 size=70
name=s value="" title="Text to QR Encode"><input type=submit
value="Show QR" name=qr>
</form>
</body>
</html>
`
