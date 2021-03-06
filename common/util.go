package common

import (
	. "github.com/nwpc-oper/hpc-model-go"
)

func SortItems(items []Item, keys []string) {
	var lessFuncs []LessFunc
	for _, key := range keys {
		lessFuncs = append(lessFuncs, CreatePropertyLessFunc(key))
	}

	sorter := CreateSorter(lessFuncs...)
	sorter.Sort(items)
}
