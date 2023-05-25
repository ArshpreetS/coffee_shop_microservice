# Golang_microservice
Learning to build microservices using Golang


## Concepts
1. Dependency Injection

This concept basically allows us to put dependency as a variable. For better understanding, lets take an example of the file "handlers/hello.go". Now in this file, we have created a new object called hello which will serve the request coming in "/" path. To create a new handler we created a function called *NewHello*, and we are passing in a logger **l**. Logger basically allows us to log the events. Now it can be the "stdout" or a file maybe, so this can be changed as per our requirements.