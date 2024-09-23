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

/** make request to server */
export function makeRequest(method, url, body, handler) {
  fetch(SERVER_URL + url, {
    method: method,
    headers: {
      Authorization: "Bearer " + getCookie(),
    },
    body: body != null ? JSON.stringify(body) : null,
  })
    .then(async (response) => {
      if (response.ok) {
        let data;
        try {
          data = await response.json();
        } catch (e) {
          console.log("failed to get response data");
        }
        handler(response.status, data);
      } else throw new Error(`status code was ${response.status}`);
    })
    .catch((error) => {
      alert("something went wrong. please try again");
      console.error("request error:", error);
    });
}
