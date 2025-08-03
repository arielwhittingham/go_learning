QUICKSTART

go mod init mod_name    #creates a go.mod file

#at this level, add a main.go file with the package declared as "main"

```
package main

import (
    "fmt"
    "test_module/folder"
)

func main() {
    fmt.Println(folder.MyFunc())
}
```
additional go files should go into `folders/new_go_file.go`

Where in the `new_go_file.go` the package declaration is:

```
package folder

func MyFunc() string {
    return "from folder"
}
```

```
    func functionName(parameter1 type1, parameter2 type2) returnType {
        // Function body - code to be executed
        return value // Optional, if returnType is specified
    }

```






-----


🧱 1. Basic Folder Structure for a Go Project
Go is opinionated and encourages a clean and minimal structure, especially when you're starting out:

go
Copy
Edit
my-go-project/
├── go.mod           <- Module definition file
├── main.go          <- Entry point if this is an executable
└── (optional) other .go files
When you start writing reusable packages or organizing larger projects:

rust
Copy
Edit
my-go-project/
├── go.mod
├── main.go
├── cmd/             <- Different CLI entry points (e.g., `cmd/appname/main.go`)
├── pkg/             <- Reusable code intended for external use
├── internal/        <- Reusable code intended for internal use only
├── test/            <- Test files or helper test data


🧾 2. Creating and Running a Go Script
Here’s a minimal runnable Go script:

main.go:

```
package main

import "fmt"

func main() {
    fmt.Println("Hello, Go!")
}
```


```
go run main.go
```
Or build and run:
```
go build -o myprogram
./myprogram
```
📦 3. Setting Up a Go Module (required outside $GOPATH)
Inside your project directory:

```
go mod init github.com/yourusername/my-go-project
```
This creates a go.mod file which tracks dependencies and the module path. If you're just playing around locally, you can use:

bash
Copy
Edit
go mod init my-go-project
🧪 4. Structure of an Exercism or HackerRank Go File
A typical main.go from those platforms:
go
Copy
Edit
package main

import "fmt"

func main() {
    // your code goes here
}
Or for unit tests (common on Exercism):

go
Copy
Edit
package hello

func HelloWorld() string {
    return "Hello, World!"
}
With a test file:


package hello

import "testing"

func TestHelloWorld(t *testing.T) {
    expected := "Hello, World!"
    if result := HelloWorld(); result != expected {
        t.Errorf("Expected %q but got %q", expected, result)
    }
}
To run the tests:
bash
Copy
Edit
go test
🧭 5. How to Identify Proper Go Boilerplate (Checklist)
Check	What to Look For
✅ package main (or package xyz)	Required at the top of each .go file. main is special: it's the entry point.
✅ func main() present	Only in the main package. This is where Go starts execution.
✅ All files in same folder share a package	All .go files in a directory must declare the same package.
✅ go.mod file (if you're using modules)	Required for modern Go projects (Go 1.11+). Run go mod init.
✅ Correct file naming	Stick to lowercase and underscores (hello_world.go, not HelloWorld.go).
✅ Valid imports	No unused imports; Go will refuse to compile.
✅ Proper formatting	Run go fmt . or gofmt to format your code idiomatically.

🧠 Bonus: Some Handy Go CLI Commands
Command	What it does
go run file.go	Compile and run a script directly
go build	Build a binary of your module
go test	Run all tests
go mod tidy	Clean up and sync go.mod and go.sum
go fmt .	Auto-format Go code in current dir

🚦Final Notes
When working on small challenges like HackerRank/Exercism:

You often don’t need go.mod if it’s a single file

Just ensure package main is set and main() exists

For testing challenges, go test assumes files end in _test.go
