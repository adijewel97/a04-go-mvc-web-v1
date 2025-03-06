---
package main

import "fmt"

func main() {
	fmt.Println("Server running on port 8080...")
}

go mod init a04-go-mvc-web-v1
go: creating new go.mod: module a04-go-mvc-web-v1
go: to add module requirements and sums:
        go mod tidy

go mod init a03-my-go-project
go get github.com/godror/godror
go get github.com/gorilla/mux
