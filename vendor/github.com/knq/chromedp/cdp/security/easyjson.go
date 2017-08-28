// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package security

import (
	json "encoding/json"
	easyjson "github.com/mailru/easyjson"
	jlexer "github.com/mailru/easyjson/jlexer"
	jwriter "github.com/mailru/easyjson/jwriter"
)

// suppress unused package warning
var (
	_ *json.RawMessage
	_ *jlexer.Lexer
	_ *jwriter.Writer
	_ easyjson.Marshaler
)

func easyjsonC5a4559bDecodeGithubComKnqChromedpCdpSecurity(in *jlexer.Lexer, out *StateExplanation) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "securityState":
			(out.SecurityState).UnmarshalEasyJSON(in)
		case "summary":
			out.Summary = string(in.String())
		case "description":
			out.Description = string(in.String())
		case "mixedContentType":
			(out.MixedContentType).UnmarshalEasyJSON(in)
		case "certificate":
			if in.IsNull() {
				in.Skip()
				out.Certificate = nil
			} else {
				in.Delim('[')
				if out.Certificate == nil {
					if !in.IsDelim(']') {
						out.Certificate = make([]string, 0, 4)
					} else {
						out.Certificate = []string{}
					}
				} else {
					out.Certificate = (out.Certificate)[:0]
				}
				for !in.IsDelim(']') {
					var v1 string
					v1 = string(in.String())
					out.Certificate = append(out.Certificate, v1)
					in.WantComma()
				}
				in.Delim(']')
			}
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonC5a4559bEncodeGithubComKnqChromedpCdpSecurity(out *jwriter.Writer, in StateExplanation) {
	out.RawByte('{')
	first := true
	_ = first
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"securityState\":")
	(in.SecurityState).MarshalEasyJSON(out)
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"summary\":")
	out.String(string(in.Summary))
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"description\":")
	out.String(string(in.Description))
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"mixedContentType\":")
	(in.MixedContentType).MarshalEasyJSON(out)
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"certificate\":")
	if in.Certificate == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
		out.RawString("null")
	} else {
		out.RawByte('[')
		for v2, v3 := range in.Certificate {
			if v2 > 0 {
				out.RawByte(',')
			}
			out.String(string(v3))
		}
		out.RawByte(']')
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v StateExplanation) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonC5a4559bEncodeGithubComKnqChromedpCdpSecurity(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v StateExplanation) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonC5a4559bEncodeGithubComKnqChromedpCdpSecurity(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *StateExplanation) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonC5a4559bDecodeGithubComKnqChromedpCdpSecurity(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *StateExplanation) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonC5a4559bDecodeGithubComKnqChromedpCdpSecurity(l, v)
}
func easyjsonC5a4559bDecodeGithubComKnqChromedpCdpSecurity1(in *jlexer.Lexer, out *SetOverrideCertificateErrorsParams) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "override":
			out.Override = bool(in.Bool())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonC5a4559bEncodeGithubComKnqChromedpCdpSecurity1(out *jwriter.Writer, in SetOverrideCertificateErrorsParams) {
	out.RawByte('{')
	first := true
	_ = first
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"override\":")
	out.Bool(bool(in.Override))
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v SetOverrideCertificateErrorsParams) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonC5a4559bEncodeGithubComKnqChromedpCdpSecurity1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v SetOverrideCertificateErrorsParams) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonC5a4559bEncodeGithubComKnqChromedpCdpSecurity1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *SetOverrideCertificateErrorsParams) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonC5a4559bDecodeGithubComKnqChromedpCdpSecurity1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *SetOverrideCertificateErrorsParams) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonC5a4559bDecodeGithubComKnqChromedpCdpSecurity1(l, v)
}
func easyjsonC5a4559bDecodeGithubComKnqChromedpCdpSecurity2(in *jlexer.Lexer, out *InsecureContentStatus) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "ranMixedContent":
			out.RanMixedContent = bool(in.Bool())
		case "displayedMixedContent":
			out.DisplayedMixedContent = bool(in.Bool())
		case "containedMixedForm":
			out.ContainedMixedForm = bool(in.Bool())
		case "ranContentWithCertErrors":
			out.RanContentWithCertErrors = bool(in.Bool())
		case "displayedContentWithCertErrors":
			out.DisplayedContentWithCertErrors = bool(in.Bool())
		case "ranInsecureContentStyle":
			(out.RanInsecureContentStyle).UnmarshalEasyJSON(in)
		case "displayedInsecureContentStyle":
			(out.DisplayedInsecureContentStyle).UnmarshalEasyJSON(in)
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonC5a4559bEncodeGithubComKnqChromedpCdpSecurity2(out *jwriter.Writer, in InsecureContentStatus) {
	out.RawByte('{')
	first := true
	_ = first
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"ranMixedContent\":")
	out.Bool(bool(in.RanMixedContent))
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"displayedMixedContent\":")
	out.Bool(bool(in.DisplayedMixedContent))
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"containedMixedForm\":")
	out.Bool(bool(in.ContainedMixedForm))
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"ranContentWithCertErrors\":")
	out.Bool(bool(in.RanContentWithCertErrors))
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"displayedContentWithCertErrors\":")
	out.Bool(bool(in.DisplayedContentWithCertErrors))
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"ranInsecureContentStyle\":")
	(in.RanInsecureContentStyle).MarshalEasyJSON(out)
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"displayedInsecureContentStyle\":")
	(in.DisplayedInsecureContentStyle).MarshalEasyJSON(out)
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v InsecureContentStatus) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonC5a4559bEncodeGithubComKnqChromedpCdpSecurity2(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v InsecureContentStatus) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonC5a4559bEncodeGithubComKnqChromedpCdpSecurity2(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *InsecureContentStatus) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonC5a4559bDecodeGithubComKnqChromedpCdpSecurity2(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *InsecureContentStatus) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonC5a4559bDecodeGithubComKnqChromedpCdpSecurity2(l, v)
}
func easyjsonC5a4559bDecodeGithubComKnqChromedpCdpSecurity3(in *jlexer.Lexer, out *HandleCertificateErrorParams) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "eventId":
			out.EventID = int64(in.Int64())
		case "action":
			(out.Action).UnmarshalEasyJSON(in)
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonC5a4559bEncodeGithubComKnqChromedpCdpSecurity3(out *jwriter.Writer, in HandleCertificateErrorParams) {
	out.RawByte('{')
	first := true
	_ = first
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"eventId\":")
	out.Int64(int64(in.EventID))
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"action\":")
	(in.Action).MarshalEasyJSON(out)
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v HandleCertificateErrorParams) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonC5a4559bEncodeGithubComKnqChromedpCdpSecurity3(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v HandleCertificateErrorParams) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonC5a4559bEncodeGithubComKnqChromedpCdpSecurity3(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *HandleCertificateErrorParams) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonC5a4559bDecodeGithubComKnqChromedpCdpSecurity3(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *HandleCertificateErrorParams) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonC5a4559bDecodeGithubComKnqChromedpCdpSecurity3(l, v)
}
func easyjsonC5a4559bDecodeGithubComKnqChromedpCdpSecurity4(in *jlexer.Lexer, out *EventSecurityStateChanged) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "securityState":
			(out.SecurityState).UnmarshalEasyJSON(in)
		case "schemeIsCryptographic":
			out.SchemeIsCryptographic = bool(in.Bool())
		case "explanations":
			if in.IsNull() {
				in.Skip()
				out.Explanations = nil
			} else {
				in.Delim('[')
				if out.Explanations == nil {
					if !in.IsDelim(']') {
						out.Explanations = make([]*StateExplanation, 0, 8)
					} else {
						out.Explanations = []*StateExplanation{}
					}
				} else {
					out.Explanations = (out.Explanations)[:0]
				}
				for !in.IsDelim(']') {
					var v4 *StateExplanation
					if in.IsNull() {
						in.Skip()
						v4 = nil
					} else {
						if v4 == nil {
							v4 = new(StateExplanation)
						}
						(*v4).UnmarshalEasyJSON(in)
					}
					out.Explanations = append(out.Explanations, v4)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "insecureContentStatus":
			if in.IsNull() {
				in.Skip()
				out.InsecureContentStatus = nil
			} else {
				if out.InsecureContentStatus == nil {
					out.InsecureContentStatus = new(InsecureContentStatus)
				}
				(*out.InsecureContentStatus).UnmarshalEasyJSON(in)
			}
		case "summary":
			out.Summary = string(in.String())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonC5a4559bEncodeGithubComKnqChromedpCdpSecurity4(out *jwriter.Writer, in EventSecurityStateChanged) {
	out.RawByte('{')
	first := true
	_ = first
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"securityState\":")
	(in.SecurityState).MarshalEasyJSON(out)
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"schemeIsCryptographic\":")
	out.Bool(bool(in.SchemeIsCryptographic))
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"explanations\":")
	if in.Explanations == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
		out.RawString("null")
	} else {
		out.RawByte('[')
		for v5, v6 := range in.Explanations {
			if v5 > 0 {
				out.RawByte(',')
			}
			if v6 == nil {
				out.RawString("null")
			} else {
				(*v6).MarshalEasyJSON(out)
			}
		}
		out.RawByte(']')
	}
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"insecureContentStatus\":")
	if in.InsecureContentStatus == nil {
		out.RawString("null")
	} else {
		(*in.InsecureContentStatus).MarshalEasyJSON(out)
	}
	if in.Summary != "" {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"summary\":")
		out.String(string(in.Summary))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v EventSecurityStateChanged) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonC5a4559bEncodeGithubComKnqChromedpCdpSecurity4(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v EventSecurityStateChanged) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonC5a4559bEncodeGithubComKnqChromedpCdpSecurity4(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *EventSecurityStateChanged) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonC5a4559bDecodeGithubComKnqChromedpCdpSecurity4(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *EventSecurityStateChanged) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonC5a4559bDecodeGithubComKnqChromedpCdpSecurity4(l, v)
}
func easyjsonC5a4559bDecodeGithubComKnqChromedpCdpSecurity5(in *jlexer.Lexer, out *EventCertificateError) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "eventId":
			out.EventID = int64(in.Int64())
		case "errorType":
			out.ErrorType = string(in.String())
		case "requestURL":
			out.RequestURL = string(in.String())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonC5a4559bEncodeGithubComKnqChromedpCdpSecurity5(out *jwriter.Writer, in EventCertificateError) {
	out.RawByte('{')
	first := true
	_ = first
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"eventId\":")
	out.Int64(int64(in.EventID))
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"errorType\":")
	out.String(string(in.ErrorType))
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"requestURL\":")
	out.String(string(in.RequestURL))
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v EventCertificateError) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonC5a4559bEncodeGithubComKnqChromedpCdpSecurity5(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v EventCertificateError) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonC5a4559bEncodeGithubComKnqChromedpCdpSecurity5(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *EventCertificateError) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonC5a4559bDecodeGithubComKnqChromedpCdpSecurity5(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *EventCertificateError) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonC5a4559bDecodeGithubComKnqChromedpCdpSecurity5(l, v)
}
func easyjsonC5a4559bDecodeGithubComKnqChromedpCdpSecurity6(in *jlexer.Lexer, out *EnableParams) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonC5a4559bEncodeGithubComKnqChromedpCdpSecurity6(out *jwriter.Writer, in EnableParams) {
	out.RawByte('{')
	first := true
	_ = first
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v EnableParams) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonC5a4559bEncodeGithubComKnqChromedpCdpSecurity6(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v EnableParams) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonC5a4559bEncodeGithubComKnqChromedpCdpSecurity6(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *EnableParams) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonC5a4559bDecodeGithubComKnqChromedpCdpSecurity6(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *EnableParams) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonC5a4559bDecodeGithubComKnqChromedpCdpSecurity6(l, v)
}
func easyjsonC5a4559bDecodeGithubComKnqChromedpCdpSecurity7(in *jlexer.Lexer, out *DisableParams) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonC5a4559bEncodeGithubComKnqChromedpCdpSecurity7(out *jwriter.Writer, in DisableParams) {
	out.RawByte('{')
	first := true
	_ = first
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v DisableParams) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonC5a4559bEncodeGithubComKnqChromedpCdpSecurity7(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v DisableParams) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonC5a4559bEncodeGithubComKnqChromedpCdpSecurity7(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *DisableParams) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonC5a4559bDecodeGithubComKnqChromedpCdpSecurity7(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *DisableParams) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonC5a4559bDecodeGithubComKnqChromedpCdpSecurity7(l, v)
}
