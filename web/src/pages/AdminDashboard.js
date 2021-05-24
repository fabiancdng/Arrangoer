import { useEffect, useState } from 'react'
import { Box, Button, Flex, Heading, Spacer, VStack } from '@chakra-ui/react'
import { ArrowForwardIcon, AtSignIcon, DragHandleIcon, EmailIcon } from '@chakra-ui/icons'

const AdminDashboard = () => {
    const [approvedApplications, setApprovedApplications] = useState([])
    const [pendingApplications, setPendingApplications] = useState([])
    var newAppApplication = approvedApplications.slice()
    var newPenApplication = pendingApplications.slice()

    useEffect(() => {
        fetch("/api/application/list")
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
                } else {
                    alert("Es gab einen Fehler beim Laden der Anmeldungen.")
                }
            })
    }, [setApprovedApplications, setPendingApplications]) // eslint-disable-line react-hooks/exhaustive-deps

    return (
        <Flex flexDirection={{base: "column", md: "row"}} mx={10}>
            <Box p={5} width={{base: "100%", md: "23%"}} mr={5} overflowX="hidden" borderWidth={1} borderRadius={8} boxShadow="lg" flexDirection="column" align="center" justifyContent="center">
                <Heading size="lg">Teams</Heading>
                {
                    pendingApplications.length < 1 ? <p>Keine Daten vorhanden.</p>
                    : pendingApplications.map(application => (
                        <Flex p={5} m={5} flexDirection="column" justifyContent="space-between" alignItems="center" borderWidth={1} borderRadius={8} boxShadow="lg">
                            <VStack mb={{base: 5, md: 0}} boxSize="100%"><DragHandleIcon boxSize="25px" /><p style={{fontSize: "17px"}}>{application.team.name === "" ? 'Kein Team' : application.team.name}</p></VStack>
                            <Button mt={3} rightIcon={<ArrowForwardIcon />} colorScheme="teal" variant="outline">Erstellen</Button>
                        </Flex>
                    ))
                }
                {
                    approvedApplications.length < 1 ? <p>Keine Daten vorhanden.</p>
                    : approvedApplications.map(application => (
                        <Flex p={5} m={5} flexDirection="column" justifyContent="space-between" alignItems="center" borderWidth={1} borderRadius={8} boxShadow="lg">
                            <VStack mb={{base: 5, md: 0}} boxSize="100%"><DragHandleIcon boxSize="25px" /><p style={{fontSize: "17px"}}>{application.team.name === "" ? 'Kein Team' : application.team.name}</p></VStack>
                        </Flex>
                    ))
                }
            </Box>
            <Spacer />
            <Box p={5} width={{base: "100%", md: "68%"}} overflowX="hidden" borderWidth={1} borderRadius={8} boxShadow="lg" flexDirection="column" align="center" justifyContent="center">
                <Heading size="lg">Ausstehende Anmeldungen</Heading>
                {
                    pendingApplications.length < 1 ? <p>Keine Daten vorhanden.</p>
                    : pendingApplications.map(application => (
                        <Flex p={5} my={5} flexDirection={{base: "column", md: "row"}} justifyContent="space-between" alignItems="center" borderWidth={1} borderRadius={8} boxShadow="lg">
                            <VStack mb={{base: 5, md: 0}} boxSize="25%"><AtSignIcon boxSize="20px" /><p>{application.name}</p></VStack>
                            <VStack mb={{base: 5, md: 0}} boxSize="25%"><EmailIcon boxSize="20px" /><p>{application.email}</p></VStack>
                            <VStack mb={{base: 5, md: 0}} boxSize="25%"><DragHandleIcon boxSize="20px" /><p>{application.team.name === "" ? 'Kein Team' : application.team.name}</p></VStack>
                            <Button rightIcon={<ArrowForwardIcon />} colorScheme="teal" variant="outline">Annehmen</Button>
                        </Flex>
                    ))
                }
                <Heading mt={50} size="lg">Angenommene Anmeldungen</Heading>
                {
                    approvedApplications.length < 1 ? <p>Keine Daten vorhanden.</p>
                    : approvedApplications.map(application => (
                        <Flex p={5} my={5} flexDirection={{base: "column", md: "row"}} justifyContent="space-between" alignItems="center" borderWidth={1} borderRadius={8} boxShadow="lg">
                            <VStack mb={{base: 5, md: 0}} boxSize="20%"><AtSignIcon boxSize="20px" /><p>{application.name}</p></VStack>
                            <VStack mb={{base: 5, md: 0}} boxSize="20%"><EmailIcon boxSize="20px" /><p>{application.email}</p></VStack>
                            <VStack mb={{base: 5, md: 0}} boxSize="20%"><DragHandleIcon boxSize="20px" /><p>{application.team.name === "" ? 'Kein Team' : application.team.name}</p></VStack>
                        </Flex>
                    ))
                }
            </Box>
            <Spacer />
        </Flex>
    )
}

export default AdminDashboard