import React, { useState } from 'react';
import { createMaterialTopTabNavigator } from '@react-navigation/material-top-tabs';
import './body.css';

import Home from './Home';
import Info from './Info';
import SecretPage from './SecretPage';
import GuestFormPage from './GuestFormPage';
import QR from './QR';

const getIsSignedIn = () => {
    let auth = false;
    const authCookie = document.cookie.split(';').find(cookie => cookie.trim().startsWith('auth='));
    if(authCookie){
        const authCookieVal = authCookie.split('=')[1];
        if (authCookieVal === "true") {
            auth = true;
        }
    }
    return auth;
};

const getIsConfirmed = () => {
    let confirmed = false;
    const confirmedCookie = document.cookie.split(';').find(cookie => cookie.trim().startsWith('confirmed='));
    if(confirmedCookie){
        const confirmedCookieVal = confirmedCookie.split('=')[1];
        if (confirmedCookieVal === "true") {
            confirmed = true;
        }
    }
    return confirmed;
};

const delay = ms => new Promise(
    resolve => setTimeout(resolve, ms)
  );

export default function Body() {
    const [isConfirmed, setIsConfirmed] = useState(false);

    const handleSubmitFromGuestFormPage = async(confirmedStatus) => {
        await delay(500);
        if(confirmedStatus==="true" || getIsConfirmed()){
            setIsConfirmed(true);
        }
    }

    const isSignedIn = getIsSignedIn();
    if(getIsConfirmed() && !isConfirmed){
        setIsConfirmed(true);
    }

    const Tab = createMaterialTopTabNavigator();

    return (
        <div className="container">
            <div className="row justify-content-center">
                <div className="col-lg-8">
                    <div className="Body">
                        <Tab.Navigator
                        initialRouteName={isSignedIn ? "Form" : "Secret"}
                        className="nav nav-tabs"
                        >
                            <Tab.Screen
                                name="Home"
                                component={Home}
                                options={{ tabBarLabel: 'Home' }}
                            />
                            <Tab.Screen
                                name="Info"
                                component={Info}
                                options={{ tabBarLabel: 'Informazioni utili' }}
                            />
                            {isSignedIn ? (
                                <Tab.Screen
                                    name="Form"
                                    component={props =>
                                        <GuestFormPage 
                                            handleSubmitFromGuestFormPage={handleSubmitFromGuestFormPage}
                                        />
                                    }
                                    options={{ tabBarLabel: 'Form di conferma' }}
                                />
                            ) : (
                                <Tab.Screen
                                    name="Secret"
                                    component={SecretPage}
                                    options={{ tabBarLabel: 'Dimmi il tuo segreto e ti dirò chi sei' }}
                                />
                            )}
                            {isSignedIn && isConfirmed && (
                                <Tab.Screen
                                    name="QR"
                                    component={QR}
                                    options={{ tabBarLabel: 'QRCode' }}
                                />
                            )}
                        </Tab.Navigator>
                    </div>
                </div>
            </div>
        </div>
    )
}
