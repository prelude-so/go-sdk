# Authentication

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/prelude-go">prelude</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/prelude-go#AuthenticationNewResponse">AuthenticationNewResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/prelude-go">prelude</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/prelude-go#AuthenticationGetResponse">AuthenticationGetResponse</a>

Methods:

- <code title="post /authentication">client.Authentication.<a href="https://pkg.go.dev/github.com/stainless-sdks/prelude-go#AuthenticationService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/prelude-go">prelude</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/prelude-go#AuthenticationNewParams">AuthenticationNewParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/prelude-go">prelude</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/prelude-go#AuthenticationNewResponse">AuthenticationNewResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /authentication/{auth_uuid}">client.Authentication.<a href="https://pkg.go.dev/github.com/stainless-sdks/prelude-go#AuthenticationService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, authUuid <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/prelude-go">prelude</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/prelude-go#AuthenticationGetResponse">AuthenticationGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Feedback

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/prelude-go">prelude</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/prelude-go#AuthenticationFeedbackNewResponse">AuthenticationFeedbackNewResponse</a>

Methods:

- <code title="post /authentication/feedback">client.Authentication.Feedback.<a href="https://pkg.go.dev/github.com/stainless-sdks/prelude-go#AuthenticationFeedbackService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/prelude-go">prelude</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/prelude-go#AuthenticationFeedbackNewParams">AuthenticationFeedbackNewParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/prelude-go">prelude</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/prelude-go#AuthenticationFeedbackNewResponse">AuthenticationFeedbackNewResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Check

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/prelude-go">prelude</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/prelude-go#CheckNewResponse">CheckNewResponse</a>

Methods:

- <code title="post /check">client.Check.<a href="https://pkg.go.dev/github.com/stainless-sdks/prelude-go#CheckService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/prelude-go">prelude</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/prelude-go#CheckNewParams">CheckNewParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/prelude-go">prelude</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/prelude-go#CheckNewResponse">CheckNewResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Retry

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/prelude-go">prelude</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/prelude-go#RetryNewResponse">RetryNewResponse</a>

Methods:

- <code title="post /retry">client.Retry.<a href="https://pkg.go.dev/github.com/stainless-sdks/prelude-go#RetryService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/prelude-go">prelude</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/prelude-go#RetryNewParams">RetryNewParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/prelude-go">prelude</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/prelude-go#RetryNewResponse">RetryNewResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Lookup

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/prelude-go">prelude</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/prelude-go#LookupGetResponse">LookupGetResponse</a>

Methods:

- <code title="get /lookup/{phone_number}">client.Lookup.<a href="https://pkg.go.dev/github.com/stainless-sdks/prelude-go#LookupService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, phoneNumber <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/prelude-go">prelude</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/prelude-go#LookupGetParams">LookupGetParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/prelude-go">prelude</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/prelude-go#LookupGetResponse">LookupGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
