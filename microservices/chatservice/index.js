

const express = require("express");
const app = express(); 

const port = process.env.PORT;
const instanceName = process.env.NAME;

//add JSON request body parsing middleware
app.use(express.json());

app.get("/v1/chat", (req, res) => {
	res.json({
		"message": "Hello from " + instanceName 
	});
});

app.listen(port, "", () => {
        //callback is executed once server is listening
        console.log(`server is listening at http://:${port}...`);
	console.log("port : " + port);
	console.log("host : " + instanceName);
});
