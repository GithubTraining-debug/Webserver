## We specify the base image we need for our
## go application
FROM golang:1.12.0-alpine3.9
## We create an /webserver directory within our
## image that will hold our application source
## files
RUN mkdir /webserver
## We copy everything in the root directory
## into our /webserver directory
ADD . /webserver
## We specify that we now wish to execute 
## any further commands inside our /webserver
## directory
WORKDIR /webserver
## we run go build to compile the binary
## executable of our Go program
RUN go build -o main .
## Our start command which kicks off
## our newly created binary executable
EXPOSE 2000
CMD ["/webserver/main"]