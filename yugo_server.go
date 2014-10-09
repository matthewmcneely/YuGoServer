package yugo_server

/*
	YugoServer. A golang-based minimalistic web server with template support, but no hub caps, no power windows.

	Start with output that redirects stdout and stderr to your logfile(s), like so:
	  $ cd <directory_to_serve_from>
	  $ ./yugo_server >access.log 2>error.log &

	Multiple host support is accomplished via the existence of host-named subdirectories (./foo.com for instance)

	Basic template support is accomplished by adding a template_data.json file containing the data object used in parsing.

	Using ApacheBench, this subcompact manages over 3,900 requests/sec (template parses), over 2,000 on a EC2 Micro Linux instance.
*/

import (
	"html/template"
	"io"
	"io/ioutil"
	"log"
	"mime"
	"net/http"
	"os"
	"path"
	"regexp"
	"strconv"
	"strings"
	"template_data"
	"time"
)

type YugoServer struct {
	port       int
	workingDir string
}

const placard string = "YugoServer/0.1"

var accessLog *log.Logger = log.New(os.Stdout, "", log.LstdFlags)
var errorLog *log.Logger = log.New(os.Stderr, "", log.LstdFlags)
var templateExpr *regexp.Regexp = regexp.MustCompile("{{[a-zA-Z0-9 ._@/\"]*}}") // very simplistic matching

func NewServer(port int, directory string) *YugoServer {
	template_data.Load(directory)
	return &YugoServer{port, directory}
}

func (yugo YugoServer) ServeHTTP(writer http.ResponseWriter, req *http.Request) {
	yugo.handleRequest(writer, req)
}

func (yugo YugoServer) Run() error {
	log.Printf(placard+" starting on port %d in folder %s\n", yugo.port, yugo.workingDir)
	if template_data.TemplateData == nil {
		accessLog.Println("No template_data file found, continuing without template support")
	}
	http.HandleFunc("/", yugo.handleRequest)
	err := http.ListenAndServe(":"+strconv.Itoa(yugo.port), nil)
	return err
}

func (yugo YugoServer) handleRequest(writer http.ResponseWriter, req *http.Request) {
	var content []byte
	var stringContent string
	var err error
	timer := time.Now()
	responseCode := 200
	writer.Header().Set("server", placard)
	server := strings.Split(req.Host, ":")[0]
	if req.URL.Path == "/" {
		req.URL.Path = "/index.html"
	}
	file := yugo.workingDir + "/" + server + path.Clean(req.URL.Path)
	if statInfo, err := os.Stat(file); err != nil {
		if os.IsNotExist(err) {
			responseCode = 404
		} else {
			responseCode = 500
		}
		goto DISCONNECT
	} else if statInfo.IsDir() {
		responseCode = 403
		goto DISCONNECT
	}
	content, err = ioutil.ReadFile(file)
	if err != nil {
		errorLog.Println(err)
		responseCode = 500
		goto DISCONNECT
	}

	if strings.Contains(mime.TypeByExtension(path.Ext(file)), "text/html"); template_data.TemplateData != nil {
		stringContent = string(content)
		if templateExpr.MatchString(stringContent) {
			t, err := template.New("").Parse(stringContent)
			if err != nil {
				errorLog.Println(err)
				goto DISCONNECT
			}
			err = t.Execute(writer, template_data.TemplateData)
			if err != nil {
				errorLog.Println(err)
				goto DISCONNECT
			}
		} else {
			io.WriteString(writer, string(stringContent))
		}
	} else {
		writer.(io.Writer).Write(content)
	}

DISCONNECT:

	if responseCode != 200 {
		http.Error(writer, strconv.Itoa(responseCode)+" error", responseCode)
	}
	accessLog.Printf("%s %s %d [%s]", req.RemoteAddr, req.RequestURI, responseCode, time.Since(timer))
}
