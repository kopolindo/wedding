import React, { useState, useEffect } from 'react';
import './guestformpage.css';

const GuestFormPage = () => {
  const [errorMessage, setErrorMessage] = useState('');
  const [guestsCount, setGuestsCount] = useState(1);
  const [guests, setGuests] = useState([{ id: '', first_name: '', last_name: '', notes: '', confirmed: false }]);
  const [prefilledGuests, setPrefilledGuests] = useState([]);
  const [formSubmitted, setFormSubmitted] = useState(false); // New state variable

  useEffect(() => {
    const fetchGuests = async () => {
      try {
        const response = await fetch(`/api/guest`);
        const data = await response.json();
        if (!response.ok) {
          throw new Error(data.errorMessage);
        }
        // Set all guests
        for (let index = 0; index < data.guests.length; index++) {
          const guest = data.guests[index];
          if(!guest.confirmed){
            guest.first_name = '';
            guest.last_name = '';
            guest.notes = '';
          }
        }
        setGuests(data.guests);
        
        // Filter and set only confirmed guests
        setPrefilledGuests(data.guests);
      } catch (error) {
        setErrorMessage(error.message);
      }
    };
    fetchGuests();
    setFormSubmitted(false);
  }, [formSubmitted]);

  useEffect(() => {
    setGuestsCount(prefilledGuests.length);
    setGuests(prefilledGuests);
  }, [prefilledGuests]);

  const handleGuestChange = (index, field, value) => {
    setGuests(prevGuests => {
      const updatedGuests = [...prevGuests];
      updatedGuests[index][field] = value;
      return updatedGuests;
    });
  };

  const handleGuestsCountChange = (event) => {
    setErrorMessage("");
    const count = parseInt(event.target.value, 10);
    if (!isNaN(count)) {
      if (count < 6) {
        setGuestsCount(count);
        setGuests(prevGuests => {
          const existingCount = prevGuests.length;
          if (count > existingCount) {
            // Add empty rows for additional guests
            const additionalGuests = Array.from(
              { length: count - existingCount }, () => (
                { id: '', first_name: '', last_name: '', notes: '', confirmed: '' }
              )
            );
            return [...prevGuests, ...additionalGuests];
          } else {
            // Remove extra rows if count is less than existing count
            return prevGuests.slice(0, count);
          }
        });
      } else {
        setErrorMessage("Inserisci un numero di ospiti non superiore a 5");
      }
    }
  };

  const handleSubmit = async () => {
    try {
      const formData = {
        people: guests.map((guest, index) => ({
          id: guest.id || index,
          first_name: document.querySelector(`input[name=first_name_${index}]`).value,
          last_name: document.querySelector(`input[name=last_name_${index}]`).value,
          notes: document.querySelector(`input[name=notes_${index}]`).value,
        })),
      };
  
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
      setFormSubmitted(true); // Trigger fetching guests after form submission
    } catch (error) {
      console.error('Error submitting form:', error);
    }
  };
  

  const handleDeleteRow = (index) => {
    const idInput = document.querySelector(`input[name=id_${index}]`);
    const id = idInput ? parseInt(idInput.value, 10) : null;
    setGuestsCount((prevCount) => prevCount - 1);
    const updatedGuests = guests.filter((_, i) => i !== index);
    if (id !== null){
      sendDelete(id);
    }
    setGuests(updatedGuests);
  };

  function sendDelete(id) {
    const formData = { id: id };

    return new Promise((resolve, reject) => {
        fetch("/api/guest", {
            method: "DELETE",
            headers: {
                "Content-Type": "application/json"
            },
            body: JSON.stringify(formData)
        })
        .then(response => {
            if (!response.ok) {
                throw new Error("Failed to delete guest");
            }
            return response.json();
        })
        .then(data => {
            resolve(data);
        })
        .catch(error => {
            reject(error.message);
        });
    });
  }

  return (
    <div className='GuestFormPage'>
      <div id="formContainer">
      {errorMessage && <p className='error'>{errorMessage}</p>}
        <div>
          <form className="guest_form" action={`/guest`} method="post">
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
                  <input
                    type="text"
                    name={`first_name_${index}`}
                    value={guest.first_name}
                    placeholder={guest.confirmed ? '' : 'Nome'}
                    onChange={(e) => handleGuestChange(index, 'first_name', e.target.value)}
                  />
                  <input
                    type="text"
                    name={`last_name_${index}`}
                    value={guest.last_name}
                    placeholder="Cognome"
                    onChange={(e) => handleGuestChange(index, 'last_name', e.target.value)}
                  />
                  <input
                    type="text"
                    name={`notes_${index}`}
                    value={guest.notes}
                    placeholder="Allergie/intolleranze"
                    onChange={(e) => handleGuestChange(index, 'notes', e.target.value)}
                  />
                  {index !== 0 && ( // Only render delete button if index is not 0
                    <button className="btn btn-danger" id="DeleteRow" type="button" onClick={() => handleDeleteRow(index)}>
                      <i className="bi bi-trash"></i>
                      Cancella
                    </button>
                  )}
                  {index === 0 && (
                    <button className="btn ghost-button" type="button" disabled title="">
                    Cancella
                    </button>
                  )}
                  {guest.id && <input type="hidden" name={`id_${index}`} value={guest.id} />}
                </div>
              ))}
            </div>
            <button type="button" className="btn btn-success" onClick={handleSubmit}>Conferma</button>
          </form>
        </div>
      </div>
    </div>
  );
};

export default GuestFormPage;
