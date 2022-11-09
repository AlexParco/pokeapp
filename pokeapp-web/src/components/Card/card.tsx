import { Flex, Image, Text } from "@chakra-ui/react";
import { NavLink } from "react-router-dom";
import { SimplePokemon } from "../../types/pagination.type";
import FollowBtn from "./FollowBtn/followbtn";

const Card = ({ id, name, picture, follow }: SimplePokemon) => {

  return (
    <Flex
      maxW={'200'}
      flexDirection={"column"}
      pos='relative'
      height="100%"
      border='1px'
      borderRadius='lg'
      paddingBottom={2}
      paddingX={3}
    >
      <NavLink to={`/${id}`}
      >
        <Image
          boxSize='200px'
          src={picture}
          alt={name}
        />
        <Flex
          justify='center'
          align='center'
          textAlign='center'
        >
          <Text
            mt={2}
            pb={1}
          >
            {name}
          </Text>
        </Flex>
      </NavLink>
      <FollowBtn
        // TODO: revisar 
        follow={follow as boolean}
        idPokemon={id as string}
      />
    </Flex >
  )
}

export default Card