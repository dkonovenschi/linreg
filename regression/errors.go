package regression

import "errors"

var errRegressionNotSolved = errors.New("Run LinearRegression.Learn() method first ")
var errWrongFeaturesAmount = errors.New("Wrong amount of features ")
var errAlreadySolved = errors.New("Regression already solved ")