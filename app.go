package app

import (
    "github.com/garyburd/twister/server"
    "github.com/garyburd/twister/web"
    "text/template"
)

var (
  page *template.Template
)

func index(request *web.Request) {
  page.Execute(
    request.Respond(web.StatusOK, web.HeaderContentType, "text/html; charset=utf-8"),
    request.URL.Host)
}

func main() {
  page = template.Must(template.New("index").Parse(body))

  server.Run(":8080",
    web.NewRouter().
      Register("/", "GET", index))
}

const body = `
<html>
  <head>
    <title>Developed using Go :)</title>
  </head>
  <body>
    <h1>I made this page using Go language</h1>
  </body>
</html> `
