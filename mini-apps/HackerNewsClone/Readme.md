# uses hacker rank client
# pulls 30 top stories
# caches top stories
# removes race conditions
# uses go routines
# uses Mutex to provide access to cache
# update cache with ticker and use reciever on a struct type 
# usage: go run main.go or go run -race main.go  
# @todo: docker and makefile; possible deployment and service file
# problem: figure out how to name go package, that's causing build failure
