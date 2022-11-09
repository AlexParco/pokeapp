import { Fav } from "@/types/fav.type";
import { Pokemon } from "@/types/pokemon.type";
import { Result, SimplePokemon } from "../types/pagination.type";

// TODO: img so slow 
export const mapPokemons = (pokemonList: Result[], favPokemons: Fav[]): SimplePokemon[] => {
    let uptPokemonList = pokemonList.map(({ name, url }) => {
        const splitUrl = url.split("/")
        const id = splitUrl[splitUrl.length - 2]
        const picture = `https://raw.githubusercontent.com/PokeAPI/sprites/master/sprites/pokemon/other/official-artwork/${id}.png`
        const follow = favPokemons.some(e => e.pokemon_id.toString() == id)
        return { id, name, picture, follow }
    })

    return uptPokemonList
}

export const setPictureUrl = (pokemon: Pokemon): Pokemon => {
    pokemon.picture = `https://raw.githubusercontent.com/PokeAPI/sprites/master/sprites/pokemon/other/official-artwork/${pokemon.id}.png`
    return pokemon

}