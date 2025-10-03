package skipcmd

import (
	"SkipAdsV2/controller/userskipadshttp/httpmodel"
	"context"
)

func (cmd *Command) HandleEventCreatePackage(ctx context.Context, h *httpmodel.CreatePackageRequest) error {

	// create package
	pkg, games := h.ConvertToPackageAndPackageGames()
	cmd.db.CreatePackage(ctx, pkg, games)
	// create package app id
	return nil
}
