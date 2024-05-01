import React, { useState, useContext } from 'react';
import './secretpage.css';

const SecretPage = () => {
  const [secret, setSecret] = useState('');

  const handleSubmit = async (e) => {
    e.preventDefault();
    
    try {
      const response = await fetch('/chisono', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json'
        },
        body: JSON.stringify({ secret })
      });
      
      if (!response.ok) {
        throw new Error('Network response was not ok');
      }

      // Handle successful response if needed
      console.log('Secret submitted successfully!');
    } catch (error) {
      // Handle error
      console.error('There was a problem submitting the secret:', error);
    }
  };

  return (
    <div className='SecretPage'>
      <h1>Condividi il tuo segreto</h1>
      <form onSubmit={handleSubmit}>
        <label htmlFor="secret">Dimmi il tuo segreto:</label>
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
  );
};

export default SecretPage;
