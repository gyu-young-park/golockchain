import { useEffect, useState } from 'react'
import WalletBalance from './components/WalletBalance/WalletBalance'
import WalletCard from './components/WalletCard/WalletCard'
import WalletSend from './components/WalletSend/WalletSend'
import './index.css'

function App() {

  useEffect( () => {
    connectWalletServer()
  }, [])
  
  const connectWalletServer = async () => {
    const res = await fetch("http://0.0.0.0:8000/wallet",{
      method: 'POST',
      mode: 'cors',
      headers: {
        'Content-Type': 'application/json',
        'Access-Control-Allow-Origin':'*',
        "Access-Control-Allow-Methods": "*"
      },
    })
    const data = await res.json()
    console.log(data["private_key"])
    console.log(data["blockchain_address"])
    console.log(data["public_key"])
  }

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
