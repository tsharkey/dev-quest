package quest

type Actionable interface {
	GetDependencies() []string
	IsComplete() bool
}

func findAvailable[V Actionable](m map[string]V) map[string]V {
	available := make(map[string]V)

mainloop:
	for key, val := range m {
		k := key
		v := val
		if v.IsComplete() {
			continue
		}

		if len(v.GetDependencies()) > 0 {
			for _, d := range v.GetDependencies() {
				dep := d
				if !m[dep].IsComplete() {
					continue mainloop
				}
			}
		}

		available[k] = v
	}

	return available
}
