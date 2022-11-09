import { Fav } from "@/types/fav.type"
import { UserLogin, UserRegister, UserWithToken } from "@/types/user.type"
import { format } from "path"
import { Comment } from "../types/comment.type"
import { PokemonPagination, SimplePokemon } from "../types/pagination.type"
import { Pokemon } from "../types/pokemon.type"

const API_URL = "http://localhost:9999/api/v1/"

export function getPokemon(keyword: string): Promise<Pokemon> {
  return fetch(`https://pokeapi.co/api/v2/pokemon/${keyword}`)
    .then(resp => {
      if (!resp.ok) {
        throw new Error(resp.statusText)
      }
      return resp.json()
    })
}

export function getPokemons(offset: string): Promise<PokemonPagination> {
  return fetch(`https://pokeapi.co/api/v2/pokemon?limit=20&offset=${offset}`)
    .then(resp => {
      if (!resp.ok) {
        throw new Error(resp.statusText)
      }
      return resp.json() as Promise<PokemonPagination>
    })
}

export function loginService({ username, password }: UserLogin): Promise<UserWithToken> {
  return fetch(API_URL + `auth/login`, {
    method: 'POST',
    body: JSON.stringify({ username, password })
  })
    .then(resp => {
      if (!resp.ok) {
        throw new Error(resp.statusText)
      }
      return resp.json() as Promise<UserWithToken>
    })
}

export function registerService({ username, password, role }: UserRegister): Promise<UserWithToken> {
  return fetch(API_URL + 'auth/register', {
    method: 'POST',
    body: JSON.stringify({ username, password, role })
  })
    .then(resp => {
      if (!resp.ok) {
        throw new Error(resp.statusText)
      }
      return resp.json() as Promise<UserWithToken>
    })
}

export function addFavPokemonService({ pokemon_id }: Fav, token: string): Promise<Fav> {
  return fetch(API_URL + 'fav/', {
    method: 'POST',
    headers: {
      Authorization: `Bearer ${token}`
    },
    body: JSON.stringify({ pokemon_id })
  })
    .then(resp => {
      if (!resp.ok) {
        throw new Error(resp.statusText)
      }
      return resp.json() as Promise<Fav>
    })
}

export function deleteFavPokemon(token: string, id: string) {
  return fetch(API_URL + `fav/${id}`, {
    method: 'DELETE',
    headers: {
      Authorization: `Bearer ${token}`
    }
  })
    .then(resp => {
      if (!resp.ok) {
        throw new Error(resp.statusText)
      }
    })
}

export function getFavPokemons(token: string): Promise<Fav[]> {
  return fetch(API_URL + 'fav/', {
    headers: {
      Authorization: `Bearer ${token}`,
      'Content-Type': 'application/json'
    }
  })
    .then(resp => {
      if (!resp.ok) {
        throw new Error(resp.statusText)
      }
      return resp.json() as Promise<Fav[]>
    })
}

export function getCommentsByPokeId(id: string): Promise<Comment[]> {
  return fetch(`http://localhost:9999/api/v1/comment?pokeid=${id}`)
    .then(resp => {
      if (!resp.ok) {
        throw new Error(resp.statusText)
      }
      return resp.json() as Promise<Comment[]>
    })
}

export function addCommentService({ pokemon_id, body }: Comment, token: string): Promise<Comment> {
  return fetch(API_URL + 'comment', {
    method: "POST",
    headers: {
      Authorization: `Bearer ${token}`
    },
    body: JSON.stringify({ pokemon_id, body }),
  })
    .then(resp => {
      if (!resp.ok) {
        throw new Error(resp.statusText)
      }
      return resp.json() as Promise<Comment>
    })
}