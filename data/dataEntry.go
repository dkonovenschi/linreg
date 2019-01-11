package data

//DataEntry contains one row of dataset and results of learning process
type DataEntry struct {
	Target     float64
	Features   [] float64
	Prediction float64
	Error      float64
}

//creates new data entry
func newDataEntry(target float64, features [] float64) DataEntry {
	return DataEntry{Target: target, Features: append([]float64{1}, features...)} //first element have to be 1
}
