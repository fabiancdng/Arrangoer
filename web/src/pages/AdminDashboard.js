import React from 'react'
import { Flex, Spacer} from '@chakra-ui/react'
import Applications from '../components/Applications'
import Teams from '../components/Teams'

const AdminDashboard = () => {
    return (
        <Flex flexDirection={{base: "column", md: "row"}} mx={10}>
            <Teams />
            <Spacer />
            <Applications />
            <Spacer />
        </Flex>
    )
}

export default AdminDashboard