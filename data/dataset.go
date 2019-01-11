package data

import (
	"log"
)

//dataset type contains data vector and information about it
type Dataset struct {
	Data *[] DataEntry
	Info *Info
}

//info type contains information about dataset
type Info struct {
	Samples  int
	Features int
}

type targetsVector *[] float64
type featuresVector [] float64
type featuresBatch *[][] float64

//Initialization of dataset create&fill procedure
func newDataset(targets targetsVector, features featuresBatch) Dataset {
	checkSize(len(*targets), len(*features))
	gauge := getFeaturesAmountGauge(features)
	return fillDataset(targets, features, gauge)
}

//initialize dataset and filling it with data vector
func fillDataset(targets targetsVector, features featuresBatch, gauge int) (ds Dataset) {
	ds.init()
	for key, target := range *targets {
		ds.addEntry(target, (*features)[key], gauge)
	}
	return
}

//initialize dataset
func (ds *Dataset) init() {
	ds.Data = &[]DataEntry{}
	ds.Info = &Info{}
}

//append new DataEntry to data vector
func (ds *Dataset) addEntry(target float64, features featuresVector, gauge int) {
	checkFeaturesAmount(features, gauge)
	de := newDataEntry(target, features)
	*ds.Data = append(*ds.Data, de)
}

//filling information about dataset
func (ds *Dataset) setInfo(features int, samples int) {
	ds.Info.Samples = samples
	ds.Info.Features = features
}

//panic if targets and features isn't same size
func checkSize(targetsSize int, featuresSize int) {
	if targetsSize != featuresSize {
		log.Panic(errDatasetLen)
	}
}

//returns feature amount of first entry
func getFeaturesAmountGauge(features featuresBatch) int {
	return len((*features)[first])
}

//panic if some of data entries has wrong len
func checkFeaturesAmount(features featuresVector, gauge int) {
	if len(features) != gauge {
		log.Panic(errDatasetFeatureLen)
	}
}

//If dataset hasn't been properly filled up function will panic
func (d *DatasetContainer) checkDatasetConsistency(featureTags int, ds *Dataset, panicOnEmpty bool) int{
	if (*ds).Data == nil {
		if panicOnEmpty {
			log.Panic(errDatasetNotEnoughData)
		}
		return 0
	}

	dsFeatures := len((*ds.Data)[first].Features) - 1
	dsSize := len(*ds.Data)

	//dataset must contain at least dsFeatures+1 samples
	if (dsSize <= dsFeatures) && (dsSize > 0) {
		log.Panic(errDatasetNotEnoughData)
	}

	//every dataset feature have to be tagged
	if (featureTags != dsFeatures) && (dsSize > 0) {
		log.Panic(errDatasetFeatureTags)
	}
	return dsSize
}