
const express = require('express');
const app = express();
const PORT = process.env.PORT || 5000;
const server = require('http').createServer(app);
const io = require('socket.io')(server);
let socket;

let amqp = require('amqplib/callback_api');

const ORDER_INFO_EXCHANGE = 'OrderInfo';
let burgersChannel;


/*
	Socket
*/
io.on('connection', function (sock) {
	console.log("Socket Connection Made");
	socket = sock;
});



function handleWebSocketPublishing(msg) {
	socket.emit('order_events', 
		JSON.parse(msg.content.toString())
	);	
}



/**
	Set up the following:
	- Connection to the RabbitMQ
	- OrderInfoChannel (if needed)
	- SocketIO connection 
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

		burgersChannel = ch;
		burgersChannel.assertExchange(ORDER_INFO_EXCHANGE, 'topic', {durable: false});

		// spin up HTTP server
		server.listen(PORT, function () {
			console.log(`Listening on port ${PORT}`);

			// start subscribing to order info events
			burgersChannel.assertQueue('', {exclusive: true}, function(err, q) {
				console.log(" [*] Waiting for messages");
				// Use '#' as topic key to receive all messages 
				burgersChannel.bindQueue(q.queue, ORDER_INFO_EXCHANGE, '#');

				burgersChannel.consume(q.queue, handleWebSocketPublishing, {
					noAck: true
				});
			});
		});
	});
});
