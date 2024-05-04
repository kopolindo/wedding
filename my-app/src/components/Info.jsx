import React from 'react';
import Map from './Map';
import './info.css';

const Info = () => {
    return(
        <div className='Info'>
            <hr/>
            <h2>Ristorante</h2>
            <Map />
            <hr/>
            <h2>Orario</h2>
            <p>Ore 12.00: inizio cerimonia</p>
            <p>Ore 12.45: inizio divertimentoooooo!🤘</p>
            <hr/>
        </div>
    );
}

export default Info;