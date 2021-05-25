import { Button } from "@chakra-ui/button"
import { ArrowForwardIcon, DragHandleIcon } from "@chakra-ui/icons"
import { Flex, VStack } from "@chakra-ui/layout"

const Team = ({ team, approved }) => {
    return (
        <Flex p={5} m={5} flexDirection="column" justifyContent="space-between" alignItems="center" borderWidth={1} borderRadius={8} boxShadow="lg">
            <VStack mb={{base: 5, md: 0}} boxSize="100%"><DragHandleIcon boxSize="25px" /><p style={{fontSize: "17px"}}>{team.name === "" ? 'Kein Team' : team.name}</p></VStack>
            { !approved && <Button mt={3} rightIcon={<ArrowForwardIcon />} colorScheme="teal" variant="outline">Erstellen</Button> }
        </Flex>
    )
}

export default Team
