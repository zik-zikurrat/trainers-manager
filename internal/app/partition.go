package app

import (
	"context"
	"time"
	"trainers-manager/internal/repo"
	"trainers-manager/pkg/logger"
)

func StartPartitionMaintainer(ctx context.Context, r repo.TrainingRepo, l logger.Interface) {
	const interval = 24 * time.Hour

	ensure := func() {
		if err := r.EnsureHistoryPartitions(ctx, 1); err != nil {
			l.Error(err, "partition maintainer")
		}
	}
	ensure()

	go func() {
		t := time.NewTicker(interval)
		defer t.Stop()
		for {
			select {
			case <-ctx.Done():
				return
			case <-t.C:
				ensure()
			}
		}
	}()
}
