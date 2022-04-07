import React, { useEffect } from "react";
import './index.css'
import {FaEthereum} from "react-icons/fa"

const WalletCard = () => {

    return (
        <div className="wallet-card-container">
            <div className="wallet-card-content-container">
                <div className="wallet-card-content-logo">
                    <p><FaEthereum/></p>    
                    <p>Golockchain</p>
                </div>
                <div className="wallet-card-content-address-container">
                    <p>#0xE65F40...B2738D296a1</p>
                    <div className="wallet-card-content-bk-pk-container">
                        <div>BK:0xE65F40...B2738D296a1</div>
                        <div className="wallet-card-content-btn-show-pk">Show PK</div>
                    </div>
                </div>
            </div>
        </div>
    )
}

export default WalletCard