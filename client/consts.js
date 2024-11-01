/** edit this file with your configuration */

// export const SERVER_HOST = "gochat.us-east-1.elasticbeanstalk.com";
// export const SERVER_PORT = null;
// export const SECURE = true;
export const SERVER_HOST = "localhost";
export const SERVER_PORT = 5000;
export const SECURE = false;

export const SERVER_URL = (protocol) =>
  `${protocol}${SECURE ? "s" : ""}://${SERVER_HOST}${SERVER_PORT ? ":" + SERVER_PORT : ""}`;
