# linreg
------------

This repository contains golang implementation of multivariable linear regression


# Installation
------------

    $ go get github.com/konovenski/linreg
    

# Usage
------------

```go

ds := data.NewContainer()
	targets := []float64{11.2, 13.4, 40.7, 5.3, 24.8, 12.7, 20.9, 35.7, 8.7, 9.6, 14.5, 26.9, 15.7, 36.2, 18.1, 28.9, 14.9, 25.8, 21.7, 25.7}
	features := [][]float64{
		{587000, 16.5, 6.2}, {643000, 20.5, 6.4},
		{635000, 26.3, 9.3}, {692000, 16.5, 5.3},
		{1248000, 19.2, 7.3}, {643000, 16.5, 5.9},
		{1964000, 20.2, 6.4}, {1531000, 21.3, 7.6},
		{713000, 17.2, 4.9}, {749000, 14.3, 6.4},
		{7895000, 18.1, 6}, {762000, 23.1, 7.4},
		{2793000, 19.1, 5.8}, {741000, 24.7, 8.6},
		{625000, 18.6, 6.5}, {854000, 24.9, 8.3},
		{716000, 17.9, 6.7}, {921000, 22.4, 8.6},
		{595000, 20.2, 8.4}, {3353000, 16.9, 6.7},
	}
	ds.NewTrainingSet(&targets, &features)
	ds.SetTarget("Murders per annum per 1,000,000 inhabitants")
	ds.SetFeature(0, "Inhabitants")
	ds.SetFeature(1, "Percent with incomes below $5000")
	ds.SetFeature(2, "Percent unemployed")

	reg := new(regression.LinearRegression)
	reg.ApplyDataset(ds)
	reg.Learn()
	reg.Predict([]float64{741000, 24.7, 8.6})

	fmt.Printf("Regression equation:\n%v\n", reg.Equation)

	//  We know this set has an R^2 above 80
	if reg.R2 < 0.8 {
		t.Errorf("R^2 was %.2f, but we expected > 80", reg.R2)
	}

```
