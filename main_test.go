package main

import (
	"flag"
	"net/url"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/urfave/cli/v2"
)

func testContext(t *testing.T, u *url.URL) *cli.Context {
	err := os.Setenv("DATABASE_URL", u.String())
	require.NoError(t, err)

	app := NewApp()
	flagset := flag.NewFlagSet(app.Name, flag.ContinueOnError)
	for _, f := range app.Flags {
		err := f.Apply(flagset)
		require.NoError(t, err)
	}

	return cli.NewContext(app, flagset, nil)
}

func TestGetDatabaseUrl(t *testing.T) {
	envURL, err := url.Parse("foo://example.org/db")
	require.NoError(t, err)
	ctx := testContext(t, envURL)

	u, err := getDatabaseURL(ctx)
	require.NoError(t, err)

	require.Equal(t, "foo", u.Scheme)
	require.Equal(t, "example.org", u.Host)
	require.Equal(t, "/db", u.Path)
}
