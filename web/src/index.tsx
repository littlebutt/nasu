import React from 'react'
import ReactDOM from 'react-dom/client'
import { HashRouter } from 'react-router-dom'
import 'antd/dist/reset.css'
import md5 from 'js-md5'
import NasuRoutes from './routes'

window.md5 = md5
const root = ReactDOM.createRoot(
  document.getElementById('root') as HTMLElement
)
root.render(
      <HashRouter>
          <NasuRoutes />
      </HashRouter>
)
