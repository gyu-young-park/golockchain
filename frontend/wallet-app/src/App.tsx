import { useEffect, useState } from 'react'
import WalletBalance from './components/WalletBalance/WalletBalance'
import WalletCard from './components/WalletCard/WalletCard'
import WalletSend from './components/WalletSend/WalletSend'
import './index.css'

function App() {
  return (
    <div className="container">
      <WalletBalance/>
      <WalletCard/>
      <WalletSend/>
      <div style={{color:'white', fontSize:'1.3em'}}>
        Lorem ipsum dolor sit amet, consectetur adipisicing elit. Asperiores molestiae vero aliquid debitis doloremque libero sit omnis impedit dolore error hic officia nihil voluptatibus unde neque, itaque nemo facilis laborum.
      </div>
    </div>
  )
}

export default App
