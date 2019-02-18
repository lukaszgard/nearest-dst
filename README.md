# nearest-dst
This is a simple REST API application used to obtain and sort nearest/fastest destinations.
Supported routes:
- sample index welcome page:
`
GET /
`
- obtain nearest location based on provided coordinates:
`
GET /routes?src={coordinates}&dst={coordinates}
`
{coordinates}: string format for 'src' and 'dst' - {longitude},{latitude}
* more than one destination 'dst' must be provided.

# Installation
`
$ go get github.com/lukaszgard/nearest-dst
`

# Usage

## runing application:
`
$ nearest-dst
`

- or with providing a path to log file and providing port (def. 8080):

`
$ nearest-dst -log-file=nearest.log -port=8080
`

## sample request:
`
$ curl -i 'localhost:8080/routes?src=13.388860,52.517037&dst=13.397634,52.529407&dst=13.428555,52.523219'
HTTP/1.1 200 OK
Content-Type: application/json
Date: Mon, 18 Feb 2019 23:37:08 GMT
Content-Length: 238

{
        "source": "13.388860,52.517037",
        "routes": [
                {
                        "destination": "13.397634,52.529407",
                        "duration": 277.5,
                        "distance": 1935.5
                },
                {
                        "destination": "13.428555,52.523219",
                        "duration": 470,
                        "distance": 4127.7
                }
        ]
}
`

