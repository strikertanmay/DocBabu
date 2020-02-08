import axios from "axios";

export default axios.create({
  baseURL: "http://d94bef41.ngrok.io",
  headers: {
    "Content-type": "application/json"
  }
});