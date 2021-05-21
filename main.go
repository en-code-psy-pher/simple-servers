package main 

import (
	api "simpleserver/pkg/server"
)

// Test with (uses pkg/handers/test.csv)
//		cd pkg/handlers 
//		go test -v

// Run with
//		go run .
// Send requests with:
//	curl -F 'file=@/path/matrix.csv' "localhost:8080/echo"
//	curl -F 'file=@/path/matrix.csv' "localhost:8080/invert"
//	curl -F 'file=@/path/matrix.csv' "localhost:8080/flatten"
//	curl -F 'file=@/path/matrix.csv' "localhost:8080/sum"
//	curl -F 'file=@/path/matrix.csv' "localhost:8080/product"

func main() {
    api.Run()
}
