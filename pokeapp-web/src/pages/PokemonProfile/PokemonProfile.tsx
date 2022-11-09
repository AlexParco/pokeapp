import { Logout } from "@/components/Logout";
import { getPokemon } from "@/service/api";
import { Pokemon } from "@/types/pokemon.type";
import { setPictureUrl } from "@/utils/mapPokemos";
import { Box, Flex, Heading, Image, Tag, Text } from "@chakra-ui/react";
import { useEffect, useState } from "react";
import { useParams } from "react-router-dom";
import { Comment } from "./Comment";


const PokemonProfile = () => {
  const { id } = useParams<{ id: string }>() as { id: string }
  const [pokemon, setPokemon] = useState<Pokemon>({
    id: "",
    name: "",
    abilities: [],
    types: [],
    stats: [],
    picture: "",
  });

  useEffect(() => {
    getPokemon(id)
      .then(data => {
        setPokemon(setPictureUrl(data))
      })
  }, [id])

  return (
    <Flex
      flexWrap='wrap'
      justify='center'
      minHeight="100vh"
      py={12}
      gap={10}
    >
      <Logout />
      <Box >
        <Heading
          size='lg'
        >
          {pokemon.name}
        </Heading>
        <Image
          boxSize='400'
          src={pokemon.picture}
          alt={pokemon.name}
        />
        <Flex gap="10" justify=''>
          <Box>
            <Tag bg='teal.100' size='lg'>Types</Tag>
            {
              pokemon.types.map((poke, i) =>
                <Text key={i} >{poke.type.name}</Text>
              )
            }
          </Box>
          <Box>
            <Tag bg='blue.100' size='lg'>Abilities</Tag>
            {
              pokemon.abilities.map((poke, i) =>
                <Text key={i}>{poke.ability.name}</Text>
              )
            }
          </Box>
          <Flex direction='column' w='50%'>
            <Tag bg='red.100' size='lg'>Stats</Tag>
            {
              pokemon.stats.map((poke, i) =>
                <Flex key={i} gap={10} pl={2}>
                  <Text>{poke.base_stat}</Text>
                  <Text>{poke.stat.name}</Text>
                </Flex>
              )
            }
          </Flex>
        </Flex>
      </Box >
      <Box>
        <Comment
          pokeId={id}
        />
      </Box>
    </Flex >
  )

}

export default PokemonProfile