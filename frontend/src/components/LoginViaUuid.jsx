import React, { useEffect, useState } from 'react';
import { useParams } from 'react-router-dom';

const getIsSignedIn = () => {
    let auth = false;
    const authCookie = document.cookie.split(';').find(cookie => cookie.trim().startsWith('auth='));
    if (authCookie){
        const authCookieVal = authCookie.split('=')[1];
        if (authCookieVal === "true") {
            auth = true;
        }
    }
    return auth;
};

export default function LoginViaUuid() {
    const { uuid } = useParams();
    const [fetchAttempted, setFetchAttempted] = useState(false); // Flag to track fetch attempts

    useEffect(() => {
        const isSignedIn = getIsSignedIn();

        if (isSignedIn) {
            window.location = "/";
        } else if (!fetchAttempted) { // Check if fetch is not attempted yet
            fetch(`/guest/${uuid}`)
                .then(response => {
                    if (!response.ok) {
                        throw new Error("Failed to get guest");
                    }
                    // Set the flag to true after a successful fetch
                    setFetchAttempted(true);
                    // Reload the page after a successful fetch
                    window.location.reload();
                })
                .catch(error => {
                    // Handle fetch errors
                    console.error("Fetch error:", error);
                    // Set the flag to true even if fetch fails to avoid infinite loop
                    setFetchAttempted(true);
                });
        }
    }, [uuid, fetchAttempted]);
    
    return (
        <div></div>
    )
}
