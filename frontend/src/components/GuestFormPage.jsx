import React, { useState, useEffect } from 'react';
import './guestformpage.css';
import AlertComponent from './alert';
import BASE_URL from '../config';

export default function GuestFormPage ({handleSubmitFromGuestFormPage}) {
  const [errorMessage, setErrorMessage] = useState('');
  const [guestsCount, setGuestsCount] = useState(1);
  const [guests, setGuests] = useState([{ id: '', first_name: '', last_name: '', notes: '', confirmed: false }]);
  const [prefilledGuests, setPrefilledGuests] = useState([]);
  const [formSubmitted, setFormSubmitted] = useState(false);
  const [isConfirmed, setIsConfirmed] = useState(false);

  useEffect(() => {
    const fetchGuests = async () => {
      try {
        const response = await fetch(`${BASE_URL}/api/guest`);
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

  const handleSubmit =  () => {
    try {
      const formData = {
        people: guests.reduce((acc, guest, index) => {
          const firstName = document.querySelector(`input[name=first_name_${index}]`).value.trim();
          const lastName = document.querySelector(`input[name=last_name_${index}]`).value.trim();
          const notes = document.querySelector(`input[name=notes_${index}]`).value.trim();
  
          if (firstName && lastName) {
            acc.push({
              id: guest.id || index,
              first_name: firstName,
              last_name: lastName,
              notes: notes,
            });
          }
          return acc;
        }, []),
      };
  
      fetch(`${BASE_URL}/api/guest`, {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify(formData),
      }).then(response => {
        if (response.status === 403) {
          setErrorMessage('PROTEGO!!');
          return;
        }
        if (!response.ok) {
          throw new Error('Failed to submit form');
        }
        var confirmedCookie = document.cookie.split(';').find(cookie => cookie.trim().startsWith('confirmed='));
        
        if (confirmedCookie) {
          var confirmedValue = confirmedCookie.split('=')[1];
          if(confirmedValue==="true"){
            setIsConfirmed(true)
          }else{
            setIsConfirmed(false)
          }
        }
        setFormSubmitted(true);
      })
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
        fetch(`${BASE_URL}/api/guest`, {
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
      <div className="container">
                <div className="card border-0">
                    <div className="card-header bg-primary text-white">
                        <h5 className="card-title mb-0">
                          Ehy! Chi siete? Cosa Fate? Cosa Portate? Dove andate?
                          <br/>
                          Sì ma quanti siete??
                        </h5>
                    </div>
                    <div className="card-body">
                      <div id="formContainer">
                        {errorMessage && <AlertComponent message={ errorMessage }/>}
                          <div>
                            <form className="form-group" action={`/guest`} method="post">
                              <input
                                type="number"
                                className="form-control"
                                style={{ width: "100px", margin: "0 auto" }}
                                id="guests"
                                name="guests"
                                min="1"
                                max="5"
                                value={guestsCount}
                                onChange={handleGuestsCountChange}
                                required
                              />
                              <br />
                              <div className="divs g-1 input-group mb-3">
                                  {guests.map((guest, index) => (
                                    <div className="row g-1">
                                      <div className="col" id="first_name">
                                        <input
                                          type="text"
                                          className="form-control"
                                          name={`first_name_${index}`}
                                          value={guest.first_name}
                                          placeholder={guest.confirmed ? '' : 'Nome'}
                                          onChange={(e) => handleGuestChange(index, 'first_name', e.target.value)}
                                        />
                                      </div>
                                      <div className="col" id="last_name">
                                        <input
                                          type="text"
                                          className="form-control"
                                          name={`last_name_${index}`}
                                          value={guest.last_name}
                                          placeholder="Cognome"
                                          onChange={(e) => handleGuestChange(index, 'last_name', e.target.value)}
                                        />
                                      </div>
                                      <div className="col" id="notes">
                                        <input
                                          type="text"
                                          className="form-control"
                                          name={`notes_${index}`}
                                          value={guest.notes}
                                          placeholder="Allergie/intolleranze"
                                          onChange={(e) => handleGuestChange(index, 'notes', e.target.value)}
                                        />
                                      </div>
                                      {index !== 0
                                        ? (
                                            <div className="col" id="delete_row">
                                              <button className="btn btn-danger text-nowrap" id="DeleteRow" type="button" onClick={() => handleDeleteRow(index)}>
                                                <i className="bi bi-trash"></i>
                                                Cancella
                                              </button>
                                            </div>
                                          )
                                        : (
                                            <div className="col" id="ghost_delete">
                                              <button className="btn ghost-button text-nowrap" style={{color: "transparent"}} type="button" disabled title="">
                                                Cancella
                                              </button>
                                            </div>
                                          )
                                      }
                                      {guest.id && <input type="hidden" name={`id_${index}`} value={guest.id} />}
                                    </div>
                                  ))}
                              </div>
                              <button
                                type="button"
                                className="btn btn-success"
                                onClick={() => {
                                    handleSubmit();
                                    handleSubmitFromGuestFormPage(isConfirmed)
                                  }
                                }
                              >Conferma</button>
                            </form>
                          </div>
                        </div>
                    </div>
                </div>
            </div>
    </div>
  );
};
