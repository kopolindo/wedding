import React, { useState, useEffect } from 'react';
import './guestformpage.css';

const QR = () => {
    const [qr,setQr] = useState('');
    const [errorMessage,seterrorMessage] = useState('');
    useEffect(() => {
    const QRGen = async () => {
        // Retrieve the guest information using the index
        console.log(`generating QR code`);
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
        {!errorMessage && <img src={`data:image/jpeg;base64,${qr}`} /> }
    </div>
    );
};

export default QR;
