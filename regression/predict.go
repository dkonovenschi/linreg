package regression

//predicts target by set of features
//panic if LR isn't solved yet
func (lr *LinearRegression) Predict(features []float64) (prediction float64, err error) {
	PanicCheck(isSolved(lr.solved))
	dataset := lr.dataset.GetTrainingSet()
	if len(features) < dataset.Info.Features {
		return 0, errWrongFeaturesAmount
	}

	prediction = lr.Weights[0]
	for i := 1; i < dataset.Info.Features+1; i++ {
		prediction += lr.Weights[i] * features[i-1]
	}

	return prediction, nil
}

//checks if LR is solved
func isSolved(solved bool) error {
	if !solved {
		return errRegressionNotSolved
	}
	return nil
}

//calculates predictions for existing DataEntries.
//using for R2 calculating
func (lr *LinearRegression) PredictionsForTrainingDataset() {
	data, samples := lr.defaultVars()
	for sample := 0; sample < samples; sample++ {
		(*data)[sample].Prediction, _ = lr.Predict((*data)[sample].Features[1:])
		(*data)[sample].Error = (*data)[sample].Prediction - (*data)[sample].Target
	}
}
