# merge-summary-go
Create a summary for a pull request based on the diffs in the pull request.

## Usage
```bash
go run main.go -owver <owner> -repo <repo> -pr <pr>

Usage of merger:
  -gpt string
    	The GPT API key (not recommended, use environment variable GPT_AUTH)
  -owner string
    	The owner of the repository (default "aws")
  -pr string
    	The pull request number (default "544")
  -repo string
    	The name of the repository (default "aws-lambda-go")
```
# Copyright for examples remains with the original authors

## Example MD for aws/aws-lambda-go pull 544 

Changes in the `lambda/invoke_loop.go` file incorporate an update in the `handleInvoke` function to include a cause for XRay when an error occurs, which is passed onto `invoke.failure` function signature. New types `xrayException` and `xrayError` have been added to improve error handling. A new function named `makeXRayError` generates an XRay error detail including the error type, message, stack trace, working directory and paths that are not repeating in the stack trace.

In `lambda/invoke_loop_test.go`, a new function `TestXRayCausePlumbing` is introduced that tests AWS X-Ray feature functionality. The function creates an array of errors and expects corresponding error messages in a certain order. A new handler cycles through these errors and the test verifies that the error messages are as expected.
In `requestRecord` struct, a new `xrayCauses` field has been added to keep track of X-Ray causes. As a result, `runtimeAPIServer` method was updated to append X-Ray error cause from the HTTP header to the `xrayCauses` field in the `record` object for each request.

In `lambda/runtime_api_client.go`, two new constants: `headerXRayErrorCause` and `xrayErrorCauseMaxSize` were added. Modifications have been made to the `success` and `failure` methods within the `invoke` type, accepting two new parameters `nil` and `causeForXRay`. The `post` method accepts an additional parameter `xrayErrorCause`, setting it in the request header if it's length is less than `xrayErrorCauseMaxSize` and not `nil`.

The changes in the `lambda/runtime_api_client_test.go` include adjusting several function calls to adapt to aforementioned `invoke.failure` function modification, impacting multiple test cases including `TestClientDoneAndError`, `TestInvalidRequestsForMalformedEndpoint`, and `TestStatusCodes`. The assertion checks within these tests remain the same.

## example for 
```bash
go run merger.go -owner ml-explore -repo mlx -pr 953 
# log output
```

## Example MD for ml-explore/mlx pull 953

- Apple Inc's copyright was updated to span from 2023 to 2024.
- A new `typename` `AccT` was introduced in `softmax` function.
- Logic was added to handle when `T` and `AccT` are not equal, checked by `std::is_same`.
- Multiple assignments were adjusted based on whether `T` and `AccT` are equal.
- In the `eval_cpu` function, an additional condition checks `precise_` variable and calls different `softmax` functions accordingly.

In the `mlx/backend/metal/kernels/softmax.metal` file, several changes took place:

- The `SOFTMAX_N_READS` template was updated to include `typename` `AccT` and `threadgroup` types were changed from `T` to `AccT`.
- Adjustments were made inside `softmax_single_row` and `softmax_looped` functions to use `AccT` instead of `T`.
- Output calculations inside `softmax_single_row` and `softmax_looped` functions were cast to `T` type.
- An extra definition for `instantiate_softmax_precise` was added for `half` and `bfloat16_t` data types.

In 'softmax.cpp' of the backend metal folder in the `mlx` project, if `dtype` of input is not `float32` and `precise_` is `true`, `precise_` is added to `op_name`.

In the `mlx/fast.cpp` file, changes in the `scaled_dot_product_attention` function simplified the creation of `scores` variable.

In `mlx/ops.cpp` file, a boolean parameter `precise` was added to the `softmax` function, and logic was added to handle when `precise` is `true`.

In the `mlx/ops.h` file, a boolean parameter `precise` was added to the `softmax` function definition.

Modifications in `mlx/primitives.cpp` involved changes in Softmax function's `vmap()` and `jvp()` method implementations and addition of a new method, `is_equivalent()`.

In `mlx/primitives.h`, modifications were made in the Softmax class, with the addition of `precise_`, a boolean parameter in the constructor and `is_equivalent()` function.

In the `python/src/ops.cpp` file, changes were made to the `init_ops` function to introduce a `precise` parameter to the `softmax` function.

In `test_ops.py`, tests evaluating the precision of the softmax function for `mx.float16` and `mx.bfloat16` types were added.

