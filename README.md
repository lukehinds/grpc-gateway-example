# grpc-gateway-example

Build protobuf / gateway

`make create`

Run the server

```
go run main.go serve
2023/03/28 20:20:49 Serving gRPC on 0.0.0.0:8080
2023/03/28 20:20:49 Serving gRPC-Gateway on http://0.0.0.0:8090
```

Use the client

```
go run main.go client
token:"1234567890"
```

```
curl -X POST -k http://localhost:8090/login -d '{"username": "luke", "password": "password"}'
{"token":"1234567890"}
```
