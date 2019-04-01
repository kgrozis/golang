
// go packages:
//   abstract out and aggregate common logic into sharable code units
//   go source files in a directory is considered to be a package
//   there are no naming requirements for a go package.  name based on pkg purpose

// multi-File packages:
//   can logical group multiple go source files together as package
//   each file must be declared with the same package name

// naming packages:
//   each package has unique fully qualified import path
//   can have as many packages as needed
//   no limit to package depth
// globally unique namespaces:
//   start pkg with unique name like organization name
// add context to path:
//   use import path to add context to the name of package
//   start generic and get more specific, left to right
//   package depth of 1 is acceptable but should be descriptive
// use short names:
//   brevity counts