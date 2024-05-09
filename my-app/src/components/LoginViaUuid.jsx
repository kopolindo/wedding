import React, { useState, useEffect } from 'react';
import { useParams } from 'react-router-dom';

import AlertComponent from './alert';

export default function LoginViaUuid() {
    const { uuid } = useParams();
    const [errorMessage, setErrorMessage] = useState('');

    useEffect(() => {
        fetch(`/guest/${uuid}`)
        .then(response => {
            if (!response.ok) {
                throw new Error("Failed to get guest");
            }
            fetch(`/api/confirmed`);
        })
        .then()
        .catch(error => {
            setErrorMessage(error.message);
        });
    }, [uuid]);

    return (
        <div className="container">
            {errorMessage && <AlertComponent message={ errorMessage }/>}
        </div>
    )
}