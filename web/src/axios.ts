import axios from 'axios'
import { getCookie } from 'typescript-cookie'
import { history } from './history'

const Axios = axios.create({
  baseURL: 'http://localhost:8080',
  timeout: 2000
})

Axios.interceptors.request.use((config) => {
  config.headers = {
    Authorization: getCookie('token')
  }
  return config
}, async (error) => {
  console.warn(error)
  return await Promise.reject(error)
})

Axios.interceptors.response.use((res) => {
  return res
}, async (error) => {
  if (error.request.status === 401) {
    history.push('/welcome')
  }
  return await Promise.reject(error)
})
export default Axios
