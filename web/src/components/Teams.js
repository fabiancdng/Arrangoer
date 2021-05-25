import { Box, Heading } from "@chakra-ui/layout"
import { useEffect, useState } from "react"
import Team from "./Team"
import { ApiAddress } from '../config'

const Teams = () => {
    const [approvedTeams, setApprovedTeams] = useState([])
    const [pendingTeams, setPendingTeams] = useState([])
    var newAppTeam = approvedTeams.slice()
    var newPenTeam = pendingTeams.slice()

    const updateData = () => {
        fetch(ApiAddress + "/api/team/list", {
            mode: 'cors',
            headers: {
                'Authorization': 'Bearer ' + localStorage.getItem('jwt')
            }
        })
            .then(async res => {
                if(res.ok) {
                    res = await res.json()
                    await res.forEach(res => {
                        console.log(res)
                        if(res.accepted > 0) newAppTeam.push(res)
                        else newPenTeam.push(res)
                    })
                setApprovedTeams(newAppTeam)
                setPendingTeams(newPenTeam)
                } else if(res.status === 401) {
                    localStorage.removeItem('jwt')
                    window.location.reload()
                } else {
                    alert("Es gab einen Fehler beim Laden der Anmeldungen.")
                }
            })
    }

    useEffect(() => {
        updateData()
    }, []) // eslint-disable-line react-hooks/exhaustive-deps

    return (
        <Box p={5} width={{base: "100%", md: "23%"}} mr={5} overflowX="hidden" borderWidth={1} borderRadius={8} boxShadow="lg" flexDirection="column" align="center" justifyContent="center">
            <Heading size="lg">Teams</Heading>
            {
                pendingTeams.length < 1 ? <p>Keine ausstehenden Teams vorhanden.</p>
                : pendingTeams.map((team, index) => (
                    <Team udpateData={updateData} key={index} team={team} approved={false} />
                ))
            }
            {
                approvedTeams.length < 1 ? <p>Keine angenommenen Teams vorhanden.</p>
                : approvedTeams.map((team, index) => (
                    <Team udpateData={updateData} key={index} team={team} approved={true} />
                ))
            }
        </Box>
    )
}

export default Teams
