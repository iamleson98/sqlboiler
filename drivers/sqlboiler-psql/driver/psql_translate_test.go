package driver

import (
	"testing"

	"github.com/aarondl/sqlboiler/v4/drivers"
)

func TestTranslateColumnType(t *testing.T) {
	p := PostgresDriver{}

	tests := []struct {
		name     string
		input    drivers.Column
		wantType string
	}{
		// integer / serial (int4) -> int32
		{name: "integer non-null", input: drivers.Column{DBType: "integer"}, wantType: "int32"},
		{name: "serial non-null", input: drivers.Column{DBType: "serial"}, wantType: "int32"},
		{name: "integer nullable", input: drivers.Column{DBType: "integer", Nullable: true}, wantType: "null.Int32"},
		{name: "serial nullable", input: drivers.Column{DBType: "serial", Nullable: true}, wantType: "null.Int32"},

		// bigint / bigserial (int8) -> int64
		{name: "bigint non-null", input: drivers.Column{DBType: "bigint"}, wantType: "int64"},
		{name: "bigserial non-null", input: drivers.Column{DBType: "bigserial"}, wantType: "int64"},
		{name: "bigint nullable", input: drivers.Column{DBType: "bigint", Nullable: true}, wantType: "null.Int64"},

		// smallint / smallserial (int2) -> int16
		{name: "smallint non-null", input: drivers.Column{DBType: "smallint"}, wantType: "int16"},
		{name: "smallserial non-null", input: drivers.Column{DBType: "smallserial"}, wantType: "int16"},
		{name: "smallint nullable", input: drivers.Column{DBType: "smallint", Nullable: true}, wantType: "null.Int16"},

		// oid -> uint32
		{name: "oid non-null", input: drivers.Column{DBType: "oid"}, wantType: "uint32"},
		{name: "oid nullable", input: drivers.Column{DBType: "oid", Nullable: true}, wantType: "null.Uint32"},

		// decimal / numeric -> types.Decimal
		{name: "decimal non-null", input: drivers.Column{DBType: "decimal"}, wantType: "types.Decimal"},
		{name: "numeric non-null", input: drivers.Column{DBType: "numeric"}, wantType: "types.Decimal"},
		{name: "decimal nullable", input: drivers.Column{DBType: "decimal", Nullable: true}, wantType: "types.NullDecimal"},

		// double precision -> float64
		{name: "double precision non-null", input: drivers.Column{DBType: "double precision"}, wantType: "float64"},
		{name: "double precision nullable", input: drivers.Column{DBType: "double precision", Nullable: true}, wantType: "null.Float64"},

		// real -> float32
		{name: "real non-null", input: drivers.Column{DBType: "real"}, wantType: "float32"},
		{name: "real nullable", input: drivers.Column{DBType: "real", Nullable: true}, wantType: "null.Float32"},

		// text -> string
		{name: "text non-null", input: drivers.Column{DBType: "text"}, wantType: "string"},
		{name: "text nullable", input: drivers.Column{DBType: "text", Nullable: true}, wantType: "null.String"},

		// uuid -> string (also covers character varying, cidr, inet, etc.)
		{name: "uuid non-null", input: drivers.Column{DBType: "uuid"}, wantType: "string"},
		{name: "uuid nullable", input: drivers.Column{DBType: "uuid", Nullable: true}, wantType: "null.String"},

		// "char" -> types.Byte
		{name: `"char" non-null`, input: drivers.Column{DBType: `"char"`}, wantType: "types.Byte"},
		{name: `"char" nullable`, input: drivers.Column{DBType: `"char"`, Nullable: true}, wantType: "null.Byte"},

		// bytea -> []byte
		{name: "bytea non-null", input: drivers.Column{DBType: "bytea"}, wantType: "[]byte"},
		{name: "bytea nullable", input: drivers.Column{DBType: "bytea", Nullable: true}, wantType: "null.Bytes"},

		// json / jsonb -> types.JSON
		{name: "json non-null", input: drivers.Column{DBType: "json"}, wantType: "types.JSON"},
		{name: "jsonb non-null", input: drivers.Column{DBType: "jsonb"}, wantType: "types.JSON"},
		{name: "json nullable", input: drivers.Column{DBType: "json", Nullable: true}, wantType: "null.JSON"},

		// boolean -> bool
		{name: "boolean non-null", input: drivers.Column{DBType: "boolean"}, wantType: "bool"},
		{name: "boolean nullable", input: drivers.Column{DBType: "boolean", Nullable: true}, wantType: "null.Bool"},

		// timestamp -> time.Time
		{name: "timestamp without time zone non-null", input: drivers.Column{DBType: "timestamp without time zone"}, wantType: "time.Time"},
		{name: "timestamp without time zone nullable", input: drivers.Column{DBType: "timestamp without time zone", Nullable: true}, wantType: "null.Time"},
		{name: "timestamp with time zone non-null", input: drivers.Column{DBType: "timestamp with time zone"}, wantType: "time.Time"},

		// USER-DEFINED -> hstore / citext
		{name: "hstore non-null", input: drivers.Column{DBType: "USER-DEFINED", UDTName: "hstore"}, wantType: "types.HStore"},
		{name: "citext non-null", input: drivers.Column{DBType: "USER-DEFINED", UDTName: "citext"}, wantType: "string"},
		{name: "citext nullable", input: drivers.Column{DBType: "USER-DEFINED", UDTName: "citext", Nullable: true}, wantType: "null.String"},

		// unknown type -> string / null.String
		{name: "unknown non-null", input: drivers.Column{DBType: "unknown_type"}, wantType: "string"},
		{name: "unknown nullable", input: drivers.Column{DBType: "unknown_type", Nullable: true}, wantType: "null.String"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := p.TranslateColumnType(tt.input)
			if got.Type != tt.wantType {
				t.Errorf("TranslateColumnType(%q, nullable=%v) = %q, want %q",
					tt.input.DBType, tt.input.Nullable, got.Type, tt.wantType)
			}
		})
	}
}
