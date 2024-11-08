# Transactional

Response Types:

- <a href="https://pkg.go.dev/github.com/prelude-so/go-sdk">prelude</a>.<a href="https://pkg.go.dev/github.com/prelude-so/go-sdk#TransactionalSendResponse">TransactionalSendResponse</a>

Methods:

- <code title="post /v2/transactional">client.Transactional.<a href="https://pkg.go.dev/github.com/prelude-so/go-sdk#TransactionalService.Send">Send</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/prelude-so/go-sdk">prelude</a>.<a href="https://pkg.go.dev/github.com/prelude-so/go-sdk#TransactionalSendParams">TransactionalSendParams</a>) (<a href="https://pkg.go.dev/github.com/prelude-so/go-sdk">prelude</a>.<a href="https://pkg.go.dev/github.com/prelude-so/go-sdk#TransactionalSendResponse">TransactionalSendResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Verification

Response Types:

- <a href="https://pkg.go.dev/github.com/prelude-so/go-sdk">prelude</a>.<a href="https://pkg.go.dev/github.com/prelude-so/go-sdk#VerificationNewResponse">VerificationNewResponse</a>
- <a href="https://pkg.go.dev/github.com/prelude-so/go-sdk">prelude</a>.<a href="https://pkg.go.dev/github.com/prelude-so/go-sdk#VerificationCheckResponse">VerificationCheckResponse</a>

Methods:

- <code title="post /v2/verification">client.Verification.<a href="https://pkg.go.dev/github.com/prelude-so/go-sdk#VerificationService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/prelude-so/go-sdk">prelude</a>.<a href="https://pkg.go.dev/github.com/prelude-so/go-sdk#VerificationNewParams">VerificationNewParams</a>) (<a href="https://pkg.go.dev/github.com/prelude-so/go-sdk">prelude</a>.<a href="https://pkg.go.dev/github.com/prelude-so/go-sdk#VerificationNewResponse">VerificationNewResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /v2/verification/check">client.Verification.<a href="https://pkg.go.dev/github.com/prelude-so/go-sdk#VerificationService.Check">Check</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/prelude-so/go-sdk">prelude</a>.<a href="https://pkg.go.dev/github.com/prelude-so/go-sdk#VerificationCheckParams">VerificationCheckParams</a>) (<a href="https://pkg.go.dev/github.com/prelude-so/go-sdk">prelude</a>.<a href="https://pkg.go.dev/github.com/prelude-so/go-sdk#VerificationCheckResponse">VerificationCheckResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Watch

Response Types:

- <a href="https://pkg.go.dev/github.com/prelude-so/go-sdk">prelude</a>.<a href="https://pkg.go.dev/github.com/prelude-so/go-sdk#WatchFeedbackResponse">WatchFeedbackResponse</a>
- <a href="https://pkg.go.dev/github.com/prelude-so/go-sdk">prelude</a>.<a href="https://pkg.go.dev/github.com/prelude-so/go-sdk#WatchPredictResponse">WatchPredictResponse</a>

Methods:

- <code title="post /v2/watch/feedback">client.Watch.<a href="https://pkg.go.dev/github.com/prelude-so/go-sdk#WatchService.Feedback">Feedback</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/prelude-so/go-sdk">prelude</a>.<a href="https://pkg.go.dev/github.com/prelude-so/go-sdk#WatchFeedbackParams">WatchFeedbackParams</a>) (<a href="https://pkg.go.dev/github.com/prelude-so/go-sdk">prelude</a>.<a href="https://pkg.go.dev/github.com/prelude-so/go-sdk#WatchFeedbackResponse">WatchFeedbackResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /v2/watch/predict">client.Watch.<a href="https://pkg.go.dev/github.com/prelude-so/go-sdk#WatchService.Predict">Predict</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/prelude-so/go-sdk">prelude</a>.<a href="https://pkg.go.dev/github.com/prelude-so/go-sdk#WatchPredictParams">WatchPredictParams</a>) (<a href="https://pkg.go.dev/github.com/prelude-so/go-sdk">prelude</a>.<a href="https://pkg.go.dev/github.com/prelude-so/go-sdk#WatchPredictResponse">WatchPredictResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
