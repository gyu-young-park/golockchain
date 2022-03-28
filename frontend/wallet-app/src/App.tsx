import { useEffect, useState } from 'react'
import WalletBalance from './components/WalletBalance/WalletBalance'
import WalletCard from './components/WalletCard/WalletCard'
import WalletSend from './components/WalletSend/WalletSend'
import './index.css'

function App() {
  const [data ,setData] = useState("hello")

  useEffect( () => {
    connectWalletServer()
  }, [])
  
  const connectWalletServer = async () => {
    const data = await fetch("http://0.0.0.0:8000",{
      mode: 'cors',
      headers: {
        'Access-Control-Allow-Origin':'*'
      }})
    setData(JSON.stringify(data))
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
