import { useContext, useEffect } from "react";
import { UserContext } from "../context/UserContext";

const Dashboard = () => {
    const userContext = useContext(UserContext)

    useEffect(() => {
        fetch("/api/auth/get/guild")
            .then(async (res) => {
                if(res.ok) {
                    res = await res.json();
                    console.log(res);
                    return
                }
            });
    }, [])

    return (
        <p>Welcome, {JSON.stringify(userContext.user)}</p>
    )
}

export default Dashboard
