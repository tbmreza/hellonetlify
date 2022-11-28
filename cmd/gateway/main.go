// package main
//
// import (
//
//	"flag"
//	"fmt"
//	"log"
//	"net/http"
//
//	"github.com/carlmjohnson/feed2json"
//
// )
//
//	func main() {
//		port := flag.Int("port", -1, "specify a port")
//		flag.Parse()
//		http.Handle("/api/feed", feed2json.Handler(
//			feed2json.StaticURLInjector("https://news.ycombinator.com/rss"), nil, nil, nil))
//		port_str := fmt.Sprint(":", *port)
//		log.Fatal(http.ListenAndServe(port_str, nil))
//	}
package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/apex/gateway"
	"github.com/carlmjohnson/feed2json"
)

func main() {
	port := flag.Int("port", -1, "specify a port to use http rather than AWS Lambda")
	flag.Parse()
	listener := gateway.ListenAndServe
	portStr := "n/a"
	if *port != -1 {
		portStr = fmt.Sprintf(":%d", *port)
		listener = http.ListenAndServe
		http.Handle("/", http.FileServer(http.Dir("./public")))
	}

	http.Handle("/api/feed", feed2json.Handler(
		feed2json.StaticURLInjector("https://news.ycombinator.com/rss"), nil, nil, nil))
	log.Fatal(listener(portStr, nil))
}
