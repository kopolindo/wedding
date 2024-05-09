import React, { useState, useEffect } from 'react';
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

export default function Body() {
    const isSignedIn = getIsSignedIn();
    const isConfirmed = getIsConfirmed();

    const Tab = createMaterialTopTabNavigator();

    return (
        <div className="container">
            <div className="row justify-content-center">
                <div className="col-lg-8">
                    <div className="Body">
                        <Tab.Navigator
                        initialRouteName={isSignedIn ? "Form" : "Secret"}
                        className="nav nav-tabs"
                        screenListeners={{
                            state: (e) => {
                              console.log('state changed', e.data);
                            },
                          }}
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
                                    component={props => <GuestFormPage {...props}/>}
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
