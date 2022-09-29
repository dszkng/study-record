module example

go 1.18

require go.uber.org/zap v1.23.0
require mymath v0.0.0
replace mymath => ./mymath

require (
	go.uber.org/atomic v1.7.0 // indirect
	go.uber.org/multierr v1.6.0 // indirect
)
