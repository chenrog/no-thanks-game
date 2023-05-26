const ws = new WebSocket("ws://localhost:8080/ws");

ws.onopen = function (event) {
    console.log("WebSocket opened");
    ws.send("gameInfoRequest")
};

ws.onmessage = function (event) {
    console.log("Received message: " + event.data);
    const gameState = document.getElementById("gameState");
    const p = document.createElement("p");
    p.innerHTML = event.data.replace(/\n/g, "<br>");
    gameState.innerHTML = p.innerHTML;
};

ws.onerror = function (event) {
    console.log("WebSocket error: " + event);
};

ws.onclose = function (event) {
    console.log("WebSocket closed");
};

const passButton = document.getElementById("pass");
const takeButton = document.getElementById("take");

passButton.addEventListener("click", function (event) {
    ws.send("pass");
});

takeButton.addEventListener("click", function (event) {
    ws.send("take");
});