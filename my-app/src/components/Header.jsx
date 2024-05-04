import React, { useState, useEffect } from 'react';
import { Tab, Tabs, TabList, TabPanel } from 'react-tabs';
import 'react-tabs/style/react-tabs.css';
import './header.css';

import Home from './Home';
import Info from './Info';
import SecretPage from './SecretPage';
import GuestFormPage from './GuestFormPage';

export default function Header() {
  const [isAuthenticated, setIsAuthenticated] = useState(false);
  const [refresh, setRefresh] = useState(false);
  const [uuid, setUuid] = useState(null); // State to hold data

  useEffect(() => {
      // Check if authentication cookie is present
      const authCookie = document.cookie.split(';').find(cookie => cookie.trim().startsWith('auth='));
      if (authCookie) {
          setIsAuthenticated(true);
      } else {
          setIsAuthenticated(false);
      }
  }, [refresh]); // Run whenever 'refresh' state changes

  // Function to handle events that might change authentication status
  const handleEventThatMightChangeAuthStatus = (d) => {
      // Logic to handle event that might change authentication status
      // For example, after form submission
      // You may need to adjust this logic based on your actual implementation
      setRefresh(prevRefresh => !prevRefresh); // Toggle 'refresh' state to trigger re-render
      setUuid(d.uuid);
  };

    return (
        <div className='Header'>
            <h1>Alex e Nadia finalmente si sposano!!</h1>
            <Tabs>
                <TabList>
                    <Tab>Home</Tab>
                    <Tab>Informazioni utili</Tab>
                    {isAuthenticated ?<Tab>Form di conferma</Tab> : <Tab>Dimmi il tuo segreto e ti dirò chi sei</Tab>}
                </TabList>
                <TabPanel><Home/></TabPanel>
                <TabPanel><Info/></TabPanel>
                <TabPanel>
                    {isAuthenticated ? <GuestFormPage uuid={uuid} /> : <SecretPage onFormSubmit={handleEventThatMightChangeAuthStatus}/>}
                </TabPanel>
            </Tabs>
        </div>
    )
}
