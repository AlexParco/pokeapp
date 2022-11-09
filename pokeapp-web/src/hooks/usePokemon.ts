import { setPictureUrl } from "@/utils/mapPokemos";
import { useEffect, useState } from "react"
import { getPokemon } from "../service/api";
import { Pokemon } from "../types/pokemon.type";

const usePokemon = (keyword: string) => {
  const [pokemon, setPokemon] = useState<Pokemon>({
    id: "",
    name: "",
    abilities: [],
    types: [],
    stats: [],
    picture: "",
  });

  useEffect(() => {
    getPokemon(keyword)
      .then(data => {
        setPokemon(setPictureUrl(data))
      })
  }, [keyword])

  return {
    pokemon
  }
}

export default usePokemon