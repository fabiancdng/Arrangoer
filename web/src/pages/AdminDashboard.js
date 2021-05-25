import React from 'react'
import { useEffect, useState } from 'react'
import { Flex, Spacer} from '@chakra-ui/react'
import { ApiAddress } from '../config'
import Applications from '../components/Applications'
import Teams from '../components/Teams'

const AdminDashboard = () => {
    const [approvedApplications, setApprovedApplications] = useState([])
    const [pendingApplications, setPendingApplications] = useState([])
    var newAppApplication = approvedApplications.slice()
    var newPenApplication = pendingApplications.slice()

    useEffect(() => {
        fetch(ApiAddress + "/api/application/list", {
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
                            if(res.accepted > 0) newAppApplication.push(res)
                            else newPenApplication.push(res)
                        })
                    setApprovedApplications(newAppApplication)
                    setPendingApplications(newPenApplication)
                } else if(res.status === 401) {
                    localStorage.removeItem('jwt')
                    window.location.reload()
                } else {
                    alert("Es gab einen Fehler beim Laden der Anmeldungen.")
                }
            })
    }, [setApprovedApplications, setPendingApplications]) // eslint-disable-line react-hooks/exhaustive-deps

    return (
        <Flex flexDirection={{base: "column", md: "row"}} mx={10}>
            <Teams approvedApplications={approvedApplications} pendingApplications={pendingApplications} />
            <Spacer />
            <Applications approvedApplications={approvedApplications} pendingApplications={pendingApplications} />
            <Spacer />
        </Flex>
    )
}

export default AdminDashboard