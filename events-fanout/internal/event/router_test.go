package event

import (
	"io/ioutil"
	"testing"
)

func Test_parseEvent(t *testing.T) {
	data, _ := ioutil.ReadFile("testdata/event_full.json")
	want := KubeEvent{}
	want.Event.Metadata.Namespace = "default"

	type args struct {
		value []byte
	}
	tests := []struct {
		name    string
		args    args
		want    func(*KubeEvent) (bool, string)
		wantErr bool
	}{
		{
			"read json event namespace",
			args{
				data,
			},
			func(e *KubeEvent) (bool, string) {
				if e.Metadata.Namespace != "default" {
					return false, "expected namespace default"
				}
				return true, "OK"
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := parseEvent(tt.args.value)
			if (err != nil) != tt.wantErr {
				t.Errorf("parseEvent() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			ok, reason := tt.want(got)
			if !ok {
				t.Errorf("parseEvent() got = %v but %s", got, reason)
			}
		})
	}
}
