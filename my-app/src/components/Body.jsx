import React, { useState, useEffect } from 'react';
import { Tab, Tabs, TabList, TabPanel } from 'react-tabs';
import { Redirect, useParams } from 'react-router-dom';
import 'react-tabs/style/react-tabs.css';
import './body.css';

import Home from './Home';
import Info from './Info';
import SecretPage from './SecretPage';
import GuestFormPage from './GuestFormPage';
import QR from './QR';

export default function Body() {
    const [isAuthenticated, setIsAuthenticated] = useState(false);
    const [refresh, setRefresh] = useState(false);
    const [uuid, setUuid] = useState(null);
    const [isConfirmed, setIsConfirmed] = useState(false);

    useEffect(() => {
        setIsAuthenticated(false);
        // Check if authentication cookie is present
        const authCookie = document.cookie.split(';').find(cookie => cookie.trim().startsWith('auth='));
        if(authCookie){
            const authCookieVal = authCookie.split('=')[1];
            if (authCookieVal === "true") {
                setIsAuthenticated(true);
            }
        }
        const confirmedCookie = document.cookie.split(';').find(cookie => cookie.trim().startsWith('confirmed='));
        if(confirmedCookie){
            const confirmedCookieVal = confirmedCookie.split('=')[1];
            if (confirmedCookieVal === "true") {
                setIsConfirmed(true);
            }
        }
        fetch(`/api/confirmed`);
    }, [refresh]);

    const secretSubmitted = (d) => {
        setRefresh(prevRefresh => !prevRefresh);
        setUuid(d.uuid);
    };

    // If guest was confirmed => render QRCode tab
    const guestConfirmed = (d) => {
        setRefresh(prevRefresh => !prevRefresh);
        setIsConfirmed(d);
    };

    const { uuidParam } = useParams();
    
    return (
        <div className="container">
            <div className="row justify-content-center">
                <div className="col-lg-8">
                    <div className="Body">
                        <Tabs defaultactivekey={uuidParam ? "form" : "home"}>
                            <TabList className="nav nav-tabs">
                                <Tab className="nav-item nav-link">Home</Tab>
                                <Tab className="nav-item nav-link">Informazioni utili</Tab>
                                {isAuthenticated ? (
                                    <Tab className="nav-item nav-link">Form di conferma</Tab>
                                ) : (
                                    <Tab className="nav-item nav-link">Dimmi il tuo segreto e ti dirò chi sei</Tab>
                                )}
                                {isAuthenticated && isConfirmed && (
                                    <Tab className="nav-item nav-link">QRCode</Tab>
                                )}
                            </TabList>
                            <TabPanel>
                                <Home />
                            </TabPanel>
                            <TabPanel>
                                <Info />
                            </TabPanel>
                            <TabPanel>
                                {isAuthenticated ? (
                                    <GuestFormPage uuid={uuidParam} onFormSubmit={guestConfirmed} />
                                ) : (
                                    <SecretPage onFormSubmit={secretSubmitted} />
                                )}
                            </TabPanel>
                            {isAuthenticated && isConfirmed && (
                                <TabPanel>
                                    <QR />
                                </TabPanel>
                            )}
                            {uuidParam && <Redirect to={`/:${uuidParam}`} />}
                        </Tabs>
                    </div>
                </div>
            </div>
        </div>
    )
}
