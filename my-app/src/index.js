import React from 'react';
import { createRoot } from 'react-dom/client';
import { BrowserRouter as Router} from "react-router-dom";

import Body from './components/Body'

const Routing = () => {
  return(
    <Router>
      <Body/>
    </Router>
  )
}


const root = createRoot(document.getElementById("root"));
root.render(
  <React.StrictMode>
    <Routing />
  </React.StrictMode>
);