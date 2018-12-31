# GO

## CLI Arguments

Verify golang installation and version:

```
root@c6dcd5d8092c:/go/src/app/01# go version
go version go1.8.7 linux/amd64
```

Compile and run a golang problem:

```
root@c6dcd5d8092c:/go/src/app/01# go run 01__01_hello-world.go
Hello, World!
```

Built-in golang test

```
go test .
```

Golang documentation tool.  Outputs to http://localhost:6001

```
root@c6dcd5d8092c:/go/src/app/01# godoc -http=":6001"
2018/12/31 19:14:49 lstat /usr/local/go/apple-touch-icon-precomposed.png: no such file or directory
2018/12/31 19:14:49 lstat /usr/local/go/apple-touch-icon.png: no such file or directory
```
