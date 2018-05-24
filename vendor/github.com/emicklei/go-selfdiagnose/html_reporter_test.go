package selfdiagnose

import (
	"html/template"
	"os"
	"testing"
)

func TestReportInHtml(t *testing.T) {
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
	{
		check := reasonInHtml{}
		reg.Register(check)
	}
	{
		check := lorum{}
		check.SetSeverity(SeverityWarning)
		reg.Register(check)
	}
	f, _ := os.Create("TestReportInHtml.html")
	defer f.Close()
	rep := HtmlReporter{f}
	reg.Run(rep)
}

type reasonInHtml struct{}

func (r reasonInHtml) Comment() string { return "<h1>reason uses HTML</h1> (comment is escaped)" }

func (r reasonInHtml) Run(ctx *Context, result *Result) {
	result.Passed = true
	result.Reason = template.HTML("<h1>Header H1</h1> (reason is unescaped)")
}

type lorum struct{ BasicTask }

func (l lorum) Comment() string { return "long lines need to be wrapped" }

func (l lorum) Run(ctx *Context, result *Result) {
	result.Passed = false // to see severity
	result.Reason = "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum."
}
