# Chapter 1 - Create the "Hello Go & Gin" app
## Objective
In this chapter, we will:
- Initialize the project and the Go module
- Create the "Hello Go & Gin" Go app
- Learn how to debug Go using VS Code


## Required tools before start
Install the following tools if they are not already installed:

- [Go](https://golang.org/doc/install). Make sure the `go version` command from a terminal window shows the version of go.  
- [Visual Studio Code](https://code.visualstudio.com/) as the code editor and debugger.
- [Git tools](https://git-scm.com/downloads) 

## Initialize the project and the Go module

Create a new folder for your project. For example, "~/src/codehouse-2022-prework" on Mac/Linux, or "c:\src\codehouse-2022-prework" on Windows.

Open a terminal window and run the following commands to initialize the project for git and Go module. 

```
cd ~/src/codehouse-2022-prework
git init
go mod init codehouse-2022-prework
```

## Create the API code

In the same terminal window, run the following commands: 
```
touch main.go
code .
```
The empty file main.go was created and VS Code window was launched. From the file browser on the left pane, open the file main.go in the editor, and type in the following code:

```
package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.String(200, "Hello Go and Gin")
	})
	r.Run(":8090")
}
```
Note this project will use the port 8090 to run the web server. Choose a different port (e.g., 8092) if it conflicts with an existing program.

Save the code in main.go, and run the following command to downlaod the Gin framework referenced by the code:
```
go get github.com/gin-gonic/gin
```
Run the following comamnd to launch the web server:
```
go run .
```
The computer will ask you to give permission to run the program on the specified port. Give the permission and you will see the terminal window outputs the following messages:

```
[GIN-debug] [WARNING] Creating an Engine instance with the Logger and Recovery middleware already attached.

[GIN-debug] [WARNING] Running in "debug" mode. Switch to "release" mode in production.
 - using env:   export GIN_MODE=release
 - using code:  gin.SetMode(gin.ReleaseMode)

[GIN-debug] GET    /                         --> main.main.func1 (3 handlers)
[GIN-debug] Listening and serving HTTP on :8090
```
Launch a web browser, and navigate to url: `http://localhost:8090`. You should see the Hello message. It's time to commit your work in the git repo

## Debug Go using VS Code

Stop the API in the terminal if it's still running (type in Ctrl + C). In VS Code, make sure main.go is the current window in editor, launch the debugger by pressing F5 or from the Run -> Start Debugging menu. Install the following **Go extensions for VS Code** if VS Code asks you:

- go (the extension) ( >= v0.26.0)
- delve ( >= v1.0.6)
- go-outline
- dlv


After installing the extensions, go back to main.go in the editor window, launch the debugger again. VS Code should enter debug mode, and the DEBUG CONSOLE tab on the bottom of the window should show debugging outputs. Now set a breakpoint at the code line:

		c.String(200, "Hello Go and Gin")

VS Code should stop at this line when you call the API from the web browser again.


Congrats! You have written an API in Go+Gin, and you are armed with some powerful tools for developing in Go.

