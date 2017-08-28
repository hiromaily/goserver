// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package audits

import (
	json "encoding/json"
	network "github.com/knq/chromedp/cdp/network"
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

func easyjsonC5a4559bDecodeGithubComKnqChromedpCdpAudits(in *jlexer.Lexer, out *GetEncodedResponseReturns) {
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
		case "body":
			out.Body = string(in.String())
		case "originalSize":
			out.OriginalSize = int64(in.Int64())
		case "encodedSize":
			out.EncodedSize = int64(in.Int64())
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
func easyjsonC5a4559bEncodeGithubComKnqChromedpCdpAudits(out *jwriter.Writer, in GetEncodedResponseReturns) {
	out.RawByte('{')
	first := true
	_ = first
	if in.Body != "" {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"body\":")
		out.String(string(in.Body))
	}
	if in.OriginalSize != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"originalSize\":")
		out.Int64(int64(in.OriginalSize))
	}
	if in.EncodedSize != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"encodedSize\":")
		out.Int64(int64(in.EncodedSize))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v GetEncodedResponseReturns) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonC5a4559bEncodeGithubComKnqChromedpCdpAudits(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v GetEncodedResponseReturns) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonC5a4559bEncodeGithubComKnqChromedpCdpAudits(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *GetEncodedResponseReturns) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonC5a4559bDecodeGithubComKnqChromedpCdpAudits(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *GetEncodedResponseReturns) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonC5a4559bDecodeGithubComKnqChromedpCdpAudits(l, v)
}
func easyjsonC5a4559bDecodeGithubComKnqChromedpCdpAudits1(in *jlexer.Lexer, out *GetEncodedResponseParams) {
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
		case "requestId":
			out.RequestID = network.RequestID(in.String())
		case "encoding":
			(out.Encoding).UnmarshalEasyJSON(in)
		case "quality":
			out.Quality = float64(in.Float64())
		case "sizeOnly":
			out.SizeOnly = bool(in.Bool())
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
func easyjsonC5a4559bEncodeGithubComKnqChromedpCdpAudits1(out *jwriter.Writer, in GetEncodedResponseParams) {
	out.RawByte('{')
	first := true
	_ = first
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"requestId\":")
	out.String(string(in.RequestID))
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"encoding\":")
	(in.Encoding).MarshalEasyJSON(out)
	if in.Quality != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"quality\":")
		out.Float64(float64(in.Quality))
	}
	if in.SizeOnly {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"sizeOnly\":")
		out.Bool(bool(in.SizeOnly))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v GetEncodedResponseParams) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonC5a4559bEncodeGithubComKnqChromedpCdpAudits1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v GetEncodedResponseParams) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonC5a4559bEncodeGithubComKnqChromedpCdpAudits1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *GetEncodedResponseParams) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonC5a4559bDecodeGithubComKnqChromedpCdpAudits1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *GetEncodedResponseParams) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonC5a4559bDecodeGithubComKnqChromedpCdpAudits1(l, v)
}
