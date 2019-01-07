# Web Server for zip

Approach this exercise by completing each of the steps described below for their respective files:

`./models/zip.go`
- This file contains starter code that you should get familiar with. You'll be using the types, `Zip`, `ZipSlice`, `ZipIndex`,
defined here, as well as the `LoadZips` function. Ask questions if you're confuse. 

`./main.go`:
- Open the `zips.csv` file inside your `./data` directory and report any errors that may occur
    - You can use `log.Fatal` to report any error as we want our server to terminate if we are unable to read in our data
    - See https://golang.org/pkg/os/#Open on how to open files in Go
- After opening `zips.csv` , you should of received a pointer to that file, `*File`. Pass this pointer to your `LoadZips` 
function inside `./models`. This function accepts two arguments, you can use `42613` for the second argument. Remember to
constantly handle any errors present. 
- Close the stream to your `*File` after you're done using it.
- Recall that `LoadZips` returns a `ZipSlice`. Iterate through this slice in order to build a map (`ZipIndex`) of city 
names to `*Zip`. Make sure the keys to your dictionary are all lower-cased.
- Read in environment variable for network address on which server should listen.
    - This address is in the form host:port, where host is the IP address or host name to bind to, and port is the port
     number to listen on.
- Create a new mux and register two handler functions for the following resource path:
    - `/` 
        - Assign the `RootHandler` to this resource path. Implementation details are listed below.
    - `/zips/city/`
        - Assign the `ZipIndexHandler` to this resource path. Implementation details are listed below.
        - Recall that you'll need to use `Handle` instead of `HandleFunc` since you'll be creating your own `Handler`.
    - Start the web server, and log any errors that can occur.

`./handlers/constants.go`
- Make sure to use these defined constants

`./handlers/zips.go`
- Here you'll need to define a new `Handler` in order to get access to some extra data. 
    - See https://drstearns.github.io/tutorials/goweb/#sechandlers 
- Define a new `ZipIndexHandler` struct that contains one field, `index` which is of type `ZipIndex`.
- Create a constructor for your `ZipIndexHandler` that accepts one argument and initializes the `index` field with the
given argument and returns a pointer to the `ZipIndexHandler` struct.
    - See https://gist.github.com/bdinh/cd350fcd36847fdea0a3d5d2a0f948d5 for OOP patterns in Go
- Implement the `ServeHTTP` method on your `ZipIndexHandler` struct.
    - Grab the last element of your request's resource path as that will represent the city of interest.
    - Make sure to convert the string to lower case 
    - See https://gist.github.com/bdinh/8ec3aa764753d687afce5bc9f5eef9ea on how to grab last element in path.
    - Add the `Content-Type` header to `application/json` to the `ResponseWriter` to let the client know the type of response being sent back.
    - Add the `Access-Control-Allow-Origin` header to the `ResponseWriter` to allow any origin. See https://gist.github.com/bdinh/a3196c532624ecbabe201a5bb1d1f5f1
    - Get the zips for the requested city of interest from your `ZipIndex` and encode it as `json` to the `ResponseWriter`.
- Implement a `RootHandler` that handles the `/` resource path which returns the given string to the client, `"Try requesting /zips/city/seattle"`

### Compile, Install and Run
Execute the following inside your terminal. Below assumes you've used `ADDR` as your environment variable name. Please
change it accordingly to what you used inside `main.go`.
```bash
export ADDR=localhost:4000
go install 
zipserver
```