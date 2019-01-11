package data

import "errors"

var errDatasetLen = errors.New("Length of Dataset entries 'Targets' and 'Features' have to be be equal ")
var errDatasetFeatureLen = errors.New("Features storage requires slices to be the same size ")
var errDatasetFeatureTags = errors.New("You have to set Tags for every feature in Dataset ")
var errDatasetTargetTag = errors.New("You have to set tag for Target variable ")
var errDatasetNotEnoughData = errors.New("Not enough data in Dataset ")