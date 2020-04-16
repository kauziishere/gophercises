package main

import (
	"flag"
	"fmt"
	"html"
	"net/http"
	"time"

	"github.com/kauziishere/gophercises/URLShortner/yamlUtil"
)

const (
	defaultYAMLFilename	= "../resources/url_data.yaml"
	yamlFlagHelpString	= "A yaml filename in form of string to get url mapping\n"
	badRequestData		= `<html><title>Kauzi URLShortner</title>
				   <body>Incorrect URL my friend :(</body>
				   </html>`
)

var urlMappings map[interface{}]interface{}
var yamlFileName *string

func mapHandler(w http.ResponseWriter, r *http.Request) {
	redirectionToken   := html.EscapeString(r.URL.Path)[1:]
	if redirectionURL, ok := urlMappings[redirectionToken] ; ok {
		http.Redirect(w, r, redirectionURL.(string), 302)
	} else {
		fmt.Fprintf(w, badRequestData)
	}
}

func init() {
	yamlFileName = flag.String("yaml", defaultYAMLFilename, yamlFlagHelpString)
}

func main() {
	var err error

	if urlMappings, err = yamlUtil.FetchMapFromYAMLFile(*yamlFileName) ; nil != err {
		panic(err)
	}


	s := &http.Server{
		Addr:		":8080",
		Handler:	http.HandlerFunc(mapHandler),
		ReadTimeout:	10 * time.Second,
		WriteTimeout:	10 * time.Second,
		MaxHeaderBytes:	1 << 20,
	}
	s.ListenAndServe()
}
