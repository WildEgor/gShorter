package router

import (
	"github.com/WildEgor/gShorter/internal/handlers"
	"github.com/google/wire"
)

var RouterSet = wire.NewSet(NewRouter, handlers.HandlersSet)
