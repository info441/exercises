// import server libraries
const express = require('express');
const path = require('path');
const app = express();
const PORT = process.env.PORT || 4000;
const server = require('http').createServer(app);

// set up middleware
const bodyParser = require('body-parser');
app.use(bodyParser.json()); // parses json and puts them in body
app.use(bodyParser.urlencoded({ extended: true }));  // parses queries and puts them in body


// import rabbit 
let amqp = require('amqplib/callback_api');
const queueName = 'NewOrders';
let burgerChannel; 
const ORDER_INFO_EXCHANGE = 'OrderInfo';
const ObjectId = require("mongodb").ObjectId;


// helper function to create an order update message
function generateOrderUpdate(id, updateType, messageStr) {
	return {
		"type": updateType,
		"orderId": id,
		"messageStr": messageStr, 
		"timestamp": (new Date()).toUTCString()
	};
}


app.get("/", ( req, res ) => {
	res.sendFile(path.join(__dirname, 'client.html'));
});



/**
	TODO: 

	Implement a POST /order that does the following:

	- Validate whether there's a request body that contains itemName and price fields (done for you)
	- If the request is valid, create a "NewOrder" message (done already for you), and send it to two places:
		- NewOrders queue (as a persistent message)
			Read this article: https://www.rabbitmq.com/tutorials/tutorial-two-javascript.html 
		- OrderInfo topic exchange 
			Read this article for more info: https://www.rabbitmq.com/tutorials/tutorial-five-javascript.html 
			Use the generateOrderUpdate helper function to create a message in the format that the GUI can understand

	- If all fails, check the solutions branch. 
*/

app.post("/order", (req, res) => {
	console.log(req.body);

	if (req.body && req.body.itemName && req.body.price) {
		console.log("Placing an order")
		// To place an order, we do the following:
		// 1) Create a new_order message
		// 2) Put the message in the NewOrdersQueue

		const orderId = ObjectId() + "";

		const newOrderMsg = {
			"type": "NewOrder",
			"orderId": orderId,
			"details": req.body,
			"timestamp": (new Date()).toUTCString()
		};

		burgerChannel.sendToQueue(
			queueName, 
			new Buffer(JSON.stringify(newOrderMsg)),
			{persistent: true}  // make sure messages are stored until ack'ed
		);

		burgerChannel.publish(ORDER_INFO_EXCHANGE, 
			orderId,
			new Buffer(JSON.stringify(generateOrderUpdate(orderId, "NewOrder", "Order Processing...")))
		);

		return res.status(201).json({
			"message": "Order initiated. Tracking id: " + newOrderMsg.orderId
		});
	} else {
		return res.status(401).json({
			"message": "Invalid Request"
		});
	}
});


/**
	Server Initialization:

	Order:
		1) Connect to Rabbit and set up burgerChannel, NewOrders queue
		2) Connect to Mongo
		3) Start listening to HTTP traffic
*/
amqp.connect('amqp://localhost:5672', function(err, conn) {
	if (err) {
		console.log("Failed to connect to Rabbit Instance from API Server.");
		process.exit(1);
	}

	conn.createChannel(function(err, ch) {
		if (err) {
			console.log("Failed to create NewOrdersChannel from API Server");
			process.exit(1);
		}

		ch.assertQueue(queueName, {durable: true});
		burgerChannel = ch;

		server.listen(PORT, function () {
			console.log(`Listening on port ${PORT}`);
		});
	});
});




