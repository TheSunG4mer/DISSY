# Structuring a Go project
There are a number of small details that may trip you up when structuring your first Go project.
If you don't actively structure your code it is easy to get off to a bad start and just end up with one large blob of code.
We will go on a brief tour of this `example-project` and explain how we can build up such a project.

## How to reconstruct this project

We start in an empty folder `example-project` at the root of our workspace.
The first thing we want to do is create the module where our main function will live.
This will be the primary entry point of our program.

1. Make a new folder called `client`
2. In your terminal go to the `client` directory and initialise the module:

        cd client
        go mod init client

3. Now you may create your first go file `client.go`

    ```golang
    package main

    func main() {
        println("Hello world!")
    }
    ```

    __Note: usually the package can be the same as the folder name, but for `client.go` to be executable it must instead be `main` here__

4. To run our program, from a terminal in the `client` directory we run:

        go run ./client.go

### More files in the same package
At some point you may want to start separating your code into smaller chunks.
The most straight forward way to do this is to add another file `./client_extended` to the `./client` directory.

Our file structure now looks like this:

    example-project
    └──  client
        ├── client.go
        ├── client_extended.go
        └── go.mod

- As `client.go` and `client_extended.go` are in the same folder they should also be in the same package, which in this case is `main`.

- Any functions in `client_extended.go` will be in the scope of `client.go` and may therefore be called as if they were in the same file.

We could for example add a function which stops the program when given an error:
```golang
package main

import "log"

func check(err error) {
    if err != nil {
        log.Fatal(err)
    }
}
```
and extend the `main` function in `client.go` with
```golang
func main() {
    //...
	strNum := "5"
	num, err := strconv.Atoi(strNum)
	//use the check function from 'client extended
	check(err)
    println(num * num)
}

```
To run our program, from a terminal in the `client` directory we may now run:

    go run ./
    go run ./client.go ./client_extended.go

This is a little inconvenient, as we need to specify the files which we wish to include. 
It may be better to add new files to a new package as described in the next section.
### Files in a new package
Having all functions in the same scope is likely to lead to high coupling within the codebase.
To combat this we may add a new package with a separate scope.
To do this we can add a new subdirectory, which we will call `helper` with a new file `helper.go`

Our file structure now looks like this:

    example-project
    └──  client
        ├── client.go
        ├── client_extended.go
        ├── go.mod
        └── helper
            └── helper.go

We can now add some code to `helper.go`
```golang
package helper

// This function needs to start with a 
// Capital letter to be visible to `client.go` 
func Add(a, b int) int {
	return a + b
}

// This function cannot be accessed from outside the helper package
func add(a, b int) int {
	return a  + b
}
```
Now we can use this functionality in `client.go`, to do this we must import helper by adding the import `"client/helper"`, where `client` specifies the module and the path `/helper` specifies the package. 
```golang
package main

import (
	"fmt"
	"client/helper"
)

func main() {
	res := helper.Add(5, 5)
	fmt.Printf("We got: %d \n", res)
}
```

## Adding another module (advanced - feel free to skip)
If you wish to work across multiple modules at once you can do the following:

In the `example-project` directory run 
    
    go work init ./client

This will set up a workspace, which you can add more modules to.
You can run `client.go` from the project root using 

    go run ./client

We can add and initialise a new module `util`, so that our file structure looks like this

    example-project
    ├── client
    │   ├── client.go
    │   ├── client_extended.go
    │   ├── go.mod
    │   └── helper
    │       └── helper.go
    ├── go.work
    └── util
        └── go.mod

Then we may add the `util` module to our workspace with:

    go work use ./util

Things from the `util` module may now be imported to `client.go` as we did for `helper`:

```golang
package main

import (
	"fmt"
	"client/helper"
	"util"
)

func main() {
	//Call the `Add` function from the helper package
	res := helper.Add(5, 5)
	fmt.Printf("We got: %d \n", res)

	// Call the `Mult` function from the util module
	res2 := util.Mult(5, 5)
	fmt.Printf("And then: %d \n", res2)
}
```

## Testing
Go has a built-in testing library, to see an example of a test see `math_test.go`.
There are a few details to be aware of for testing:
- Test files must have names which end with `_test`
- All test functions must be named `TestXxx`, where `Xxx` is an arbitrary name starting with a __capital__ letter.

For useful guides to testing in Go try looking through one of the following
- https://blog.jetbrains.com/go/2022/11/22/comprehensive-guide-to-testing-in-go/#WritingFuzzTests
- https://www.digitalocean.com/community/tutorials/how-to-write-unit-tests-in-go-using-go-test-and-the-testing-package


## Error handling
Go does not force you to handle errors like java does with its Exceptions.
This may tempt you to ignore errors by writing,

```golang
numString, _ := reader.ReadString('\n')
```
instead of 
```golang
numString, err := reader.ReadString('\n')
if err != nil {
    ...
}
```
This is a bad idea!
It is extremely helpful to detect errors early during development, and avoid strange unpredictable behaviour.

## More useful links

- https://www.digitalocean.com/community/tutorials/how-to-use-go-modules
- https://go.dev/doc/tutorial/workspaces#create-the-workspace
