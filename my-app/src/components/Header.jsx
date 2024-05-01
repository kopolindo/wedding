import React from 'react'
import {Link} from 'react-router-dom'
import { Tab, Tabs, TabList, TabPanel } from 'react-tabs';
import 'react-tabs/style/react-tabs.css';
import './header.css'

import Home from './Home';
import Info from './Info';
import SecretPage from './SecretPage';

export default function header() {
    return (
        <div className='Header'>
            <h1>Alex e Nadia finalmente si sposano!!</h1>
            <Tabs>
                <TabList>
                    <Tab>Home</Tab>
                    <Tab>Info</Tab>
                    <Tab>Dimmi il tuo segreto e ti dirò chi sei</Tab>
                </TabList>
                <TabPanel>
                    <Home/>
                </TabPanel>
                <TabPanel>
                    <Info/>
                </TabPanel>
                <TabPanel>
                    <SecretPage/>
                </TabPanel>
            </Tabs>
        </div>
    )
}