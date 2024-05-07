import React from 'react';
import Map from './Map';
import { Clock } from 'react-bootstrap-icons';

const Info = () => {
    return(
        <div className='Info'>
            <Map />
            <hr/>
            <div className="container">
                <div className="card">
                    <div className="card-header bg-primary text-white">
                        <h5 className="card-title mb-0"><Clock className="mr-2" /> Orario</h5>
                    </div>
                    <div className="card-body">
                        <p>Ore 12.00: Inizio cerimonia</p>
                        <p>Ore 12.45: Inizio divertimento🤘</p>
                    </div>
                </div>
            </div>
        </div>
    );
}

export default Info;