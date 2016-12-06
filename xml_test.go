package reactor

import "testing"
import "github.com/stretchr/testify/require"

func TestParser(t *testing.T) {
	m := MustParseDisplayModel(`<test id="abc" htmlID="def" bool:testt="true" bool:testf="false" reportEvents="click input"> test </test>`)
	require.NotNil(t, m)
	require.Equal(t, "abc", m.ID)
	require.Equal(t, true, m.Attributes["testt"])
	require.Equal(t, false, m.Attributes["testf"])
	require.Equal(t, false, m.Attributes["testf"])
	require.Equal(t, "def", m.Attributes["id"])
	require.Equal(t, []ReportEvent{ReportEvent{Name: "click", ExtraValues: []string{}}, ReportEvent{Name: "input", ExtraValues: []string{}}}, m.ReportEvents)
	require.Equal(t, 1, len(m.Children))
	require.Equal(t, " test ", m.Children[0].Text)
}

func TestParserWithIntermittentTags(t *testing.T) {
	m := MustParseDisplayModel(`<test id="abc" htmlID="def" bool:testt="true" bool:testf="false" reportEvents="click input"> test <span>1</span> 2 </test>`)
	require.NotNil(t, m)
	require.Equal(t, "abc", m.ID)
	require.Equal(t, true, m.Attributes["testt"])
	require.Equal(t, false, m.Attributes["testf"])
	require.Equal(t, false, m.Attributes["testf"])
	require.Equal(t, "def", m.Attributes["id"])
	require.Equal(t, []ReportEvent{
		ReportEvent{
			Name:        "click",
			ExtraValues: []string{},
		},
		ReportEvent{
			Name:        "input",
			ExtraValues: []string{},
		},
	}, m.ReportEvents)
	require.Equal(t, 3, len(m.Children))
	require.Equal(t, " 2 ", m.Children[2].Text)
	require.Equal(t, " test ", m.Children[0].Text)

}

func TestParserWithReportEvent(t *testing.T) {
	m := MustParseDisplayModel(`<test id="abc" htmlID="def" bool:testt="true" bool:testf="false" reportEvents="click:PD input:SP:X-screenX"/>`)
	require.NotNil(t, m)
	require.Equal(t, []ReportEvent{
		ReportEvent{
			PreventDefault: true,
			Name:           "click",
			ExtraValues:    []string{},
		},
		ReportEvent{
			StopPropagation: true,
			Name:            "input",
			ExtraValues:     []string{"screenX"},
		},
	}, m.ReportEvents)

}
