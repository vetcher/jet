package sqlbuilder

import (
	"bytes"
	"gotest.tools/assert"
	"testing"
)

func TestNewBoolColumn(t *testing.T) {
	boolColumn := NewBoolColumn("col", Nullable)

	out := bytes.Buffer{}
	err := boolColumn.Serialize(&out)

	assert.NilError(t, err)
	assert.Equal(t, out.String(), "col")

	out.Reset()
	err = boolColumn.Serialize(&out, FOR_PROJECTION)
	assert.NilError(t, err)
	assert.Equal(t, out.String(), "col")

	out.Reset()
	err = boolColumn.setTableName("table1")
	assert.NilError(t, err)
	err = boolColumn.Serialize(&out, FOR_PROJECTION)
	assert.NilError(t, err)
	assert.Equal(t, out.String(), `table1.col AS "table1.col"`)

	out.Reset()
	err = boolColumn.setTableName("table1")
	assert.NilError(t, err)
	aliasedBoolColumn := boolColumn.As("alias1")
	err = aliasedBoolColumn.SerializeForProjection(&out)
	assert.NilError(t, err)
	assert.Equal(t, out.String(), `table1.col AS "alias1"`)
}

func TestNewIntColumn(t *testing.T) {
	integerColumn := NewIntegerColumn("col", Nullable)

	out := bytes.Buffer{}
	err := integerColumn.Serialize(&out)

	assert.NilError(t, err)
	assert.Equal(t, out.String(), "col")

	out.Reset()
	err = integerColumn.Serialize(&out, FOR_PROJECTION)
	assert.NilError(t, err)
	assert.Equal(t, out.String(), "col")

	out.Reset()
	err = integerColumn.setTableName("table1")
	assert.NilError(t, err)
	err = integerColumn.Serialize(&out, FOR_PROJECTION)
	assert.NilError(t, err)
	assert.Equal(t, out.String(), `table1.col AS "table1.col"`)

	out.Reset()
	err = integerColumn.setTableName("table1")
	assert.NilError(t, err)
	aliasedBoolColumn := integerColumn.As("alias1")
	err = aliasedBoolColumn.SerializeForProjection(&out)
	assert.NilError(t, err)
	assert.Equal(t, out.String(), `table1.col AS "alias1"`)
}

func TestNewNumericColumnColumn(t *testing.T) {
	numericColumn := NewNumericColumn("col", Nullable)

	out := bytes.Buffer{}
	err := numericColumn.Serialize(&out)

	assert.NilError(t, err)
	assert.Equal(t, out.String(), "col")

	out.Reset()
	err = numericColumn.Serialize(&out)
	assert.NilError(t, err)
	assert.Equal(t, out.String(), "col")

	out.Reset()
	err = numericColumn.setTableName("table1")
	assert.NilError(t, err)
	err = numericColumn.Serialize(&out, FOR_PROJECTION)
	assert.NilError(t, err)
	assert.Equal(t, out.String(), `table1.col AS "table1.col"`)

	out.Reset()
	err = numericColumn.setTableName("table1")
	assert.NilError(t, err)
	aliasedBoolColumn := numericColumn.As("alias1")
	err = aliasedBoolColumn.SerializeForProjection(&out)
	assert.NilError(t, err)
	assert.Equal(t, out.String(), `table1.col AS "alias1"`)
}
