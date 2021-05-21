import { createContext, useState } from 'react'

export const UserContext = createContext();

export const UserProvider = ({ children }) => {
    const [loggedIn, setLoggedIn] = useState(false);
    const [user, setUser] = useState({});

    fetch("/api/auth/get/user")
        .then(async (res) => {
            if(res.ok) {
                res = await res.json();
                console.log(res);
                setLoggedIn(true);
                setUser(res);
                return
            }
        })

    return (
        <UserContext.Provider
            value={{
                loggedIn,
                user,
                setLoggedIn,
                setUser
            }}
        >
            {children}
        </UserContext.Provider>
    )
}