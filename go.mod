module github.com/TrueBlocks/tokenomics.io

go 1.18

require (
	github.com/TrueBlocks/trueblocks-core/src/apps/chifra v0.0.0-20220331004001-55a0099f6ede
	github.com/ethereum/go-ethereum v1.10.16
	github.com/spf13/cobra v1.4.0
	github.com/spf13/pflag v1.0.5
)

require (
	github.com/alecthomas/participle/v2 v2.0.0-alpha7 // indirect
	github.com/araddon/dateparse v0.0.0-20210429162001-6b43995a97de // indirect
	github.com/btcsuite/btcd v0.22.0-beta // indirect
	github.com/deckarep/golang-set v1.8.0 // indirect
	github.com/fsnotify/fsnotify v1.5.1 // indirect
	github.com/go-ole/go-ole v1.2.6 // indirect
	github.com/go-stack/stack v1.8.1 // indirect
	github.com/gorilla/websocket v1.4.2 // indirect
	github.com/hashicorp/hcl v1.0.0 // indirect
	github.com/inconshreveable/mousetrap v1.0.0 // indirect
	github.com/magiconair/properties v1.8.5 // indirect
	github.com/mitchellh/mapstructure v1.4.2 // indirect
	github.com/panjf2000/ants/v2 v2.4.6 // indirect
	github.com/pelletier/go-toml v1.9.4 // indirect
	github.com/shirou/gopsutil v3.21.11+incompatible // indirect
	github.com/spf13/afero v1.6.0 // indirect
	github.com/spf13/cast v1.4.1 // indirect
	github.com/spf13/jwalterweatherman v1.1.0 // indirect
	github.com/spf13/viper v1.9.0 // indirect
	github.com/subosito/gotenv v1.2.0 // indirect
	github.com/tklauser/go-sysconf v0.3.9 // indirect
	github.com/tklauser/numcpus v0.4.0 // indirect
	github.com/yusufpapurcu/wmi v1.2.2 // indirect
	golang.org/x/crypto v0.0.0-20220213190939-1e6e3497d506 // indirect
	golang.org/x/sys v0.0.0-20220209214540-3681064d5158 // indirect
	golang.org/x/term v0.0.0-20210927222741-03fcf44c2211 // indirect
	golang.org/x/text v0.3.7 // indirect
	gopkg.in/ini.v1 v1.63.2 // indirect
	gopkg.in/natefinch/npipe.v2 v2.0.0-20160621034901-c1b8fa8bdcce // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
)

replace (
	github.com/TrueBlocks/tokenomics.io/pkg/types => ./pkg/types
	github.com/TrueBlocks/tokenomics.io/internal => ./internal
)
