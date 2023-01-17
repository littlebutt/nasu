import axios from 'axios';
import {getCookie} from "typescript-cookie";


const Axios = axios.create({
    baseURL: 'http://localhost:8080',
    timeout: 2000,
})

Axios.interceptors.request.use((config) => {
    config.headers = {
        'Authorization': getCookie('token')
    }
    return config;
}, (error) => {
    console.warn(error);
    return Promise.reject(error);
})


Axios.interceptors.response.use((res) => {

    return res;
}, (error) => {
    if (error.request.status === 401) {
        window.location.replace('http://localhost:3000/welcome');
    }
    return Promise.reject(error);
})
export default Axios