import { useContext, useEffect } from "react";
import { UserContext } from "../context/UserContext";
import AdminDashboard from '../components/AdminDashboard'
import UserDashboard from '../components/UserDashboard'

const Dashboard = () => {
    const { guild, setGuild } = useContext(UserContext)

    useEffect(() => {
        fetch("/api/auth/get/guild")
            .then(async res => {
                if(res.ok) {
                    res = await res.json();
                    console.log(res);
                    setGuild(res)
                }
            });
    }, [setGuild])

    if(guild.user_is_admin) {
        return <AdminDashboard />
    } else {
        return <UserDashboard />
    }

}

export default Dashboard
