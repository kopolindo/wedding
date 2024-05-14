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
              <div className="d-flex justify-content-center">
                    <div id="carousel" className="carousel slide carousel-fade" data-bs-ride="ride">
                      <div className="carousel-inner">
                        <div className="carousel-item active">
                          <img src="/images/0.png" className="d-block w-100" alt="first"/>
                        </div>
                        <div className="carousel-item">
                          <img src="/images/1.png" className="d-block w-100" alt="second"/>
                        </div>
                        <div className="carousel-item">
                          <img src="/images/2.png" className="d-block w-100" alt="third"/>
                        </div>
                        <div className="carousel-item active">
                          <img src="/images/3.png" className="d-block w-100" alt="fourth"/>
                        </div>
                        {/*<div className="carousel-item">
                          <img src="/images/4.png" className="d-block w-100" alt="fifth"/>
                        </div> */}
                        <div className="carousel-item">
                          <img src="/images/5.png" className="d-block w-100" alt="sixth"/>
                        </div>
                        <div className="carousel-item active">
                          <img src="/images/6.png" className="d-block w-100" alt="seventh"/>
                        </div>
                        <div className="carousel-item">
                          <img src="/images/7.png" className="d-block w-100" alt="eighth"/>
                        </div>
                        {/*<div className="carousel-item">
                          <img src="/images/8.png" className="d-block w-100" alt="nineth"/>
                        </div> 
                        <div className="carousel-item active">
                          <img src="/images/9.png" className="d-block w-100" alt="tenth"/>
                        </div>
                        <div className="carousel-item">
                          <img src="/images/10.png" className="d-block w-100" alt="eleventh"/>
                        </div>
                        <div className="carousel-item">
                          <img src="/images/11.png" className="d-block w-100" alt="twelveth"/>
                        </div>
                        <div className="carousel-item active">
                          <img src="/images/12.png" className="d-block w-100" alt="thirteenth"/>
                        </div>
                        <div className="carousel-item">
                          <img src="/images/13.png" className="d-block w-100" alt="fourteenth"/>
                        </div>
                        <div className="carousel-item">
                          <img src="/images/14.png" className="d-block w-100" alt="fifteenth"/>
                        </div>
                        <div className="carousel-item active">
                          <img src="/images/15.png" className="d-block w-100" alt="sixteenth"/>
                        </div>
                        <div className="carousel-item">
                          <img src="/images/16.png" className="d-block w-100" alt="seventeenth"/>
                        </div>
                        <div className="carousel-item">
                          <img src="/images/17.png" className="d-block w-100" alt="eighteenth"/>
                        </div>
                        <div className="carousel-item active">
                          <img src="/images/18.png" className="d-block w-100" alt="nineteenth"/>
                        </div>
                        <div className="carousel-item">
                          <img src="/images/19.png" className="d-block w-100" alt="twentyth"/>
                        </div>
                        <div className="carousel-item">
                          <img src="/images/20.png" className="d-block w-100" alt="twenty-first"/>
                        </div>
                        <div className="carousel-item active">
                          <img src="/images/21.png" className="d-block w-100" alt="twenty-second"/>
                        </div>*/}
                        <div className="carousel-item">
                          <img src="/images/22.png" className="d-block w-100" alt="twenty-third"/>
                        </div>
                        {/*<div className="carousel-item">
                          <img src="/images/23.png" className="d-block w-100" alt="twenty-fourth"/>
                        </div>
                        <div className="carousel-item">
                          <img src="/images/24.png" className="d-block w-100" alt="twenty-fifth"/>
                        </div>
                        <div className="carousel-item">
                          <img src="/images/25.png" className="d-block w-100" alt="twenty-sixth"/>
                        </div>*/}
                        <div className="carousel-item active">
                          <img src="/images/26.png" className="d-block w-100" alt="twenty-seventh"/>
                        </div>
                        {/*<div className="carousel-item">
                          <img src="/images/27.png" className="d-block w-100" alt="twenty-eighth"/>
                        </div>
                        <div className="carousel-item">
                          <img src="/images/28.png" className="d-block w-100" alt="twenty-nineth"/>
                        </div>*/}
                      </div>
                      <button className="carousel-control-prev" type="button" data-bs-target="#carousel" data-bs-slide="prev">
                        <span className="carousel-control-prev-icon" aria-hidden="true"></span>
                        <span className="visually-hidden">Previous</span>
                      </button>
                      <button className="carousel-control-next" type="button" data-bs-target="#carousel" data-bs-slide="next">
                        <span className="carousel-control-next-icon" aria-hidden="true"></span>
                        <span className="visually-hidden">Next</span>
                      </button>
                </div>
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
);
}

export default Home;
