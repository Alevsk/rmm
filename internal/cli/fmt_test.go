package cli

import (
	"encoding/json"
	"reflect"
	"strconv"
	"testing"

	"github.com/Alevsk/rmm/internal/mindmap"
)

func convertEscapeSequences(text string) string {
	converted, err := strconv.Unquote(`"` + text + `"`)
	if err != nil {
		panic(err)
	}

	return converted
}

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

	var treeTest1 mindmap.Node
	err := json.Unmarshal([]byte(dataTest1), &treeTest1)
	if err != nil {
		panic(err)
	}
	type args struct {
		data     mindmap.Node
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
			want: `mx\n\tedu.mx\n\t\titesm.edu.mx\n\t\ttecmilenio.edu.mx\n\titesm.mx\n\t\tadmision.itesm.mx\n\t\tadmisionprepatec.itesm.mx\n\t\tags.itesm.mx\n\t\tapps.itesm.mx\n\t\tbtec.itesm.mx\n\t\tcdj.itesm.mx\n\t\tcegs.itesm.mx\n\t\tchi.itesm.mx\n\t\tdm.itesm.mx\n\t\texatec1.itesm.mx\n\t\tlag.itesm.mx\n\t\tmty.itesm.mx\n\t\t\tweb8.mty.itesm.mx\n\t\tnet.itesm.mx\n\t\tqueretaro.itesm.mx\n\t\t\tcomunicacionypublicidad.queretaro.itesm.mx\n\t\t\tidentidad.queretaro.itesm.mx\n\t\truv.itesm.mx\n\t\trzn.itesm.mx\n\t\tsal.itesm.mx\n\t\tsistema.itesm.mx\n\t\tsitios.itesm.mx\n\t\tslp.itesm.mx\n\t\tsorteotec.itesm.mx\n\t\ttecreview.itesm.mx\n\t\tzac.itesm.mx\n\ttecreview.mx\nsoy\n\tprepatec.soy\n`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := TreeToList(tt.args.data, tt.args.markdown); !reflect.DeepEqual(got, convertEscapeSequences(tt.want)) {
				t.Errorf("PrintMarkDownList() = %v, want %v", got, tt.want)
			}
		})
	}
}
