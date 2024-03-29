package main

type BenchOutput struct {
	Date     string
	Commit   string
	Subject  string
	Author   string
	Arch     string
	TaskName string
	Result   []BenchResult
}

type BenchResult struct {
	Name        string
	NsPerOp     int64
	AllocsPerOp int64
	BytesPerOp  int64
}
