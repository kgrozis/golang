
// go build [<package import path>]
//   can explicitly specify import path or omit
//   import path can be fully qualified or relative paths
//
//   go build ./...
//     build all pkgs and subpkgs
//   go build .
//     build pkg in dir
//   go build 06.2/volt
//     build pkg in dir path

// installing a package
//   generates a usable artifact
//   leaves a compiled object files
//   saves object in $GOPATH/pkg dir with .a extension
//
//   go install ./...
//     install all pkgs and subpkgs
//   go install .
//     install pkg in dir
//   go install 06.2/volt
//     install pkg in dir path