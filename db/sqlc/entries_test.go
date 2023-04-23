package db

import (
	"context"
	"testing"
	"time"

	"github.com/elayadeismail/simple_bank/utils"
	"github.com/stretchr/testify/require"
)

// createRandomEntry Takes Account and Create a random entry for that account
func createRandomEntry(t *testing.T, account Account) Entry {
	args := CreateEntryParams{AccountID: account.ID, Amount: utils.RandomMoney()}

	entry, err := testQueries.CreateEntry(context.Background(), args)

	require.NoError(t, err)
	require.NotEmpty(t, entry)
	require.Equal(t, entry.AccountID, args.AccountID)
	require.Equal(t, entry.Amount, args.Amount)
	require.NotZero(t, entry.ID)

	return entry
}

func CreateEntryTest(t *testing.T) {
	account := createRandomAccount(t)
	createRandomEntry(t, account)
}

func GetEntryTest(t *testing.T) {
	account := createRandomAccount(t)
	entry := createRandomEntry(t, account)

	e, err := testQueries.GetEntry(context.Background(), entry.ID)

	require.NoError(t, err)

	require.Equal(t, e.AccountID, entry.AccountID)
	require.Equal(t, e.Amount, entry.AccountID)
	require.WithinDuration(t, e.CreatedAt, account.CreatedAt, time.Second)
}

func ListEntriesTest(t *testing.T) {
	account := createRandomAccount(t)
	for i := 0; i < 10; i++ {
		createRandomEntry(t, account)
	}
	args := ListEntriesParams{AccountID: account.ID, Limit: 5, Offset: 5}
	es, err := testQueries.ListEntries(context.Background(), args)

	require.NoError(t, err)

	require.Equal(t, len(es), 5)

	for _, e := range es {
		require.Equal(t, e.AccountID, account.ID)
	}
}
