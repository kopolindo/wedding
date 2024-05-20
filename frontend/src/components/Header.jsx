import React from 'react';
import './header.css';

export default function Header() {
    return(
        <div className="container text-center mt-5">
            <h1 className="Header display-3 text-primary">
                <span className="font-italic">Alex</span> <span className="text-danger">&</span> <span className="font-italic">Nadia</span>
                <br/>
                <small className="text-muted">26 Settembre 2024</small>
            </h1>
        </div>
    );
}