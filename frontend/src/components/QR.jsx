import React, { useState, useEffect } from 'react';
import AlertComponent from './alert';
import './guestformpage.css';

const QR = () => {
    const [qr,setQr] = useState('');
    const [errorMessage,seterrorMessage] = useState('');
    useEffect(() => {
    const QRGen = async () => {
        try {
            const response = await fetch(`/api/qr`);
            const data = await response.json();
            if (!response.ok) {
                seterrorMessage(data.errorMessage);
            }else{
                setQr(data.qrcode);
            }
        } catch (error) {
            console.error(error);
        }
    }
    QRGen(); // {SCHEMA}://{DOMAIN}:{PORT}/{UUID}
    }, []);

    return (
        <div className='QR'>
            {errorMessage && <AlertComponent message={ errorMessage }/>}
            <div className="container">
                <div className="card border-0">
                    <div className="card-header bg-primary text-white">
                        <h5 className="card-title mb-0">Smistamento</h5>
                    </div>
                    <div className="card-body">
                        <p className="card-text">
                            Il giorno della cerimonia presentati allo smistamento con questo codice 🪄
                        </p>
                        {!errorMessage && <img src={`data:image/jpeg;base64,${qr}`} alt="" /> }
                    </div>
                </div>
            </div>
        </div>
    );
};

export default QR;
