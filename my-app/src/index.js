import React from 'react';
import ReactDOM from 'react-dom';
import { BrowserRouter as Router, Switch, Route} from "react-router-dom";

import Home from './components/Home';
import Header from './components/Header'
import SecretPage from './components/SecretPage'
import GuestFormPage from './components/GuestFormPage'
import Info from './components/Info'

const Routing = () => {
  return(
    <Router>
      <Header/>
      <Switch>
        <Route path="/chisono" component={SecretPage} />
        <Route path="/guest/:uuid" component={GuestFormPage} />
      </Switch>
    </Router>
  )
}


ReactDOM.render(
  <React.StrictMode>
    <Routing />
  </React.StrictMode>,
  document.getElementById('root')
);