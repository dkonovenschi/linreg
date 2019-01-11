package regression

import (
	"fmt"
	"github.com/dkonovenschi/linreg/data"
)

//Contains current dataset and results of learning process
type LinearRegression struct {
	dataset  *data.DatasetContainer
	Weights  [] float64
	R2       float64
	Equation string
	solved   bool
}

//locks dataset and push it to LR container
func (lr *LinearRegression) ApplyDataset(ds *data.DatasetContainer) {
	ds.Lock()
	lr.dataset = ds
}

func (lr *LinearRegression) buildEquation() {
	equation := fmt.Sprintf("y = %.3g ", lr.Weights[0])
	for i := 1; i < len(lr.Weights); i++ {
		equation += fmt.Sprintf("+ (%s * %.3g) ", lr.dataset.Tags.Features[i-1], lr.Weights[i])
	}
	lr.Equation = equation
}

//returns most frequently used vars
func (lr *LinearRegression) defaultVars() (data *[]data.DataEntry, samples int) {
	dataset := lr.dataset.GetTrainingSet()
	data = dataset.Data
	samples = dataset.Info.Samples
	return
}
