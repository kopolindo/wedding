import React from 'react';
import { createRoot } from 'react-dom/client';
import { BrowserRouter as Router} from "react-router-dom";

import Header from './components/Header'

const Routing = () => {
  return(
    <Router>
      <Header/>
    </Router>
  )
}


const root = createRoot(document.getElementById("root"));
root.render(
  <React.StrictMode>
    <Routing />
  </React.StrictMode>
);