package main

import(
	"fmt"
)

var (
	id int = -1
)

type State struct {
	id int
	states []*State
	accepts rune
}

func NewState() *State {
	// reset global states
	id = -1
	return newState(-1)	
}

func newState(accepts rune) *State {
	stateArr := make([]*State, 0)
	id++

	return &State{
		id: id,
		states: stateArr,
		accepts: accepts,
	}
}

func (s *State) InsertKeyword(kw []rune) {
	currState := s

	// iterate over runes in keyword
	for _, r := range kw {
		foundValidState := false

		// iterate over states
		for _, state := range currState.states {
			// try insert substring into valid state
			if state.isValidTransition(r) {
				currState = state	
				foundValidState = true
				fmt.Printf("%s\n", state.ToString())		
				break
			}
		}
		// no valid transition for c make new state and insert
		if (foundValidState == false) {
			// create new state and append to currStates list of valid states
			newState := newState(r)
			currState.states = append(currState.states, newState)
			currState = newState
		}
	}
}

func (s *State) GetKeywordStates(kw []rune) []int {
	currState := s
	arr := []int{0} 
	for _, r := range kw {
		if currState.states != nil { 
			for _, state := range currState.states {
				if state.isValidTransition(r) {
					fmt.Printf("%s\n", state.ToString())
					arr = append(arr, state.id)
					currState = state
					break
				} else {
					fmt.Println("no valid transition")
				}
			} 
		}
	}	
	return arr
}


func (s *State) isValidTransition(r rune) bool {
	return s.accepts == r
}

func (s *State) ToString() string {
	return fmt.Sprintf("id:%d, accepts:%s", s.id, string(s.accepts))
}

func main() {
	fmt.Println("Aho Corasick Algorithm")
	
	root := NewState()

	root.InsertKeyword([]rune("car"))
	root.InsertKeyword([]rune("tim"))
	root.InsertKeyword([]rune("cards"))

	fmt.Printf("transitions: %v\n", root.GetKeywordStates([]rune("car")))
	fmt.Printf("transitions: %v\n", root.GetKeywordStates([]rune("tim")))
	fmt.Printf("transitions: %v\n", root.GetKeywordStates([]rune("cards")))
}
