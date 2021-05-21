import { createContext, useState } from 'react'

export const UserContext = createContext();

export const UserProvider = ({ children }) => {
    const [loggedIn, setLoggedIn] = useState(false);
    const [user, setUser] = useState({});
    const [guild, setGuild] = useState({});

    return (
        <UserContext.Provider
            value={{
                loggedIn,
                user,
                guild,
                setLoggedIn,
                setUser,
                setGuild
            }}
        >
            {children}
        </UserContext.Provider>
    )
}