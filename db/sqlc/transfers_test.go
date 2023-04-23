package db

import (
	"context"
	"testing"
	"time"

	"github.com/elayadeismail/simple_bank/utils"
	"github.com/stretchr/testify/require"
)

func createRandomTransfer(t *testing.T, fa, ta Account) Transfer {
	args := CreateTransferParams{FromAccountID: fa.ID, ToAccountID: ta.ID, Amount: utils.RandomMoney()}

	tr, err := testQueries.CreateTransfer(context.Background(), args)

	require.NoError(t, err)

	require.NotEmpty(t, tr)
	require.Equal(t, tr.FromAccountID, args.FromAccountID)
	require.Equal(t, tr.ToAccountID, args.ToAccountID)
	require.Equal(t, tr.Amount, args.Amount)
	require.NotZero(t, tr.ID)
	require.NotZero(t, tr.CreatedAt)
	return tr
}

func TestCreateTransfer(t *testing.T) {
	// From_Account
	fa := createRandomAccount(t)
	// To Account
	ta := createRandomAccount(t)

	createRandomTransfer(t, fa, ta)
}

func TestGetTransfer(t *testing.T) {
	// From_Account
	fa := createRandomAccount(t)
	// To Account
	ta := createRandomAccount(t)

	tr1 := createRandomTransfer(t, fa, ta)

	tr2, err := testQueries.GetTransfer(context.Background(), tr1.ID)

	require.NoError(t, err)

	require.Equal(t, tr1.ID, tr2.ID)
	require.Equal(t, tr1.FromAccountID, tr2.FromAccountID)
	require.Equal(t, tr1.ToAccountID, tr2.ToAccountID)
	require.WithinDuration(t, tr1.CreatedAt, tr2.CreatedAt, time.Second)

}

func TestListTransfer(t *testing.T) {
	fa := createRandomAccount(t)
	// To Account
	ta := createRandomAccount(t)

	for i := 0; i < 10; i++ {
		createRandomTransfer(t, fa, ta)
	}

	args := ListTransfersParams{
		FromAccountID: fa.ID,
		ToAccountID:   ta.ID,
		Limit:         5,
		Offset:        5,
	}

	trs, err := testQueries.ListTransfers(context.Background(), args)

	require.NoError(t, err)

	require.Equal(t, len(trs), 5)

	for _, tr := range trs {
		require.Equal(t, tr.FromAccountID, fa.ID)
		require.Equal(t, tr.ToAccountID, ta.ID)

		require.NotZero(t, tr.ID)
	}
}
