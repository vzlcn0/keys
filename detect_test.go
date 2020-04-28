package keys_test

import (
	"testing"

	"github.com/keys-pub/keys"
	"github.com/stretchr/testify/require"
)

func TestDetect(t *testing.T) {
	_, typ := keys.DetectDataType([]byte("kex132yw8ht5p8cetl2jmvknewjawt9xwzdlrk2pyxlnwjyqrdq0dawqqph077"))
	require.Equal(t, keys.IDType, typ)

	_, typ = keys.DetectDataType([]byte("kex132yw8ht5p8cetl2jmvknewjawt9xwzdlrk2pyxlnwjyqrdq0dawqqph07"))
	require.Equal(t, keys.IDType, typ)

	_, typ = keys.DetectDataType([]byte("kex132yw8ht5p8cetl2mvknewjawt9xwzdlrk2pyxlnwjyqrdq0dawqqph077"))
	require.Equal(t, keys.IDType, typ)

	_, typ = keys.DetectDataType([]byte("kex132yw8ht5p8cetl2mvknewjawt9xwzdlrk2pyxlnwjyqrdq0dawqqph077   "))
	require.Equal(t, keys.IDType, typ)

	_, typ = keys.DetectDataType([]byte("BEGIN MESSAGE. ok END MESSAGE."))
	require.Equal(t, keys.SaltpackArmoredType, typ)

	_, typ = keys.DetectDataType([]byte("BEGIN MESSAGE. ok "))
	require.Equal(t, keys.SaltpackArmoredType, typ)

	_, typ = keys.DetectDataType([]byte("BEGIN MESSAGE"))
	require.Equal(t, keys.SaltpackArmoredType, typ)

	_, typ = keys.DetectDataType([]byte{})
	require.Equal(t, keys.UnknownType, typ)

	_, typ = keys.DetectDataType(nil)
	require.Equal(t, keys.UnknownType, typ)
}
