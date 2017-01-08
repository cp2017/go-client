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
[smoke-test:L0:0s] + go run main.go http --iterations 10 --delay 1000 --url http://localhost:8080
[smoke-test:L1:0s] 2017/01/08 12:29:18 Fetching 'http://localhost:8080' 10 times with a delay of 1000
[smoke-test:L2:0s] 2017/01/08 12:29:18 [  1] 200 OK
[smoke-test:L3:1s] 2017/01/08 12:29:19 [  2] 200 OK
[smoke-test:L4:2s] 2017/01/08 12:29:20 [  3] 200 OK
[smoke-test:L5:3s] 2017/01/08 12:29:21 [  4] 200 OK
[smoke-test:L6:4s] 2017/01/08 12:29:22 [  5] 500 Internal Server Error
[smoke-test:L7:5s] 2017/01/08 12:29:23 [  6] 200 OK
[smoke-test:L8:6s] 2017/01/08 12:29:24 [  7] 200 OK
[smoke-test:L9:7s] 2017/01/08 12:29:25 [  8] 200 OK
[smoke-test:L10:8s] 2017/01/08 12:29:26 [  9] 200 OK
[smoke-test:L11:9s] 2017/01/08 12:29:27 [ 10] 500 Internal Server Error
[smoke-test:L12:10s] 2017/01/08 12:29:28 Status: 200 Count: 8
[smoke-test:L13:10s] 2017/01/08 12:29:28 Status: 500 Count: 2
```
