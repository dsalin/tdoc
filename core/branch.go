// Represents a branch of the project
package core

type BranchType int8

const (
	InvalidBranch  BranchType = 0
	HashBranch     BranchType = 1
	SymbolicBranch BranchType = 2
)

func (r BranchType) String() string {
	switch r {
	case InvalidBranch:
		return "invalid-branch"
	case HashBranch:
		return "hash-branch"
	case SymbolicBranch:
		return "symbolic-branch"
	}

	return ""
}

type Branch struct {
	Parent *Branch
	Name   string
	Type   BranchType
}
