module github.com/giantswarm/microendpoint

go 1.23.0

toolchain go1.24.5

require (
	github.com/giantswarm/microerror v0.4.1
	github.com/giantswarm/micrologger v1.1.2
	github.com/giantswarm/versionbundle v1.1.0
	github.com/go-kit/kit v0.13.0
	github.com/prometheus/client_golang v1.22.0
)

require (
	github.com/beorn7/perks v1.0.1 // indirect
	github.com/cespare/xxhash/v2 v2.3.0 // indirect
	github.com/coreos/go-semver v0.3.1 // indirect
	github.com/go-kit/log v0.2.1 // indirect
	github.com/go-logfmt/logfmt v0.6.0 // indirect
	github.com/go-logr/logr v1.4.2 // indirect
	github.com/go-stack/stack v1.8.1 // indirect
	github.com/munnerz/goautoneg v0.0.0-20191010083416-a7dc8b61c822 // indirect
	github.com/prometheus/client_model v0.6.1 // indirect
	github.com/prometheus/common v0.62.0 // indirect
	github.com/prometheus/procfs v0.15.1 // indirect
	golang.org/x/net v0.41.0 // indirect
	golang.org/x/sync v0.16.0 // indirect
	golang.org/x/sys v0.34.0 // indirect
	google.golang.org/protobuf v1.36.5 // indirect
	gopkg.in/resty.v1 v1.12.0 // indirect
)

replace golang.org/x/net => golang.org/x/net v0.42.0

replace google.golang.org/protobuf v1.31.0 => google.golang.org/protobuf v1.33.0
