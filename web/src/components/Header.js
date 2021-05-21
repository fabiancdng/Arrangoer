import { MoonIcon, SunIcon } from "@chakra-ui/icons"
import { Box } from "@chakra-ui/layout"
import { IconButton, useColorMode } from '@chakra-ui/react'

const Header = () => {
    const { colorMode, toggleColorMode } = useColorMode();

    return (
        <Box textAlign="right" p={8}>
            <IconButton
            icon={colorMode === 'light' ? <MoonIcon /> : <SunIcon />}
            onClick={toggleColorMode}
            variant="ghost"
            />
        </Box>
    )
}

export default Header