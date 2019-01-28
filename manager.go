package sync

import "sync"

type Manager struct {
	D <-chan struct{}
	m *manager
}

type manager struct {
	d            chan<- struct{}
	signaled     bool
	workAmount   int
	workAmountMu sync.Mutex
}

func New() *Manager {
	doneC := make(chan struct{})
	return &Manager{
		D: doneC,
		m: &manager{
			d:          doneC,
			workAmount: 0,
		},
	}
}

func (m *Manager) Set(amount int) {
	m.m.workAmountMu.Lock()
	defer m.m.workAmountMu.Unlock()

	m.m.workAmount = amount
}

func (m *Manager) Add(amount int) {
	m.m.workAmountMu.Lock()
	defer m.m.workAmountMu.Unlock()

	m.m.workAmount += amount
}

func (m *Manager) Done() {
	m.m.workAmountMu.Lock()
	defer m.m.workAmountMu.Unlock()

	m.m.workAmount--
	if m.m.workAmount == 0 && !m.m.signaled {
		go m.m.signalDone()
		m.m.signaled = true
	}
}

func (m *manager) signalDone() {
	m.d <- struct{}{}
}
