package operation

import (
	"github.com/google/go-cmp/cmp"
	"github.com/kanade0404/tenhou-log/serdes/xml"
	"testing"
)

func TestReach(t *testing.T) {
	type args struct {
		who  string
		step string
	}
	tests := []struct {
		name    string
		args    args
		want    *Reach
		wantErr bool
	}{
		{
			name: "reach not complete",
			args: args{
				who:  "3",
				step: "1",
			},
			want: &Reach{
				playerIndex: xml.PlayerIndex(3),
				isComplete:  false,
			},
		},
		{
			name: "reach complete",
			args: args{
				who:  "1",
				step: "2",
			},
			want: &Reach{
				playerIndex: xml.PlayerIndex(1),
				isComplete:  true,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewReach(tt.args.who, tt.args.step)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewReach() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if diff := cmp.Diff(got.PlayerIndex(), tt.want.PlayerIndex()); diff != "" {
				t.Errorf("(-got+want)\n%v", diff)
			}
			if diff := cmp.Diff(got.IsComplete(), tt.want.IsComplete()); diff != "" {
				t.Errorf("(-got+want)\n%v", diff)
			}
		})
	}
}
