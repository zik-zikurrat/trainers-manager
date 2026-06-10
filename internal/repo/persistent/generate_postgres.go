package persistent

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"trainers-manager/internal/entity"
	"trainers-manager/internal/repo"
	"trainers-manager/pkg/postgres"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

// GeneratorRepo -.
type GeneratorRepo struct {
	*postgres.Posgtres
}

// New -.
func NewGeneratorRepo(pg *postgres.Posgtres) *GeneratorRepo {
	return &GeneratorRepo{pg}
}

func (r *GeneratorRepo) CreateTraining(ctx context.Context) (uuid.UUID, error) {
	var id uuid.UUID
	if err := r.Pool.QueryRow(ctx, insertTrainingQuery).Scan(&id); err != nil {
		return uuid.Nil, fmt.Errorf("create training: %w", err)
	}
	return id, nil
}
func (r *GeneratorRepo) RecentPlans(ctx context.Context, groupID uuid.UUID, limit int) ([]entity.TrainingPlan, error) {
	rows, err := r.Pool.Query(ctx, getRecentPlans, groupID, limit)
	if err != nil {
		return nil, fmt.Errorf("recent plans: %v", err)
	}
	defer rows.Close()
	out := make([]entity.TrainingPlan, 0, limit)
	for rows.Next() {
		var t entity.TrainingPlan
		if err := rows.Scan(&t.Accent, &t.Skills, &t.CreatedAt); err != nil {
			return nil, fmt.Errorf("recent plans scan: %w", err)
		}
		out = append(out, t)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("recent plans rows: %w", err)
	}
	return out, nil
}
func (r *GeneratorRepo) GetGroupByName(ctx context.Context, name string) (entity.TrainingGroup, error) {
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
func (r *GeneratorRepo) GetStructure(ctx context.Context, structureID uuid.UUID) (entity.TrainingStructure, error) {
	var s entity.TrainingStructure
	err := r.Pool.QueryRow(ctx, getTrainingStructure, structureID).Scan(&s.ID, &s.Structure, &s.CreatedAt, &s.UpdatedAt)
	if errors.Is(err, pgx.ErrNoRows) {
		return entity.TrainingStructure{}, repo.ErrNotFound
	}
	if err != nil {
		return entity.TrainingStructure{}, fmt.Errorf("get training structure: %w", err)
	}
	return s, nil
}
func (r *GeneratorRepo) ListExercises(ctx context.Context) ([]entity.Exercise, error) {
	rows, err := r.Pool.Query(ctx, listExercisesQuery)
	if err != nil {
		return nil, fmt.Errorf("list exercises: %w", err)
	}
	defer rows.Close()

	out := make([]entity.Exercise, 0, _defaultEntityCap)
	for rows.Next() {
		var e entity.Exercise
		if err := rows.Scan(&e.ID, &e.Muscle, &e.Description, &e.CreatedAt, &e.UpdatedAt); err != nil {
			return nil, fmt.Errorf("list exercises scan: %w", err)
		}
		out = append(out, e)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("list exercises rows: %w", err)
	}
	return out, nil
}
func (r *GeneratorRepo) LinkExercises(ctx context.Context, trainingID uuid.UUID, exerciseIDs []uuid.UUID) error {
	if len(exerciseIDs) == 0 {
		return nil
	}
	if _, err := r.Pool.Exec(ctx, linkExercisesQuery, trainingID, exerciseIDs); err != nil {
		return fmt.Errorf("link exercises: %w", err)
	}
	return nil
}

func (r *GeneratorRepo) StoreTrainingPlan(ctx context.Context, p entity.TrainingPlan) (uuid.UUID, error) {
	tx, err := r.Pool.Begin(ctx)
	if err != nil {
		return uuid.Nil, fmt.Errorf("store plan begin: %w", err)
	}
	defer tx.Rollback(ctx)

	err = tx.QueryRow(ctx, insertTrainingPlanQuery,
		p.Plan, p.Status, p.TrainID, p.GroupID, p.Accent, p.Skills, p.TrainingStructureID,
	).Scan(&p.ID, &p.CreatedAt, &p.UpdatedAt)
	if err != nil {
		return uuid.Nil, fmt.Errorf("store plan insert: %w", err)
	}

	snapshot, err := json.Marshal(p)
	if err != nil {
		return uuid.Nil, fmt.Errorf("store plan marshal: %w", err)
	}
	if _, err := tx.Exec(ctx, insertPlanHistoryQuery, p.ID, entity.HistoryActionCreate, snapshot); err != nil {
		return uuid.Nil, fmt.Errorf("store plan history: %w", err)
	}

	if err := tx.Commit(ctx); err != nil {
		return uuid.Nil, fmt.Errorf("store plan commit: %w", err)
	}
	return p.ID, nil
}
