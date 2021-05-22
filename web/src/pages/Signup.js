import { useContext } from 'react'
import { Box, Divider, Flex, Heading, Input, Text } from '@chakra-ui/react'
import { UserContext } from '../context/UserContext'

const Signup = () => {
    const { user } = useContext(UserContext)

    return (
        <Flex mt={40} mb={5} flexDirection="column" align="center" justifyContent="center">
            <Flex p={8} maxW={{ base: "90%", md: "800px" }} borderWidth={1} borderRadius={8} boxShadow="lg" flexDirection="column" align="center" justifyContent="center" >
                <Box textAlign="center" width="80%">
                    <Heading size="xl">Für den Wettbewerb anmelden</Heading>
                    <Text p={5}>Hier kannst du dich anmelden, {user.username}. Bitte gib deine richtigen Daten an und melde dich bitte nur ein Mal an. Teams können zur Not über den Discord nochmal abgesprochen und geändert werden.</Text>
                </Box>
                <Divider mt={3} mb={3} />
                <Box my={4} width="full">
                    <form>
                        <Text mb={3}>Dein Name</Text>
                        <Input variant="outline" placeholder="Vor- und Nachname" type="text" required />
                        <Text my={3}>Dein E-Mail</Text>
                        <Input variant="outline" placeholder="E-Mail Adresse" type="email" required />
                        <Text my={3}>Dein Team-Name (optional)</Text>
                        <Input variant="outline" placeholder="Team-Name" type="text" required />
                        <Text mt={5} fontSize="sm">
                            <b>Infos zu Teams: </b>
                            Du kannst deinen Team-Namen frei wählen. Wenn du dich schon mit jemandem abgesprochen hast, könnt ihr den gleichen Team-Namen hier eingeben.
                            Der Spielleiter muss deinen Team-Namen erst bestätigen und dann wird er dir automatisch zugewiesen.
                            Wenn du ein Einzelteilnehmer bist, kannst du dieses Feld auch leer lassen. Wenn du noch ein Team suchst, kannst du in den Discord schreiben.
                        </Text>
                        <Input _hover={{cursor: "pointer"}} mt={5} type="submit" value="Anmeldung abschicken" />
                    </form>
                </Box>
            </Flex>
        </Flex>
    )
}

export default Signup
