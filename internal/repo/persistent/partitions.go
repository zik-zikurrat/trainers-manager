package persistent

import (
	"context"
	"fmt"

	"trainers-manager/pkg/postgres"
)

// PartitionRepo -.
type PartitionRepo struct {
	*postgres.Posgtres
}

// New -.
func NewPartitionRepo(pg *postgres.Posgtres) *PartitionRepo {
	return &PartitionRepo{pg}
}

func (r *PartitionRepo) EnsureHistoryPartitions(ctx context.Context, ahead int) error {
	if _, err := r.Pool.Exec(ctx, "SELECT ensure_history_partitions($1)", ahead); err != nil {
		return fmt.Errorf("ensure history partitions: %w", err)
	}
	return nil
}
