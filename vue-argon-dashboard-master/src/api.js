import axios from "axios";

export default axios.create({
  baseURL: "http://6e022ec2.ngrok.io",
  headers: {
    "Content-Type": "application/json",
    "Access-Control-Allow-Origin" : "*"
  }
});
