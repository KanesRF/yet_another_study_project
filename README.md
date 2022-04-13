# yet_another_study_project
This is a "blank" server, that has only JWT (or kind of) authentication with bare minimum use of PostgreSQL. Also there is no frontend, so for testing curl is used.
# How to build and run

Just use 
```sh
go build -o server ./server/server.go
./server/server
```

# Examples

There is simple example of request, used by curl in file 
[request example](test.sh)