import { useAuth } from "@/context/auth.context"
import { Button } from "@chakra-ui/react"

const Logout = () => {
  const { logout } = useAuth()

  return (
    <Button position='absolute' top='2' left='2' onClick={() => logout()}>Log out</Button>
  )
}

export default Logout