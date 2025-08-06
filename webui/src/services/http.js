import axios from 'axios';

const http = axios.create({
  baseURL: import.meta.env.VITE_APP_API_BASE_URL || '',
  timeout: 50000,
});

http.interceptors.request.use(
  config => {
    config.headers = {
      "Content-Type": "application/json;charset=utf-8",
    };
    config.withCredentials = true;
    return config;
  },
  error => {
    return Promise.reject(error);
  }
);

http.interceptors.response.use(
  response => {
    if (response && response.data.code == 10) {
      if (window.location.pathname.substring(0,5) != "/user") {
        window.location.href = "/user/login";
      }
    }
    return response.data;
  },
  error => {
    console.log("axios err：" + error)
    return Promise.reject(error)
  },
);

export default http;
