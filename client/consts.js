/** edit this file with your configuration */

export const SERVER_HOST = "gochat.us-east-1.elasticbeanstalk.com";
export const SERVER_PORT = null;

export const SERVER_URL = (protocol) =>
  `${protocol}://${SERVER_HOST}${SERVER_PORT ? ":" + SERVER_PORT : ""}`;
