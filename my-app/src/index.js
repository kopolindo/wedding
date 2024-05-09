import React from 'react';
import { createRoot } from 'react-dom/client';
import { BrowserRouter as Router, Route, Switch } from "react-router-dom";

import Body from './components/Body'
import Header from './components/Header'
import Layout from './components/Layout';
import LoginViaUuid from './components/LoginViaUuid';

const Routing = () => {
  return(
    <Router>
      <Layout>
        <Header/>
        <Body/>
      </Layout>
      <Switch>
        <Route path="/guest/:uuid" component={LoginViaUuid} />
      </Switch>
    </Router>
  )
}


const root = createRoot(document.getElementById("root"));
root.render(
  <React.StrictMode>
    <Routing />
  </React.StrictMode>
);