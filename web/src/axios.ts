import axios from 'axios';

const Axios = axios.create({
    baseURL: 'http://localhost:8080',
    timeout: 2000,
})

Axios.defaults.headers.post['Content-Type'] = 'application/x-www-form-urlencoded;charset=UTF-8';
Axios.defaults.headers.common['Authorization'] = window.token;
// TODO: interceptor
export default Axios