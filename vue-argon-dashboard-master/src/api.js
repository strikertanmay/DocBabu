import axios from "axios";

export default axios.create({
  baseURL: "http://fae2f7cc.ngrok.io",
  headers: {
    "Content-Type": "application/json",
    "Access-Control-Allow-Origin" : "*"
  }
});
