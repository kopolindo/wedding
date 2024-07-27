import React, { useState, useEffect } from 'react';
import { Alert } from 'react-bootstrap'; // Import Alert component from React Bootstrap

const AlertComponent = ({ message }) => {
  const [showAlert, setShowAlert] = useState(false);

  // Set showAlert to true when a new message is received
  useEffect(() => {
    setShowAlert(true);
  }, [message]);

  // Set a timer to hide the alert after 5 seconds
  useEffect(() => {
    const timer = setTimeout(() => {
      setShowAlert(false);
    }, 5000);

    // Cleanup function to clear the timer
    return () => clearTimeout(timer);
  }, [showAlert]); // Run effect whenever showAlert state changes

  return (
    <div>
      {showAlert && (
        <Alert variant="danger">
          <p>{message}</p>
        </Alert>
      )}
    </div>
  );
};

const SuccessAlertComponent = ({ message }) => {
  const [showAlert, setShowAlert] = useState(false);

  // Set showAlert to true when a new message is received
  useEffect(() => {
    setShowAlert(true);
  }, [message]);

  // Set a timer to hide the alert after 5 seconds
  useEffect(() => {
    const timer = setTimeout(() => {
      setShowAlert(false);
    }, 5000);

    // Cleanup function to clear the timer
    return () => clearTimeout(timer);
  }, [showAlert]); // Run effect whenever showAlert state changes

  return (
    <div>
      {showAlert && (
        <Alert variant="success">
          <p>{message}</p>
        </Alert>
      )}
    </div>
  );
};

export { AlertComponent, SuccessAlertComponent };
