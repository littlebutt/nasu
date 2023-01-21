import { Route, Routes } from 'react-router-dom'
import App from './app'
import Welcome from './welcome'
import React from 'react'

const NasuRoutes: React.FC = () => {
  return (
        <Routes>
            <Route path='/' element={<App/>}/>
            <Route path='/welcome' element={<Welcome />}/>
        </Routes>
  )
}

export default NasuRoutes
