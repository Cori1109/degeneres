//go:generate rm -rf data helpers server main_test.go
//go:generate bash -c "go run `go env GOPATH`/src/github.com/Cori1109/degeneres/main.go generate -f {{.ProtoFilePath}} -o `pwd`"

package main
