import axios from "axios";

export const http = axios.create({
  baseURL: "http://localhost:8080/api/v1",
  timeout: 500,
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
      localStorage.removeItem("todo");
      window.location.href = "/login";
      return;
    }

    return Promise.reject(err);
  }
);
