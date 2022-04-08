import React from "react";
import './index.css'
import {SiBitcoinsv} from "react-icons/si"

interface TransactionData {
    senderPrivateKey: string;

}

const WalletSend = () => {

    const processSendMoneyHandler = (event : React.MouseEvent<HTMLElement>) => {
        event.preventDefault()
        const confirm_text = "Are you sure to send?"
        const confirm_result = confirm(confirm_text)
        if(confirm_result !== true){
            alert("Canceled")
            return
        }
        const transaction_data = {

        }
    } 

    return (
        <div className="wallet-send-container">
            <div className="wallet-send-logo-container">
                <SiBitcoinsv/>
            </div>
            <div className="wallet-send-content-container">
                <div>Coin</div>
                <div className="wallet-send-content-coin-input-container">
                    <input type="text"/>
                </div>
                <div>Address</div>
                <div className="wallet-send-content-address-input-container">
                    <input type="text"/>
                </div>
            </div>
            <div className="wallet-send-btn" onClick={processSendMoneyHandler}>
                <div>SEND</div>
            </div>
        </div>
    )
}

export default WalletSend