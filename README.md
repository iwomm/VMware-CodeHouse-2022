# VMware CodeHouse 2022

This project will serve as guidance for attendees and mentors of VMware CodeHouse 2022.

The goal of VMware CodeHouse 2022 will be to create an application that furthers STEM eduction and/or Diversity & Inclusion.

This project contains several incremental chapters with hands-on coding exercises. A chapter contains a few small steps, with each step introduces a bite-size new concept or a tool. The expected code for each chapter is provided in the `codehouse-2022-prework` sub-folder of the chapter. The finished product is a TODO web application written in Go, Gin and Vue. 

There are two main components in the finished app:

- **Server side** -  a Rest API written in Go (the base language) and Gin (a web framework for Go).
- **Client side** -  a Javascript application running in the browser that interacts with the API. Vue is used in this project because it's relatively easier to understand for someone not familiar with any front-end frameworks. Students are encouraged to replace it with other frameworks they feel comfortable with, such as AngularJS or React. 

Table of content:

- Chapter 1 - Create the "Hello Go & Gin" app
- Chapter 2 - Create the Todo API in Go & Gin
- Chapter 3 - Scaffold the "Hello Vue" app
- Chapter 4 - Display the todo list and make a Vue component
- Chapter 5 - Post a new todo item

Given the step-by-step nature, you can jump right into chapter one and start learning by coding. Use the recommended readings below if you find some parts need more introduction. These materials were selected for their brevity and clarity in order to save time.     

## Recommended reading

To make sure you are comfortable with the concepts and tools we will be using, please read through the following documents well in advance:

- Introduction to **Go**. [Learning Go](https://www.miek.nl/go/) is one online free book featuring numerous execises. There are [many resources for learning Go](https://github.com/dariubs/GoBooks).
- Intruduction to **Rest API**. [This page ](https://www.sitepoint.com/rest-api/)covers many aspects of Rest API with simple exercises.
- Introduction to **Gin**. [This tutorial ](https://blog.logrocket.com/how-to-build-a-rest-api-with-golang-using-gin-and-gorm/)builds a simple CRUD API using Gin. It uses Sqlite for database. 
- Introduction to **Vue**. [A turtorial for Vue like this one](https://www.taniarascia.com/getting-started-with-vue/) provides a good overview of the concepts and will make you feel more comfortable working with Vue.
  
## Completion of pre-req tasks

Once you have completed these tasks, make a pull request to this repository with a Markdown file in this format:

`[initials of your first and last name].md`
```
I've finished reading through the pre-req docs and ran the final result on my laptop!
```
