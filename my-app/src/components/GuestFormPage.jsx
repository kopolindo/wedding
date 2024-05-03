import React, { useState, useEffect } from 'react';
import './guestformpage.css';

const GuestFormPage = () => {
  const [errorMessage, setErrorMessage] = useState('');
  const [guestsCount, setGuestsCount] = useState(1);
  const [guests, setGuests] = useState([{ firstName: '', lastName: '', notes: '' }]);
  const [submitSuccess, setSubmitSuccess] = useState(false);
  const [prefilledGuests, setPrefilledGuests] = useState([]);

  useEffect(() => {
    const fetchGuests = async () => {
      try {
        const response = await fetch(`/api/guest`);
        const data = await response.json();
        if (!response.ok) {
          throw new Error(data.errorMessage);
        }
        setGuests(data);
      } catch (error) {
        setErrorMessage(error.message);
      }
    };

    fetchGuests();
  }, []);

  useEffect(() => {
    // Set guests count based on prefilledGuests length
    setGuestsCount(prefilledGuests.length);
    // Set prefilled guests data
    setGuests(prefilledGuests);
  }, [prefilledGuests]);

  const handleSubmit = async (e) => {
    e.preventDefault();
    try {
      const formData = guests.map((guest) => ({
        id: guest.ID,
        firstName: e.target.elements[`first_name_${guest.ID}`].value,
        lastName: e.target.elements[`last_name_${guest.ID}`].value,
        notes: e.target.elements[`notes_${guest.ID}`].value,
      }));
      const response = await fetch(`/api/guest`, {
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
      console.log('Form submitted successfully');
      // Optionally handle success (redirect, show confirmation, etc.)
      setPrefilledGuests(data.guests.filter(guest => guest.confirmed));
      setSubmitSuccess(true);
    } catch (error) {
      console.error('Error submitting form:', error);
      // Optionally handle error (display error message, retry, etc.)
    }
  };

  const handleGuestChange = (index, field, value) => {
    const updatedGuests = [...guests];
    updatedGuests[index][field] = value;
    setGuests(updatedGuests);
  };

  const handleGuestsCountChange = (event) => {
    const count = parseInt(event.target.value, 10);
    if (!isNaN(count)) {
      setGuestsCount(count);
      const updatedGuests = Array.from({ length: count }, () => ({ firstName: '', lastName: '', notes: '' }));
      setGuests(updatedGuests);
    }
  };

  const submitForm = async () => {
    try {
      const formData = {
        guests: guestsCount,
        people: guests
      };

      // Convert form data to JSON
      const jsonData = JSON.stringify(formData);

      // Example of sending JSON data using fetch API
      const response = await fetch(`/api/guest/${guests[0].id}`, {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json'
        },
        body: jsonData
      });

      if (!response.ok) {
        throw new Error('Network response was not ok');
      }

      // Remove the SendSecret div after successful submission
      const sendSecretDiv = document.querySelector('.SendSecret');
      if (sendSecretDiv) {
        sendSecretDiv.remove();
      }
    } catch (error) {
      setErrorMessage('There was a problem submitting the form: ' + error.message);
    }
  };

  return (
    <div className='GuestFormPage'>
{/*       <h1>Guest Form</h1>
      <form onSubmit={handleSubmit}>
        {guests.map((guest) => (
          <div key={guest.ID}>
            <input
              type="text"
              name={`first_name_${guest.ID}`}
              placeholder="First Name"
              defaultValue={guest.FirstName || ''}
              readOnly={guest.Confirmed}
            />
            <input
              type="text"
              name={`last_name_${guest.ID}`}
              placeholder="Last Name"
              defaultValue={guest.LastName || ''}
              readOnly={guest.Confirmed}
            />
            <input
              type="text"
              name={`notes_${guest.ID}`}
              placeholder="Notes"
              defaultValue={guest.Notes || ''}
              readOnly={guest.Confirmed}
            />
          </div>
        ))}
        <button type="submit">Submit</button>
      </form> */}
      <div id="formContainer">
      {errorMessage && <p className='error'>{errorMessage}</p>}
        {submitSuccess && (
          <div>
            <h1>Form di conferma</h1>
            <form className="guest_form" action={`/guest/${guests[0].id}`} method="post" onSubmit={handleSubmit}>
              <label htmlFor="guests">Numero <b>totale</b> di partecipanti:</label>
              <input
                type="number"
                id="guests"
                name="guests"
                min="1"
                max="5"
                value={guestsCount}
                onChange={handleGuestsCountChange}
                required
              />
              <br />
              <div className="divs">
                {guests.map((guest, index) => (
                  <div key={index} id="row">
                    <button className="btn btn-danger" id="DeleteRow" type="button">
                      <i className="bi bi-trash"></i>
                      Cancella
                    </button>
                    <input
                      type="text"
                      name={`first_name_${index}`}
                      value={guest.firstName}
                      placeholder="Nome"
                      onChange={(e) => handleGuestChange(index, 'firstName', e.target.value)}
                    />
                    <input
                      type="text"
                      name={`last_name_${index}`}
                      value={guest.lastName}
                      placeholder="Cognome"
                      onChange={(e) => handleGuestChange(index, 'lastName', e.target.value)}
                    />
                    <input
                      type="text"
                      name={`notes_${index}`}
                      value={guest.notes}
                      placeholder="Allergie/intolleranze"
                      onChange={(e) => handleGuestChange(index, 'notes', e.target.value)}
                    />
                    {guest.id && <input type="hidden" name={`id_${index}`} value={guest.id} />}
                  </div>
                ))}
              </div>
              <button type="button" onClick={submitForm}>Conferma</button>
            </form>
          </div>
        )}
      </div>
    </div>
  );
};

export default GuestFormPage;
