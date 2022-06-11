package quest

type Actionable interface {
	GetDependencies() []string
	IsComplete() bool
}
