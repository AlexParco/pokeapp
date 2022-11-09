import { Card } from "@/components"
import { usePokemons } from "@/context/pokemon.context"
import { SimplePokemon } from "@/types/pagination.type"
import { ArrowBackIcon, ArrowForwardIcon } from "@chakra-ui/icons"
import { Box, Flex, IconButton } from "@chakra-ui/react"

const CardList = ({ pokemons }: { pokemons: SimplePokemon[] }) => {
  const { page, setPage } = usePokemons()

  return (
    <>
      <Box m={8}  >
        {
          (page != 0) &&
          <IconButton
            aria-label="prev page"
            onClick={() => setPage(page - 20)}
            type="submit"
            icon={<ArrowBackIcon />}
            mx={2}
          />
        }
        <IconButton
          mx={2}
          aria-label="next page"
          onClick={() => setPage(page + 20)}
          type="submit"
          icon={<ArrowForwardIcon />}
        />
      </Box>

      <Flex
        mb={8}
        flexWrap={'wrap'}
        justifyContent={'center'}
        gap={30}
        w="100%"
      >
        {
          pokemons.map(pokemon =>
            <Card
              key={pokemon.id}
              {...pokemon}
            />
          )
        }
      </Flex>
    </>
  )

}

export default CardList


