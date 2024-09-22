import net from "net";
import { random, getCookie, AjaxPost } from "./common.js";
import { SERVER_HOST, SERVER_PORT, SERVER_URL } from "./consts.js";

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
timeoutIds = {};

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
  AjaxPost("/broadcast", { message }, () => {});
});

// get all users to draw initial gophers on the screen
AjaxGet("/allUsers", (status, body) => {
  body.users.forEach((user) => {
    addGopher(user.id, user.username);
  });
});

// create websocket client
// https://cs.lmu.edu/~ray/notes/jsnetexamples/
const client = new net.Socket();
client.connect({ port: SERVER_PORT, host: SERVER_HOST });

// handle received data
client.on("data", (data) => {
  // must first convert to string and then to json
  dataJson = JSON.parse(data.toString("utf-8"));
  console.log(dataJson);
  handleMessage(dataJson);
});
