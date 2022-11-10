package main

import(
	"fmt"
)

var (
	id int = -1
	statesMap = make(map[int]*State, 0)
	outputMap = make(map[int][]string, 0)
	failureMap = make(map[int]int, 0)
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

	s := &State{
		id: id,
		states: stateArr,
		accepts: accepts,
	}

	// append to stateMapo
	statesMap[id] = s

	return s
}

func AhoCorasick(keywords []string, text string) {

	// build trie
	root := NewState()

	for _, kw := range keywords {
		root.InsertKeyword([]rune(kw))	
	}

	// build failure function
	root.BuildFailureFunction()

	// iterate over text
	state := 0
	for i, c := range []rune(text) {
		for {
			if GoTo(state, c) != -1 {
				break
			}

			state = failureMap[state]
		}			
		state = GoTo(state, c)

		if len(outputMap[state]) != 0 {
			fmt.Printf("i: %d\n", i)
			fmt.Printf("output: %v", outputMap[state])
		}  
	}
}

func GoTo(state int, char rune) int {
	s := statesMap[state]
	for _, state := range s.states {
		if state.isValidTransition(char) {
			return state.id
		} 
	} 

	// we should add a loop on state 0
	if state == 0 {
		return 0
	}

	// fail - we can formalise this later
	return -1
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

	// add word to output
	outputMap[currState.id] = append(outputMap[currState.id], string(kw))
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

func (s *State) BuildFailureFunction() {
	fmt.Println("outputmap: %v", outputMap)

	queue := make([]*State, 0)

	// initialise with d=1
	for _, s1 := range s.states {
		failureMap[s1.id] = 0
		queue = append(queue, s1)
		fmt.Printf("id: %d, s: %s\n", s1.id, string(s1.accepts))
	}

	fmt.Printf("queue: %v\n", queue)

	for {
		// terminate
		if len(queue) == 0 {
			break
		}

		// get first element, and pop from queue
		r := queue[0]
		queue = queue[1:]

		fmt.Printf("queue: %v\n", queue)


		// iterate though valid transitions for r
		for _, ss := range r.states {
			
			a := ss.accepts

			// append to queue
			queue = append(queue, ss)
			state := failureMap[r.id]

			
			fmt.Printf("a: %s\n", string(a))
			fmt.Printf("r: %d\n", r.id)
			fmt.Printf("s: %d\n", ss.id)
			fmt.Printf("state: %d, a: %s, goto: %d\n", state, string(a), GoTo(state, a))

			for {
				if GoTo(state, a) != -1 {
					break
				}  

				state = failureMap[state] 
			}

			failureMap[ss.id] = GoTo(state, a)
			outputMap[ss.id] = append(outputMap[ss.id], outputMap[failureMap[ss.id]]...)
		}
	}

	fmt.Println("failureMap: %v", failureMap)
	fmt.Println("outputmap: %v", outputMap)
}


func (s *State) isValidTransition(r rune) bool {
	return s.accepts == r
}

func (s *State) ToString() string {
	return fmt.Sprintf("id:%d, accepts:%s", s.id, string(s.accepts))
}

func main() {
	fmt.Println("Aho Corasick Algorithm")
	
	//root := NewState()

//root.InsertKeyword([]rune("he"))
//root.InsertKeyword([]rune("she"))
//root.InsertKeyword([]rune("his"))
//root.InsertKeyword([]rune("hers"))
//
//fmt.Printf("transitions: %v\n", root.GetKeywordStates([]rune("he")))
//fmt.Printf("transitions: %v\n", root.GetKeywordStates([]rune("she")))
//fmt.Printf("transitions: %v\n", root.GetKeywordStates([]rune("his")))
//fmt.Printf("transitions: %v\n", root.GetKeywordStates([]rune("hers")))
//
//
//// goto
//state := 0
//char := 'h'
//fmt.Printf("state: %d, char: %s, nextState: %d\n", state, string(char), GoTo(state, char))
//
//// output
//for key, elem := range outputMap {
//	fmt.Printf("key: %d, elements: %v\n", key, elem)
//}
//
//root.BuildFailureFunction()
	
	arr := []string{"he","she", "his", "hers"}
	text := "ushers"

	AhoCorasick(arr, text)

}
