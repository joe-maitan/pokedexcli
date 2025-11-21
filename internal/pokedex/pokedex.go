package pokedex

import (
	// "time"
	"sync"
)

type Pokedex struct {
	data map[string]Pokemon
	mux   *sync.Mutex
} // End Pokedex struct

type Pokemon struct {
	Name           string `json:"name"`
	BaseExperience int    `json:"base_experience"`
	Height		   int 	  `json:"height"`
	Weight	       int    `json:"weight"`
} // End Pokemon struct

func (p *Pokedex) AddPokemon(name string, newPokemon Pokemon) {
	p.mux.Lock()
	defer p.mux.Unlock()

	p.data[name] = newPokemon
} // End AddPokemon() func

func (p *Pokedex) GetPokemon(pokemonName string) (Pokemon, bool) {
	p.mux.Lock()
	defer p.mux.Unlock()

	pokemon, ok :=  p.data[pokemonName]
	return pokemon, ok
} // End GetPokemon() func

func (p *Pokedex) DeletePokemon(pokemonName string) bool {
	return true
} // End DeletePokemon() func