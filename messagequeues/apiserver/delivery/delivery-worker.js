var amqp = require('amqplib/callback_api');

const WORKER_NAME = process.env.NAME;

if (!WORKER_NAME) {
	console.log("Please name this worker and try again.");
	process.exit(1);
}


const APPROVED_ORDER_EXCHANGE_NAME = 'ApprovedOrders';
const ORDER_INFO_EXCHANGE = 'OrderInfo';
let burgersChannel;


// These two variables allow multiple delivery-workers to coordinate without double-booking
let isDeciding; 
let approvedOrdersCache = {}; 


function generateOrderUpdate(id, updateType, messageStr) {
	return {
		"type": updateType,
		"orderId": id,
		"messageStr": messageStr,
		"acceptee": WORKER_NAME,
		"timestamp": (new Date()).toUTCString()
	};
}


/**
	Handle delivery requests. 
*/
function handleDeliveryRequests(msg) {
	msg = JSON.parse(msg.content.toString());

	if (!isDeciding) {
		// sleep anywhere from 0 - 5 seconds
		isDeciding = setTimeout( () => {
			console.log("Trying to decide about order %s...", msg.orderId)
			if ( (msg.orderId in approvedOrdersCache) && approvedOrdersCache[msg.orderId].acceptee ){
				// reset the isDeciding flag to free this delivery worker
				isDeciding = null;
			} else {
				// broadcast that this delivery worker won the order 
				burgersChannel.publish(ORDER_INFO_EXCHANGE, 
					msg.orderId, 
					new Buffer(JSON.stringify(generateOrderUpdate(msg.orderId, "OrderAccepted", "Order accepted by worker: " + WORKER_NAME)))
				);

				setTimeout( () => {
					// broadcast that this delivery worker is on the way in 0-3 seconds
					burgersChannel.publish(ORDER_INFO_EXCHANGE, 
						msg.orderId, 
						new Buffer(JSON.stringify(generateOrderUpdate(msg.orderId, "InProgress", "Delivery on the way via worker: " + WORKER_NAME)))
					);				

					setTimeout( () => {
						// broadcast that this delivery worker has arrived
						burgersChannel.publish(ORDER_INFO_EXCHANGE, 
							msg.orderId, 
							new Buffer(JSON.stringify(generateOrderUpdate(msg.orderId, "Arrived", "Order delivered by worker: " + WORKER_NAME)))
						);				

						// reset the isDeciding flag to free this delivery worker
						isDeciding = null;
					}, Math.random() * 3000);
				}, Math.random() * 3000);
			} 
		}, Math.random() * 5000); 
	} else {
		// worker is busy deciding
	}
}


/**
	When an order update arrives, store the update into the approvedOrdersCache.
	
	approvedOrdersCache is a key value store where:
	 - Key: orderId
	 - Value: last order update msg

	delivery-workers use approvedOrdersCache to check whether they should accept the order or not. 

	There are the following order update types:
	 - accepted
	 - delivery on its way 
	 - delivered
*/
function handleOrderUpdate(msg) {
	console.log(" [handleOrderUpdate] %s", msg.content.toString());

  	const payload = JSON.parse(msg.content.toString());

	approvedOrdersCache[payload.orderId] = payload;
}



/**
	- Setup connections to our Rabbit instance. 
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
		burgersChannel.assertExchange(APPROVED_ORDER_EXCHANGE_NAME, 'fanout', {durable: false});
		burgersChannel.assertExchange(ORDER_INFO_EXCHANGE, 'topic', {durable: false});

		// subscribe to newly approved order events, so the delivery worker can try to win the job
		burgersChannel.assertQueue('', {exclusive: true}, function(err, q) {
			console.log(" [*] Waiting for messages in %s in APPROVED_ORDERS exchange", q.queue);
			burgersChannel.bindQueue(q.queue, APPROVED_ORDER_EXCHANGE_NAME, '');

			burgersChannel.consume(q.queue, handleDeliveryRequests, {
				noAck: true
			});
		});


		// start subscribing to order info events
		burgersChannel.assertQueue('', {exclusive: true}, function(err, q) {
			console.log(" [*] Waiting for messages in OrderInfo exchange");
			// Use '#' as topic key to receive all messages 
			burgersChannel.bindQueue(q.queue, ORDER_INFO_EXCHANGE, '#');

			burgersChannel.consume(q.queue, handleOrderUpdate, {
				noAck: true
			});
		});
	});
});



