// Factory how does it work.
package caddyshack

type Config struct {
	Store map[string]Store //interface values do not req pointers.
}

func NewConfig() (config *Config) {

	config = &Config{
		Store: make(map[string]Store),
	}
	return
}

func (config *Config) GetConfig(name string) (status bool, store Store) {

	store, status = config.Store[name]
	return
}

func (config *Config) AddStore(store Store) bool {

	config.Store[store.GetName()] = store
	return true
}
