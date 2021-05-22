import { Button } from "@chakra-ui/button"
import { DragHandleIcon, EditIcon } from "@chakra-ui/icons"
import { Image } from "@chakra-ui/image"
import { Box, Divider, Flex, Heading, Text } from "@chakra-ui/layout"
import { useContext } from "react"
import { UserContext } from "../context/UserContext"

const UserDashboard = () => {
    const { user, guild } = useContext(UserContext)

    return (
        <Flex mt={40} flexDirection="column" align="center" justifyContent="center">
            <Flex p={8} maxW={{ base: "90%", md: "600px" }} borderWidth={1} borderRadius={8} boxShadow="lg" flexDirection="column" align="center" justifyContent="center" >
                <Box textAlign="center" width="80%">
                    <Heading>Hallo, {user.username}!</Heading>
                    {guild.user_is_member && <Text p={5}>Du bist bereits auf unserem Discord! Sehr gut! Bitte wähle aus, ob du dich für den Wettbewerb anmelden möchtest oder dir dein Team zuweisen möchtest.</Text>}
                    {!guild.user_is_member && <Text p={5}>Du scheinst noch nicht auf dem Discord-Server zu sein... Um dich (über den Bot) anmelden zu können und/oder dein Team wählen zu können, tritt bitte dem Discord-Server bei und logge dich mit einem Discord-Account hier ein, der auf dem Server ist.</Text>}
                </Box>
                <Divider mt={3} mb={3} />
                {guild.user_is_member &&
                    (<Box my={4}>
                        <Button width="100%" color="#fff" onClick={e => {window.location.href= "http://localhost:5000/api/auth"}} size="md">
                            <EditIcon mr={3} />
                            Für den Wettbewerb anmelden
                        </Button>
                        <Divider mt={3} mb={3} />
                        <Button width="100%" color="#fff" onClick={e => {window.location.href= "http://localhost:5000/api/auth"}} size="md">
                            <DragHandleIcon mr={3} />
                            Team wählen
                        </Button>
                    </Box>)
                }
                {!guild.user_is_member &&
                    (<Box my={4}>
                        <Button width="100%" color="#fff" onClick={e => {window.location.href= "http://localhost:5000/api/auth"}} size="md">
                            <Image boxSize="30px" mr={3} alt="" src="/assets/Discord-Logo-White.png" />
                            Dem Discord-Server beitreten
                        </Button>
                    </Box>)
                }
            </Flex>
        </Flex>
    )
}

export default UserDashboard
