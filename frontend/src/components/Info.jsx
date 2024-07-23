import React from 'react';
import Map from './Map';
import { Clock, Gift } from 'react-bootstrap-icons';
import { google, outlook, yahoo, office365, ics } from "calendar-link";

const Info = () => {
    const event = {
        title: 'Matrimonio Alex&Nadia',
        description: 'Il lungo viaggio inizia dalla Franciacorta',
        location: 'Via Campiani, 9, 25060 Collebeato BS',
        start: '2024-09-26T10:00:00.000Z',
        duration: [10, 'hour'],
    };

    const googleUrl = google(event);
    const outlookUrl = outlook(event);
    const office365Url = office365(event); 
    const yahooUrl = yahoo(event);
    const icsUrl = ics(event); 

    return(
        <div className='Info'>
            <Map />
            <hr/>
            <div className="container">
                <div className="card border-0">
                    <div className="card-header bg-primary text-white d-flex justify-content-center">
                        <h5 className="card-title mb-0"><Clock className="mr-2" /> Orario</h5>
                        <div className="dropdown dropend">
                            <button
                                type="button"
                                className="btn btn-primary dropdown-toggle btn-sm"
                                data-bs-toggle="dropdown"
                                aria-expanded="false">#SaveTheDate
                            </button>
                            <ul className="dropdown-menu">
                                <li><a className="dropdown-item" href={icsUrl}>Add to Apple (iCal)</a></li>
                                <li><a className="dropdown-item" href={googleUrl}>Add to Google Calendar</a></li>
                                <li><a className="dropdown-item" href={outlookUrl}>Add to Outlook</a></li>
                                <li><a className="dropdown-item" href={office365Url}>Add to Office365</a></li>
                                <li><a className="dropdown-item" href={yahooUrl}>Add to Yahoo</a></li>
                            </ul>
                        </div>
                    </div>
                    <div className="card-body">
                        <div className="col d-flex justify-content-center">   
                            <p align="left">
                                Ore 11.30: Accoglienza<br/>
                                Ore 12.00: Inizio cerimonia<br/>
                                Ore 12.45: Inizio divertimento🤘
                            </p>
                        </div>
                    </div>
                </div>
            </div>
            <div className="container">
                <div className="card border-0">
                    <div className="card-header bg-primary text-white">
                        <h5 className="card-title mb-0"><Gift className="mr-2" /> Kickstart</h5>
                    </div>
                    <div className="card-body">
                        <p align="justify">Per quanto sia bello ricevere e spacchettare regali, sappiamo anche quanto sia difficile farne!<br/>
                        Vogliamo quindi semplificarvi la vita e risparmiare al contempo un sacco di carta ;)</p>
                        <p><i>#SaveTheEarth</i></p>
                        <p align="justify">Se volete potete aiutarci nei primi passi della nostra nuova avventura!<br/>
                        Intestatario: Alex Conti<br/>
                        IBAN: IT69R0301503200000002893665</p>
                    </div>
                </div>
            </div>
        </div>
    );
}

export default Info;