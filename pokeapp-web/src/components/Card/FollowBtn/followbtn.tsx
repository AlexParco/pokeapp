import { usePokemons } from '@/context/pokemon.context'
import { AddIcon, CheckCircleIcon } from '@chakra-ui/icons'
import { IconButton } from '@chakra-ui/react'

const FollowBtn = ({ follow, idPokemon }: { follow: boolean, idPokemon: string }) => {
  const { followPokemon, unfollowPokemon } = usePokemons()

  const handleClick = (): void => {
    if (!follow) {
      followPokemon(idPokemon)
    } else {
      unfollowPokemon(idPokemon)
    }
  }

  return <IconButton
    position='absolute'
    bottom={1}
    right={2}
    aria-label='Follow Icon'
    borderRadius={'50%'}
    onClick={handleClick}
    icon={!follow ? <AddIcon /> : <CheckCircleIcon />}
  />

}

export default FollowBtn