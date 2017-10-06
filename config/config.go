// Config is put into a different package to prevent cyclic imports in case
// it is needed in several locations

package config

type TracebeatConfig struct {
	Period  *int64
	Host    *string
	MaxHops *int
}

type ConfigSettings struct {
	Input TracebeatConfig
}
