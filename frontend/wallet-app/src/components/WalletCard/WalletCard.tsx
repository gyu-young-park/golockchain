import React, { useContext, useEffect, useState } from "react";
import './index.css'
import {FaEthereum} from "react-icons/fa"
import { WalletContext } from "../../context/WalletContext";

const WalletCard = () => {
    const {senderPrivateKey, senderPublicKey, senderBlockchainAddress} = useContext(WalletContext)
    return (
        <div className="wallet-card-container">
            <div className="wallet-card-content-container">
                <div className="wallet-card-content-logo">
                    <p><FaEthereum/></p>    
                    <p>Golockchain</p>
                </div>
                <div className="wallet-card-content-address-container">
                    <p>{senderPublicKey.slice(0,10)}...</p>
                    <div className="wallet-card-content-bk-pk-container">
                        <div>BK:{senderBlockchainAddress.slice(0,10)}...</div>
                        <div className="wallet-card-content-btn-show-pk">{senderPrivateKey.slice(0,10)}...</div>
                    </div>
                </div>
            </div>
        </div>
    )
}

export default WalletCard