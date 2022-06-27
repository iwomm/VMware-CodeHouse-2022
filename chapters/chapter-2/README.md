# Chapter 2 - Create the Todo API in Go & Gin

## Objective

To continue with the project created in chapter 1, we will create API for managing TODO items. The API will have endpoints for

- GET /api/todos - get the list of all todo items
- POST /api/todos - add a new todo item
- DELETE /api/todos/:id - delete a todo item by its id

## Required tools before start

We are going to use [curl](https://www.postman.com/) for API testing. Although other tools such as Postman can also be used for testing, curl is a command-line tool so it's easy to share the command text with other people, such as using Slack or in this document. To install curl:

- Windows - the [Git tools](https://git-scm.com/downloads) comes with the curl command, which can be run from [the Git Bash terminal window](https://gitforwindows.org/)
- Mac - use [Homebrew](https://formulae.brew.sh/formula/curl)
- Linux - use the package manager such as `apt install curl`

## Add a Go struct type for the todo data

The API will send and receive data of todo items using JSON format. For cleaner code organization, we will use a separate Go file to define data structure of a single todo item.

Create a new file named **todo.go** next to main.go, and type in the following code in todo.go:

```go
package main

type Todo struct {
 ID      int    `json:id`
 Value   string `json:"value"`
 DueDate string `json:"due_date"`
}
```

In addition to the three data members in each todo item, the `json:...` tells the JSON formatter how to map each field in the go data structure to the field in the JSON object, and vice versa. For example, the following raw JSON input provided through API call

```js
{
 "id": 10, "value": "some-value", "due_date": "Jun 31 2022"
}
```

will become

```go
Todo{ID: 10, Value: "some-value", DueDate: "Jun 31 2022" }
```

## Add API for getting all todo items

Modify the code in main.go to this state:

```go
package main

import (
 "net/http"
 "github.com/gin-gonic/gin"
)

var nextId int = 0
var todos []Todo

func GetNextId() int {
 value := nextId
 nextId++
 return value
}

func GetTodos(c *gin.Context) {
 c.JSON(http.StatusOK, gin.H{"list": todos})
}

func main() {
 todos = append(todos, Todo{Id: GetNextId(), Value: "CodeHouse", DueDate: "7/31/2022"})

 r := gin.Default()
 r.GET("/api/todos", GetTodos)
 r.Run(":8090")
}
```

The code declares a Go slice (`var todos []TODO`) of Todo type. The main() function starts by adding the first item to this slice, and it maps the `GET /api/todo` API endpoint to the `GetTodos` method as the API endpoint handler. The handler method returns a JSON object using [the Context object's JSON() method](https://pkg.go.dev/github.com/gin-gonic/gin#Context.JSON). **Note you can use the link to navigate the Gin's help documents, and search for helps for other Gin methods**.

Launch the server with the `go run .` command. From another terminal window, run the following command, which will output the result as:

```sh
curl -X GET http://localhost:8090/api/todos
{"list":[{"Id":0,"value":"CodeHouse","due_date":"7/31/2022"}]}
```

## Add API for uploading a new todo item

Rest API uses the HTTP POST method to receive a request for uploading a new item in the API resource collection (e.g., todos).

Create a new handler method named `PostTodo()` in main.go:

```go
func PostTodo(c *gin.Context) {
 var item Todo
 if err := c.ShouldBindJSON(&item); err != nil {
  c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
  return
 }

 item.Id = GetNextId()
 todos = append(todos, item)
 c.String(http.StatusCreated, c.FullPath()+"/"+strconv.Itoa(item.Id))
}
```

The ShouldBindJSON() method takes the HTTP request body data, assumes the data is in JSON format, and attemts to bind (translate) the data to the `item` varable, which was declared as Todo type. If anything went wrong during binding, the returned error object will not be nil, and the method returns HTTP status code 400 Bad Request to the gin context with error details. After successful binding, the `PostTodo` method assigns the new item with an id and appends the new todo item to the `todos` slice, and then returns the HTTP status code 201 Created.  

Next, connect the new handler in the main() mthod by adding this line:

`r.POST("/api/todos", PostTodo)`

To test the new API, re-run the new API code using `go run .`, and in another terminal window, type in the following command:

`curl -X POST -d '{"value":"complete schoolwork","due_date":"7/29/2022"}' http://localhost:8090/api/todos`

The curl command uploads a JSON data body (`-d '{the json}'`) using the POST method (`-X POST`) to the API server. Successful response should be a url value for the new todo item:

`/api/todos/1`

Now you can check what items are on the server by running the above curl -x GET command again:

```sh
curl -X GET http://localhost:8090/api/todos                                                                                     
{"list":[{"Id":0,"value":"CodeHouse","due_date":"7/31/2022"},{"Id":1,"value":"complete schoolwork","due_date":"7/29/2022"},{"Id":2,"value":"complete schoolwork","due_date":"7/29/2022"},{"Id":3,"value":"complete schoolwork","due_date":"7/29/2022"}]}

```

The above output example has 4 todo items because I run the same curl POST command 3 times.

## Add the API for deleting a todo item

Add the following method as the DELETE API handler to main.go:

```go
func DeleteTodo(c *gin.Context) {
 idString := c.Param("id")

 if id, err := strconv.Atoi(idString); err == nil {
  for index := range todos {
   if todos[index].Id == id {
    todos = append(todos[:index], todos[index+1:]...)
    c.Writer.WriteHeader(http.StatusNoContent)
    return
   }
  }
 }

 c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
}
```

The code searches for the todo item with the specified id, and the item is deleted with the line of code if the item is found:

`todos = append(todos[:index], todos[index+1:]...)`

and the following new API endpoint to the main() method:

`r.DELETE("/api/todos/:id", DeleteTodo)`

Run the new API code, and test the DELETE API with command:

`curl -X DELETE http://localhost:8090/api/todos/[actual id for testing]`

With existing id's, the response to the DELETE request has no body. When no todo item with the specified id exists, the response should be `{"error":"Invalid Id"}`.

## Food for thought

Delete the last item in the list, then post a new one, the id of the new item will not assume the id that was just deleted. This means although the slice of todo items was compacted after removing an item, the next id keeps growing. This is by design - what is the reason for not reusing the id?
