package yugo_server

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
)

func getWorkingDir(t *testing.T) string {
	workingDir, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}
	return workingDir + "/test_fixtures"
}

func Test200(t *testing.T) {
	yugo := NewServer(-1, getWorkingDir(t))
	ts := httptest.NewServer(yugo)
	defer ts.Close()
	res, err := http.Get(ts.URL)
	if err != nil || res.StatusCode != 200 {
		t.Fail()
	}
}

func Test404(t *testing.T) {
	yugo := NewServer(-1, getWorkingDir(t))
	ts := httptest.NewServer(yugo)
	defer ts.Close()
	res, err := http.Get(ts.URL + "/non-existent-page.html")
	if err != nil || res.StatusCode != 404 {
		t.Fail()
	}
}

func Test403(t *testing.T) {
	yugo := NewServer(-1, getWorkingDir(t))
	ts := httptest.NewServer(yugo)
	defer ts.Close()
	res, err := http.Get(ts.URL + "/empty_directory")
	if err != nil || res.StatusCode != 403 {
		t.Fail()
	}
}

func TestTemplate(t *testing.T) {
	yugo := NewServer(-1, getWorkingDir(t))
	ts := httptest.NewServer(yugo)
	defer ts.Close()
	res, err := http.Get(ts.URL + "/template_page.html")
	if err != nil || res.StatusCode != 200 {
		t.Fail()
	}
	body, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		t.Fail()
	}
	if strings.TrimSpace(string(body)) != strings.TrimSpace(parsedResult) {
		t.Fail()
	}
}

const parsedResult string = `<html>
<body>
<h1>Hello World</h1>

	<p>Hello world: (English)</p>

	<p>हैलो दुनिया: (Hindi)</p>

	<p>你好ा: (Chinese)</p>

	<p>Ciao mondo: (Italian)</p>

	<p>Kaixo mundu: (Basque)</p>

</body>
</html>`
