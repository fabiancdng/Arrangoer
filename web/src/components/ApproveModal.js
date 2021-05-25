import { ArrowForwardIcon } from '@chakra-ui/icons'
import { Button, FormControl, FormLabel, Input, Modal, ModalBody, ModalCloseButton, ModalContent, ModalFooter, ModalHeader, ModalOverlay, useDisclosure } from '@chakra-ui/react'

const ApproveModal = ({ isApplication, application }) => {
    const { isOpen, onOpen, onClose } = useDisclosure()
  
    return (
      <>
        <Button mt={3} onClick={onOpen} rightIcon={<ArrowForwardIcon />} colorScheme="teal" variant="outline">Erstellen</Button>

        <Modal isOpen={isOpen} onClose={onClose}>
          <ModalOverlay />
          <ModalContent>
            <ModalHeader>{ isApplication ? 'Anmeldung' : 'Team ' } annehmen/bearbeiten</ModalHeader>
            <ModalCloseButton />
            <ModalBody pb={6}>
              <FormControl>
                <FormLabel>Name</FormLabel>
                <Input value={isApplication ? application.name : application.team.name} />
              </FormControl>
  
            { isApplication && 
                (<FormControl mt={4}>
                    <FormLabel>E-Mail</FormLabel>
                    <Input readOnly={true} value={application.email}  />
                </FormControl>)
            }

            </ModalBody>
  
            <ModalFooter>
              <Button colorScheme="blue" mr={3}>Speichern & Annehmen</Button>
              <Button>Ablehnen</Button>
            </ModalFooter>
          </ModalContent>
        </Modal>
      </>
    )
}

export default ApproveModal
