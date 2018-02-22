package main

import "fmt"

func main() {
	fmt.Println("Hi there!")
}

/*

1. How do we run the code in our project?
2. What does "package" mean?
3. What does "import fmt" mean?
4. What's that "func" thing?
5. How is the main.go file organized?

--

1. go run [filename.go]
>> go build: compules a bunch of go source code files
>> go run: compiles and executes one or two files
>> >> build just compiles, does not execute (type go build, then ls - can do ./[filename] to then execute)
>> go fmt: formats all the code in each file in the current directory
>> go install: compiles and installs a package
>> go get: downloads the raw source code of someone else's package
>> go test: runs any tests associated with the current project

2. package == project == workspace
any associated files must declare the package name for the package they belong to
two types of pkgs:
executable: makes a "runnable" file
reusable: code used as "helpers"; good place to put reusable logic (modules)
name of package that you use determines which type - "main" makes executable package; any other name would not spit out an executable file
package main MUST have a func called "main"

3. gives our package access to code written in another package
>> give main access to all the functionality in the fmt package
>> fmt is standard library package incl'd with go, used to print out info to terminal, etc.
golang.org/pkg for docs

4. functions, hooray! declaration, name, argument/s, body

5. always organized like this:
>> 1. package declartion
>> 2. import other packages
>> 3. declare functions, tell Go to do things

*/
