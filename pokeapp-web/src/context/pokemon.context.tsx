import { addFavPokemonService, deleteFavPokemon, getFavPokemons, getPokemons } from "@/service/api";
import { Fav } from "@/types/fav.type";
import { SimplePokemon } from "@/types/pagination.type";
import { mapPokemons } from "@/utils/mapPokemos";
import { createContext, ReactNode, useContext, useEffect, useState } from "react";
import { useAuth } from "./auth.context";

interface IPokeContext {
  pokemons: SimplePokemon[];
  followPokemon: (id: string) => void;
  unfollowPokemon: (id: string) => void;
  setPage: (page: number) => void
  page: number;
}

const PokeContext = createContext<IPokeContext>({
  pokemons: [],
  followPokemon: () => null,
  unfollowPokemon: () => null,
  setPage: () => null,
  page: 0
})

export const PokeProvider = ({ children }: { children: ReactNode, token: String }) => {
  const [pokemons, setPokemons] = useState<SimplePokemon[]>([])
  const [page, setPage] = useState<number>(0)
  const { token } = useAuth()

  useEffect(() => {
    async function fetchData() {
      try {
        const data = await getPokemons(page.toString())
        // 
        let resp: Fav[] = []
        if (token) {
          resp = await getFavPokemons(token)
        }
        console.log(resp)
        const resPoke = mapPokemons(data.results, resp)

        setPokemons(resPoke)
      } catch (error) {
        console.log(error)
      }
    }
    fetchData()
  }, [setPokemons, token, page])

  const followPokemon = (id: string) => {
    addFavPokemonService({ pokemon_id: parseInt(id) }, token)
      .then(data => {
        const resPoke = pokemons.map(pokemon => {
          if (pokemon.id == data.pokemon_id.toString()) {
            pokemon.follow = true
          }
          return pokemon
        })
        setPokemons(resPoke)
      })
      .catch(err => console.log(err))
  }

  const unfollowPokemon = (id: string) => {
    deleteFavPokemon(token, id)
      .catch(err => console.log(err))
    const resPokemon = pokemons.map(pokemon => {
      if (pokemon.id == id) {
        pokemon.follow = false
      }
      return pokemon
    })
    setPokemons(resPokemon)
  }

  return (
    <PokeContext.Provider
      value={{
        pokemons,
        followPokemon,
        unfollowPokemon,
        setPage,
        page
      }}>
      {children}

    </PokeContext.Provider>
  )
}



export const usePokemons = () => {
  const { pokemons, followPokemon, unfollowPokemon, setPage, page } = useContext(PokeContext)
  return {
    pokemons,
    followPokemon,
    unfollowPokemon,
    setPage,
    page
  }
}

