package dao

import "github.com/hawkingrei/bms/model"

func (d *DAO) InsertBenchmark(benchOutput *model.BenchOutput) error {
	stmt, err := d.mysql.Prepare("INSERT INTO BenchData (Date, Commit, Branch, CompilerVersion, Arch, TestName, NsPerOp, AllocsPerOp, BytesPerOp) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)")
	if err != nil {
		return err
	}

	for _, result := range benchOutput.Result {
		_, err = stmt.Exec(benchOutput.Date, benchOutput.Commit, benchOutput.Branch, benchOutput.CompilerVersion, benchOutput.Arch, result.Name, result.NsPerOp, result.AllocsPerOp, result.BytesPerOp)
		if err != nil {
			return err
		}
	}
	return nil
}
