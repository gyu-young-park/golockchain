import React from 'react'
import ReactDOM from 'react-dom'
import './index.css'
import App from './App'
import { WalletContextProvider } from './context/WalletContext'

ReactDOM.render(
  <WalletContextProvider>
    <React.StrictMode>
      <App />
    </React.StrictMode>
  </WalletContextProvider>,
  document.getElementById('root')
)
