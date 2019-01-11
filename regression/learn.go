package regression

import (
	"github.com/dkonovenschi/linreg/data"
	"gonum.org/v1/gonum/mat"
)

//initialize learning process
func (lr *LinearRegression) Learn() error {
	if lr.solved {
		return errAlreadySolved
	}
	lr.learn()
	lr.solved = true
	lr.calculateR2()
	lr.buildEquation()
	return nil
}

//leaning algorythm:
//- make QR factorization for feature matrix (A=Q*R)
//- create qTy - Transposed Q Matrix (Q^(-1))
//- multiply qTy with target vector
//- calculate weights for regression equation with back substitution of qTy and R
func (lr *LinearRegression) learn() {
	dataset := lr.dataset.GetTrainingSet()

	targetDense, featureDense := makeDenses(dataset)
	q, r := calculateQR(featureDense)
	qty := calculateQTY(q, targetDense)

	_, features := featureDense.Dims()
	lr.Weights = calculateWeights(features, qty, r)
}

//returns matrix of features
func makeFeatureDense(x int, y int, data *[]data.DataEntry) *mat.Dense {
	dense := mat.NewDense(x, y, nil)
	for row := 0; row < x; row++ {
		for column := 0; column < y; column++ {
			dense.Set(row, column, (*data)[row].Features[column])
		}
	}
	return dense
}

//returns matrix of targets
func makeTargetDense(x int, data *[]data.DataEntry) *mat.Dense {
	dense := mat.NewDense(x, 1, nil)
	for row := 0; row < x; row++ {
		dense.Set(row, 0, (*data)[row].Target)
	}
	return dense
}

//returns matrices of targets and features
func makeDenses(dataset *data.Dataset) (targetDense *mat.Dense, featureDense *mat.Dense) {
	targetDense = makeTargetDense(dataset.Info.Samples, dataset.Data)
	featureDense = makeFeatureDense(dataset.Info.Samples, dataset.Info.Features+1, dataset.Data)
	return
}

//calculate QR factorization of feature matrix. (A = Q*R)
//    β=(X^(T) * X)^(−1) * X^(T) * y
// => β=((QR)^(T) * (QR))^(−1) * (QR)^(T) * y
func calculateQR(featureDense *mat.Dense) (*mat.Dense, *mat.Dense) {
	qr := new(mat.QR)
	qr.Factorize(featureDense)
	q := qr.QTo(nil)
	r := qr.RTo(nil)
	return q, r
}

//calculates (QR)^(T) * y
func calculateQTY(q *mat.Dense, targetDense *mat.Dense) *mat.Dense {
	transposedQ := q.T()
	qty := new(mat.Dense)
	qty.Mul(transposedQ, targetDense)
	return qty
}

//calculate weights by back substitution method
func calculateWeights(featureCount int, qty *mat.Dense, r *mat.Dense) []float64 {
	weights := make([]float64, featureCount)
	for row := featureCount - 1; row >= 0; row-- {
		bi := qty.At(row, 0)
		weights[row] = backSubstitutionStep(weights, r, bi, row, featureCount)
	}
	return weights
}

//one step of back substitution
func backSubstitutionStep(X []float64, A *mat.Dense, bi float64, step int, columns int) (xi float64) {
	xi = bi
	for column := step + 1; column < columns; column++ {
		ai := A.At(step, column)
		xi -= X[column] * ai
	}
	return xi / A.At(step, step)
}


