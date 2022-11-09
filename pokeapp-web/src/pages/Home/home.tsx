import { Logout } from "@/components/Logout"
import { usePokemons } from "@/context/pokemon.context"
import { Box } from "@chakra-ui/react"
import { useEffect } from "react"
import { CardList } from "./CardList"

const Home = () => {
  const { pokemons } = usePokemons()

  return (
    <Box>
      <Logout />
      <CardList pokemons={pokemons} />
    </Box>
  )
}

export default Home