import axios from 'axios';

const apiUrl = process.env.NODE_ENV === 'production' ? "https://melodica.tech/api" : "http://localhost:8080";

axios.defaults.withCredentials = true;

export const api = axios.create({
  baseURL: apiUrl,
});
