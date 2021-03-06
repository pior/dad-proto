package taskapi

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestTaskConfigIsHash(t *testing.T) {
	config := &TaskConfig{
		payload: map[interface{}]interface{}{},
	}

	require.True(t, config.IsHash())

	config = &TaskConfig{
		payload: "",
	}

	require.False(t, config.IsHash())
}

func TestTaskConfigProperty(t *testing.T) {
	config := &TaskConfig{
		name:    "test",
		payload: map[interface{}]interface{}{"key": []interface{}{"one", "two"}},
	}

	val, err := config.GetListOfStringsPropertyDefault("key", []string{})
	require.NoError(t, err)
	require.Equal(t, []string{"one", "two"}, val)

	val, err = config.GetListOfStringsPropertyDefault("nope", []string{"three"})
	require.NoError(t, err)
	require.Equal(t, []string{"three"}, val)
}

func TestTaskConfigPropertyInvalid(t *testing.T) {
	config := &TaskConfig{name: "test", payload: "a string"}

	_, err := config.GetListOfStringsPropertyDefault("key", []string{})
	require.Error(t, err)
	require.Equal(t, "not a hash: string (a string)", err.Error())
}

func TestTaskConfigStringOrStringProperty(t *testing.T) {
	config := &TaskConfig{name: "test", payload: "value"}

	val, err := config.GetStringPropertyAllowSingle("key")
	require.NoError(t, err)
	require.Equal(t, val, "value")
}

func TestTaskConfigGetBooleanPropertyDefault(t *testing.T) {
	payload := map[interface{}]interface{}{"flag": true}
	config := &TaskConfig{name: "test", payload: payload}

	val, err := config.GetBooleanPropertyDefault("flag", false)
	require.NoError(t, err)
	require.Equal(t, val, true)

	val, err = config.GetBooleanPropertyDefault("nope", false)
	require.NoError(t, err)
	require.Equal(t, val, false)
}

func TestTaskConfigStringProperty(t *testing.T) {
	value := map[interface{}]interface{}{"key": "val"}
	config := &TaskConfig{name: "test", payload: value}

	val, err := config.GetStringProperty("key")
	require.NoError(t, err)
	require.Equal(t, val, "val")

	val, err = config.GetStringProperty("nope")
	require.Error(t, err)
	require.Equal(t, "property \"nope\" not found", err.Error())
}

func TestTaskConfigStringPropertyInvalid(t *testing.T) {
	config := &TaskConfig{name: "test", payload: 42}
	_, err := config.GetStringPropertyAllowSingle("key")
	require.Error(t, err)
	require.Contains(t, err.Error(), "not a hash")
	require.Contains(t, err.Error(), "int (42)")

	config = &TaskConfig{name: "test", payload: false}
	_, err = config.GetStringPropertyAllowSingle("key")
	require.Error(t, err)
	require.Contains(t, err.Error(), "not a hash")
	require.Contains(t, err.Error(), "bool (false)")

	config = &TaskConfig{name: "test", payload: "thisisastring"}
	_, err = config.GetStringProperty("key1")
	require.Error(t, err)
	require.Contains(t, err.Error(), "not a hash")
	require.Contains(t, err.Error(), "thisisastring")

	payload := map[interface{}]interface{}{"version": 3.6}
	config = &TaskConfig{name: "test", payload: payload}
	_, err = config.GetStringProperty("version")
	require.Error(t, err)
	require.Contains(t, err.Error(), "not a string")
	require.Contains(t, err.Error(), "float64 (3.6)")
}

func TestTaskConfigListOfStrings(t *testing.T) {
	value := []interface{}{"one", "two"}
	config := &TaskConfig{name: "test", payload: value}

	result, err := config.GetListOfStrings()
	require.NoError(t, err)
	require.Equal(t, []string{"one", "two"}, result)
}

func TestTaskConfigListOfStringsEmpty(t *testing.T) {
	config := &TaskConfig{name: "test", payload: []interface{}{}}

	result, err := config.GetListOfStrings()
	require.NoError(t, err)
	require.Equal(t, []string{}, result)
}

func TestTaskConfigListOfStringsInvalidElement(t *testing.T) {
	config := &TaskConfig{name: "test", payload: []interface{}{"one", 2}}

	_, err := config.GetListOfStrings()
	require.Error(t, err)
	require.Equal(t, "not a list of strings: invalid element: type int (2)", err.Error())
}

func TestTaskConfigListOfStringsInvalidType(t *testing.T) {
	config := &TaskConfig{name: "test", payload: "plop"}

	_, err := config.GetListOfStrings()
	require.Error(t, err)
	require.Equal(t, "not a list of strings: type string (plop)", err.Error())
}
