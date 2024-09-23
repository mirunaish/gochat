import { SERVER_URL } from "./consts.js";

export function random(start, end) {
  return start + Math.floor(Math.random() * (end - start));
}

/** save the given jwt into a cookie */
export function setCookie(jwt) {
  // https://www.w3schools.com/js/js_cookies.asp
  // https://tkacz.pro/how-to-securely-store-jwt-tokens/
  // i removed httpOnly because it was already set
  document.cookie = "jwt=" + jwt + "; secure; sameSite=Lax;";
}

export function getCookie() {
  const cookies = document.cookie.split(";"); // split cookies
  // construct object of key:value
  const cookieMap = {};
  cookies.forEach((cookie) => {
    if (cookie.indexOf("=") == -1) return;
    const [key, value] = cookie.split("=");
    cookieMap[key] = value;
  });
  // now simply get my jwt cookie
  return cookieMap.jwt;
}

/** make get request to server */
export function AjaxGet(url, handler) {
  try {
    const xhr = new XMLHttpRequest();
    xhr.open("GET", SERVER_URL + url, true);
    xhr.setRequestHeader("Authorization", "Bearer " + getCookie());
    xhr.send();

    xhr.onload = () => handler(xhr.status, JSON.parse(xhr.response));
  } catch {
    alert("something went wrong. please try again");
  }
}

/** make post request to server */
export function AjaxPost(url, body, handler) {
  try {
    let xhr = new XMLHttpRequest();
    xhr.open("POST", SERVER_URL + url, true);
    xhr.setRequestHeader("Content-type", "application/json; charset=UTF-8");
    xhr.setRequestHeader("Authentication", "Bearer " + getCookie());
    xhr.send(JSON.stringify(body));

    xhr.onload = () => handler(xhr.status, JSON.parse(xhr.response));
  } catch {
    alert("something went wrong. please try again");
  }
}
