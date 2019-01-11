package data

import "log"

//Container for Training/Testing datasets and its tags
type DatasetContainer struct {
	Tags     tagsContainer
	Training Dataset
	Testing  Dataset
	lock     bool
}

//Container for data labels(tags)
type tagsContainer struct {
	Target   string
	Features map[int]string
}

//Initialize and return new DatasetContainer
func NewContainer() *DatasetContainer {
	dc := DatasetContainer{}
	dc.Tags = tagsContainer{Features: make(map[int]string)}
	return &dc
}

//Create and fill new training Dataset
func (d *DatasetContainer) NewTrainingSet(targets *[] float64, features *[][] float64) {
	d.Training = newDataset(targets, features)
}

//Returns training Dataset
func (d *DatasetContainer) GetTrainingSet() *Dataset {
	return &d.Training
}

//Create and fill new testing Dataset
func (d *DatasetContainer) NewTestingSet(targets *[] float64, features *[][] float64) {
	d.Testing = newDataset(targets, features)
}

//Returns testing Dataset
func (d *DatasetContainer) GetTestingSet() *Dataset {
	return &d.Testing
}

//Create new feature tag
func (d *DatasetContainer) SetFeature(column int, tag string) {
	d.Tags.Features[column] = tag
}

//Returns feature tag by its column id
func (d *DatasetContainer) GetFeature(column int) string {
	return d.Tags.Features[column]
}

//Create new target tag
func (d *DatasetContainer) SetTarget(tag string) {
	d.Tags.Target = tag
}

//Returns target tag
func (d *DatasetContainer) GetTarget() string {
	return d.Tags.Target
}

//Check DatasetContainer Consistency and lock it
func (d *DatasetContainer) Lock() {
	d.checkConsistency()
	d.lock = true
}

//If container hasn't been properly filled up function will panic
func (d *DatasetContainer) checkConsistency() {
	featureTags := len(d.Tags.Features)
	trainingSize := d.checkDatasetConsistency(featureTags, &d.Training, true)
	testingSize := d.checkDatasetConsistency(featureTags, &d.Testing, false)
	//target tag have to be defined
	if d.Tags.Target == "" {
		log.Panic(errDatasetTargetTag)
	}

	d.setInfo(featureTags, trainingSize, testingSize)
}

//Fill info about datasets
func (d *DatasetContainer) setInfo(featureTags int, trainingSize int, testingSize int) {
	d.Training.setInfo(featureTags, trainingSize)

	if testingSize != 0 {
		d.Testing.setInfo(featureTags, trainingSize)
	}
}
