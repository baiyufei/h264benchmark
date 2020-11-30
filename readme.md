```bash
go test -run TestCurrentH264Reader -cpuprofile currentcpu.out
go tool pprof -http=":8081" ./h264benchmark.test.exe currentcpu.out

go test -run TestNewH264Reader -cpuprofile newcpu.out
go tool pprof -http=":8081" ./h264benchmark.test.exe newcpu.out
```
