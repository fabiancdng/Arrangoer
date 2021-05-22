import { MoonIcon, SunIcon } from "@chakra-ui/icons"
import { Heading, Flex } from "@chakra-ui/layout"
import { IconButton, Image, useColorMode } from '@chakra-ui/react'
import { useContext } from "react"
import { UserContext } from "../context/UserContext"

const Header = () => {
    const userContext = useContext(UserContext)
    const { colorMode, toggleColorMode } = useColorMode();

    return (
        <Flex justifyContent="space-between" p={6} width="full">
            <Flex alignItems="center">
                <Image mr={2} rounded={10} boxSize="50px" src={`/assets/favicon-32x32.png`} />
                <Heading fontSize="2xl">Arrangør</Heading>
            </Flex>
            <Flex alignItems="center">
                {userContext.user.loggedIn ? <Image mr={5} justifyContent="flex-end" rounded={10} boxSize="40px" src={`https://cdn.discordapp.com/avatars/${userContext.user.id}/${userContext.user.avatar}.png`} /> : ''}
                {userContext.user.loggedIn ? <Heading fontSize="1xl">{userContext.user.username}t</Heading> : ''}
                <IconButton
                    icon={colorMode === 'light' ? <MoonIcon /> : <SunIcon />}
                    onClick={toggleColorMode}
                    variant="ghost"
                />
            </Flex>
        </Flex>
    )
}

export default Header