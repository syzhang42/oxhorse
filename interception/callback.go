package interception

/*
#include "interception.h"
*/
import "C"

import (
	"sync"
)

var funcMap sync.Map

//export InterceptionPredicateGo
func InterceptionPredicateGo(device C.InterceptionDevice) int {
	predicatecb, ok := funcMap.Load("predicateFunc")
	if ok {

		predicate, ok := predicatecb.(InterceptionPredicate)
		if ok {
			return predicate(InterceptionDevice(device))
		}
	}
	return 0
}
