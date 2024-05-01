import React, { useState, useEffect } from 'react';
import { useParams } from "react-router-dom";
import './guestformpage.css';

const GuestFormPage = ({ match }) => {
  const [guests, setGuests] = useState([]);
  const [error, setError] = useState(null);

  const {uuid} = useParams(); // Extract route parameter 'uuid'
  useEffect(() => {
    console.log('UUID:', uuid);
    const fetchGuests = async () => {
      try {
        const response = await fetch(`/guest/${uuid}`);
        if (!response.ok) {
          throw new Error('User not found');
        }
        const data = await response.json();
        setGuests(data);
      } catch (error) {
        setError(error.message);
      }
    };

    fetchGuests();
  }, [uuid]);

  const handleSubmit = async (e) => {
    e.preventDefault();
    try {
      const formData = guests.map((guest) => ({
        id: guest.ID,
        firstName: e.target.elements[`first_name_${guest.ID}`].value,
        lastName: e.target.elements[`last_name_${guest.ID}`].value,
        notes: e.target.elements[`notes_${guest.ID}`].value,
      }));
      const response = await fetch(`/guest/${uuid}`, {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify(formData),
      });
      if (!response.ok) {
        throw new Error('Failed to submit form');
      }
      console.log('Form submitted successfully');
      // Optionally handle success (redirect, show confirmation, etc.)
    } catch (error) {
      console.error('Error submitting form:', error);
      // Optionally handle error (display error message, retry, etc.)
    }
  };

  if (error) {
    return <div>{error}</div>;
  }

  return (
    <div className='GuestFormPage'>
      <h1>Guest Form</h1>
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
      </form>
    </div>
  );
};

export default GuestFormPage;
