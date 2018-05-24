package selfdiagnose

import "testing"
import "os"

func TestReportInXML(t *testing.T) {
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
	f, _ := os.Create("TestReportInXML.xml")
	defer f.Close()
	rep := XMLReporter{f}
	reg.Run(rep)
}
