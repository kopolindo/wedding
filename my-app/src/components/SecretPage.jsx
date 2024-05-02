import React, { useState} from 'react';
import './secretpage.css';

const SecretPage = () => {
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

      const guests = data.guests;
      const formContainer = document.getElementById('formContainer');

      guests.forEach((guest, index) => {
        const inputContainer = document.createElement('div');
        inputContainer.classList.add('guest-input-container');
        
        const firstNameInput = document.createElement('input');
        firstNameInput.setAttribute('type', 'text');
        firstNameInput.setAttribute('placeholder', 'Nome');
        firstNameInput.setAttribute('name', `guest_${index}_firstName`);
        firstNameInput.value = guest.confirmed ? guest.firstName : '';

        const lastNameInput = document.createElement('input');
        lastNameInput.setAttribute('type', 'text');
        lastNameInput.setAttribute('placeholder', 'Cognome');
        lastNameInput.setAttribute('name', `guest_${index}_lastName`);
        lastNameInput.value = guest.confirmed ? guest.lastName : '';

        const notesInput = document.createElement('input');
        notesInput.setAttribute('type', 'text');
        notesInput.setAttribute('placeholder', 'Allergie/Intolleranze');
        notesInput.setAttribute('name', `guest_${index}_notes`);
        notesInput.value = guest.confirmed ? guest.notes : '';

        inputContainer.appendChild(firstNameInput);
        inputContainer.appendChild(lastNameInput);
        inputContainer.appendChild(notesInput);
        formContainer.appendChild(inputContainer);
      });
      // Submit guest data
      const formData = new FormData();
      guests.forEach((guest, index) => {
        formData.append(`guest_${index}_firstName`, guest.firstName);
        formData.append(`guest_${index}_lastName`, guest.lastName);
        formData.append(`guest_${index}_notes`, guest.notes);
      });
      // Remove the SendSecret div after successful submission
      const sendSecretDiv = document.querySelector('.SendSecret');
      if (sendSecretDiv) {
        sendSecretDiv.remove();
      }
      const errorDiv = document.querySelector('.error');
      if (errorDiv) {
        errorDiv.remove();
      }
    } catch (error) {
      setErrorMessage(error.message);
    }
  };

  return (
    <div className='SecretPage'>
      {errorMessage && <p className='error'>{errorMessage}</p>}
      <div id="formContainer"></div>
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

export default SecretPage;
