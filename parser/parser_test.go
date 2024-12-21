package parser

import (
	"bytes"
	"encoding/hex"
	"testing"
)

type TestStructLogin struct {
	name    string // Название теста
	buf     bytes.Buffer
	want    string
	wantErr bool // Ожидание ошибки
}

func TestHandleCompressionLogin(t *testing.T) {

	data1, err := hex.DecodeString("FF1B00780153F65136D233B0CECC4DCDB4F673B476B4343602002FF404E6")
	if err != nil {
		t.Fatal(err)
	}

	tests := []TestStructLogin{
		{
			name:    "Test valid all data",
			buf:     *bytes.NewBuffer(data1),
			want:    "imei",
			wantErr: false,
		},
	}

	for _, tt := range tests {
		p := parser{}
		t.Run(tt.name, func(t *testing.T) {
			res, err := p.HandleCompressionLogin(tt.buf)
			if tt.wantErr == true && err != nil {
				return
			}

			if (err != nil) != tt.wantErr {
				t.Errorf("ParseHandleCompressionLogin() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if res != tt.want {
				t.Errorf("ParseHandleCompressionLogin() res = %s, want %s", res, tt.want)
				return
			}
			// if !reflect.DeepEqual(res, tt.want) {
			// 	t.Errorf("StringToStruct() = %+v, %T, want %+v, %T", tt.itf, tt.itf, tt.want, tt.want)
			// }
		})
	}
}
