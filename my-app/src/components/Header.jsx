import React, { useState, useEffect } from 'react';
import { Tab, Tabs, TabList, TabPanel } from 'react-tabs';
import 'react-tabs/style/react-tabs.css';
import './header.css';

import Home from './Home';
import Info from './Info';
import SecretPage from './SecretPage';
import GuestFormPage from './GuestFormPage';
import QR from './QR';

export default function Header() {
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

    return (
        <div className='Header'>
            <h1>Alex e Nadia finalmente si sposano!!</h1>
            <Tabs>
                <TabList>
                    <Tab>Home</Tab>
                    <Tab>Informazioni utili</Tab>
                    {isAuthenticated
                        ? <Tab>Form di conferma</Tab>
                        : <Tab>Dimmi il tuo segreto e ti dirò chi sei</Tab>
                    }
                    {isAuthenticated && isConfirmed && <Tab>QRCode</Tab>}
                </TabList>
                <TabPanel><Home/></TabPanel>
                <TabPanel><Info/></TabPanel>
                <TabPanel>
                    {isAuthenticated
                        ? <GuestFormPage uuid={uuid} onFormSubmit={guestConfirmed}/>
                        : <SecretPage onFormSubmit={secretSubmitted}/>
                    }
                </TabPanel>
                {isAuthenticated && isConfirmed && <TabPanel><QR/></TabPanel>}
            </Tabs>
        </div>
    )
}
