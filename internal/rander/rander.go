package rander

import "math/rand"

// AdvanceRander package the base rand.Rand, and provide some advance func to get object rand.
type AdvanceRander struct {
	randInstance *rand.Rand
}

// NewAdvanceRander return new AdvanceRander with special seed
func NewAdvanceRander(seed int64) *AdvanceRander {
	return &AdvanceRander{
		randInstance: rand.New(rand.NewSource(seed)),
	}
}

// RandInterface will return random one element in value map by weigh, the input is element with weigh.
// The weigh will recalculate. The weigh will reset to 'one weigh / sum(weighs)'
func (a AdvanceRander) RandInterface(valueMap map[interface{}]int) interface{} {
	objList := make([]interface{}, 0, 100)

	for obj, value := range valueMap {
		for i := 0; i < value; i++ {
			objList = append(objList, obj)
		}
	}

	index := a.randInstance.Intn(len(objList))

	return objList[index]
}
