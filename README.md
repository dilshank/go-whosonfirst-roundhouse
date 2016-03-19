# go-whosonfirst-roundhouse

Given a Who's On First ID resolve its absolute URL.

## Caveats

Currently the default root for all absolute URLs is `https://whosonfirst.mapzen.com/data/`.

It is not possible to specify alternatives. Yet.

## Build

The easiest thing is just to use the handy `build` target in the Makefile. This will fetch all the dependencies and build all the tools, described below. Like this:

```
$> make build
```

## Tools

### wof-roundhouse

```
$> ./bin/wof-roundhouse 85784831 102020079
https://whosonfirst.mapzen.com/data/857/848/31/85784831.geojson
https://whosonfirst.mapzen.com/data/102/020/079/102020079.geojson
```

### wof-roundhouse-server

```
./bin/wof-roundhouse-server -h
Usage of ./bin/wof-roundhouse-server:
  -host string
    	The hostname to listen for requests on (default "localhost")
  -port int
    	The port number to listen for requests on (default 8080)
```

For example:

```
$> curl -v http://localhost:8080/85784831
*   Trying ::1...
* connect to ::1 port 8080 failed: Connection refused
*   Trying 127.0.0.1...
* Connected to localhost (127.0.0.1) port 8080 (#0)
> GET /85784831 HTTP/1.1
> Host: localhost:8080
> User-Agent: curl/7.43.0
> Accept: */*
> 
< HTTP/1.1 303 See Other
< Location: https://whosonfirst.mapzen.com/data/857/848/31/85784831.geojson
< Date: Sat, 19 Mar 2016 06:32:00 GMT
< Content-Length: 90
< Content-Type: text/html; charset=utf-8
< 
<a href="https://whosonfirst.mapzen.com/data/857/848/31/85784831.geojson">See Other</a>.
```

## See also

* https://github.com/whosonfirst/go-whosonfirst-utils
