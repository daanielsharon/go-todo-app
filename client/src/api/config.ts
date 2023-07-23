import axios from "axios";

export const http = axios.create({
  baseURL: "http://localhost:8080/api/v1",
  timeout: 1000,
  headers: {
    "Content-Type": "application/json",
    Accept: "application/json",
  },
  responseType: "json",
  withCredentials: true,
});

http.interceptors.response.use(
  (res) => {
    return res;
  },
  (err) => {
    if (err.response.data.code === 401) {
      sessionStorage.removeItem("todo");
      window.location.href = "/login";
    }
  }
);
