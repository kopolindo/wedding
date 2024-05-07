import React from 'react';
import { Card } from 'react-bootstrap';
import { GeoAltFill } from 'react-bootstrap-icons';

const Map = () => {
    const KEY = process.env.REACT_APP_GMAPS_API_KEY
    return (
        <div className="container">
            <Card>
                <Card.Header className="bg-primary text-white">
                    <Card.Title className="mb-0"><GeoAltFill/> Ristorante</Card.Title>
                </Card.Header>
                <Card.Body>
                    <div className="embed-responsive embed-responsive-16by9">
                        <iframe
                            className="embed-responsive-item"
                            title="restaurant"
                            width="600"
                            height="450"
                            loading="lazy"
                            allowFullScreen
                            referrerPolicy="no-referrer-when-downgrade"
                            src={`https://www.google.com/maps/embed/v1/place?key=${KEY}&q=Via Campiani, 9, 25060 Collebeato BS`}
                        />
                    </div>
                </Card.Body>
            </Card>
        </div>
    );
};

export default Map;
