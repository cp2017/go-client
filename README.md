# go-client [![Build Status](http://ec2-54-194-144-141.eu-west-1.compute.amazonaws.com/api/badges/cp2017/go-client/status.svg)](http://ec2-54-194-144-141.eu-west-1.compute.amazonaws.com/cp2017/go-client)
GOLANG client to interact with marketplace

Able to run tests against:

- `--url` Plain http url
- `--swagger-path` path to swagger.json file from which the url is extracted
- `--swagger-id` IPFS if to swagger.json file from which the url is extracted

## DroneCI

```
drone exec --build.event "tag"
Running Matrix job #0
[smoke-test:L0:0s] + go run main.go http --iterations 10 --delay 100 --url http://localhost:8081
[smoke-test:L1:4s] 2017/01/15 12:34:07 Fetching 'http://localhost:8081' 10 times with a delay of 100; validate:''
[smoke-test:L2:4s] 2017/01/15 12:34:07 [  1] 200 OK
[smoke-test:L3:4s] 2017/01/15 12:34:07 [  2] 200 OK
[smoke-test:L4:4s] 2017/01/15 12:34:07 [  3] 200 OK
[smoke-test:L5:4s] 2017/01/15 12:34:07 [  4] 200 OK
[smoke-test:L6:4s] 2017/01/15 12:34:07 [  5] 500 Internal Server Error
[smoke-test:L7:4s] 2017/01/15 12:34:07 [  6] 200 OK
[smoke-test:L8:4s] 2017/01/15 12:34:08 [  7] 200 OK
[smoke-test:L9:5s] 2017/01/15 12:34:08 [  8] 200 OK
[smoke-test:L10:5s] 2017/01/15 12:34:08 [  9] 200 OK
[smoke-test:L11:5s] 2017/01/15 12:34:08 [ 10] 500 Internal Server Error
[smoke-test:L12:5s] 2017/01/15 12:34:08 Status: 200 Count: 8
[smoke-test:L13:5s] 2017/01/15 12:34:08 Status: 500 Count: 2
[smoke-test:L14:5s] + go run main.go http --iterations 10 --delay 100 --url http://localhost:8081 --validate 200:80,500:20
[smoke-test:L15:9s] 2017/01/15 12:34:12 Fetching 'http://localhost:8081' 10 times with a delay of 100; validate:'200:80,500:20'
[smoke-test:L16:9s] 2017/01/15 12:34:12 [  1] 200 OK
[smoke-test:L17:9s] 2017/01/15 12:34:12 [  2] 200 OK
[smoke-test:L18:9s] 2017/01/15 12:34:12 [  3] 200 OK
[smoke-test:L19:9s] 2017/01/15 12:34:12 [  4] 200 OK
[smoke-test:L20:9s] 2017/01/15 12:34:12 [  5] 500 Internal Server Error
[smoke-test:L21:9s] 2017/01/15 12:34:13 [  6] 200 OK
[smoke-test:L22:9s] 2017/01/15 12:34:13 [  7] 200 OK
[smoke-test:L23:10s] 2017/01/15 12:34:13 [  8] 200 OK
[smoke-test:L24:10s] 2017/01/15 12:34:13 [  9] 200 OK
[smoke-test:L25:10s] 2017/01/15 12:34:13 [ 10] 500 Internal Server Error
[smoke-test:L26:10s] 2017/01/15 12:34:13 Status: 200 Count: 8
[smoke-test:L27:10s] 2017/01/15 12:34:13 Status: 500 Count: 2
[smoke-test:L28:10s] + go run main.go http --iterations 8 --delay 100 --url http://localhost:8081 --validate 200:80,500:20 --negate-validate
[smoke-test:L29:14s] 2017/01/15 12:34:17 Fetching 'http://localhost:8081' 8 times with a delay of 100; validate:'200:80,500:20'
[smoke-test:L30:14s] 2017/01/15 12:34:17 [  1] 200 OK
[smoke-test:L31:14s] 2017/01/15 12:34:17 [  2] 200 OK
[smoke-test:L32:14s] 2017/01/15 12:34:17 [  3] 200 OK
[smoke-test:L33:14s] 2017/01/15 12:34:17 [  4] 200 OK
[smoke-test:L34:14s] 2017/01/15 12:34:17 [  5] 500 Internal Server Error
[smoke-test:L35:14s] 2017/01/15 12:34:17 [  6] 200 OK
[smoke-test:L36:14s] 2017/01/15 12:34:17 [  7] 200 OK
[smoke-test:L37:14s] 2017/01/15 12:34:18 [  8] 200 OK
[smoke-test:L38:15s] 2017/01/15 12:34:18 Status: 200 Count: 7
[smoke-test:L39:15s] 2017/01/15 12:34:18 Status: 500 Count: 1
[smoke-test:L40:15s] 2017/01/15 12:34:18 Expected 200:80, found 88
[smoke-test:L41:15s] 2017/01/15 12:34:18 Validation was false, but as 'negate-validate' was set turn it around
[smoke-test:L42:15s] + go run main.go ipfs --iterations 10 --delay 100 --validate 200:80,500:20 --swagger-path resources/services/hello-service/swagger.json
[smoke-test:L43:19s] 2017/01/15 12:34:22 Start IPFS Monitor
[smoke-test:L44:19s] QmeX24C4mcpvTJo64DREzpWEpHKgodY7BahG7oK3JPS4jB
[smoke-test:L45:19s] 2017/01/15 12:34:22 Fetching 'http://localhost:8081' 10 times with a delay of 100; validate:'200:80,500:20'
[smoke-test:L46:19s] 2017/01/15 12:34:22 [  1] 200 OK
[smoke-test:L47:19s] 2017/01/15 12:34:22 [  2] 500 Internal Server Error
[smoke-test:L48:19s] 2017/01/15 12:34:22 [  3] 200 OK
[smoke-test:L49:19s] 2017/01/15 12:34:22 [  4] 200 OK
[smoke-test:L50:19s] 2017/01/15 12:34:22 [  5] 200 OK
[smoke-test:L51:19s] 2017/01/15 12:34:22 [  6] 200 OK
[smoke-test:L52:19s] 2017/01/15 12:34:22 [  7] 500 Internal Server Error
[smoke-test:L53:19s] 2017/01/15 12:34:23 [  8] 200 OK
[smoke-test:L54:20s] 2017/01/15 12:34:23 [  9] 200 OK
[smoke-test:L55:20s] 2017/01/15 12:34:23 [ 10] 200 OK
[smoke-test:L56:20s] 2017/01/15 12:34:23 Status: 500 Count: 2
[smoke-test:L57:20s] 2017/01/15 12:34:23 Status: 200 Count: 8
[smoke-test:L58:20s] + go run main.go ipfs --iterations 10 --delay 100 --validate 200:80,500:20 --swagger-id QmeX24C4mcpvTJo64DREzpWEpHKgodY7BahG7oK3JPS4jB
[smoke-test:L59:24s] 2017/01/15 12:34:27 Start IPFS Monitor
[smoke-test:L60:24s] 2017/01/15 12:34:27 Fetching 'http://localhost:8081' 10 times with a delay of 100; validate:'200:80,500:20'
[smoke-test:L61:24s] 2017/01/15 12:34:27 [  1] 200 OK
[smoke-test:L62:24s] 2017/01/15 12:34:27 [  2] 500 Internal Server Error
[smoke-test:L63:24s] 2017/01/15 12:34:27 [  3] 200 OK
[smoke-test:L64:24s] 2017/01/15 12:34:27 [  4] 200 OK
[smoke-test:L65:24s] 2017/01/15 12:34:27 [  5] 200 OK
[smoke-test:L66:24s] 2017/01/15 12:34:27 [  6] 200 OK
[smoke-test:L67:24s] 2017/01/15 12:34:27 [  7] 500 Internal Server Error
[smoke-test:L68:24s] 2017/01/15 12:34:28 [  8] 200 OK
[smoke-test:L69:25s] 2017/01/15 12:34:28 [  9] 200 OK
[smoke-test:L70:25s] 2017/01/15 12:34:28 [ 10] 200 OK
[smoke-test:L71:25s] 2017/01/15 12:34:28 Status: 200 Count: 8
[smoke-test:L72:25s] 2017/01/15 12:34:28 Status: 500 Count: 2
[smoke-test] exit code 0```
