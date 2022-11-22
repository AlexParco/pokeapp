import { useAuth } from "@/context/auth.context"
import { Box, Button, Flex, FormControl, FormLabel, Heading, Input } from "@chakra-ui/react"
import { useState } from "react"
import { useNavigate } from "react-router-dom"

const Register = () => {
  const [username, setUsername] = useState<string>("")
  const [password, setPassword] = useState<string>("")
  const { register } = useAuth()
  const navegate = useNavigate()


  const handleSubmit = (e: React.FormEvent<HTMLFormElement>) => {
    e.preventDefault()
    register({ username, role: 'ROLE_USER', password })
    navegate("/login")

    setUsername("")
    setPassword("")
  }

  return (
    <>
      <Flex
        minH='100vh'
        align='center'
        justify='center'
      >
        <form
          onSubmit={handleSubmit}
        >
          <Heading>Register</Heading>
          <FormControl mb={10}>
            <Box mt={10}>
              <FormLabel>Username</FormLabel>
              <Input
                type='text'
                value={username}
                onChange={(e) => setUsername(e.target.value)}
              />
            </Box>
          </FormControl>
          <FormControl>
            <Box my={10}>
              <FormLabel mt='8'>Password</FormLabel>
              <Input
                type='password'
                value={password}
                onChange={(e) => setPassword(e.target.value)}
              />
            </Box>
          </FormControl>
          <Button type="submit" >Enviar</Button>
        </form>
      </Flex>
    </>
  )
}

export default Register
