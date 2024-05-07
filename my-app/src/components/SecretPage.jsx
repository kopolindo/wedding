import React, { useState } from 'react';
import AlertComponent from './alert';

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
      
      onFormSubmit(data);  //uuid
    } catch (error) {
      setErrorMessage(error.message);
    }
  };

  const handleSecretChange = (e) => {
    setErrorMessage('');
    setSecret(e.target.value);
  };

  return (
    <div className='SecretPage'>
      {errorMessage && <AlertComponent message={ errorMessage }/>}
      <div className='SendSecret'>
        <h1>Parola d'ordine?</h1>
        <form className="form-group" onSubmit={ submitSecret }>
          <div className="divs">
            <input
              type="text"
              id="secret"
              className="form-control"
              value={secret}
              onChange={ handleSecretChange }
              style={{ width: "300px", margin: "0 auto" }}
              required
            />
          </div>
          <button type="submit" className="btn btn-success" >Conferma</button>
        </form>
      </div>
    </div>
  );
};
