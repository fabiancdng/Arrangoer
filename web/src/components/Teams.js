import { Box, Heading } from "@chakra-ui/layout"
import Team from "./Team"

const Teams = ({ pendingApplications, approvedApplications }) => {
    return (
        <Box p={5} width={{base: "100%", md: "23%"}} mr={5} overflowX="hidden" borderWidth={1} borderRadius={8} boxShadow="lg" flexDirection="column" align="center" justifyContent="center">
            <Heading size="lg">Teams</Heading>
            {
                pendingApplications.length < 1 ? <p>Keine ausstehenden Teams vorhanden.</p>
                : pendingApplications.map((application, index) => (
                    <Team key={index} team={application.team} approved={false} />
                ))
            }
            {
                approvedApplications.length < 1 ? <p>Keine angenommenen Teams vorhanden.</p>
                : approvedApplications.map((application, index) => (
                    <Team key={index} team={application.team} approved={true} />
                ))
            }
        </Box>
    )
}

export default Teams
