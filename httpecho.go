package main

// Stdlib
import "fmt"
import "log"
import "flag"
import "time"
import "strconv"
import "strings"
import "net/http"
import "net/http/httputil"

func RequestDumpHandler(resp http.ResponseWriter, req *http.Request) {
	request_dump, err := httputil.DumpRequest(req, true)

	t := time.Now()
	fmt.Printf("[%s] %s, with dump:\n", t.Format(time.StampMilli), req.URL.String())

	if err == nil {
		fmt.Printf("---\n")
		fmt.Printf("%s", request_dump)
		fmt.Printf("---\n")
		resp.Write([]byte("Request received\n"))
	} else {
		fmt.Printf("Request unable to be output\n")
		resp.Write([]byte("Could not handle request\n"))
	}
}

func HttpResponseCodeHandler(resp http.ResponseWriter, req *http.Request) {
	parts := strings.Split(req.URL.Path, "/")

	if len(parts) == 3 && len(parts[2]) > 0 {
		if i, err := strconv.Atoi(parts[2]); err == nil {
			resp.WriteHeader(i)
		} else {
			resp.Write([]byte(fmt.Sprintf("Unable to parse int from '%s'\n", parts[2])))
		}
	} else {
		resp.Write([]byte("Format for response-code manipulation is: '/code/{CODE}'\n"))
	}

	t := time.Now()
	fmt.Printf("[%s] %s\n", t.Format(time.StampMilli), req.URL.String())
}

func main() {
	addr := flag.String("address", "", "the address to bind to. Default is 0.0.0.0")
	port := flag.String("port", "8080", "the local port to bind to")

	flag.Parse()

	http.HandleFunc("/code/", HttpResponseCodeHandler)
	http.HandleFunc("/", RequestDumpHandler)

	fmt.Printf("httpecho starting up on %s:%s...", *addr, *port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%s", *addr, *port), nil))
}
