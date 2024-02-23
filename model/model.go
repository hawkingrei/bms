package model

type BenchOutput struct {
	Date            string
	Commit          string
	Branch          string
	CompilerVersion string
	Arch            string
	CPUInfo         string
	Flag            string
	Result          []BenchResult
}

type BenchResult struct {
	Name        string
	NsPerOp     int64
	AllocsPerOp int64
	BytesPerOp  int64
}
