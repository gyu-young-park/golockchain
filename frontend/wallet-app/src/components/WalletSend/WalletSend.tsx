import React, {useContext, useState}  from "react";
import './index.css'
import {SiBitcoinsv} from "react-icons/si"
import { WalletContextInterface, WalletContext } from '../../context/WalletContext'

interface TransactionDataInterface extends WalletContextInterface{
    recipientBlockchainAddress: string
    value: string
}

const WalletSend = () => {
    const [recipientBlockchainAddress, setRecipientBlockchainAddress] = useState("")
    const [value, setValue] = useState(0)
    const {senderPrivateKey, senderPublicKey, senderBlockchainAddress} = useContext(WalletContext)
    
    const processSendMoneyHandler = async (event : React.MouseEvent<HTMLElement>) => {
        event.preventDefault()
        const confirm_text = "Are you sure to send?"
        const confirm_result = confirm(confirm_text)
        if(confirm_result !== true){
            alert("Canceled")
            return
        }
        const strValue = value + ''
        const transactionData : TransactionDataInterface = {
            senderPrivateKey, senderPublicKey, senderBlockchainAddress, recipientBlockchainAddress, value:strValue
        }
        console.log(transactionData)
        const jsonTransactionData = {
            sender_private_key: transactionData.senderPrivateKey,
            sender_blockchain_address: transactionData.senderBlockchainAddress,
            sender_public_key: transactionData.senderPublicKey,
            recipient_blockchain_address: transactionData.recipientBlockchainAddress,
            value: transactionData.value,
        }

        const res = await fetch("http://0.0.0.0:8000/transaction",{
            method: 'POST',
            mode: 'cors',
            headers: {
                'Content-Type': 'application/json',
                'Access-Control-Allow-Origin':'*',
                "Access-Control-Allow-Methods": "*"
            },
            body: JSON.stringify(jsonTransactionData)
        })

        const data = await res.json()
        console.log(data)
    }

    const recipientAddressChangeHandler = (e: React.FormEvent<HTMLInputElement>) => {
        setRecipientBlockchainAddress(e.currentTarget.value)
    }
    const valueChangeHandler = (e: React.FormEvent<HTMLInputElement>) => {
        let coin : string | Number = Number(e.currentTarget.value)
        if(Number.isNaN(coin)){
            alert("Please input Number")
            coin = 0
        }
        setValue(Number(coin))
    }

    return (
        <div className="wallet-send-container">
            <div className="wallet-send-logo-container">
                <SiBitcoinsv/>
            </div>
            <div className="wallet-send-content-container">
                <div>Coin</div>
                <div className="wallet-send-content-coin-input-container">
                    <input type="text" value={value} onChange={valueChangeHandler} />
                </div>
                <div>Address</div>
                <div className="wallet-send-content-address-input-container">
                    <input type="text" value={recipientBlockchainAddress} onChange={recipientAddressChangeHandler} />
                </div>
            </div>
            <div className="wallet-send-btn" onClick={processSendMoneyHandler}>
                <div>SEND</div>
            </div>
        </div>
    )
}

export default WalletSend