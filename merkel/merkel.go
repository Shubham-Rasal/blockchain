package merkel

import "fmt"

//Merkel Tree functions
// set of functions to create a merkel tree
// 1. insert
// 2. lookup
// 3. delete
// 4. update

type MerkelNode struct {
	hash string
}

type MerkelTree struct {
	LevelCount int
	Levels     [][]MerkelNode
}

type MerkelTreeInterface interface {
	Init(values []string) bool
	Build()
	Insert(val string) bool
	Lookup(key string) string
	Delete(key string) bool
	Update(key string, val string) bool
}

// Init initializes the merkel tree based on the values passed
func (m *MerkelTree) Init(values []string) bool {
	// create a merkel tree
	var baseLevel []MerkelNode
	for _, val := range values {
		baseLevel = append(baseLevel, MerkelNode{hash: val})
	}

	//duplicate the last node if odd number of nodes
	if len(baseLevel)%2 != 0 {
		baseLevel = append(baseLevel, baseLevel[len(baseLevel)-1])
	}

	m.Levels = append(m.Levels, baseLevel)
	m.LevelCount++
	return true

}

// Build builds the merkel tree
func (m *MerkelTree) Build() {
	// build the merkel tree
	// 1. create a new level
	// 2. create new nodes from the previous level
	// 3. append the new level to the merkel tree
	// 4. increment the level count

	// create a new level
	var newLevel []MerkelNode

	// create new nodes from the previous level
	prevLevelLen := len(m.Levels[m.LevelCount-1])

	if prevLevelLen == 1 {
		return
	}

	//if odd number of nodes, duplicate the last node
	if prevLevelLen%2 != 0 {
		prevNode := m.Levels[m.LevelCount-1][prevLevelLen-1]
		fmt.Println(prevNode.hash)
		m.Levels[m.LevelCount-1] = append(m.Levels[m.LevelCount-1], prevNode)
		fmt.Println(m.Levels)
		prevLevelLen++
	}

	prevLevel := m.LevelCount - 1
	for i := 0; i < prevLevelLen; i += 2 {
		// create a new node
		var newNode MerkelNode

		// hash the two nodes
		newNode.hash = m.Levels[prevLevel][i].hash + m.Levels[prevLevel][i+1].hash

		// append the new node to the new level
		newLevel = append(newLevel, newNode)
	}

	// append the new level to the merkel tree
	m.Levels = append(m.Levels, newLevel)

	// increment the level count
	m.LevelCount++

	//call build recursively until the length of the last level is 1
	m.Build()

}

func (m *MerkelTree) Insert(val string) bool {
	// insert a value into the merkel tree
	// 1. insert the value into the base level
	// 2. build the merkel tree

	// insert the value into the base level
	//check if the last node is a duplicate
	baseLevelLen := len(m.Levels[0])
	last := m.Levels[0][baseLevelLen-1]
	secondLast := m.Levels[0][baseLevelLen-2]
	if last.hash == secondLast.hash {
		m.Levels[0][baseLevelLen-1].hash = val
	} else {
		m.Levels[0] = append(m.Levels[0], MerkelNode{hash: val})
	}

	//clear the levels except the base level
	m.Levels = m.Levels[:1]
	m.LevelCount = 1

	// build the merkel tree
	m.Build()

	return true

}

func (m *MerkelTree) Lookup(key string) string {

	// lookup a value in the merkel tree
	// 1. lookup the value in the base level
	// 2. if found, return the value
	// 3. if not found, return nil

	// lookup the value in the base level
	for _, node := range m.Levels[0] {
		if node.hash == key {
			return node.hash
		}
	}

	// if not found, return nil
	return "nf"
}

func (m *MerkelTree) Delete(key string) bool {
	// delete a value from the merkel tree
	// 1. lookup the value in the base level
	// 2. if found, delete the value
	// 3. if not found, return false

	// lookup the value in the base level
	for i := len(m.Levels[0]) - 1; i >= 0; i-- {
		node := m.Levels[0][i]
		if node.hash == key {
			// delete the value from the base level
			//if the node is the last node, delete the node and the previous node
			if i == len(m.Levels[0])-1 {
				fmt.Println("last node")
				m.Levels[0] = m.Levels[0][:i-1]
			} else {
				m.Levels[0] = append(m.Levels[0][:i], m.Levels[0][i+1:]...)
			}

			//check if tree is empty
			if len(m.Levels[0]) == 0 {
				m.LevelCount = 0
				m.Levels = nil
				return true
			}

			//clear the levels except the base level
			m.Levels = m.Levels[:1]
			m.LevelCount = 1

			// build the merkel tree
			m.Build()


			return true
		}
	}
	// if not found, return false
	return false

}
