import { random, getCookie, makeRequest } from "./common.js";
import { SERVER_HOST, SERVER_PORT } from "./consts.js";

// if no jwt, go to login page
if (getCookie() == undefined || getCookie() == "")
  window.location.assign("./login.html");

// add a user gopher on the screen
function addGopher(id, name) {
  const gopher = document.createElement("div"); // div that contains everything
  gopher.id = id;
  gopher.classList.add("gopher");

  const messageBox = document.createElement("div");
  messageBox.classList.add("messageBox");
  const messageBubble = document.createElement("div");
  messageBubble.id = `${id}-message`;
  messageBubble.classList.add("messageBubble");
  messageBox.appendChild(messageBubble);
  gopher.appendChild(messageBox);

  const usernameBox = document.createElement("div"); // contains user's name
  usernameBox.innerHTML = name;
  gopher.appendChild(usernameBox);

  const sprite = document.createElement("img"); // gopher image
  sprite.style.height = "150px";
  sprite.style.width = "110px";
  sprite.src = "./res/gophersprite.png";
  gopher.appendChild(sprite);

  // use the calculated size of the canvas and the gopher to find a good location on the screen
  const canvas = document.getElementById("canvas");
  const { width, height } = canvas.getBoundingClientRect();
  canvas.appendChild(gopher);

  const { width: gWidth, height: gHeight } = gopher.getBoundingClientRect();
  const maxX = width - gWidth;
  const maxY = height - gHeight;

  gopher.style.left = `${random(0, maxX)}px`;
  gopher.style.top = `${random(0, maxY)}px`;
}

// message hiding timeouts - so i can clear them
const timeoutIds = {};

// show message bubble above user that sent it
function handleMessage(message) {
  const { userId, message: text } = message;

  if (timeoutIds[userId]) {
    clearTimeout(timeoutIds[userId]); // if a timeout was set, clear it
  }

  const messageBox = document.getElementById(`${id}-message`);
  messageBox.innerHTML = text;
  messageBox.style.display = "block";

  // https://developer.mozilla.org/en-US/docs/Web/API/setTimeout
  // wait 5 seconds then hide the message
  timeoutIds[userId] = setTimeout(() => {
    messageBox.style.display = "none";
    delete timeoutIds[userId]; // delete this timeout id
  }, 5 * 1000);
}

// for testing only. TODO remove
document.getElementById("addgopher").addEventListener("click", (e) => {
  // add a gopher
  addGopher(random(0, 1000), "go gopher");
});

// handle send message button TODO
document.getElementById("sendMessage").addEventListener("click", (e) => {
  makeRequest("POST", "/broadcast", { message }, () => {});
});

// get all users to draw initial gophers on the screen
makeRequest("GET", "/allUsers", null, (status, body) => {
  // TODO add all users to backend
  body.users.forEach((user) => {
    addGopher(user.id, user.username);
  });
});

console.log("opening socket connection...");
// create websocket client
const socket = new WebSocket(
  `ws://${SERVER_HOST}:${SERVER_PORT}/subscribe?Authorization=${encodeURI(
    getCookie()
  )}`
);

// handle received data
socket.onmessage = (event) => {
  const data = JSON.parse(event.data);
  console.log(data);

  switch (data.messageType) {
    case "text":
      handleMessage(data);
      break;
    case "joined":
      addGopher(data.senderId, data.message);
      break;
    case "left":
      // remove gopher
      document.getElementById(data.senderId).remove();
      break;
    default:
      console.log("received unknown message type:", event.data);
  }
};
