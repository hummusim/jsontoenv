package jsontoenv

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestJSONToEnv(t *testing.T) {
	t.Run("FromBytes", func(t *testing.T) {
		defer os.Clearenv()

		env := New(Opts{})
		data := []byte(`{"key1":"value1","key2":2,"key3":3.14,"key4":true,"key5":["a","b","c"],"key6":{"subkey1":"subvalue1","subkey2":"subvalue2"}}`)
		require.NoError(t, env.FromBytes(data))

		for key, value := range map[string]string{
			"key1":    "value1",
			"key2":    "2",
			"key3":    "3.14",
			"key4":    "true",
			"key5":    "a,b,c",
			"subkey1": "subvalue1",
			"subkey2": "subvalue2",
		} {
			envValue, ok := os.LookupEnv(key)
			require.True(t, ok)
			require.NotEmpty(t, envValue)
			require.Equal(t, value, envValue)
		}
	})

	t.Run("FromBytesWithUpperCase", func(t *testing.T) {
		defer os.Clearenv()

		env := New(Opts{
			UseUpperCase: true,
		})
		data := []byte(`{"key1":"value1","key2":2,"key3":3.14,"key4":true,"key5":["a","b","c"],"key6":{"subkey1":"subvalue1","subkey2":"subvalue2"}}`)
		require.NoError(t, env.FromBytes(data))

		for key, value := range map[string]string{
			"KEY1":    "value1",
			"KEY2":    "2",
			"KEY3":    "3.14",
			"KEY4":    "true",
			"KEY5":    "a,b,c",
			"SUBKEY1": "subvalue1",
			"SUBKEY2": "subvalue2",
		} {
			envValue, ok := os.LookupEnv(key)
			require.True(t, ok)
			require.NotEmpty(t, envValue)
			require.Equal(t, value, envValue)
		}
	})

	t.Run("FromBytesWithOmitKeys", func(t *testing.T) {
		defer os.Clearenv()

		env := New(Opts{})
		env.OmitKeys("key1")

		data := []byte(`{"key1":"value1","key2":2,"key3":3.14,"key4":true,"key5":["a","b","c"],"key6":{"subkey1":"subvalue1","subkey2":"subvalue2"}}`)

		require.NoError(t, env.FromBytes(data))

		for key, value := range map[string]string{
			"key1":    "",
			"key2":    "2",
			"key3":    "3.14",
			"key4":    "true",
			"key5":    "a,b,c",
			"subkey1": "subvalue1",
			"subkey2": "subvalue2",
		} {
			if value == "" {
				_, ok := os.LookupEnv(key)
				require.False(t, ok)
			} else {
				envValue, ok := os.LookupEnv(key)
				require.True(t, ok)
				require.NotEmpty(t, envValue)
				require.Equal(t, value, envValue)
			}
		}
	})
}
