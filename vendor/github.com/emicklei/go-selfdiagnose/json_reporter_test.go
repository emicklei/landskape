package selfdiagnose

import "testing"
import "os"

func TestReportInJSON(t *testing.T) {
	reg := &Registry{}
	{
		check := ReportMessage{}
		check.SetComment("test critical")
		reg.Register(check)
	}
	{
		check := ReportMessage{}
		check.SetSeverity(SeverityWarning)
		check.SetComment("test warning")
		reg.Register(check)
	}
	{
		check := ReportMessage{}
		check.SetComment("test none")
		reg.Register(check)
	}
	{
		check := ReportMessage{}
		check.SetComment("test odd/even")
		reg.Register(check)
	}
	f, _ := os.Create("TestReportInJSON.json")
	defer f.Close()
	rep := JSONReporter{f}
	reg.Run(rep)
}
