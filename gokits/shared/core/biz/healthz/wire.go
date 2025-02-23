package healthz

import "github.com/google/wire"

var ProviderSet = wire.NewSet(NewHealthZBiz)
