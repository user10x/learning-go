
FROM golang:1.16

# build the code
RUN mkdir /app
WORKDIR /app
COPY . ./
#RUN go build -o hnc main.go
RUN go run main.go

# Expose the application on port 8081
EXPOSE 4000

# Set the entry point of the container
CMD ["./hnc"]