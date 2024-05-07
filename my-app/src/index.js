import React from 'react';
import { createRoot } from 'react-dom/client';
import { BrowserRouter as Router} from "react-router-dom";

import Body from './components/Body'
import Header from './components/Header'
import Layout from './components/Layout';

const Routing = () => {
  return(
    <Router>
      <Layout>
        <Header/>
        <Body/>
      </Layout>
    </Router>
  )
}


const root = createRoot(document.getElementById("root"));
root.render(
  <React.StrictMode>
    <Routing />
  </React.StrictMode>
);