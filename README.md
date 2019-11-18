# godebridaria

set GOARCH=amd64
set GOOS=linux
go tool dist install -v pkg/runtime
go install -v -a std


go env
$env:GOOS = "linux"
$env:GOARCH = "arm"
go build