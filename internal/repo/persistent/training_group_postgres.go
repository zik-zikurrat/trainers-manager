package persistent

import (
	"context"
	"errors"
	"fmt"

	"trainers-manager/internal/entity"
	"trainers-manager/internal/repo"
	"trainers-manager/internal/usecase/dto"
	"trainers-manager/pkg/postgres"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
)

const uniqueViolation = "23505"

// GroupRepo -.
type GroupRepo struct {
	*postgres.Posgtres
}

// New -.
func NewGroupRepo(pg *postgres.Posgtres) *GroupRepo {
	return &GroupRepo{pg}
}

func (r *GroupRepo) CreateGroup(ctx context.Context, g entity.TrainingGroup) (uuid.UUID, error) {
	var id uuid.UUID
	err := r.Pool.QueryRow(ctx, insertTrainingGroupQuery, g.Name, g.AccentCycle, g.SkillCycle).Scan(&id)
	if isUniqueViolation(err) {
		return uuid.Nil, repo.ErrAlreadyExists
	}
	if err != nil {
		return uuid.Nil, fmt.Errorf("insert training group: %w", err)
	}
	return id, nil
}

func (r *GroupRepo) ListGroups(ctx context.Context) ([]entity.TrainingGroup, error) {
	rows, err := r.Pool.Query(ctx, listTrainingGroupsQuery)
	if err != nil {
		return nil, fmt.Errorf("list training groups: %w", err)
	}
	defer rows.Close()

	out := make([]entity.TrainingGroup, 0, _defaultEntityCap)
	for rows.Next() {
		var g entity.TrainingGroup
		if err := rows.Scan(&g.ID, &g.Name, &g.AccentCycle, &g.SkillCycle); err != nil {
			return nil, fmt.Errorf("list training groups scan: %w", err)
		}
		out = append(out, g)
	}
	return out, rows.Err()
}

func (r *GroupRepo) UpdateGroup(ctx context.Context, g dto.UpdateGroupInput) error {
	ct, err := r.Pool.Exec(ctx, updateTrainingGroupQuery, g.Name, g.AccentCycle, g.SkillCycle, g.ID)
	if isUniqueViolation(err) {
		return repo.ErrAlreadyExists
	}
	if err != nil {
		return fmt.Errorf("update training group: %w", err)
	}
	if ct.RowsAffected() == 0 {
		return repo.ErrNotFound
	}
	return nil
}

func (r *GroupRepo) DeleteGroup(ctx context.Context, id uuid.UUID) error {
	ct, err := r.Pool.Exec(ctx, deleteTrainingGroupQuery, id)
	if err != nil {
		return fmt.Errorf("delete training group: %w", err)
	}
	if ct.RowsAffected() == 0 {
		return repo.ErrNotFound
	}
	return nil
}

func (r *GroupRepo) GetGroupByName(ctx context.Context, name string) (entity.TrainingGroup, error) {
	var g entity.TrainingGroup
	err := r.Pool.QueryRow(ctx, getTrainingGroupByNameQuery, name).
		Scan(&g.ID, &g.Name, &g.AccentCycle, &g.SkillCycle)
	if errors.Is(err, pgx.ErrNoRows) {
		return entity.TrainingGroup{}, repo.ErrNotFound
	}
	if err != nil {
		return entity.TrainingGroup{}, fmt.Errorf("get training group by name: %w", err)
	}
	return g, nil
}

func isUniqueViolation(err error) bool {
	var pgErr *pgconn.PgError
	return errors.As(err, &pgErr) && pgErr.Code == uniqueViolation
}
