// Process
package main

type Process struct {
	Import []Process
	Mount  []Service
}

func (process *Process) hasServices() bool {
	return len(process.Mount) != 0
}

func (process *Process) hasProcess() bool {
	return len(process.Import) != 0
}

func (process *Process) servicex() map[string]Service {

	servs := make(map[string]Service)

	if process.hasServices() {
		for _, service := range process.Mount {
			servs[service.Name] = service
		}
	}

	// Con proceso
	if process.hasProcess() {
		for _, proc := range process.Import {
			for id, srv := range proc.servicex() {
				servs[id] = srv
			}
		}
	}

	return servs
}
