import { Box, Heading } from "@chakra-ui/layout"
import Application from "./Application"

const Applications = ({ approvedApplications, pendingApplications }) => {
    return (
        <Box p={5} width={{base: "100%", md: "68%"}} overflowX="hidden" borderWidth={1} borderRadius={8} boxShadow="lg" flexDirection="column" align="center" justifyContent="center">
                <Heading size="lg">Ausstehende Anmeldungen</Heading>
                {
                    pendingApplications.length < 1 ? <p>Keine ausstehenden Anmeldungen vorhanden.</p>
                    : pendingApplications.map((application, index) => (
                        <Application key={index} application={application} accepted={false} />
                    ))
                }
                <Heading mt={50} size="lg">Angenommene Anmeldungen</Heading>
                {
                    approvedApplications.length < 1 ? <p>Keine angenommenen Anmeldungen vorhanden.</p>
                    : approvedApplications.map((application, index) => (
                        <Application key={index} application={application} accepted={true} />
                    ))
                }
        </Box>
    )
}

export default Applications
