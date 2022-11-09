export interface Pokemon {
  id: string;
  name: string;
  abilities: Ability[];
  types: Type[];
  stats: Stat[];
  picture: string;
}

export interface Stat {
  base_stat: number;
  stat: Specie;
}

export interface Ability {
  ability: Specie;
}

export interface Type {
  slot: number;
  type: Specie;
}

export interface Specie {
  name: string;
  url: string;
}
