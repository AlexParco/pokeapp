export interface PokemonPagination {
  results: Result[];
}

export interface Result {
  name: string
  url: string
}

export interface SimplePokemon {
  id?: string;
  name?: string;
  picture?: string;
  follow?: boolean;
}
