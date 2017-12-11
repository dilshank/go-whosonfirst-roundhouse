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

### wof-roundhoused

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

### wof-roundhouse-repod

```
./bin/wof-roundhouse-repod /usr/local/mapzen/whosonfirst-data*
2017/07/03 15:50:47 start indexing whosonfirst-data-venue-us-ca at 2017-07-03 15:50:47.808249449 +0000 UTC
2017/07/03 15:50:47 start indexing whosonfirst-data at 2017-07-03 15:50:47.809042249 +0000 UTC
2017/07/03 15:50:47 start indexing whosonfirst-data-venue-ca at 2017-07-03 15:50:47.810867862 +0000 UTC
2017/07/03 15:51:27 time to index whosonfirst-data-venue-ca: 39.240279572s
2017/07/03 15:51:32 time to index whosonfirst-data: 44.511392089s
2017/07/03 15:51:51 time to index whosonfirst-data-venue-us-ca: 1m3.898184345s
2017/07/03 15:51:51 time to index all: 1m3.898692508s
2017/07/03 15:51:51 indexed 2577907 pairs
```

## Docker

[Yes](Dockerfile). Well, at least for `wof-roundhoused`. I'm not sure how (or where) to define Dockerfiles for multiple services yet. If you know, [I'd love to hear about it](https://github.com/whosonfirst/go-whosonfirst-roundhouse/issues).

```
docker build -t wof-roundhoused .
docker run -p 6161:8080 -e BASE='https://example.com/' -e HOST='0.0.0.0' wof-roundhoused

curl localhost:6161/1234
<a href="https://example.com/123/4/1234.geojson">See Other</a>.
```

## See also

* https://github.com/whosonfirst/go-whosonfirst-utils
