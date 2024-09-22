// adapted from https://medium.com/swlh/how-to-create-your-first-login-page-with-html-css-and-javascript-602dd71144f1

/** save the given jwt into a cookie */
function setCookie(jwt) {
  // https://www.w3schools.com/js/js_cookies.asp
  // https://tkacz.pro/how-to-securely-store-jwt-tokens/
  // i removed httpOnly because it was already set
  document.cookie = "jwt=" + jwt + "; secure; sameSite=Lax;";
}

/** make request to server */
function AjaxPost(url, body, handler) {
  try {
    let xhr = new XMLHttpRequest();
    xhr.open("POST", url, true);
    xhr.setRequestHeader("Content-type", "application/json; charset=UTF-8");
    xhr.send(JSON.stringify(body));

    xhr.onload = () => handler(xhr.status, xhr.response);
  } catch {
    alert("something went wrong. please try again");
  }
}

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
  AjaxPost(
    "http://localhost:5000/signup",
    { username, email, password },
    (status, response) => {
      console.log(status, response);
      if (status == 200) {
        setCookie(response);
        window.location.assign("./index.html");
      } else alert("something went wrong. please try again");
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
  AjaxPost("localhost:5000/login", { email, password }, (status, response) => {
    console.log(status, response);
    if (status == 200) {
      setCookie(response);
      window.location.assign("./index.html");
    } else alert("something went wrong. please try again");
  });
});
