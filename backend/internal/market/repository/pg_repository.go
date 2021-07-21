package repository

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/opentracing/opentracing-go"
	"github.com/pkg/errors"

	"github.com/seregaa020292/capitalhub/internal/market"
	"github.com/seregaa020292/capitalhub/internal/market/model"
	"github.com/seregaa020292/capitalhub/pkg/utils"
)

// Market Repository
type marketRepo struct {
	db *sqlx.DB
}

// Market repository constructor
func NewMarketRepository(db *sqlx.DB) market.Repository {
	return &marketRepo{db: db}
}

// Create market
func (repository *marketRepo) Create(ctx context.Context, market *model.Market) (*model.Market, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "marketRepo.Create")
	defer span.Finish()

	var n model.Market
	if err := repository.db.QueryRowxContext(
		ctx,
		create,
		&market.Title,
		&market.Ticker,
		&market.Content,
		&market.ImageURL,
		&market.CurrencyID,
		&market.InstrumentID,
	).StructScan(&n); err != nil {
		return nil, errors.Wrap(err, "marketRepo.Create.QueryRowxContext")
	}

	return &n, nil
}

// Update market item
func (repository *marketRepo) Update(ctx context.Context, market *model.Market) (*model.Market, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "marketRepo.Update")
	defer span.Finish()

	var n model.Market
	if err := repository.db.QueryRowxContext(
		ctx,
		update,
		&market.Title,
		&market.Ticker,
		&market.Content,
		&market.ImageURL,
		&market.InstrumentID,
		&market.MarketID,
	).StructScan(&n); err != nil {
		return nil, errors.Wrap(err, "marketRepo.Update.QueryRowxContext")
	}

	return &n, nil
}

// Get single market by id
func (repository *marketRepo) GetByID(ctx context.Context, marketID uuid.UUID) (*model.MarketBase, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "marketRepo.GetByID")
	defer span.Finish()

	n := &model.MarketBase{}
	if err := repository.db.GetContext(ctx, n, getByID, marketID); err != nil {
		return nil, errors.Wrap(err, "marketRepo.GetByID.GetContext")
	}

	return n, nil
}

// Get single market by user id
func (repository *marketRepo) GetByUserID(ctx context.Context, userID uuid.UUID) (*[]model.MarketRegister, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "marketRepo.GetByUserID")
	defer span.Finish()

	marketModel := &[]model.MarketRegister{}
	if err := repository.db.SelectContext(ctx, marketModel, getByUserID, userID); err != nil {
		return nil, errors.Wrap(err, "marketRepo.GetByUserID.GetContext")
	}

	return marketModel, nil
}

// Delete market by id
func (repository *marketRepo) Delete(ctx context.Context, marketID uuid.UUID) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "marketRepo.Delete")
	defer span.Finish()

	result, err := repository.db.ExecContext(ctx, deleteById, marketID)
	if err != nil {
		return errors.Wrap(err, "marketRepo.Delete.ExecContext")
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return errors.Wrap(err, "marketRepo.Delete.RowsAffected")
	}
	if rowsAffected == 0 {
		return errors.Wrap(sql.ErrNoRows, "marketRepo.Delete.rowsAffected")
	}

	return nil
}

// Get markets
func (repository *marketRepo) GetAll(ctx context.Context, pq *utils.PaginationQuery) (*model.MarketList, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "marketRepo.GetAll")
	defer span.Finish()

	var totalCount int
	if err := repository.db.GetContext(ctx, &totalCount, getTotalCount); err != nil {
		return nil, errors.Wrap(err, "marketRepo.GetAll.GetContext.totalCount")
	}

	if totalCount == 0 {
		return &model.MarketList{
			TotalCount: totalCount,
			TotalPages: utils.GetTotalPages(totalCount, pq.GetSize()),
			Page:       pq.GetPage(),
			Size:       pq.GetSize(),
			HasMore:    utils.GetHasMore(pq.GetPage(), totalCount, pq.GetSize()),
			Markets:    make([]*model.MarketBase, 0),
		}, nil
	}

	var marketList = make([]*model.MarketBase, 0, pq.GetSize())
	rows, err := repository.db.QueryxContext(ctx, getAll, pq.GetOffset(), pq.GetLimit())
	if err != nil {
		return nil, errors.Wrap(err, "marketRepo.GetAll.QueryxContext")
	}
	defer rows.Close()

	for rows.Next() {
		n := &model.MarketBase{}
		if err = rows.StructScan(n); err != nil {
			return nil, errors.Wrap(err, "marketRepo.GetAll.StructScan")
		}
		marketList = append(marketList, n)
	}

	if err = rows.Err(); err != nil {
		return nil, errors.Wrap(err, "marketRepo.GetAll.rows.Err")
	}

	return &model.MarketList{
		TotalCount: totalCount,
		TotalPages: utils.GetTotalPages(totalCount, pq.GetSize()),
		Page:       pq.GetPage(),
		Size:       pq.GetSize(),
		HasMore:    utils.GetHasMore(pq.GetPage(), totalCount, pq.GetSize()),
		Markets:    marketList,
	}, nil
}

// Find market by title
func (repository *marketRepo) SearchByTitle(ctx context.Context, title string, query *utils.PaginationQuery) (*model.MarketList, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "marketRepo.SearchByTitle")
	defer span.Finish()

	var totalCount int
	if err := repository.db.GetContext(ctx, &totalCount, findByTitleCount, title); err != nil {
		return nil, errors.Wrap(err, "marketRepo.SearchByTitle.GetContext")
	}
	if totalCount == 0 {
		return &model.MarketList{
			TotalCount: totalCount,
			TotalPages: utils.GetTotalPages(totalCount, query.GetSize()),
			Page:       query.GetPage(),
			Size:       query.GetSize(),
			HasMore:    utils.GetHasMore(query.GetPage(), totalCount, query.GetSize()),
			Markets:    make([]*model.MarketBase, 0),
		}, nil
	}

	var marketList = make([]*model.MarketBase, 0, query.GetSize())
	rows, err := repository.db.QueryxContext(ctx, findByTitle, title, query.GetOffset(), query.GetLimit())
	if err != nil {
		return nil, errors.Wrap(err, "marketRepo.SearchByTitle.QueryxContext")
	}
	defer rows.Close()

	for rows.Next() {
		n := &model.MarketBase{}
		if err = rows.StructScan(n); err != nil {
			return nil, errors.Wrap(err, "marketRepo.SearchByTitle.StructScan")
		}
		marketList = append(marketList, n)
	}

	if err = rows.Err(); err != nil {
		return nil, errors.Wrap(err, "marketRepo.SearchByTitle.rows.Err")
	}

	return &model.MarketList{
		TotalCount: totalCount,
		TotalPages: utils.GetTotalPages(totalCount, query.GetSize()),
		Page:       query.GetPage(),
		Size:       query.GetSize(),
		HasMore:    utils.GetHasMore(query.GetPage(), totalCount, query.GetSize()),
		Markets:    marketList,
	}, nil
}
