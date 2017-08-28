// Package debugger provides the Chrome Debugging Protocol
// commands, types, and events for the Debugger domain.
//
// Debugger domain exposes JavaScript debugging capabilities. It allows
// setting and removing breakpoints, stepping through execution, exploring stack
// traces, etc.
//
// Generated by the chromedp-gen command.
package debugger

// Code generated by chromedp-gen. DO NOT EDIT.

import (
	"context"

	cdp "github.com/knq/chromedp/cdp"
	"github.com/knq/chromedp/cdp/runtime"
)

// EnableParams enables debugger for the given page. Clients should not
// assume that the debugging has been enabled until the result for this command
// is received.
type EnableParams struct{}

// Enable enables debugger for the given page. Clients should not assume that
// the debugging has been enabled until the result for this command is received.
func Enable() *EnableParams {
	return &EnableParams{}
}

// Do executes Debugger.enable against the provided context and
// target handler.
func (p *EnableParams) Do(ctxt context.Context, h cdp.Handler) (err error) {
	return h.Execute(ctxt, cdp.CommandDebuggerEnable, nil, nil)
}

// DisableParams disables debugger for given page.
type DisableParams struct{}

// Disable disables debugger for given page.
func Disable() *DisableParams {
	return &DisableParams{}
}

// Do executes Debugger.disable against the provided context and
// target handler.
func (p *DisableParams) Do(ctxt context.Context, h cdp.Handler) (err error) {
	return h.Execute(ctxt, cdp.CommandDebuggerDisable, nil, nil)
}

// SetBreakpointsActiveParams activates / deactivates all breakpoints on the
// page.
type SetBreakpointsActiveParams struct {
	Active bool `json:"active"` // New value for breakpoints active state.
}

// SetBreakpointsActive activates / deactivates all breakpoints on the page.
//
// parameters:
//   active - New value for breakpoints active state.
func SetBreakpointsActive(active bool) *SetBreakpointsActiveParams {
	return &SetBreakpointsActiveParams{
		Active: active,
	}
}

// Do executes Debugger.setBreakpointsActive against the provided context and
// target handler.
func (p *SetBreakpointsActiveParams) Do(ctxt context.Context, h cdp.Handler) (err error) {
	return h.Execute(ctxt, cdp.CommandDebuggerSetBreakpointsActive, p, nil)
}

// SetSkipAllPausesParams makes page not interrupt on any pauses (breakpoint,
// exception, dom exception etc).
type SetSkipAllPausesParams struct {
	Skip bool `json:"skip"` // New value for skip pauses state.
}

// SetSkipAllPauses makes page not interrupt on any pauses (breakpoint,
// exception, dom exception etc).
//
// parameters:
//   skip - New value for skip pauses state.
func SetSkipAllPauses(skip bool) *SetSkipAllPausesParams {
	return &SetSkipAllPausesParams{
		Skip: skip,
	}
}

// Do executes Debugger.setSkipAllPauses against the provided context and
// target handler.
func (p *SetSkipAllPausesParams) Do(ctxt context.Context, h cdp.Handler) (err error) {
	return h.Execute(ctxt, cdp.CommandDebuggerSetSkipAllPauses, p, nil)
}

// SetBreakpointByURLParams sets JavaScript breakpoint at given location
// specified either by URL or URL regex. Once this command is issued, all
// existing parsed scripts will have breakpoints resolved and returned in
// locations property. Further matching script parsing will result in subsequent
// breakpointResolved events issued. This logical breakpoint will survive page
// reloads.
type SetBreakpointByURLParams struct {
	LineNumber   int64  `json:"lineNumber"`             // Line number to set breakpoint at.
	URL          string `json:"url,omitempty"`          // URL of the resources to set breakpoint on.
	URLRegex     string `json:"urlRegex,omitempty"`     // Regex pattern for the URLs of the resources to set breakpoints on. Either url or urlRegex must be specified.
	ColumnNumber int64  `json:"columnNumber,omitempty"` // Offset in the line to set breakpoint at.
	Condition    string `json:"condition,omitempty"`    // Expression to use as a breakpoint condition. When specified, debugger will only stop on the breakpoint if this expression evaluates to true.
}

// SetBreakpointByURL sets JavaScript breakpoint at given location specified
// either by URL or URL regex. Once this command is issued, all existing parsed
// scripts will have breakpoints resolved and returned in locations property.
// Further matching script parsing will result in subsequent breakpointResolved
// events issued. This logical breakpoint will survive page reloads.
//
// parameters:
//   lineNumber - Line number to set breakpoint at.
func SetBreakpointByURL(lineNumber int64) *SetBreakpointByURLParams {
	return &SetBreakpointByURLParams{
		LineNumber: lineNumber,
	}
}

// WithURL uRL of the resources to set breakpoint on.
func (p SetBreakpointByURLParams) WithURL(url string) *SetBreakpointByURLParams {
	p.URL = url
	return &p
}

// WithURLRegex regex pattern for the URLs of the resources to set
// breakpoints on. Either url or urlRegex must be specified.
func (p SetBreakpointByURLParams) WithURLRegex(urlRegex string) *SetBreakpointByURLParams {
	p.URLRegex = urlRegex
	return &p
}

// WithColumnNumber offset in the line to set breakpoint at.
func (p SetBreakpointByURLParams) WithColumnNumber(columnNumber int64) *SetBreakpointByURLParams {
	p.ColumnNumber = columnNumber
	return &p
}

// WithCondition expression to use as a breakpoint condition. When specified,
// debugger will only stop on the breakpoint if this expression evaluates to
// true.
func (p SetBreakpointByURLParams) WithCondition(condition string) *SetBreakpointByURLParams {
	p.Condition = condition
	return &p
}

// SetBreakpointByURLReturns return values.
type SetBreakpointByURLReturns struct {
	BreakpointID BreakpointID `json:"breakpointId,omitempty"` // Id of the created breakpoint for further reference.
	Locations    []*Location  `json:"locations,omitempty"`    // List of the locations this breakpoint resolved into upon addition.
}

// Do executes Debugger.setBreakpointByUrl against the provided context and
// target handler.
//
// returns:
//   breakpointID - Id of the created breakpoint for further reference.
//   locations - List of the locations this breakpoint resolved into upon addition.
func (p *SetBreakpointByURLParams) Do(ctxt context.Context, h cdp.Handler) (breakpointID BreakpointID, locations []*Location, err error) {
	// execute
	var res SetBreakpointByURLReturns
	err = h.Execute(ctxt, cdp.CommandDebuggerSetBreakpointByURL, p, &res)
	if err != nil {
		return "", nil, err
	}

	return res.BreakpointID, res.Locations, nil
}

// SetBreakpointParams sets JavaScript breakpoint at a given location.
type SetBreakpointParams struct {
	Location  *Location `json:"location"`            // Location to set breakpoint in.
	Condition string    `json:"condition,omitempty"` // Expression to use as a breakpoint condition. When specified, debugger will only stop on the breakpoint if this expression evaluates to true.
}

// SetBreakpoint sets JavaScript breakpoint at a given location.
//
// parameters:
//   location - Location to set breakpoint in.
func SetBreakpoint(location *Location) *SetBreakpointParams {
	return &SetBreakpointParams{
		Location: location,
	}
}

// WithCondition expression to use as a breakpoint condition. When specified,
// debugger will only stop on the breakpoint if this expression evaluates to
// true.
func (p SetBreakpointParams) WithCondition(condition string) *SetBreakpointParams {
	p.Condition = condition
	return &p
}

// SetBreakpointReturns return values.
type SetBreakpointReturns struct {
	BreakpointID   BreakpointID `json:"breakpointId,omitempty"`   // Id of the created breakpoint for further reference.
	ActualLocation *Location    `json:"actualLocation,omitempty"` // Location this breakpoint resolved into.
}

// Do executes Debugger.setBreakpoint against the provided context and
// target handler.
//
// returns:
//   breakpointID - Id of the created breakpoint for further reference.
//   actualLocation - Location this breakpoint resolved into.
func (p *SetBreakpointParams) Do(ctxt context.Context, h cdp.Handler) (breakpointID BreakpointID, actualLocation *Location, err error) {
	// execute
	var res SetBreakpointReturns
	err = h.Execute(ctxt, cdp.CommandDebuggerSetBreakpoint, p, &res)
	if err != nil {
		return "", nil, err
	}

	return res.BreakpointID, res.ActualLocation, nil
}

// RemoveBreakpointParams removes JavaScript breakpoint.
type RemoveBreakpointParams struct {
	BreakpointID BreakpointID `json:"breakpointId"`
}

// RemoveBreakpoint removes JavaScript breakpoint.
//
// parameters:
//   breakpointID
func RemoveBreakpoint(breakpointID BreakpointID) *RemoveBreakpointParams {
	return &RemoveBreakpointParams{
		BreakpointID: breakpointID,
	}
}

// Do executes Debugger.removeBreakpoint against the provided context and
// target handler.
func (p *RemoveBreakpointParams) Do(ctxt context.Context, h cdp.Handler) (err error) {
	return h.Execute(ctxt, cdp.CommandDebuggerRemoveBreakpoint, p, nil)
}

// GetPossibleBreakpointsParams returns possible locations for breakpoint.
// scriptId in start and end range locations should be the same.
type GetPossibleBreakpointsParams struct {
	Start              *Location `json:"start"`                        // Start of range to search possible breakpoint locations in.
	End                *Location `json:"end,omitempty"`                // End of range to search possible breakpoint locations in (excluding). When not specified, end of scripts is used as end of range.
	RestrictToFunction bool      `json:"restrictToFunction,omitempty"` // Only consider locations which are in the same (non-nested) function as start.
}

// GetPossibleBreakpoints returns possible locations for breakpoint. scriptId
// in start and end range locations should be the same.
//
// parameters:
//   start - Start of range to search possible breakpoint locations in.
func GetPossibleBreakpoints(start *Location) *GetPossibleBreakpointsParams {
	return &GetPossibleBreakpointsParams{
		Start: start,
	}
}

// WithEnd end of range to search possible breakpoint locations in
// (excluding). When not specified, end of scripts is used as end of range.
func (p GetPossibleBreakpointsParams) WithEnd(end *Location) *GetPossibleBreakpointsParams {
	p.End = end
	return &p
}

// WithRestrictToFunction only consider locations which are in the same
// (non-nested) function as start.
func (p GetPossibleBreakpointsParams) WithRestrictToFunction(restrictToFunction bool) *GetPossibleBreakpointsParams {
	p.RestrictToFunction = restrictToFunction
	return &p
}

// GetPossibleBreakpointsReturns return values.
type GetPossibleBreakpointsReturns struct {
	Locations []*BreakLocation `json:"locations,omitempty"` // List of the possible breakpoint locations.
}

// Do executes Debugger.getPossibleBreakpoints against the provided context and
// target handler.
//
// returns:
//   locations - List of the possible breakpoint locations.
func (p *GetPossibleBreakpointsParams) Do(ctxt context.Context, h cdp.Handler) (locations []*BreakLocation, err error) {
	// execute
	var res GetPossibleBreakpointsReturns
	err = h.Execute(ctxt, cdp.CommandDebuggerGetPossibleBreakpoints, p, &res)
	if err != nil {
		return nil, err
	}

	return res.Locations, nil
}

// ContinueToLocationParams continues execution until specific location is
// reached.
type ContinueToLocationParams struct {
	Location         *Location                          `json:"location"` // Location to continue to.
	TargetCallFrames ContinueToLocationTargetCallFrames `json:"targetCallFrames,omitempty"`
}

// ContinueToLocation continues execution until specific location is reached.
//
// parameters:
//   location - Location to continue to.
func ContinueToLocation(location *Location) *ContinueToLocationParams {
	return &ContinueToLocationParams{
		Location: location,
	}
}

// WithTargetCallFrames [no description].
func (p ContinueToLocationParams) WithTargetCallFrames(targetCallFrames ContinueToLocationTargetCallFrames) *ContinueToLocationParams {
	p.TargetCallFrames = targetCallFrames
	return &p
}

// Do executes Debugger.continueToLocation against the provided context and
// target handler.
func (p *ContinueToLocationParams) Do(ctxt context.Context, h cdp.Handler) (err error) {
	return h.Execute(ctxt, cdp.CommandDebuggerContinueToLocation, p, nil)
}

// StepOverParams steps over the statement.
type StepOverParams struct{}

// StepOver steps over the statement.
func StepOver() *StepOverParams {
	return &StepOverParams{}
}

// Do executes Debugger.stepOver against the provided context and
// target handler.
func (p *StepOverParams) Do(ctxt context.Context, h cdp.Handler) (err error) {
	return h.Execute(ctxt, cdp.CommandDebuggerStepOver, nil, nil)
}

// StepIntoParams steps into the function call.
type StepIntoParams struct{}

// StepInto steps into the function call.
func StepInto() *StepIntoParams {
	return &StepIntoParams{}
}

// Do executes Debugger.stepInto against the provided context and
// target handler.
func (p *StepIntoParams) Do(ctxt context.Context, h cdp.Handler) (err error) {
	return h.Execute(ctxt, cdp.CommandDebuggerStepInto, nil, nil)
}

// StepOutParams steps out of the function call.
type StepOutParams struct{}

// StepOut steps out of the function call.
func StepOut() *StepOutParams {
	return &StepOutParams{}
}

// Do executes Debugger.stepOut against the provided context and
// target handler.
func (p *StepOutParams) Do(ctxt context.Context, h cdp.Handler) (err error) {
	return h.Execute(ctxt, cdp.CommandDebuggerStepOut, nil, nil)
}

// PauseParams stops on the next JavaScript statement.
type PauseParams struct{}

// Pause stops on the next JavaScript statement.
func Pause() *PauseParams {
	return &PauseParams{}
}

// Do executes Debugger.pause against the provided context and
// target handler.
func (p *PauseParams) Do(ctxt context.Context, h cdp.Handler) (err error) {
	return h.Execute(ctxt, cdp.CommandDebuggerPause, nil, nil)
}

// ScheduleStepIntoAsyncParams steps into next scheduled async task if any is
// scheduled before next pause. Returns success when async task is actually
// scheduled, returns error if no task were scheduled or another
// scheduleStepIntoAsync was called.
type ScheduleStepIntoAsyncParams struct{}

// ScheduleStepIntoAsync steps into next scheduled async task if any is
// scheduled before next pause. Returns success when async task is actually
// scheduled, returns error if no task were scheduled or another
// scheduleStepIntoAsync was called.
func ScheduleStepIntoAsync() *ScheduleStepIntoAsyncParams {
	return &ScheduleStepIntoAsyncParams{}
}

// Do executes Debugger.scheduleStepIntoAsync against the provided context and
// target handler.
func (p *ScheduleStepIntoAsyncParams) Do(ctxt context.Context, h cdp.Handler) (err error) {
	return h.Execute(ctxt, cdp.CommandDebuggerScheduleStepIntoAsync, nil, nil)
}

// ResumeParams resumes JavaScript execution.
type ResumeParams struct{}

// Resume resumes JavaScript execution.
func Resume() *ResumeParams {
	return &ResumeParams{}
}

// Do executes Debugger.resume against the provided context and
// target handler.
func (p *ResumeParams) Do(ctxt context.Context, h cdp.Handler) (err error) {
	return h.Execute(ctxt, cdp.CommandDebuggerResume, nil, nil)
}

// SearchInContentParams searches for given string in script content.
type SearchInContentParams struct {
	ScriptID      runtime.ScriptID `json:"scriptId"`                // Id of the script to search in.
	Query         string           `json:"query"`                   // String to search for.
	CaseSensitive bool             `json:"caseSensitive,omitempty"` // If true, search is case sensitive.
	IsRegex       bool             `json:"isRegex,omitempty"`       // If true, treats string parameter as regex.
}

// SearchInContent searches for given string in script content.
//
// parameters:
//   scriptID - Id of the script to search in.
//   query - String to search for.
func SearchInContent(scriptID runtime.ScriptID, query string) *SearchInContentParams {
	return &SearchInContentParams{
		ScriptID: scriptID,
		Query:    query,
	}
}

// WithCaseSensitive if true, search is case sensitive.
func (p SearchInContentParams) WithCaseSensitive(caseSensitive bool) *SearchInContentParams {
	p.CaseSensitive = caseSensitive
	return &p
}

// WithIsRegex if true, treats string parameter as regex.
func (p SearchInContentParams) WithIsRegex(isRegex bool) *SearchInContentParams {
	p.IsRegex = isRegex
	return &p
}

// SearchInContentReturns return values.
type SearchInContentReturns struct {
	Result []*SearchMatch `json:"result,omitempty"` // List of search matches.
}

// Do executes Debugger.searchInContent against the provided context and
// target handler.
//
// returns:
//   result - List of search matches.
func (p *SearchInContentParams) Do(ctxt context.Context, h cdp.Handler) (result []*SearchMatch, err error) {
	// execute
	var res SearchInContentReturns
	err = h.Execute(ctxt, cdp.CommandDebuggerSearchInContent, p, &res)
	if err != nil {
		return nil, err
	}

	return res.Result, nil
}

// SetScriptSourceParams edits JavaScript source live.
type SetScriptSourceParams struct {
	ScriptID     runtime.ScriptID `json:"scriptId"`         // Id of the script to edit.
	ScriptSource string           `json:"scriptSource"`     // New content of the script.
	DryRun       bool             `json:"dryRun,omitempty"` //  If true the change will not actually be applied. Dry run may be used to get result description without actually modifying the code.
}

// SetScriptSource edits JavaScript source live.
//
// parameters:
//   scriptID - Id of the script to edit.
//   scriptSource - New content of the script.
func SetScriptSource(scriptID runtime.ScriptID, scriptSource string) *SetScriptSourceParams {
	return &SetScriptSourceParams{
		ScriptID:     scriptID,
		ScriptSource: scriptSource,
	}
}

// WithDryRun If true the change will not actually be applied. Dry run may be
// used to get result description without actually modifying the code.
func (p SetScriptSourceParams) WithDryRun(dryRun bool) *SetScriptSourceParams {
	p.DryRun = dryRun
	return &p
}

// SetScriptSourceReturns return values.
type SetScriptSourceReturns struct {
	CallFrames       []*CallFrame              `json:"callFrames,omitempty"`       // New stack trace in case editing has happened while VM was stopped.
	StackChanged     bool                      `json:"stackChanged,omitempty"`     // Whether current call stack  was modified after applying the changes.
	AsyncStackTrace  *runtime.StackTrace       `json:"asyncStackTrace,omitempty"`  // Async stack trace, if any.
	ExceptionDetails *runtime.ExceptionDetails `json:"exceptionDetails,omitempty"` // Exception details if any.
}

// Do executes Debugger.setScriptSource against the provided context and
// target handler.
//
// returns:
//   callFrames - New stack trace in case editing has happened while VM was stopped.
//   stackChanged - Whether current call stack  was modified after applying the changes.
//   asyncStackTrace - Async stack trace, if any.
//   exceptionDetails - Exception details if any.
func (p *SetScriptSourceParams) Do(ctxt context.Context, h cdp.Handler) (callFrames []*CallFrame, stackChanged bool, asyncStackTrace *runtime.StackTrace, exceptionDetails *runtime.ExceptionDetails, err error) {
	// execute
	var res SetScriptSourceReturns
	err = h.Execute(ctxt, cdp.CommandDebuggerSetScriptSource, p, &res)
	if err != nil {
		return nil, false, nil, nil, err
	}

	return res.CallFrames, res.StackChanged, res.AsyncStackTrace, res.ExceptionDetails, nil
}

// RestartFrameParams restarts particular call frame from the beginning.
type RestartFrameParams struct {
	CallFrameID CallFrameID `json:"callFrameId"` // Call frame identifier to evaluate on.
}

// RestartFrame restarts particular call frame from the beginning.
//
// parameters:
//   callFrameID - Call frame identifier to evaluate on.
func RestartFrame(callFrameID CallFrameID) *RestartFrameParams {
	return &RestartFrameParams{
		CallFrameID: callFrameID,
	}
}

// RestartFrameReturns return values.
type RestartFrameReturns struct {
	CallFrames      []*CallFrame        `json:"callFrames,omitempty"`      // New stack trace.
	AsyncStackTrace *runtime.StackTrace `json:"asyncStackTrace,omitempty"` // Async stack trace, if any.
}

// Do executes Debugger.restartFrame against the provided context and
// target handler.
//
// returns:
//   callFrames - New stack trace.
//   asyncStackTrace - Async stack trace, if any.
func (p *RestartFrameParams) Do(ctxt context.Context, h cdp.Handler) (callFrames []*CallFrame, asyncStackTrace *runtime.StackTrace, err error) {
	// execute
	var res RestartFrameReturns
	err = h.Execute(ctxt, cdp.CommandDebuggerRestartFrame, p, &res)
	if err != nil {
		return nil, nil, err
	}

	return res.CallFrames, res.AsyncStackTrace, nil
}

// GetScriptSourceParams returns source for the script with given id.
type GetScriptSourceParams struct {
	ScriptID runtime.ScriptID `json:"scriptId"` // Id of the script to get source for.
}

// GetScriptSource returns source for the script with given id.
//
// parameters:
//   scriptID - Id of the script to get source for.
func GetScriptSource(scriptID runtime.ScriptID) *GetScriptSourceParams {
	return &GetScriptSourceParams{
		ScriptID: scriptID,
	}
}

// GetScriptSourceReturns return values.
type GetScriptSourceReturns struct {
	ScriptSource string `json:"scriptSource,omitempty"` // Script source.
}

// Do executes Debugger.getScriptSource against the provided context and
// target handler.
//
// returns:
//   scriptSource - Script source.
func (p *GetScriptSourceParams) Do(ctxt context.Context, h cdp.Handler) (scriptSource string, err error) {
	// execute
	var res GetScriptSourceReturns
	err = h.Execute(ctxt, cdp.CommandDebuggerGetScriptSource, p, &res)
	if err != nil {
		return "", err
	}

	return res.ScriptSource, nil
}

// SetPauseOnExceptionsParams defines pause on exceptions state. Can be set
// to stop on all exceptions, uncaught exceptions or no exceptions. Initial
// pause on exceptions state is none.
type SetPauseOnExceptionsParams struct {
	State ExceptionsState `json:"state"` // Pause on exceptions mode.
}

// SetPauseOnExceptions defines pause on exceptions state. Can be set to stop
// on all exceptions, uncaught exceptions or no exceptions. Initial pause on
// exceptions state is none.
//
// parameters:
//   state - Pause on exceptions mode.
func SetPauseOnExceptions(state ExceptionsState) *SetPauseOnExceptionsParams {
	return &SetPauseOnExceptionsParams{
		State: state,
	}
}

// Do executes Debugger.setPauseOnExceptions against the provided context and
// target handler.
func (p *SetPauseOnExceptionsParams) Do(ctxt context.Context, h cdp.Handler) (err error) {
	return h.Execute(ctxt, cdp.CommandDebuggerSetPauseOnExceptions, p, nil)
}

// EvaluateOnCallFrameParams evaluates expression on a given call frame.
type EvaluateOnCallFrameParams struct {
	CallFrameID           CallFrameID `json:"callFrameId"`                     // Call frame identifier to evaluate on.
	Expression            string      `json:"expression"`                      // Expression to evaluate.
	ObjectGroup           string      `json:"objectGroup,omitempty"`           // String object group name to put result into (allows rapid releasing resulting object handles using releaseObjectGroup).
	IncludeCommandLineAPI bool        `json:"includeCommandLineAPI,omitempty"` // Specifies whether command line API should be available to the evaluated expression, defaults to false.
	Silent                bool        `json:"silent,omitempty"`                // In silent mode exceptions thrown during evaluation are not reported and do not pause execution. Overrides setPauseOnException state.
	ReturnByValue         bool        `json:"returnByValue,omitempty"`         // Whether the result is expected to be a JSON object that should be sent by value.
	GeneratePreview       bool        `json:"generatePreview,omitempty"`       // Whether preview should be generated for the result.
	ThrowOnSideEffect     bool        `json:"throwOnSideEffect,omitempty"`     // Whether to throw an exception if side effect cannot be ruled out during evaluation.
}

// EvaluateOnCallFrame evaluates expression on a given call frame.
//
// parameters:
//   callFrameID - Call frame identifier to evaluate on.
//   expression - Expression to evaluate.
func EvaluateOnCallFrame(callFrameID CallFrameID, expression string) *EvaluateOnCallFrameParams {
	return &EvaluateOnCallFrameParams{
		CallFrameID: callFrameID,
		Expression:  expression,
	}
}

// WithObjectGroup string object group name to put result into (allows rapid
// releasing resulting object handles using releaseObjectGroup).
func (p EvaluateOnCallFrameParams) WithObjectGroup(objectGroup string) *EvaluateOnCallFrameParams {
	p.ObjectGroup = objectGroup
	return &p
}

// WithIncludeCommandLineAPI specifies whether command line API should be
// available to the evaluated expression, defaults to false.
func (p EvaluateOnCallFrameParams) WithIncludeCommandLineAPI(includeCommandLineAPI bool) *EvaluateOnCallFrameParams {
	p.IncludeCommandLineAPI = includeCommandLineAPI
	return &p
}

// WithSilent in silent mode exceptions thrown during evaluation are not
// reported and do not pause execution. Overrides setPauseOnException state.
func (p EvaluateOnCallFrameParams) WithSilent(silent bool) *EvaluateOnCallFrameParams {
	p.Silent = silent
	return &p
}

// WithReturnByValue whether the result is expected to be a JSON object that
// should be sent by value.
func (p EvaluateOnCallFrameParams) WithReturnByValue(returnByValue bool) *EvaluateOnCallFrameParams {
	p.ReturnByValue = returnByValue
	return &p
}

// WithGeneratePreview whether preview should be generated for the result.
func (p EvaluateOnCallFrameParams) WithGeneratePreview(generatePreview bool) *EvaluateOnCallFrameParams {
	p.GeneratePreview = generatePreview
	return &p
}

// WithThrowOnSideEffect whether to throw an exception if side effect cannot
// be ruled out during evaluation.
func (p EvaluateOnCallFrameParams) WithThrowOnSideEffect(throwOnSideEffect bool) *EvaluateOnCallFrameParams {
	p.ThrowOnSideEffect = throwOnSideEffect
	return &p
}

// EvaluateOnCallFrameReturns return values.
type EvaluateOnCallFrameReturns struct {
	Result           *runtime.RemoteObject     `json:"result,omitempty"`           // Object wrapper for the evaluation result.
	ExceptionDetails *runtime.ExceptionDetails `json:"exceptionDetails,omitempty"` // Exception details.
}

// Do executes Debugger.evaluateOnCallFrame against the provided context and
// target handler.
//
// returns:
//   result - Object wrapper for the evaluation result.
//   exceptionDetails - Exception details.
func (p *EvaluateOnCallFrameParams) Do(ctxt context.Context, h cdp.Handler) (result *runtime.RemoteObject, exceptionDetails *runtime.ExceptionDetails, err error) {
	// execute
	var res EvaluateOnCallFrameReturns
	err = h.Execute(ctxt, cdp.CommandDebuggerEvaluateOnCallFrame, p, &res)
	if err != nil {
		return nil, nil, err
	}

	return res.Result, res.ExceptionDetails, nil
}

// SetVariableValueParams changes value of variable in a callframe.
// Object-based scopes are not supported and must be mutated manually.
type SetVariableValueParams struct {
	ScopeNumber  int64                 `json:"scopeNumber"`  // 0-based number of scope as was listed in scope chain. Only 'local', 'closure' and 'catch' scope types are allowed. Other scopes could be manipulated manually.
	VariableName string                `json:"variableName"` // Variable name.
	NewValue     *runtime.CallArgument `json:"newValue"`     // New variable value.
	CallFrameID  CallFrameID           `json:"callFrameId"`  // Id of callframe that holds variable.
}

// SetVariableValue changes value of variable in a callframe. Object-based
// scopes are not supported and must be mutated manually.
//
// parameters:
//   scopeNumber - 0-based number of scope as was listed in scope chain. Only 'local', 'closure' and 'catch' scope types are allowed. Other scopes could be manipulated manually.
//   variableName - Variable name.
//   newValue - New variable value.
//   callFrameID - Id of callframe that holds variable.
func SetVariableValue(scopeNumber int64, variableName string, newValue *runtime.CallArgument, callFrameID CallFrameID) *SetVariableValueParams {
	return &SetVariableValueParams{
		ScopeNumber:  scopeNumber,
		VariableName: variableName,
		NewValue:     newValue,
		CallFrameID:  callFrameID,
	}
}

// Do executes Debugger.setVariableValue against the provided context and
// target handler.
func (p *SetVariableValueParams) Do(ctxt context.Context, h cdp.Handler) (err error) {
	return h.Execute(ctxt, cdp.CommandDebuggerSetVariableValue, p, nil)
}

// SetAsyncCallStackDepthParams enables or disables async call stacks
// tracking.
type SetAsyncCallStackDepthParams struct {
	MaxDepth int64 `json:"maxDepth"` // Maximum depth of async call stacks. Setting to 0 will effectively disable collecting async call stacks (default).
}

// SetAsyncCallStackDepth enables or disables async call stacks tracking.
//
// parameters:
//   maxDepth - Maximum depth of async call stacks. Setting to 0 will effectively disable collecting async call stacks (default).
func SetAsyncCallStackDepth(maxDepth int64) *SetAsyncCallStackDepthParams {
	return &SetAsyncCallStackDepthParams{
		MaxDepth: maxDepth,
	}
}

// Do executes Debugger.setAsyncCallStackDepth against the provided context and
// target handler.
func (p *SetAsyncCallStackDepthParams) Do(ctxt context.Context, h cdp.Handler) (err error) {
	return h.Execute(ctxt, cdp.CommandDebuggerSetAsyncCallStackDepth, p, nil)
}

// SetBlackboxPatternsParams replace previous blackbox patterns with passed
// ones. Forces backend to skip stepping/pausing in scripts with url matching
// one of the patterns. VM will try to leave blackboxed script by performing
// 'step in' several times, finally resorting to 'step out' if unsuccessful.
type SetBlackboxPatternsParams struct {
	Patterns []string `json:"patterns"` // Array of regexps that will be used to check script url for blackbox state.
}

// SetBlackboxPatterns replace previous blackbox patterns with passed ones.
// Forces backend to skip stepping/pausing in scripts with url matching one of
// the patterns. VM will try to leave blackboxed script by performing 'step in'
// several times, finally resorting to 'step out' if unsuccessful.
//
// parameters:
//   patterns - Array of regexps that will be used to check script url for blackbox state.
func SetBlackboxPatterns(patterns []string) *SetBlackboxPatternsParams {
	return &SetBlackboxPatternsParams{
		Patterns: patterns,
	}
}

// Do executes Debugger.setBlackboxPatterns against the provided context and
// target handler.
func (p *SetBlackboxPatternsParams) Do(ctxt context.Context, h cdp.Handler) (err error) {
	return h.Execute(ctxt, cdp.CommandDebuggerSetBlackboxPatterns, p, nil)
}

// SetBlackboxedRangesParams makes backend skip steps in the script in
// blackboxed ranges. VM will try leave blacklisted scripts by performing 'step
// in' several times, finally resorting to 'step out' if unsuccessful. Positions
// array contains positions where blackbox state is changed. First interval
// isn't blackboxed. Array should be sorted.
type SetBlackboxedRangesParams struct {
	ScriptID  runtime.ScriptID  `json:"scriptId"` // Id of the script.
	Positions []*ScriptPosition `json:"positions"`
}

// SetBlackboxedRanges makes backend skip steps in the script in blackboxed
// ranges. VM will try leave blacklisted scripts by performing 'step in' several
// times, finally resorting to 'step out' if unsuccessful. Positions array
// contains positions where blackbox state is changed. First interval isn't
// blackboxed. Array should be sorted.
//
// parameters:
//   scriptID - Id of the script.
//   positions
func SetBlackboxedRanges(scriptID runtime.ScriptID, positions []*ScriptPosition) *SetBlackboxedRangesParams {
	return &SetBlackboxedRangesParams{
		ScriptID:  scriptID,
		Positions: positions,
	}
}

// Do executes Debugger.setBlackboxedRanges against the provided context and
// target handler.
func (p *SetBlackboxedRangesParams) Do(ctxt context.Context, h cdp.Handler) (err error) {
	return h.Execute(ctxt, cdp.CommandDebuggerSetBlackboxedRanges, p, nil)
}
