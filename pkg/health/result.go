package health

import "sync"

type Result map[string]error

func (r Result) IsHealthy() (healthy bool) {
	for _, e := range r {
		if e != nil {
			return false
		}
	}
	return true
}

func (r Result) CheckResult(name string) (knownCheck bool, result error) {
	result, knownCheck = r[name]
	return
}

type ResultWriter interface {
	WriteResult(checkName string, result error)
	GetResult() Result
}

func NewResultWriter() ResultWriter {
	return &resultWriter{
		lock:   new(sync.Mutex),
		result: Result{},
	}
}

type resultWriter struct {
	lock   sync.Locker
	result Result
}

func (r *resultWriter) WriteResult(checkName string, result error) {
	r.lock.Lock()
	defer r.lock.Unlock()

	r.result[checkName] = result
}

func (r resultWriter) GetResult() Result {
	return r.result
}
