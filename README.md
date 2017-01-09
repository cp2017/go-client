# go-client [![Build Status](http://ec2-54-194-144-141.eu-west-1.compute.amazonaws.com/api/badges/cp2017/go-client/status.svg)](http://ec2-54-194-144-141.eu-west-1.compute.amazonaws.com/cp2017/go-client)
GOLANG client to interact with marketplace

## DroneCI

```
drone exec --build.event "tag"                                                                                                                                                                                              git:(master|✚4…
Running Matrix job #0
[testhttpd:L0:0s] 2017/01/08 11:29:13 Start Testserver. Verbose:false, Seq:[200 200 200 200 500]
[unit-test:L0:0s] + ./test.sh
[unit-test:L1:0s] + govendor fetch +missing
[unit-test:L2:0s] + echo '> govendor remove +unused'
[unit-test:L3:0s] + govendor remove +unused
[unit-test:L4:0s] > govendor remove +unused
[unit-test:L5:0s] + echo '> govendor update +local'
[unit-test:L6:0s] + govendor update +local
[unit-test:L7:0s] > govendor update +local
[unit-test:L8:1s] + echo '> govendor sync +external'
[unit-test:L9:1s] + govendor sync +external
[unit-test:L10:1s] > govendor sync +external
[unit-test:L11:1s] + '[' '!' -d resources/coverity ']'
[unit-test:L12:1s] + go test -cover -coverprofile=go-client.cover
[unit-test:L13:2s] testing: warning: no tests to run
[unit-test:L14:2s] + COVER_FILES=go-client.cover
[unit-test:L15:2s] ++ find . -maxdepth 1 -type d
[unit-test:L16:2s] ++ egrep -v '(\.$|\.git|vendor|bin|resources|deploy)'
[unit-test:L17:2s] + for x in '$(find . -maxdepth 1 -type d |egrep -v "(\.$|\.git|vendor|bin|resources|deploy)")'
[unit-test:L18:2s] + go test -cover -coverprofile=resources/coverity/./cmd.cover ./cmd
[unit-test:L19:3s] + COVER_FILES='go-client.cover resources/coverity/./cmd.cover'
[unit-test:L20:3s] + coveraggregator -o coverage-all.out go-client.cover resources/coverity/./cmd.cover
[unit-test] exit code 0
[smoke-test:L0:0s] + go run main.go http --iterations 10 --delay 100 --url http://localhost:8080
[smoke-test:L1:1s] 2017/01/09 21:03:18 Fetching 'http://localhost:8080' 10 times with a delay of 100; validate:''
[smoke-test:L2:1s] 2017/01/09 21:03:18 [  1] 200 OK
[smoke-test:L3:1s] 2017/01/09 21:03:18 [  2] 200 OK
[smoke-test:L4:1s] 2017/01/09 21:03:18 [  3] 200 OK
[smoke-test:L5:1s] 2017/01/09 21:03:18 [  4] 200 OK
[smoke-test:L6:1s] 2017/01/09 21:03:18 [  5] 500 Internal Server Error
[smoke-test:L7:1s] 2017/01/09 21:03:18 [  6] 200 OK
[smoke-test:L8:1s] 2017/01/09 21:03:18 [  7] 200 OK
[smoke-test:L9:2s] 2017/01/09 21:03:18 [  8] 200 OK
[smoke-test:L10:2s] 2017/01/09 21:03:18 [  9] 200 OK
[smoke-test:L11:2s] 2017/01/09 21:03:19 [ 10] 500 Internal Server Error
[smoke-test:L12:2s] 2017/01/09 21:03:19 Status: 200 Count: 8
[smoke-test:L13:2s] 2017/01/09 21:03:19 Status: 500 Count: 2
[smoke-test:L14:2s] + go run main.go http --iterations 10 --delay 100 --url http://localhost:8080 --validate 200:80,500:20
[smoke-test:L15:3s] 2017/01/09 21:03:20 Fetching 'http://localhost:8080' 10 times with a delay of 100; validate:'200:80,500:20'
[smoke-test:L16:3s] 2017/01/09 21:03:20 [  1] 200 OK
[smoke-test:L17:3s] 2017/01/09 21:03:20 [  2] 200 OK
[smoke-test:L18:3s] 2017/01/09 21:03:20 [  3] 200 OK
[smoke-test:L19:3s] 2017/01/09 21:03:20 [  4] 200 OK
[smoke-test:L20:3s] 2017/01/09 21:03:20 [  5] 500 Internal Server Error
[smoke-test:L21:3s] 2017/01/09 21:03:20 [  6] 200 OK
[smoke-test:L22:4s] 2017/01/09 21:03:20 [  7] 200 OK
[smoke-test:L23:4s] 2017/01/09 21:03:20 [  8] 200 OK
[smoke-test:L24:4s] 2017/01/09 21:03:21 [  9] 200 OK
[smoke-test:L25:4s] 2017/01/09 21:03:21 [ 10] 500 Internal Server Error
[smoke-test:L26:4s] 2017/01/09 21:03:21 Status: 500 Count: 2
[smoke-test:L27:4s] 2017/01/09 21:03:21 Status: 200 Count: 8
[smoke-test:L28:4s] + go run main.go http --iterations 8 --delay 100 --url http://localhost:8080 --validate 200:80,500:20 --negate-validate
[smoke-test:L29:5s] 2017/01/09 21:03:22 Fetching 'http://localhost:8080' 8 times with a delay of 100; validate:'200:80,500:20'
[smoke-test:L30:5s] 2017/01/09 21:03:22 [  1] 200 OK
[smoke-test:L31:5s] 2017/01/09 21:03:22 [  2] 200 OK
[smoke-test:L32:5s] 2017/01/09 21:03:22 [  3] 200 OK
[smoke-test:L33:5s] 2017/01/09 21:03:22 [  4] 200 OK
[smoke-test:L34:5s] 2017/01/09 21:03:22 [  5] 500 Internal Server Error
[smoke-test:L35:5s] 2017/01/09 21:03:22 [  6] 200 OK
[smoke-test:L36:6s] 2017/01/09 21:03:22 [  7] 200 OK
[smoke-test:L37:6s] 2017/01/09 21:03:22 [  8] 200 OK
[smoke-test:L38:6s] 2017/01/09 21:03:23 Status: 200 Count: 7
[smoke-test:L39:6s] 2017/01/09 21:03:23 Status: 500 Count: 1
[smoke-test:L40:6s] 2017/01/09 21:03:23 Expected 200:80, found 88
[smoke-test:L41:6s] 2017/01/09 21:03:23 Validation was false, but as 'negate-validate' was set turn it around
[smoke-test] exit code 0
```
