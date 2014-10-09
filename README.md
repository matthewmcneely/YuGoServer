Yu(Go)Server
=========
A golang-based minimalistic web server with template support, but no hub caps, no power windows.

Start Yu(Go)Server with output that redirects stdout and stderr to your logfile(s), like so:

    $ cd <directory_to_serve_from>
    $ ./yugo_server >access.log 2>error.log &

Yu(Go)Server can also be embedded easily; have a look at the unit test. Also supported is the ServeHTTP interface for testing.

- Multiple host support is accomplished via the existence of host-named folders (./foo.com for instance).
- Basic template support is accomplished by adding a template_data.json file containing the data object used in parsing. See the test_fixtures directory for an example.
- Using ApacheBench, YugoServer manages over 3,900 requests/sec (template parses) on OSX, over 2,000 on a Linux Micro EC2 instance on Amazon.

> This software is a very minimal implementation of HTTP. There are tons of features that *should* be implemented before it would be suitable for production environments.

Example of running from the command line:

	package main
	
	import (
		"os"
		"yugo_server"
	)
	
	var server *yugo_server.YugoServer
	
	func main() {
		workingDir, _ := os.Getwd()
		server = yugo_server.NewServer(8090, workingDir)
		if err := server.Run(); err != nil {
			panic(err.Error())
		}
	}


Made available via an MIT-style license (see the ./LICENSE file)—use this as you see fit.
