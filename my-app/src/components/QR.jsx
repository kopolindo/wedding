import React, { useState, useEffect } from 'react';
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
    QRGen();
    }, []);

    return (
    <div className='QR'>
        {errorMessage}
        {!errorMessage && <img src={`data:image/jpeg;base64,${qr}`} alt="" /> }
    </div>
    );
};

export default QR;
