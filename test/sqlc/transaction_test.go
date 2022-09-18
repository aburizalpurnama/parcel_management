package db

import (
	"context"
	"database/sql"
	"fmt"
	db "parcel-management/db/sqlc"
	"parcel-management/util"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func TestCreateTransaction(t *testing.T) {
	createRandomTransaction(t)
}

func TestGetTransactionById(t *testing.T) {
	trans1 := createRandomTransaction(t)

	trans2, err := testQueries.GetTransactionById(context.Background(), trans1.ID)
	require.NoError(t, err)

	require.NotEmpty(t, trans2)
	require.Equal(t, trans1, trans2)
}

func TestGetListAllTransactions(t *testing.T) {
	for i := 0; i < 3; i++ {
		time.Sleep(1 * time.Millisecond)
		createRandomTransaction(t)
	}

	listTrans, err := testQueries.GetListAllTransactions(context.Background())
	require.NoError(t, err)
	require.NotEmpty(t, listTrans)

	for _, tr := range listTrans {
		fmt.Printf("tr: %v\n", tr)
	}
}

func TestGetListDoneTransactions(t *testing.T) {
	var listTransId []uuid.UUID
	for i := 0; i < 3; i++ {
		time.Sleep(1 * time.Millisecond)
		tr := createRandomTransaction(t)
		require.Zero(t, tr.PickedAt)
		require.Zero(t, tr.UserOutID)
		require.Zero(t, tr.UpdatedAt)

		listTransId = append(listTransId, tr.ID)
	}

	for _, u := range listTransId {
		trans, err := testQueries.GetTransactionById(context.Background(), u)
		require.NoError(t, err)
		require.NotEmpty(t, trans)

		updateTransaction(t, trans)
	}
	listDoneTrans, err := testQueries.GetListDoneTransactions(context.Background())
	require.NoError(t, err)
	require.NotEmpty(t, listDoneTrans)

	for _, tr := range listDoneTrans {
		require.NotZero(t, tr.PickedAt)
		require.NotZero(t, tr.UserOutID)
		require.NotZero(t, tr.UpdatedAt)
	}
}

func TestGetPendingTransaction(t *testing.T) {
	for i := 0; i < 3; i++ {
		createRandomTransaction(t)
	}

	listPendingTransaction, err := testQueries.GetListPendingTransactions(context.Background())
	require.NoError(t, err)
	require.NotEmpty(t, listPendingTransaction)

	for _, tr := range listPendingTransaction {
		require.Zero(t, tr.PickedAt)
		require.Zero(t, tr.PickedBy)
		require.Zero(t, tr.UpdatedAt)
	}
}

func TestUpdateTransaction(t *testing.T) {
	user := createRandomUnit(t)
	trans1 := createRandomTransaction(t)

	require.Zero(t, trans1.PickedAt)
	require.Zero(t, trans1.UpdatedAt)
	require.Zero(t, trans1.PickedBy)
	require.Zero(t, trans1.UserOutID)

	qty_picked := util.RandomInt(1, int64(trans1.Qty))
	arg := db.UpdateTransactionParams{
		Qty:       trans1.Qty - int32(qty_picked),
		UserOutID: uuid.NullUUID{UUID: user.ID, Valid: true},
		PickedBy:  sql.NullString{String: util.RandomString(10), Valid: true},
		ID:        trans1.ID,
	}

	err := testQueries.UpdateTransaction(context.Background(), arg)
	require.NoError(t, err)

	trans2, err := testQueries.GetTransactionById(context.Background(), trans1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, trans2)

	require.NotZero(t, trans2.PickedAt)
	require.NotZero(t, trans2.UpdatedAt)
	require.NotZero(t, trans2.PickedBy)
	require.NotZero(t, trans2.UserOutID)
}

func createRandomTransaction(t *testing.T) db.Transaction {
	user := createRandomUser(t)
	unit := createRandomUnit(t)
	arg := db.CreateTransactionParams{
		UnitID:      unit.ID,
		DeliveredBy: util.RandomString(10),
		Type:        db.ProductTypesGoods,
		Qty:         int32(util.RandomInt(3, 10)),
		Owner:       util.RandomString(10),
		Phone:       util.RandomPhone(),
		UserInID:    user.ID,
	}

	trans, err := testQueries.CreateTransaction(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, trans)

	require.NotZero(t, trans.ID)
	require.NotZero(t, trans.CreatedAt)
	require.Equal(t, arg.UnitID, trans.UnitID)
	require.Equal(t, arg.DeliveredBy, trans.DeliveredBy)
	require.Equal(t, arg.Type, trans.Type)
	require.Equal(t, arg.Qty, trans.Qty)
	require.Equal(t, arg.Owner, trans.Owner)
	require.Equal(t, arg.Phone, trans.Phone)
	require.Equal(t, arg.UserInID, trans.UserInID)

	return trans
}

func updateTransaction(t *testing.T, trans1 db.Transaction) db.Transaction {
	user := createRandomUnit(t)

	qty_picked := util.RandomInt(1, int64(trans1.Qty))
	arg := db.UpdateTransactionParams{
		Qty:       trans1.Qty - int32(qty_picked),
		UserOutID: uuid.NullUUID{UUID: user.ID, Valid: true},
		PickedBy:  sql.NullString{String: util.RandomString(10), Valid: true},
		ID:        trans1.ID,
	}

	err := testQueries.UpdateTransaction(context.Background(), arg)
	require.NoError(t, err)

	trans2, err := testQueries.GetTransactionById(context.Background(), trans1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, trans2)

	require.NotZero(t, trans2.PickedAt)
	require.NotZero(t, trans2.UpdatedAt)
	require.NotZero(t, trans2.PickedBy)
	require.NotZero(t, trans2.UserOutID)

	return trans2
}
