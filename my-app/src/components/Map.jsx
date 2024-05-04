import React from 'react';
//https://developers.google.com/maps/documentation/embed/get-started?hl=it  


const Map = () => {
    const KEY = process.env.REACT_APP_GMAPS_API_KEY
    return (
        <div className="map">
            <iframe
                title="restaurant"
                width="600"
                height="450"
                loading="lazy"
                allowFullScreen
                referrerPolicy="no-referrer-when-downgrade"
                src={`https://www.google.com/maps/embed/v1/place?key=${KEY}&q=Via Campiani, 9, 25060 Collebeato BS`}>
            </iframe>
        </div>
    );
};

export default Map;