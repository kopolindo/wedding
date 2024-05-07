import React, { useState, useEffect } from 'react';

function Home() {
  const [countdown, setCountdown] = useState('');

  useEffect(() => {
    const intervalId = setInterval(() => {
      const targetDate = new Date('2024-09-26T00:00:00+02:00'); // September 26th, 12:00 AM Italian time
      const now = new Date(); // Current date and time

      const difference = targetDate - now; // Difference in milliseconds

      if (difference > 0) {
        const days = Math.floor(difference / (1000 * 60 * 60 * 24));
        const hours = Math.floor((difference % (1000 * 60 * 60 * 24)) / (1000 * 60 * 60));
        const minutes = Math.floor((difference % (1000 * 60 * 60)) / (1000 * 60));
        const seconds = Math.floor((difference % (1000 * 60)) / 1000);

        setCountdown(`${days} giorni, ${hours} ore, ${minutes} minuti, ${seconds} secondi`);
      } else {
        setCountdown('Countdown expired');
        clearInterval(intervalId); // Stop the interval when countdown expires
      }
    }, 1000); // Update every second

    // Clean up the interval when the component unmounts
    return () => clearInterval(intervalId);
  }, []);

  return (
    <div className="Home container">
      <div className="row">
        <div className="col-md-6">
          <p className="text">Vogliamo collezionare ogni singolo ricordo...</p>
          <p className="text">Aiutateci caricando le vostre foto della festa su questo album!</p>
          <a href="https://photos.app.goo.gl/zRJfPDHPipjQ1b3z8" className="btn btn-primary">Visualizza Album</a>
        </div>
        <div className="col-md-6">
          <div className="card">
            <div className="card-body">
              <h5 className="card-title">Countdown</h5>
              <p className="card-text">{countdown}</p>
            </div>
          </div>
        </div>
      </div>
    </div>
  );
}

export default Home;
