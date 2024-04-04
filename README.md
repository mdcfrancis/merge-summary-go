# merge-summary-go
Create a summary for a pull request based on the diffs in the pull request.

## Usage
```bash
go run main.go -owner <owner> -repo <repo> -pr <pr>

Usage of merger:
  -gpt string
    	The GPT API key (not recommended, use environment variable GPT_AUTH)
  -owner string
    	The owner of the repository (default "aws")
  -pr string
    	The pull request number (default "544")
  -repo string
    	The name of the repository (default "aws-lambda-go")
  -qualitative
    	Include a qualitative analysis of the changes
```
# Copyright for examples remains with the original authors
## Example MD for mdcfrancis/merge-summary-go pull 1
# File: `merger.go` - Modifications

Several updates were made to `merger.go` largely aimed at enhancing summarization of changes.

## Function: 'chunkToSummary'

The 'chunkToSummary' function underwent a few alterations. Programmer comments were removed and the presentation of the 'prompt' variable has been changed.

## New Functions: 'getFromGPT' and 'summarizeSummary'

Two new functions were incorporated, 'getFromGPT' and 'summarizeSummary'. 'getFromGPT' takes over fetching data from the GPT API, which was previously a part of the 'chunkToSummary' function. 'SummarizeSummary', on the other hand, builds a 'prompt' variable by summarizing an existing summary, with the aim to offer a GPT API response later on.

## Modification: Prompt setup change in 'chunkToSummary'

The prompt setup in the 'chunkToSummary' function was updated to optionally allow a quality analysis of the amendments.

## Script-Level Changes

At the script's top level, a new 'qualitative' flag has been added that dictates whether a qualitative summary is provided or not. The 'body' variable's display in the main function is now commented, and the 'chunkToSummary' function has been superseded by 'summarizeSummary'.

# Quality Analysis

The updates made to the `merger.go` file have been intended to improve the processes of fetching data and summarizing changes. With the cleaned coder comments and new functions, the file is expected to perform more effectively. The quality analysis inclusion in the summary provides deeper insight into the changes. However, the functionality and impact of commenting out the 'body' variable display are yet to be evaluated.

## Example MD for aws/aws-lambda-go pull 544 

Changes in the `lambda/invoke_loop.go` file incorporate an update in the `handleInvoke` function to include a cause for XRay when an error occurs, which is passed onto `invoke.failure` function signature. New types `xrayException` and `xrayError` have been added to improve error handling. A new function named `makeXRayError` generates an XRay error detail including the error type, message, stack trace, working directory and paths that are not repeating in the stack trace.

In `lambda/invoke_loop_test.go`, a new function `TestXRayCausePlumbing` is introduced that tests AWS X-Ray feature functionality. The function creates an array of errors and expects corresponding error messages in a certain order. A new handler cycles through these errors and the test verifies that the error messages are as expected.
In `requestRecord` struct, a new `xrayCauses` field has been added to keep track of X-Ray causes. As a result, `runtimeAPIServer` method was updated to append X-Ray error cause from the HTTP header to the `xrayCauses` field in the `record` object for each request.

In `lambda/runtime_api_client.go`, two new constants: `headerXRayErrorCause` and `xrayErrorCauseMaxSize` were added. Modifications have been made to the `success` and `failure` methods within the `invoke` type, accepting two new parameters `nil` and `causeForXRay`. The `post` method accepts an additional parameter `xrayErrorCause`, setting it in the request header if it's length is less than `xrayErrorCauseMaxSize` and not `nil`.

The changes in the `lambda/runtime_api_client_test.go` include adjusting several function calls to adapt to aforementioned `invoke.failure` function modification, impacting multiple test cases including `TestClientDoneAndError`, `TestInvalidRequestsForMalformedEndpoint`, and `TestStatusCodes`. The assertion checks within these tests remain the same.