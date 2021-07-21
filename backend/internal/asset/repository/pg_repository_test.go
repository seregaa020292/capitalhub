package repository

import (
	"context"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/require"

	"github.com/seregaa020292/capitalhub/internal/asset/model"
)

func TestAssetRepo_Create(t *testing.T) {
	t.Parallel()

	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	require.NoError(t, err)
	defer db.Close()

	sqlxDB := sqlx.NewDb(db, "sqlmock")
	defer sqlxDB.Close()

	assetRepo := NewAssetRepository(sqlxDB)

	t.Run("Create", func(t *testing.T) {
		userUID := uuid.New()
		assetID := uuid.New()
		amount := 20e3

		rows := sqlmock.NewRows([]string{"user_id", "asset_id", "amount"}).AddRow(userUID, assetID, amount)

		asset := &model.Asset{
			UserID:  userUID,
			AssetID: assetID,
			Amount:  amount,
		}

		mock.ExpectQuery(createAsset).WithArgs(asset.UserID, &asset.AssetID, asset.Amount).WillReturnRows(rows)

		createdAsset, err := assetRepo.Create(context.Background(), asset)

		require.NoError(t, err)
		require.NotNil(t, createdAsset)
		require.Equal(t, createdAsset, asset)
	})

	t.Run("Create ERR", func(t *testing.T) {
		assetID := uuid.New()
		amount := 20e3
		createErr := errors.New("Create asset error")

		asset := &model.Asset{
			AssetID: assetID,
			Amount:  amount,
		}

		mock.ExpectQuery(createAsset).WithArgs(asset.UserID, &asset.AssetID, asset.Amount).WillReturnError(createErr)

		createdAsset, err := assetRepo.Create(context.Background(), asset)

		require.Nil(t, createdAsset)
		require.NotNil(t, err)
	})
}

func TestAssetRepo_Update(t *testing.T) {
	t.Parallel()

	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	require.NoError(t, err)
	defer db.Close()

	sqlxDB := sqlx.NewDb(db, "sqlmock")
	defer sqlxDB.Close()

	assetRepo := NewAssetRepository(sqlxDB)

	t.Run("Update", func(t *testing.T) {
		userUID := uuid.New()
		assetUID := uuid.New()
		amount := 20e3

		rows := sqlmock.NewRows([]string{"user_id", "asset_id", "amount"}).AddRow(userUID, assetUID, amount)

		asset := &model.Asset{
			AssetID: assetUID,
			Amount:  amount,
		}

		mock.ExpectQuery(updateAsset).WithArgs(asset.Amount, asset.AssetID).WillReturnRows(rows)

		createdAsset, err := assetRepo.Update(context.Background(), asset)

		require.NoError(t, err)
		require.NotNil(t, createdAsset)
		require.Equal(t, createdAsset.Amount, asset.Amount)
	})

	t.Run("Update ERR", func(t *testing.T) {
		assetUID := uuid.New()
		amount := 20e3
		updateErr := errors.New("Create asset error")

		asset := &model.Asset{
			AssetID: assetUID,
			Amount:  amount,
		}

		mock.ExpectQuery(updateAsset).WithArgs(asset.Amount, asset.AssetID).WillReturnError(updateErr)

		createdAsset, err := assetRepo.Update(context.Background(), asset)

		require.NotNil(t, err)
		require.Nil(t, createdAsset)
	})
}

func TestAssetRepo_Delete(t *testing.T) {
	t.Parallel()

	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	require.NoError(t, err)
	defer db.Close()

	sqlxDB := sqlx.NewDb(db, "sqlmock")
	defer sqlxDB.Close()

	assetRepo := NewAssetRepository(sqlxDB)

	t.Run("Delete", func(t *testing.T) {
		assetUID := uuid.New()
		mock.ExpectExec(deleteAsset).WithArgs(assetUID).WillReturnResult(sqlmock.NewResult(1, 1))
		err := assetRepo.Delete(context.Background(), assetUID)

		require.NoError(t, err)
	})

	t.Run("Delete Err", func(t *testing.T) {
		assetUID := uuid.New()

		mock.ExpectExec(deleteAsset).WithArgs(assetUID).WillReturnResult(sqlmock.NewResult(1, 0))

		err := assetRepo.Delete(context.Background(), assetUID)
		require.NotNil(t, err)
	})
}
