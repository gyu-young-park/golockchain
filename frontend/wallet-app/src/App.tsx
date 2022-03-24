import { useEffect, useState } from 'react'
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
    <div className="App">
      {data}
    </div>
  )
}

export default App
