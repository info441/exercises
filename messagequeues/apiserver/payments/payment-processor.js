var amqp = require('amqplib/callback_api');

const ObjectId = require("mongodb").ObjectId;
const QUEUE_NAME = 'NewOrders';
const APPROVED_ORDERS_EXCHANGE = 'ApprovedOrders';
const ORDER_INFO_EXCHANGE = 'OrderInfo';
let burgersChannel;


/**
	TODO: This is a payment processer worker that handles NewOrder events sent by the API Server.

	Implement the code that does the following:

	1) Approve 60% of new orders and reject the rest to simulate whether credit cards are going "through" or not.
	2) If the order is approved, then send an "ApprovedOrder" event to the ApprovedOrders exchange, and broadcast on OrderInfo exchange. 
	3) If not, send a "DeniedOrder" event to the OrderInfo exchange. 
*/
function handleNewOrder(msg) {
	console.log(" [x] Received %s", msg.content.toString());

	newOrderMsg = JSON.parse(msg.content.toString());

	const shouldApprove = Math.random() <= 0.6 ? true : false;

	// this statement makes sure that Rabbit deletes the message and doesn't redeliver
	burgersChannel.ack(msg);
}


/**
	Initialize the following:
	- Setup connections to our Rabbit instance. 
	- Declare exchanges and queues
*/
amqp.connect('amqp://localhost:5672', function(err, conn) {
	if (err) {
		console.log("Failed to connect to Rabbit Instance from payment processor.");
		process.exit(1);
	}

	conn.createChannel(function(err, ch) {
		if (err) {
			console.log("Failed to create NewOrdersChannel from payment processor");
			process.exit(1);
		}

		burgersChannel = ch;
		burgersChannel.assertQueue(QUEUE_NAME, {durable: true});
		burgersChannel.prefetch(1);

		// Set up the worker callback for NewOrders queue
		burgersChannel.consume(QUEUE_NAME, handleNewOrder, {noAck: false});

		// Create the exchange onto which we publish ApprovedOrders events (if needed)
		burgersChannel.assertExchange(APPROVED_ORDERS_EXCHANGE, 'fanout', {durable: false});

		// Create the exchange onto which we publish OrderInfo events (if needed)
		burgersChannel.assertExchange(ORDER_INFO_EXCHANGE, 'topic', {durable: false});
	});
});



