Run server and client separately

run server
 go run server.go

run client
 go run client.go

The service would add the two supplied request parameter and return the sum
c = a + b


The "gopls" command is not available. Run "go install -v golang.org/x/tools/gopls@latest" to install.

fix by running
go clean -modcache
go install -v golang.org/x/tools/gopls@latest

DO brew install gopls when the above failed.

USE goenv as your go version manager.

brew update
brew install goenv

Follow this https://github.com/go-nv/goenv/blob/master/INSTALL.md


Generating client and server code
$ protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    routeguide/route_guide.proto