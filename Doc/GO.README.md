# GO

## CLI Arguments

Verify golang installation and version:

```
root@c6dcd5d8092c:/go/src/app/01# go version
go version go1.8.7 linux/amd64
```

Compile and run a golang problem (Compiling a go program creates a permanent 
executable file, and compile and run creates a temporary executable file):

```
Compile:
root@c6dcd5d8092c:/go/src/app/01# go build 01__01_hello-world.go

Run:
root@c6dcd5d8092c:/go/src/app/01# ./01__01_hello-world.go
Hello, World!

Compile and Run:
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

Downloading and compiling go packages

```
Download:
root@2ff69b7b88d7:/go/src/app# go get -v github.com/mactsouk/go/simpleGitHub
github.com/mactsouk/go (download)
github.com/mactsouk/go/simpleGitHub

Source Code:
root@2ff69b7b88d7:/go/src/app# ls -l /go/src/github.com/mactsouk/go/simpleGitHub/
total 4
-rw-r--r-- 1 root root 66 Mar 26 14:18 simpleGitHub.go

Compiled Code:
root@2ff69b7b88d7:/go/src/app# ls -l /go/pkg/linux_amd64/github.com/mactsouk/go/simpleGitHub.a 
-rw-r--r-- 1 root root 1122 Mar 26 14:18 /go/pkg/linux_amd64/github.com/mactsouk/go/simpleGitHub.a

Remove Intermediate Packages:
root@2ff69b7b88d7:/go/src/app# go clean -i -v -x /go/src/github.com/mactsouk/go/simpleGitHub/
cd /go/src/github.com/mactsouk/go/simpleGitHub
rm -f simpleGitHub.test simpleGitHub.test.exe
rm -f /go/pkg/linux_amd64/github.com/mactsouk/go/simpleGitHub.a
```
