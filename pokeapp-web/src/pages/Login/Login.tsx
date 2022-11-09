import { useAuth } from "@/context/auth.context"
import { Box, Button, Flex, FormControl, FormLabel, Heading, Input, Link } from "@chakra-ui/react"
import { useEffect, useState } from "react"
import { Link as ReachLink, useNavigate } from "react-router-dom"

const Login = () => {
  const [username, setUsername] = useState<string>("")
  const [password, setPassword] = useState<string>("")
  const { login, isLogged, token } = useAuth()
  const navegate = useNavigate()

  useEffect(() => {
    if (isLogged) {
      navegate("/")
    }
  }, [token, isLogged])

  const handleSubmit = (e: React.FormEvent<HTMLFormElement>) => {
    e.preventDefault()
    login({ username, password })
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
        <form onSubmit={handleSubmit}>
          <Heading>Login</Heading>
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
          <Box w='77%' textAlign='right' color='blue'>
            <Link as={ReachLink} to="/register" position='absolute'>
              register
            </Link>
          </Box>
        </form >
      </Flex >
    </>
  )
}

export default Login