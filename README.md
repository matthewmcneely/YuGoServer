Yu(Go)Server
=========
A golang-based minimalistic web server with template support, but no hub caps, no power windows.

Start Yu(Go)Server with output that redirects stdout and stderr to your logfile(s), like so:

    $ cd <directory_to_serve_from>
    $ ./yugo_server >access.log 2>error.log &

Yu(Go)Server can also be embedded easily; have a look at the unit test. Also supported is the ServeHTTP interface for testing.

- Mulitple host support is accommplished via the existance of host-named subdirectories (./foo.com for instance).
- Basic template support is accomplished by adding a template_data.json file containing the data object used in parsing. See the test_fixtures directory for an example.
- Using ApacheBench, she manages over 3900 requests/sec (template parses) on OSX, over 2000 on a Linux Micro EC2 instance on Amazon.

> This server is a very mimimal implementation of HTTP. There are tons of features that *should* be implemented before it would be suitable for production environments.

Hit me up (matthew.mcneely@gmail.com) for the project files if you want to try it in Eclipse.

Made available via an MIT-style license (see the ./LICENSE file)—use this as you see fit.
