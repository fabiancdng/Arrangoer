import { createContext, useState } from 'react'

export const UserContext = createContext();

export const UserProvider = ({ children }) => {
    const [loggedIn, setLoggedIn] = useState(false);
    const [user, setUser] = useState({});

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