package test

import (
	"context"
	"database/sql"
	"fmt"
	db "parcel-management/db/sqlc"
	"parcel-management/util"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestCreateUnit(t *testing.T) {
	createRandomUnit(t)
}

func TestGetById(t *testing.T) {
	u1 := createRandomUnit(t)

	u2, err := testQueries.GetUnitById(context.Background(), u1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, u2)

	require.Equal(t, u1.ID, u2.ID)
	require.WithinDuration(t, u1.CreatedAt.Time, u2.CreatedAt.Time, time.Second)
	require.Equal(t, u1.Email, u2.Email)
	require.Equal(t, u1.ItemPendingQty, u2.ItemPendingQty)
	require.Equal(t, u1.No, u2.No)
	require.Equal(t, u1.Phone, u2.Phone)
}

func TestGetListUnit(t *testing.T) {
	for i := 0; i < 5; i++ {
		createRandomUnit(t)
	}

	u, err := testQueries.GetListUnits(context.Background())
	require.NoError(t, err)
	require.NotEmpty(t, u)

	for _, u2 := range u {
		fmt.Printf("u2: %v\n", u2)
	}

	fmt.Printf("len(u): %v\n", len(u))
}

func TestUpdatePhoneUnit(t *testing.T) {
	unit1 := createRandomUnit(t)
	newPhone := util.RandomPhone()
	arg := db.UpdatePhoneUnitParams{Phone: newPhone, ID: unit1.ID}

	err := testQueries.UpdatePhoneUnit(context.Background(), arg)
	require.NoError(t, err)

	unit2, err := testQueries.GetUnitById(context.Background(), unit1.ID)
	require.NoError(t, err)

	require.NotEqual(t, unit1.Phone, unit2.Phone)
	require.Equal(t, newPhone, unit2.Phone)
}

func TestDeleteUnit(t *testing.T) {
	unit := createRandomUnit(t)

	err := testQueries.DeleteUnitById(context.Background(), unit.ID)
	require.NoError(t, err)

	unit2, err := testQueries.GetUnitById(context.Background(), unit.ID)
	require.Error(t, err)
	require.EqualError(t, sql.ErrNoRows, err.Error())
	require.Empty(t, unit2)
}

func createRandomUnit(t *testing.T) db.Unit {
	arg := db.CreateUnitParams{
		No:    util.RandomUnitNo(),
		Email: util.RandomEmail(),
		Phone: util.RandomPhone(),
	}

	unit, err := testQueries.CreateUnit(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, unit)

	require.Equal(t, arg.Email, unit.Email)
	require.Equal(t, arg.No, unit.No)
	require.Equal(t, arg.Phone, unit.Phone)

	require.NotZero(t, unit.ID)
	require.NotZero(t, unit.CreatedAt)

	return unit
}
