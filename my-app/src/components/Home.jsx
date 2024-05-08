import React, { useState, useEffect } from 'react';

function Home({ albumLink, title, description }) {
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

  useEffect(() => {
    const script = document.createElement('script');
    script.src = 'https://cdn.jsdelivr.net/npm/publicalbum@latest/embed-ui.min.js';
    script.async = true;
    document.body.appendChild(script);

    return () => {
      document.body.removeChild(script);
    };
  }, []);

  return (
    <div className="Home container">
      <div className="container">
        <div className="card border-0">
            <div className="card-header bg-primary text-white">
                <h5 className="card-title mb-0">Countdown</h5>
            </div>
            <div className="card-body">
                <p className="card-text">{countdown}</p>
            </div>
        </div>
      </div>
      <hr/>
      <div className="container-fluid">
        <div className="row">
          <div className="container">
            <div className="card border-0">
              <div className="card-header bg-primary text-white">
                <h5 className="card-title mb-0">Sharing is caring</h5>
              </div>
            </div>
          </div>
          <div className="col-sm-6 d-flex justify-content-center">
            <div className="card border-0" style={{width: '20em'}}>
              <div className="card-body">
                <h5 className="card-title">Foto</h5>
                <p className="text">Vogliamo collezionare ogni singolo ricordo...</p>
                <p className="text">Aiutateci caricando le vostre foto della festa su questo album!</p>
                <a href="https://photos.app.goo.gl/zRJfPDHPipjQ1b3z8" className="btn btn-primary">Visualizza Album</a>
              </div>
            </div>
          </div>
          <div className="col-sm-6 d-flex justify-content-center">
            <div className="card border-0" style={{width: '20em'}}>
              <div className="card-body">
                <h5 className="card-title">Jukebox</h5>
                <p className="text">Let's play some music!!</p>
                <p className="text">Aggiungete qui le vostre canzoni e vediamo se riusciamo a farvi sentire qualcosa di bello!</p>
                <a href="https://open.spotify.com/playlist/3BEJVi97rLZCB7XgKF4dqX?si=6f8229ee4ec74982&pt=89435809e1ad5fb15750c98f4dc5950d" className="btn btn-primary">Visualizza Playlist</a>
              </div>
            </div>
          </div>
          <div className="col-sm-6 d-flex justify-content-center">
            <div className="card border-0" style={{width: '20em'}}>
              <div className="card-body">
                <div className="pa-carousel-widget" style={{ width: '100%', display: 'none' }}
                  data-link={albumLink}
                  data-title={title}
                  data-description={description}>
                  <object data="https://lh3.googleusercontent.com/pw/AP1GczOIhJ-HGKw6soEHl7uOHmWEwG95EKRrE-MmQ0Q-ltqmbDDAT_Kh2hEv9SKRs3lkbAoTC2XkefKLBmu00--iQcAbOMkirbZi-S-uE2DQXn9Uvp79P-Be=w1920-h1080"></object>
                </div>
              </div>
            </div>
          </div>
          <div className="col-sm-6 d-flex justify-content-center">
            <div className="card border-0" style={{width: '20em'}}>
              <div className="card-body">
                <iframe
                  title="Spotify Playlist"
                  src="https://open.spotify.com/embed/playlist/3BEJVi97rLZCB7XgKF4dqX?utm_source=generator"
                  width="100%"
                  height="152"
                  allow="autoplay; clipboard-write; encrypted-media; fullscreen; picture-in-picture"
                  allowFullScreen
                  loading="lazy"
                ></iframe>
              </div>
            </div>
          </div>
        </div>
      </div>

    </div>
);
}

export default Home;
