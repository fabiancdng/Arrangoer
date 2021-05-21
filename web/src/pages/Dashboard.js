const Dashboard = ({ user }) => {
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
        <p>Welcome, {JSON.stringify(user)}</p>
    )
}

export default Dashboard
