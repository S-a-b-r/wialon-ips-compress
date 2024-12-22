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

	data1, _ := hex.DecodeString("FF1B00780153F65136D233B0CECC4DCDB4F673B476B4343602002FF404E6")
	data2, _ := hex.DecodeString("ff1b00789c52f65136d233b03606036b3f476b474b632340000000ffff2f310474")
	data3, _ := hex.DecodeString("ff1b00789c52f65136d233b036b2b4303537313331373533b1f673b476b4343602040000ffff4e1d05c5")

	tests := []TestStructLogin{
		{
			name:    "Test valid all data 1",
			buf:     *bytes.NewBuffer(data1),
			want:    "imei",
			wantErr: false,
		},
		{
			name:    "Test valid all data 2",
			buf:     *bytes.NewBuffer(data2),
			want:    "333333",
			wantErr: false,
		},
		{
			name:    "Test valid all data 3",
			buf:     *bytes.NewBuffer(data3),
			want:    "298574647564",
			wantErr: false,
		},
	}

	for _, tt := range tests {
		p := parser{}
		t.Run(tt.name, func(t *testing.T) {
			res, err := p.HandleCompressionLogin(tt.buf)
			if tt.wantErr == true && err != nil {
				t.Fatal("ParseHandleCompressionLogin() should return error")
			}

			if (err != nil) != tt.wantErr {
				t.Fatalf("ParseHandleCompressionLogin() error = %v, wantErr %v", err, tt.wantErr)
			}
			if res != tt.want {
				t.Fatalf("ParseHandleCompressionLogin() res = %s, want %s", res, tt.want)
			}
			// if !reflect.DeepEqual(res, tt.want) {
			// 	t.Errorf("StringToStruct() = %+v, %T, want %+v, %T", tt.itf, tt.itf, tt.want, tt.want)
			// }
		})
	}
}
