import React from "react";
import './index.css'
import {SiBitcoinsv} from "react-icons/si"

const WalletSend = () => {
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
            <div className="wallet-send-btn">
                <div>SEND</div>
            </div>
        </div>
    )
}

export default WalletSend