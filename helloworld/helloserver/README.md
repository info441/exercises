# Web Server for helloworld
Approach this exercise by completing each of the steps described below:
- Since this repository came with starter code, you should get into the habit of seeking to understand what your 
provided code does. For this exercise, you're provided with `identicon.go`. Comments have been included to 
help you read and understand the code. NOTE: You DO NOT need to add code to this file.
- Read in environment variable for network address on which server should listen. 
    -  This address is in the form `host:port`, where `host` is the IP address or host name to bind to, and `port` is the
     port number to listen on.

- Create a new `mux` and register two handler functions for the following resource path:
    - `/` (root)
        - For this handler, the provided request will contain a query string with parameter name being `name`.
          An example would be `GET /?name=Everyone`, where the query string is `name=Everyone`, with `name` being the 
          parameter name and `Everyone` being the parameter value. 
        - If no parameter value is provided, default to the value `World`.
        - Add the `Access-Control-Allow-Origin` header to the `ResponseWriter` to allow any origin. See 
        https://gist.github.com/bdinh/a3196c532624ecbabe201a5bb1d1f5f1
        - Respond to the client with a formatted string representing the expression `Hello, {string}!`, where `{string}`
        represents the parameter value from the query string.
    - `/identicon/`
        - Since this resource path ends with a `/`, we can use unique identifiers after the initial path to provide 
        additional information in our request. Here, we'll look to provide a name that we can use within our handler for 
        this resource path. An example would be `GET /identicon/Everyone`, where `Everyone` becomes our unique identifier.
        Grab this last element of the path. See https://gist.github.com/bdinh/8ec3aa764753d687afce5bc9f5eef9ea
        - If no identifier is provided, default to the value `World`.
        - Add the `Access-Control-Allow-Origin` header to the `ResponseWriter` to allow any origin. See above
        - Add the `Content-Type` header to `image/png` to the `ResponseWriter` to let the client know the type of 
        response being sent back.
        - Call the `identicon` function with the given identifier as the argument and return the response back to the 
        client. See https://golang.org/pkg/image/png/#Encode

- Start the web server, and log any errors that can occur. 

### Compile, Install and Run
Execute the following inside your terminal. Below assumes you've used `ADDR` as your environment variable name. Please
change it accordingly to what you used inside `main.go`.
```bash
export ADDR=localhost:4000
go install 
helloserver
```



If you are lost, please refer to the following tutorial, [Go Web Servers](https://drstearns.github.io/tutorials/goweb/).