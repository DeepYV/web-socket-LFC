# websocket-simple-feed-chat
simple web-sockets concepts used to make live feed and chat between two client 

How to Run
---
1. go run . 
2. Open inspect in browser:
2. Run command in console `let socket = new WebSocket("ws://localhost:8000/connect")`
3. Run command in console `socket.onmessage=(event) => {console.log("received from the server:",event.data)}`
4. repeat these step in another browser tab 
5. from one tab write `socket.send("your msg")`
6. if you want to subscribe to feed run this command  `let socket = new WebSocket("ws://localhost:8000/orderBook")`
7. Run command in console `socket.onmessage=(event) => {console.log("received from the server:",event.data)}`
8. Demo live feed