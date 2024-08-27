import React, { useState, useEffect } from 'react';
import './guestformpage.css';
import { AlertComponent, SuccessAlertComponent } from './alert';
import InputSpinner from 'react-bootstrap-input-spinner';
import { Form, Button, Row, Col } from 'react-bootstrap';

export default function GuestFormPage({ handleSubmitFromGuestFormPage }) {
  const [errorMessage, setErrorMessage] = useState('');
  const [okMessage, setOkMessage] = useState('');
  const [guestsCount, setGuestsCount] = useState(1);
  const [guests, setGuests] = useState([{ id: '', first_name: '', last_name: '', notes: '', confirmed: false, type: 0 }]);
  const [prefilledGuests, setPrefilledGuests] = useState([]);
  const [formSubmitted, setFormSubmitted] = useState(false);
  const [isConfirmed, setIsConfirmed] = useState(false);

  useEffect(() => {
    const fetchGuests = async () => {
      try {
        const response = await fetch('/api/guest');
        const data = await response.json();
        if (!response.ok) {
          throw new Error(data.errorMessage);
        }
        for (let index = 0; index < data.guests.length; index++) {
          const guest = data.guests[index];
          if (!guest.confirmed) {
            guest.first_name = '';
            guest.last_name = '';
            guest.notes = '';
            guest.type = 0; // Default to "Adult" (integer)
          }
        }
        setGuests(data.guests);
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
      if (field === 'type') {
        // Convert value to integer for type field
        updatedGuests[index][field] = parseInt(value, 10);
      } else {
        updatedGuests[index][field] = value;
      }
      return updatedGuests;
    });
  };

  const handleGuestsCountChange = (event) => {
    setErrorMessage("");
    const count = parseInt(event, 10);
    if (!isNaN(count)) {
      if (count < 6) {
        setGuestsCount(count);
        setGuests(prevGuests => {
          const existingCount = prevGuests.length;
          if (count > existingCount) {
            const additionalGuests = Array.from(
              { length: count - existingCount }, () => (
                { id: '', first_name: '', last_name: '', notes: '', confirmed: false, type: 0 } // Default type as integer
              )
            );
            return [...prevGuests, ...additionalGuests];
          } else {
            return prevGuests.slice(0, count);
          }
        });
      } else {
        setErrorMessage("Inserisci un numero di ospiti non superiore a 5");
      }
    }
  };

  const handleSubmit = () => {
    try {
      const formData = {
        people: guests.reduce((acc, guest, index) => {
          const firstName = document.querySelector(`input[name=first_name_${index}]`).value.trim();
          const lastName = document.querySelector(`input[name=last_name_${index}]`).value.trim();
          const notes = document.querySelector(`input[name=notes_${index}]`).value.trim();
          const type = parseInt(document.querySelector(`select[name=type_${index}]`).value, 10); // Ensure type is an integer

          if (firstName && lastName) {
            acc.push({
              id: guest.id || index,
              first_name: firstName,
              last_name: lastName,
              notes: notes,
              type: type
            });
          }
          return acc;
        }, []),
      };

      fetch('/api/guest', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify(formData),
      }).then(response => {
        window.scrollTo({ top: 0, behavior: 'smooth' });
        if (response.status === 403) {
          setErrorMessage('PROTEGO!!');
          return;
        }
        if (response.status === 400) {
          return response.json().then(data => {
            if (data.errors && Array.isArray(data.errors)) {
              const errorMessages = data.errors.map(error => error.message).join(' e ');
              setErrorMessage(errorMessages);
            } else {
              throw new Error('Invalid error format received');
            }
          });
        }
        if (!response.ok) {
          throw new Error('Failed to submit form');
        }
        setOkMessage("Grazie per aver confermato!");
        const confirmedCookie = document.cookie.split(';').find(cookie => cookie.trim().startsWith('confirmed='));

        if (confirmedCookie) {
          const confirmedValue = confirmedCookie.split('=')[1];
          setIsConfirmed(confirmedValue === "true");
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
    if (id !== null) {
      sendDelete(id);
    }
    setGuests(updatedGuests);
  };

  function sendDelete(id) {
    const formData = { id: id };
    return new Promise((resolve, reject) => {
      fetch('/api/guest', {
        method: "DELETE",
        headers: {
          "Content-Type": "application/json"
        },
        body: JSON.stringify(formData)
      })
        .then(response => {
          window.scrollTo({ top: 0, behavior: 'smooth' });
          if (!response.ok) {
            throw new Error("Failed to delete guest");
          }
          setOkMessage("Partecipante cancellato");
          return response.json();
        })
        .then(data => {
          resolve(data);
        })
        .catch(error => {
          setErrorMessage(error.message);
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
              <br />
              Sì ma quanti siete???
            </h5>
          </div>
          <div className="card-body">
            <div id="formContainer">
              {errorMessage && <AlertComponent message={errorMessage} />}
              {okMessage && <SuccessAlertComponent message={okMessage} />}
              <div>
                <Form className="form-group">
                  <div className='divs' style={{ width: "170px", margin: "0 auto" }}>
                    <InputSpinner
                      type={'int'}
                      id="guests"
                      name="guests"
                      required
                      max={5}
                      min={1}
                      step={1}
                      value={guestsCount}
                      onChange={handleGuestsCountChange}
                      variant={'primary'}
                      size="sm"
                    />
                  </div>
                  <br />
                  <div className="divs input-group mb-3">
                    {guests.map((guest, index) => (
                      <Row key={index} className="mb-4"> {/* Adjust the margin between groups */}
                        {/* First Column: Input Fields Stacked Vertically */}
                        <Col>
                          {/* ID */}
                          <Form.Control
                            type="hidden"
                            name={`id_${index}`}
                            value={guest.id}
                          />

                          {/* First Name */}
                          <Form.Control
                            type="text"
                            name={`first_name_${index}`}
                            value={guest.first_name}
                            placeholder={guest.confirmed ? '' : 'Nome'}
                            onChange={(e) => handleGuestChange(index, 'first_name', e.target.value)}
                            className="mb-2" // Space between fields within the group
                          />

                          {/* Last Name */}
                          <Form.Control
                            type="text"
                            name={`last_name_${index}`}
                            value={guest.last_name}
                            placeholder="Cognome"
                            onChange={(e) => handleGuestChange(index, 'last_name', e.target.value)}
                            className="mb-2"
                          />

                          {/* Notes */}
                          <Form.Control
                            type="text"
                            name={`notes_${index}`}
                            value={guest.notes}
                            placeholder="Allergie/intolleranze"
                            onChange={(e) => handleGuestChange(index, 'notes', e.target.value)}
                            className="mb-2"
                          />

                          {/* Type */}
                          <Form.Select
                            name={`type_${index}`}
                            value={guest.type}
                            onChange={(e) => handleGuestChange(index, 'type', e.target.value)}
                          >
                            <option value={0}>Adulto</option>
                            <option value={1}>12 anni o meno</option>
                            <option value={2}>3 anni o meno</option>
                          </Form.Select>
                        </Col>

                        {/* Second Column: Trash Button */}
                        {index !== 0 && (
                          <Col xs="auto" className="d-flex align-items-center ">
                            <Button variant="danger" onClick={() => handleDeleteRow(index)}>
                              <i className="bi bi-trash"></i> Cancella
                            </Button>
                          </Col>
                        )}
                      </Row>
                    ))}
                  </div>
                  <Button
                    type="button"
                    variant="success"
                    onClick={() => {
                      handleSubmit();
                      handleSubmitFromGuestFormPage(isConfirmed);
                    }}
                  >
                    Conferma
                  </Button>
                </Form>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  );
}
