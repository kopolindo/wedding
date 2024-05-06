import React, { useState, useEffect } from 'react';
import './guestformpage.css';

const QR = () => {
  const [errorMessage, setErrorMessage] = useState('');
  const [guestsCount, setGuestsCount] = useState(1);
  const [guests, setGuests] = useState([{ id: '', first_name: '', last_name: '', notes: '', confirmed: false }]);
  const [prefilledGuests, setPrefilledGuests] = useState([]);
  const [formSubmitted, setFormSubmitted] = useState(false); // New state variable

  useEffect(() => {
    const QRGen = async (index) => {
        // Retrieve the guest information using the index
        const guest = guests[index];
        console.log(`generating QR code ${guest.id}`);
        try {
            const formData = {
            id: guest.id
            };

            const response = await fetch(`/api/qr`, {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify(formData),
            });

            const data = await response.json();

            if (!response.ok) {
            throw new Error(data.errorMessage);
            }
        } catch (error) {
            console.error('Error submitting form:', error);
        }
    }
  }, [guests]);

  return (
    <div className='QR'>
      QR here
    </div>
  );
};

export default QR;
