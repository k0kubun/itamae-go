package resource

type Execute struct {
	Base
	Command string
}

func (e *Execute) Apply() {
	for _, action := range e.Action {
		if action == "run" {
			e.actionRun()
		}
	}
}

func (e *Execute) actionRun() {
	e.notifyApply()
	e.run(e.Command)
}

func (e *Execute) DryRun() {
	for _, action := range e.Action {
		if action == "run" {
			e.notifyApply()
		}
	}
}
