package test_functions

import (
	"reflect"
	"testing"
)

func IsErrorEqual(t *testing.T, err error, wantErr error) {
	if err != wantErr {
		t.Errorf("expected error (%v), got error (%v)", wantErr, err)
	}
}

func IsObjectEqual(t *testing.T, obj interface{}, wantObj interface{}) {
	if !reflect.DeepEqual(wantObj, obj) {
		t.Errorf("expected (%v), got (%v)", wantObj, obj)
	}
}

func IsEqual(t *testing.T, obj interface{}, wantObj interface{}) {
	if wantObj != obj {
		t.Errorf("expected (%v), got (%v)", wantObj, obj)
	}
}

func IsTypeEqual(t *testing.T, obj interface{}, wantObj interface{}) {
	wantType := reflect.TypeOf(wantObj)
	gotType := reflect.TypeOf(obj)
	if wantType != gotType {
		t.Errorf("expected type (%v), got (%v)", wantType, gotType)
	}
}
