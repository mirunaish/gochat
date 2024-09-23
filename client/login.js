// adapted from https://medium.com/swlh/how-to-create-your-first-login-page-with-html-css-and-javascript-602dd71144f1

import { setCookie, makeRequest } from "./common.js";

// -------- signup ---------

const signupForm = document.getElementById("signup-form");
const signupButton = document.getElementById("signup-submit");

// TODO report errors
// const signupErrorMsg = document.getElementById("signup-error-msg");

signupButton.addEventListener("click", (e) => {
  e.preventDefault();
  const username = signupForm.username.value;
  const email = signupForm.email.value;
  const password = signupForm.password.value;

  // make signup request to server
  makeRequest(
    "POST",
    "/signup",
    { username, email, password },
    (status, response) => {
      setCookie(response.token);
      window.location.assign("./index.html");
    }
  );
});

// -------- login ---------

const loginForm = document.getElementById("login-form");
const loginButton = document.getElementById("login-submit");

loginButton.addEventListener("click", (e) => {
  e.preventDefault();
  const email = loginForm.email.value;
  const password = loginForm.password.value;

  // make signup request to server
  makeRequest("POST", "/login", { email, password }, (status, response) => {
    console.log(status, response);

    setCookie(response.token);
    window.location.assign("./index.html");
  });
});
