import React, { useState } from 'react';
import './secretpage.css';

export default function SecretPage({ onFormSubmit }) {
  const [secret, setSecret] = useState('');
  const [errorMessage, setErrorMessage] = useState('');

  const submitSecret = async (event) => {
    event.preventDefault();
    try {
      const response = await fetch('/chisono', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json'
        },
        body: JSON.stringify({ secret })
      });
      const data = await response.json();
      
      if (!response.ok) {
        throw new Error(data.errorMessage);
      }
      
      onFormSubmit(data);
    } catch (error) {
      setErrorMessage(error.message);
    }
  };

  return (
    <div className='SecretPage'>
      {errorMessage && <p className='error'>{errorMessage}</p>}
      <div className='SendSecret'>
        <h1>Parola d'ordine?</h1>
        <form onSubmit={ submitSecret }>
          <input
            type="text"
            id="secret"
            value={secret}
            onChange={(e) => setSecret(e.target.value)}
            required
          />
          <button type="submit">Conferma</button>
        </form>
      </div>
    </div>
  );
};
