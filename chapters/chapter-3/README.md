# Chapter 3 - Scaffold the "Hello Vue" app

## Objective
We will add UI to the Go project using the Vue.js framework. 
- Scaffold a Vue app named todo-vue
- Unify the Go and Vue apps
- Debug the Vue app


## Required tools before start
- [Install nodejs and npm](https://nodejs.org/en/download/)
- Install vue-cli using npm by running command:

	`npm install -g @vue/cli`


## Scaffold the "Hello Vue" app

Run the following command in a terminal:
```
cd ~/src/codehouse-2022-prework
vue create todo-vue
```
Select `Default ([Vue 2] babel, eslink)` from the command-line options using the up/down arrow keys and the enter key. We choose Vue 2.x instead of Vue 3.x because many current learning materials are still for Vue 2, and most concept2 from Vue 2 still apply to Vue 3.

The last command created a subfolder named `vue-todo` under your project folder. Next run these commands in the terminal:
```
cd todo-vue
npm run serve
```
This launched another web app on local port 8081. Open a browser and navigate to http://localhost:8081. The "Welcome to Your Vue.js App" page will display. This is a good time to take a detour to go over [**a Vue tutorial**](https://www.taniarascia.com/getting-started-with-vue/) to get familiar with the Vue framework.


## Unify the Go and Vue apps
Note that the Vue app is running on port 8081 because we launched inside its own folder. Now we are going to embed the Vue project in the Go project, and serve requests to both parts from the same port (8090).

In the previous terminal, first stop the Vue server launched from the last step (ctrl+C), then run the following commands:
```
rm -rf dist; npm run build
```
This compiled the Vue project and saved the result in the `dist` folder, which contains the static contents (html/css/image/js) for browser to run. In order to serve this static content from the Go app, first run the following commands in terminal:
```
cd ~/src/codehouse-2022-prework
go get github.com/gin-contrib/static
```
Next, in VS Code, edit the main.go file to add a new import:

`"github.com/gin-contrib/static" `

and add a new route url ('/') before the API endpoint handlers:
  
`	r.Use(static.Serve("/", static.LocalFile("./todo-vue/dist", false))) 
`

The new main.go should look like:

```
package main

import (
	"net/http"
	"strconv"

	"github.com/gin-contrib/static"
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

func main() {
	todos = append(todos, Todo{Id: GetNextId(), Value: "CodeHouse", DueDate: "7/31/2022"})

	r := gin.Default()
	r.Use(static.Serve("/", static.LocalFile("./todo-vue/dist", false)))
	r.GET("/api/todos", GetTodos)
	r.POST("/api/todos", PostTodo)
	r.DELETE("/api/todos/:id", DeleteTodo)
	r.Run(":8090")
}
```

Use the `go run .` comamnd to run the Go app, then open the browser to navigate to http://localhost:8090. You should now see the same "Welcome to Your Vue.js App" page. 

Now the Go app and Vue app are unified,  we will modify the Vue app to consume the TODO API in the next chapter.

## Debug the Vue app
Debugging the Vue app is very easy with the Vue.js devtools extension. To install:
- [Chrome](https://chrome.google.com/webstore/detail/vuejs-devtools/nhdogjmejiglipccpnnnanhbledajbpd?hl=en)
- [Firefox](https://addons.mozilla.org/en-US/firefox/addon/vue-js-devtools/)

In the borwser, open the browser's developer tool, and refresh the page `http://lcaolhost:8090`. In the **Source** tab of Chrome's dev tool or the **Debugger** tab of Firefox's dev tool, navigate to *Webpack->src* folder and open App.Vue. You should be able to set a breakpoint in the script section of App.vue and debug the code. 


