package mindmap

import (
	"encoding/json"
	"errors"
	"reflect"
	"testing"
)

type scannerInputMock struct {
}

func (s scannerInputMock) ReadLines() <-chan LineResult {
	return scannerInputReadLinesFunc()
}

var scannerInputReadLinesFunc func() <-chan LineResult

func TestCreateMindMap(t *testing.T) {
	type args struct {
		source scannerInputMock
	}
	// mocking
	scannerMock := scannerInputMock{}
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

	var treeTest1 Node
	err := json.Unmarshal([]byte(dataTest1), &treeTest1)
	if err != nil {
		panic(err)
	}

	dataTest3 := `{
		"com": {
			"google.com": {
				"www.google.com": {}
			},
			"host.com": {
				"www.host.com": {}
			}
		}
	}`

	var treeTest3 Node
	err = json.Unmarshal([]byte(dataTest3), &treeTest3)
	if err != nil {
		panic(err)
	}

	tests := []struct {
		name          string
		args          args
		readLinesFunc func() <-chan LineResult
		want          Node
		wantErr       bool
	}{
		{
			name: "Test 1: Correctly parsing domains",
			args: args{
				source: scannerMock,
			},
			readLinesFunc: func() <-chan LineResult {
				linesCh := make(chan LineResult)

				go func() {
					defer close(linesCh)

					lines := []string{
						"chi.itesm.mx",
						"itesm.mx",
						"ags.itesm.mx",
						"slp.itesm.mx",
						"tecreview.mx",
						"rzn.itesm.mx",
						"mty.itesm.mx",
						"web8.mty.itesm.mx",
						"sistema.itesm.mx",
						"sorteotec.itesm.mx",
						"prepatec.soy",
						"zac.itesm.mx",
						"ruv.itesm.mx",
						"itesm.edu.mx",
						"lag.itesm.mx",
						"dm.itesm.mx",
						"cegs.itesm.mx",
						"tecreview.itesm.mx",
						"exatec1.itesm.mx",
						"btec.itesm.mx",
						"tecmilenio.edu.mx",
						"net.itesm.mx",
						"comunicacionypublicidad.queretaro.itesm.mx",
						"apps.itesm.mx",
						"sitios.itesm.mx",
						"admision.itesm.mx",
						"cdj.itesm.mx",
						"queretaro.itesm.mx",
						"identidad.queretaro.itesm.mx",
						"admisionprepatec.itesm.mx",
						"sal.itesm.mx",
					}

					for _, line := range lines {
						linesCh <- LineResult{Line: line}
					}

				}()

				return linesCh
			},
			want:    treeTest1,
			wantErr: false,
		},
		{
			name: "Test 2: Error while reading input",
			args: args{
				source: scannerMock,
			},
			readLinesFunc: func() <-chan LineResult {
				linesCh := make(chan LineResult)
				go func() {
					defer close(linesCh)
					linesCh <- LineResult{Error: errors.New("something went wrong")}
				}()
				return linesCh
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Test 3: Correctly parsing domains that contain paths",
			args: args{
				source: scannerMock,
			},
			readLinesFunc: func() <-chan LineResult {
				linesCh := make(chan LineResult)

				go func() {
					defer close(linesCh)

					lines := []string{
						"www.host.com/path1",
						"www.host.com/path2",
						"www.host.com/path3",
						"www.google.com",
					}

					for _, line := range lines {
						linesCh <- LineResult{Line: line}
					}

				}()

				return linesCh
			},
			want:    treeTest3,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.readLinesFunc != nil {
				scannerInputReadLinesFunc = tt.readLinesFunc
			}
			got, err := CreateMindMap(tt.args.source)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateMindMap() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateMindMap() = %v, want %v", got, tt.want)
			}
		})
	}
}
