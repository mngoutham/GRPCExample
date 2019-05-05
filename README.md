# A Sample gRPC Web Server

## Things to do

- [ ] - A simple golang gRPC web server which returns pong response for ping request.

- [ ] - Dockerize above service and run as a docker container.

- [ ] - Use go dep for dependency management.

- [ ] - Use go tool chain for profiling. https://www.youtube.com/watch?v=uBjoTxosSys&t=405s

- [ ] - Add thrift interface for another simple request. 

- [ ] - Add uber fx dependency injection framework.
 
- [ ] - Add uber logging library called zap and metrics library called tally.
 
- [ ] - Update business logic to make use of channels & mutexes. 

- [ ] - Deploy this on Google cloud.



## Commands used for getting this project up and running

### To generate grpc code given a protobuf file

1. Install `protoc` from
```
https://github.com/protocolbuffers/protobuf/releases
```
2. Install `protoc-gen-go` using 
```
go get -d -u github.com/golang/protobuf/protoc-gen-go
go install github.com/golang/protobuf/protoc-gen-go
```
3. Generate grpc code using
```
protoc -I .  todo.proto  --go_out=plugins=grpc:../../build/gen/
```