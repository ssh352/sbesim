package service

import "go.uber.org/fx"

// Module ...
var Module = fx.Provide(
	New,
)
