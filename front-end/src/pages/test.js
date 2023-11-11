const WebSocket = require('ws');

const socket = new WebSocket('ws://localhost:9999/ws');

socket.onmessage = (event) => {
    // Handle received messages from the server (RabbitMQ)
    console.log(event.data);
    // const message = JSON.parse(event.data);
    // Update your React components or state with the received message
};

socket.onclose = (event) => {
    // Handle WebSocket connection closed
};