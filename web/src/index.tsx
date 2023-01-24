import React from 'react'
import ReactDOM from 'react-dom/client'
import { unstable_HistoryRouter as HistoryRouter } from 'react-router-dom'
import 'antd/dist/reset.css'
import md5 from 'js-md5'
import NasuRoutes from './routes'
import { history } from './history'

window.md5 = md5
const root = ReactDOM.createRoot(
  document.getElementById('root') as HTMLElement
)
root.render(
  // @ts-expect-error
      <HistoryRouter history={history}>
          <NasuRoutes />
      </HistoryRouter>
)
