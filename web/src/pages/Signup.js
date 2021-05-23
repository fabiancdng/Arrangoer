import { useState, useContext } from 'react'
import { Box, Button, Divider, Flex, Heading, Input, Text } from '@chakra-ui/react'
import { UserContext } from '../context/UserContext'

const Signup = () => {
    const { user } = useContext(UserContext)

    const [name, setName] = useState('')
    const [email, setEmail] = useState('')
    const [team, setTeam] = useState('')
    const [submitted, setSubmitted] = useState(1)

    const submitHandler = (e) => {
        e.preventDefault()
        setSubmitted(2)

        fetch("/api/application/submit", {
            method: 'POST',
            headers: {
              'Accept': 'application/json',
              'Content-Type': 'application/json'
            },
            body: JSON.stringify({name: name, email: email, team: team})
        })
            .then(res => {
                if(res.ok) {
                    setSubmitted(3)
                } else {
                    setSubmitted(4)
                }
            })
    }

    return (
        <Flex mt={10} mb={5} flexDirection="column" align="center" justifyContent="center">
            <Flex p={8} maxW={{ base: "90%", md: "800px" }} borderWidth={1} borderRadius={8} boxShadow="lg" flexDirection="column" align="center" justifyContent="center" >
                <Box textAlign="center" width="80%">
                    <Heading size="xl">Für den Wettbewerb anmelden</Heading>
                    <Text p={5}>Hier kannst du dich anmelden, {user.username}. Bitte gib deine richtigen Daten an und melde dich bitte nur ein Mal an. Teams können zur Not über den Discord nochmal abgesprochen und geändert werden.</Text>
                </Box>
                <Divider mt={3} mb={3} />
                <Box my={4} width="full">
                    <form onSubmit={e => submitHandler(e)}>
                        <Text mb={3}>Dein Name</Text>
                        <Input value={name} variant="outline" placeholder="Vor- und Nachname" type="text" onChange={e => setName(e.target.value)} required />
                        <Text my={3}>Dein E-Mail</Text>
                        <Input value={email} variant="outline" placeholder="E-Mail Adresse" type="email" onChange={e => setEmail(e.target.value)} required />
                        <Text my={3}>Dein Team-Name (optional)</Text>
                        <Input value={team} variant="outline" placeholder="Team-Name" type="text" onChange={e => setTeam(e.target.value)} />
                        <Text mt={5} fontSize="sm">
                            <b>Infos zu Teams: </b>
                            Du kannst deinen Team-Namen frei wählen. Wenn du dich schon mit jemandem abgesprochen hast, könnt ihr den gleichen Team-Namen hier eingeben.
                            Der Spielleiter muss deinen Team-Namen erst bestätigen und dann wird er dir automatisch zugewiesen.
                            Wenn du ein Einzelteilnehmer bist, kannst du dieses Feld auch leer lassen. Wenn du noch ein Team suchst, kannst du in den Discord schreiben.
                        </Text>
                        {
                            submitted === 1 ? <Input _hover={{cursor: "pointer"}} mt={5} type="submit" value="Anmeldung abschicken" />
                             : submitted === 2 ? <Button mt={5} isLoading loadingText="Wird gesendet...">...</Button>
                             : submitted === 3 ? <p mt={3} style={{color: "green"}}>Deine Anmeldung wurde abgeschickt!</p>
                             :  <p mt={3} style={{color: "red"}}>Deine Anmeldung konnte nicht abgeschickt werden! Warte einen Moment und versuche es erneut oder kontaktiere uns über Discord!</p>
                        }
                    </form>
                </Box>
            </Flex>
        </Flex>
    )
}

export default Signup