import { MoonIcon, SunIcon } from "@chakra-ui/icons"
import { Heading, Flex } from "@chakra-ui/layout"
import { IconButton, Image, useColorMode } from '@chakra-ui/react'
import { useContext } from "react"
import { UserContext } from "../context/UserContext"

const Header = () => {
    const userContext = useContext(UserContext)
    const { colorMode, toggleColorMode } = useColorMode();

    return (
        <Flex alignItems="center" p={8} width="full">
            <Flex alignItems="center">
                <Image mb={1} rounded={10}  boxSize="50px" src={`/assets/favicon-32x32.png`} />
                <Heading fontSize="2xl">Arrangør</Heading>
            </Flex>
            <Flex alignItems="center" textAlign="right">
                <Image justifyContent="flex-end" rounded={10} boxSize="40px" src={`https://cdn.discordapp.com/avatars/${userContext.user.id}/${userContext.user.avatar}.png`} />
                <Heading fontSize="1xl">{userContext.user.username}</Heading>
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