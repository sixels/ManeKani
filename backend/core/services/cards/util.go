package cards

import (
	"context"
	"fmt"
)

func checkResourceOwner[T any](
	ctx context.Context,
	resourceID T,
	userID string,
	queryOwner func(context.Context, T) (string, error),
) (bool, error) {
	owner, err := queryOwner(ctx, resourceID)
	if err != nil {
		return false, fmt.Errorf(
			"could not check wether the specified user owns the resource: %w", err,
		)
	}

	return owner == userID, nil
}
