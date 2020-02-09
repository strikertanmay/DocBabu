import axios from "axios";

export default axios.create({
  baseURL: "http://68643f4f.ngrok.io",
  headers: {
    "Content-Type": "application/json",
    "Access-Control-Allow-Origin" : "*"
  }
});
