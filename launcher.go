package udp

import (
	"sync"

	"github.com/elojah/services"
)

// Namespaces maps configs used for udp service with config file namespaces.
type Namespaces struct {
	UDP services.Namespace
}

// Launcher represents a udp launcher.
type Launcher struct {
	*services.Configs
	ns Namespaces

	mux *Mux
	m   sync.Mutex
}

// NewLauncher returns a new udp Launcher.
func (mux *Mux) NewLauncher(ns Namespaces, nsRead ...services.Namespace) *Launcher {
	return &Launcher{
		Configs: services.NewConfigs(nsRead...),
		mux:     mux,
		ns:      ns,
	}
}

// Up starts the udp service with new configs.
func (l *Launcher) Up(configs services.Configs) error {
	l.m.Lock()
	defer l.m.Unlock()

	cfg := Config{}
	if err := cfg.Dial(configs[l.ns.UDP]); err != nil {
		// Add namespace key when returning error with logrus
		return err
	}
	return l.mux.Dial(cfg)
}

// Down stops the udp service.
func (l *Launcher) Down(configs services.Configs) error {
	l.m.Lock()
	defer l.m.Unlock()

	return l.mux.Close()
}
