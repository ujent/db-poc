package settings

type EnvType int

const (
	Invalid EnvType = iota
	OnPrem
	Kuber
)

type ServerSettings struct {
	Port    string
	EnvType EnvType
}

// ToDo - remove from here!
type DBType int

const (
	InvalidDialectDB DBType = iota
	MySQLDBType
	PostgresDBType
	SqliteDBType
)
