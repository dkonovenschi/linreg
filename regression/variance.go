package regression

import (
	"math"
)

//calculates mean of target and predicted vectors
func (lr *LinearRegression) calculateMean() (targetMean, predictionMean float64) {
	data, samples := lr.defaultVars()
	var targetSum, predictionSum float64

	for sample := 0; sample < samples; sample++ {
		targetSum += (*data)[sample].Target
		predictionSum += (*data)[sample].Prediction
	}
	targetMean = targetSum / float64(samples)
	predictionMean = predictionSum / float64(samples)
	return
}

//calculates variance of target and predicted vectors
//using for R2
func (lr *LinearRegression) calculateVariance() (targetVariance, predictionVariance float64) {
	data, samples := lr.defaultVars()

	targetMean, predictionMean := lr.calculateMean()

	for i := 0; i < samples; i++ {
		target := (*data)[i].Target
		prediction := (*data)[i].Prediction

		targetVariance += math.Pow(target-targetMean, 2)
		predictionVariance += math.Pow(prediction-predictionMean, 2)
	}

	return (targetVariance / float64(samples)),
		(predictionVariance / float64(samples))
}

//calculate R2
func (lr *LinearRegression) calculateR2() {
	lr.PredictionsForTrainingDataset()
	targetVariance, predictionVariance := lr.calculateVariance()
	lr.R2 = predictionVariance / targetVariance
}
