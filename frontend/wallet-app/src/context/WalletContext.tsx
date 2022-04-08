import React, { useState , useEffect} from "react";

export interface WalletContextInterface {
    senderPrivateKey: string
    senderBlockchainAddress: string
    senderPublicKey: string
}

const defaultData : WalletContextInterface = {
    senderPrivateKey: "",
    senderBlockchainAddress: "",
    senderPublicKey: "",
}
export const WalletContext : React.Context<WalletContextInterface> = React.createContext<WalletContextInterface>(defaultData)

export const WalletContextProvider = ({children} : {children: React.ReactNode}) => {
    const [senderPrivateKey, setSenderPrivateKey] = useState<string>("")
    const [senderBlockchainAddress, setSenderBlockchainAddress] = useState<string>("")
    const [senderPublicKey, setSenderPublicKey] = useState<string>("")
    
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

        setSenderPrivateKey(data["private_key"])
        setSenderBlockchainAddress(data["blockchain_address"])
        setSenderPublicKey(data["public_key"])
    }

    return (
        <WalletContext.Provider value={{senderPrivateKey,senderBlockchainAddress,senderPublicKey}}>
            {children}
        </WalletContext.Provider>
    )
}