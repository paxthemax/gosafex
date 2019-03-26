package store

import (
	"testing"
)

var testVector = []KV{
	{
		Key:   []byte{0},
		Value: []byte{0},
	},
	{
		Key:   []byte{1},
		Value: []byte{1},
	},
	{
		Key:   []byte{2},
		Value: []byte{2},
	},
	{
		Key:   []byte{3},
		Value: []byte{3},
	},
	{
		Key:   []byte{4},
		Value: []byte{4},
	},
	{
		Key:   []byte{5},
		Value: []byte{5},
	},
}

func TestKVIterator_Valid(t *testing.T) {
	type fields struct {
		data []KV
		idx  int
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			name: "true, cursor in range",
			fields: fields{
				data: testVector,
				idx:  0,
			},
			want: true,
		},
		{
			name: "false, cursor out of range",
			fields: fields{
				data: testVector,
				idx:  999,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			it := &KVIterator{
				data: tt.fields.data,
				idx:  tt.fields.idx,
			}
			if got := it.Valid(); got != tt.want {
				t.Errorf("KVIterator.Valid() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestKVIterator_Next(t *testing.T) {
	type fields struct {
		data []KV
		idx  int
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "passes, in range",
			fields: fields{
				data: testVector,
				idx:  0,
			},
		},
		{
			name: "fails, out of range",
			fields: fields{
				data: testVector,
				idx:  5,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			it := &KVIterator{
				data: tt.fields.data,
				idx:  tt.fields.idx,
			}
			if err := it.Next(); (err != nil) != tt.wantErr {
				t.Errorf("KVIterator.Next() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
