package doctor

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/krish0723/ait/internal/profile"
)

const jsonSchemaVersion = 1

type jsonReport struct {
	SchemaVersion int           `json:"schema_version"`
	AitVersion    string        `json:"ait_version"`
	Profile       string        `json:"profile"`
	Preset        string        `json:"preset"`
	CWD           string        `json:"cwd"`
	Findings      []jsonFinding `json:"findings"`
}

type jsonFinding struct {
	Code      string `json:"code"`
	Severity  string `json:"severity"`
	Message   string `json:"message"`
	Path      string `json:"path,omitempty"`
	Hint      string `json:"hint,omitempty"`
	DocAnchor string `json:"doc_anchor,omitempty"`
}

func writeJSONReport(out io.Writer, aitVersion string, cwd string, rp *profile.ResolvedProfile, findings []Finding) error {
	prof := ""
	preset := ""
	if rp != nil {
		prof = rp.ProfileID
		preset = rp.PresetID
	}
	if aitVersion == "" {
		aitVersion = "dev"
	}
	jf := make([]jsonFinding, 0, len(findings))
	for _, f := range findings {
		jf = append(jf, jsonFinding{
			Code:      f.Code,
			Severity:  string(f.Severity),
			Message:   f.Message,
			Path:      f.Path,
			Hint:      f.Hint,
			DocAnchor: f.DocAnchor,
		})
	}
	rep := jsonReport{
		SchemaVersion: jsonSchemaVersion,
		AitVersion:    aitVersion,
		Profile:       prof,
		Preset:        preset,
		CWD:           cwd,
		Findings:      jf,
	}
	b, err := json.MarshalIndent(rep, "", "  ")
	if err != nil {
		return err
	}
	_, err = fmt.Fprintln(out, string(b))
	return err
}
