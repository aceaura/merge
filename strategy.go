package merge

import "fmt"

type SliceStrategy int

const (
	SliceStrategyIgnore SliceStrategy = iota
	SliceStrategyAppend
	SliceStrategyRefer
	SliceStrategyReplace
	SliceStrategyReplaceElem
	SliceStrategyReplaceElemDynamic
	SliceStrategyReplaceDeep
	SliceStrategyReplaceDeepDynamic
)

var sliceStrategyNames = map[SliceStrategy]string{
	SliceStrategyIgnore:             "Ignore",
	SliceStrategyAppend:             "Append",
	SliceStrategyRefer:              "Refer",
	SliceStrategyReplace:            "Replace",
	SliceStrategyReplaceElem:        "ReplaceElem",
	SliceStrategyReplaceElemDynamic: "ReplaceElemDynamic",
	SliceStrategyReplaceDeep:        "ReplaceDeep",
	SliceStrategyReplaceDeepDynamic: "ReplaceDeepDynamic",
}

func (s SliceStrategy) String() string {
	if v, ok := sliceStrategyNames[s]; ok {
		return v
	}
	return fmt.Errorf("%w: %d", ErrInvalidStrategy, s).Error()
}

type StructStrategy int

const (
	StructStrategyIgnore StructStrategy = iota
	StructStrategyReplace
	StructStrategyReplaceElem
	StructStrategyReplaceDeep
)

var structStrategyNames = map[StructStrategy]string{
	StructStrategyIgnore:      "Ignore",
	StructStrategyReplace:     "Replace",
	StructStrategyReplaceElem: "ReplaceElem",
	StructStrategyReplaceDeep: "ReplaceDeep",
}

func (s StructStrategy) String() string {
	if v, ok := structStrategyNames[s]; ok {
		return v
	}
	return fmt.Errorf("%w: %d", ErrInvalidStrategy, s).Error()
}

type ArrayStrategy int

const (
	ArrayStrategyIgnore ArrayStrategy = iota
	ArrayStrategyReplace
	ArrayStrategyReplaceElem
	ArrayStrategyReplaceDeep
)

var arrayStrategyNames = map[ArrayStrategy]string{
	ArrayStrategyIgnore:      "Ignore",
	ArrayStrategyReplace:     "Replace",
	ArrayStrategyReplaceElem: "ReplaceElem",
	ArrayStrategyReplaceDeep: "ReplaceDeep",
}

func (s ArrayStrategy) String() string {
	if v, ok := arrayStrategyNames[s]; ok {
		return v
	}
	return fmt.Errorf("%w: %d", ErrInvalidStrategy, s).Error()
}

type ChanStrategy int

const (
	ChanStrategyIgnore ChanStrategy = iota
	ChanStrategyRefer
	ChanStrategyAppend
	ChanStrategyReplace
	ChanStrategyReplaceElem
	ChanStrategyReplaceElemDynamic
	ChanStrategyReplaceDeep
	ChanStrategyReplaceDeepDynamic
)

var chanStrategyNames = map[ChanStrategy]string{
	ChanStrategyIgnore:             "Ignore",
	ChanStrategyAppend:             "Append",
	ChanStrategyRefer:              "Refer",
	ChanStrategyReplace:            "Replace",
	ChanStrategyReplaceElem:        "ReplaceElem",
	ChanStrategyReplaceElemDynamic: "ReplaceElemDynamic",
	ChanStrategyReplaceDeep:        "ReplaceDeep",
	ChanStrategyReplaceDeepDynamic: "ReplaceDeepDynamic",
}

type MapStrategy int

const (
	MapStrategyIgnore MapStrategy = iota
	MapStrategyRefer
	MapStrategyReplace
	MapStrategyReplaceElem
	MapStrategyReplaceDeep
)

var mapStrategyNames = map[MapStrategy]string{
	MapStrategyIgnore:      "Ignore",
	MapStrategyRefer:       "Refer",
	MapStrategyReplace:     "Replace",
	MapStrategyReplaceElem: "ReplaceElements",
	MapStrategyReplaceDeep: "ReplaceDeep",
}
