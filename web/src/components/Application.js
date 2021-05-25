import { Button } from "@chakra-ui/button"
import { ArrowForwardIcon, AtSignIcon, DragHandleIcon, EmailIcon } from "@chakra-ui/icons"
import { Flex, VStack } from "@chakra-ui/layout"

const Application = ({ application, accepted }) => {
    return (
        <Flex p={5} my={5} flexDirection={{base: "column", md: "row"}} justifyContent="space-between" alignItems="center" borderWidth={1} borderRadius={8} boxShadow="lg">
            <VStack mb={{base: 5, md: 0}} boxSize="25%"><AtSignIcon boxSize="20px" /><p>{application.name}</p></VStack>
            <VStack mb={{base: 5, md: 0}} boxSize="25%"><EmailIcon boxSize="20px" /><p>{application.email}</p></VStack>
            <VStack mb={{base: 5, md: 0}} boxSize="25%"><DragHandleIcon boxSize="20px" /><p>{application.team.name === "" ? 'Kein Team' : application.team.name}</p></VStack>
            { !accepted && <Button rightIcon={<ArrowForwardIcon />} colorScheme="teal" variant="outline">Annehmen</Button> }
        </Flex>
    )
}

export default Application
