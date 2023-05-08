package cli

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestPrintMarkDownList(t *testing.T) {
	dataTest1 := `{
        "mx": {
            "edu.mx": {
                "itesm.edu.mx": {},
                "tecmilenio.edu.mx": {}
            },
            "itesm.mx": {
                "admision.itesm.mx": {},
                "admisionprepatec.itesm.mx": {},
                "ags.itesm.mx": {},
                "apps.itesm.mx": {},
                "btec.itesm.mx": {},
                "cdj.itesm.mx": {},
                "cegs.itesm.mx": {},
                "chi.itesm.mx": {},
                "dm.itesm.mx": {},
                "exatec1.itesm.mx": {},
                "lag.itesm.mx": {},
                "mty.itesm.mx": {
                    "web8.mty.itesm.mx": {}
                },
                "net.itesm.mx": {},
                "queretaro.itesm.mx": {
                    "comunicacionypublicidad.queretaro.itesm.mx": {},
                    "identidad.queretaro.itesm.mx": {}
                },
                "ruv.itesm.mx": {},
                "rzn.itesm.mx": {},
                "sal.itesm.mx": {},
                "sistema.itesm.mx": {},
                "sitios.itesm.mx": {},
                "slp.itesm.mx": {},
                "sorteotec.itesm.mx": {},
                "tecreview.itesm.mx": {},
                "zac.itesm.mx": {}
            },
            "tecreview.mx": {}
        },
        "soy": {
            "prepatec.soy": {}
        }
    }`

	var treeTest1 map[string]interface{}
	err := json.Unmarshal([]byte(dataTest1), &treeTest1)
	if err != nil {
		panic(err)
	}
	type args struct {
		data     map[string]interface{}
		markdown bool
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Test 1: Convert tree to list",
			args: args{
				data:     treeTest1,
				markdown: false,
			},
			want: ``,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := TreeToList(tt.args.data, tt.args.markdown); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PrintMarkDownList() = %v, want %v", got, tt.want)
			}
		})
	}
}
