# Lookup

Response Types:

- <a href="https://pkg.go.dev/github.com/prelude-so/go-sdk">prelude</a>.<a href="https://pkg.go.dev/github.com/prelude-so/go-sdk#LookupLookupResponse">LookupLookupResponse</a>

Methods:

- <code title="get /v2/lookup/{phone_number}">client.Lookup.<a href="https://pkg.go.dev/github.com/prelude-so/go-sdk#LookupService.Lookup">Lookup</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, phoneNumber <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/prelude-so/go-sdk">prelude</a>.<a href="https://pkg.go.dev/github.com/prelude-so/go-sdk#LookupLookupParams">LookupLookupParams</a>) (<a href="https://pkg.go.dev/github.com/prelude-so/go-sdk">prelude</a>.<a href="https://pkg.go.dev/github.com/prelude-so/go-sdk#LookupLookupResponse">LookupLookupResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

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

# VerificationManagement

Response Types:

- <a href="https://pkg.go.dev/github.com/prelude-so/go-sdk">prelude</a>.<a href="https://pkg.go.dev/github.com/prelude-so/go-sdk#VerificationManagementListSenderIDsResponse">VerificationManagementListSenderIDsResponse</a>
- <a href="https://pkg.go.dev/github.com/prelude-so/go-sdk">prelude</a>.<a href="https://pkg.go.dev/github.com/prelude-so/go-sdk#VerificationManagementSubmitSenderIDResponse">VerificationManagementSubmitSenderIDResponse</a>

Methods:

- <code title="get /v2/verification/management/sender-id">client.VerificationManagement.<a href="https://pkg.go.dev/github.com/prelude-so/go-sdk#VerificationManagementService.ListSenderIDs">ListSenderIDs</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (<a href="https://pkg.go.dev/github.com/prelude-so/go-sdk">prelude</a>.<a href="https://pkg.go.dev/github.com/prelude-so/go-sdk#VerificationManagementListSenderIDsResponse">VerificationManagementListSenderIDsResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /v2/verification/management/sender-id">client.VerificationManagement.<a href="https://pkg.go.dev/github.com/prelude-so/go-sdk#VerificationManagementService.SubmitSenderID">SubmitSenderID</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/prelude-so/go-sdk">prelude</a>.<a href="https://pkg.go.dev/github.com/prelude-so/go-sdk#VerificationManagementSubmitSenderIDParams">VerificationManagementSubmitSenderIDParams</a>) (<a href="https://pkg.go.dev/github.com/prelude-so/go-sdk">prelude</a>.<a href="https://pkg.go.dev/github.com/prelude-so/go-sdk#VerificationManagementSubmitSenderIDResponse">VerificationManagementSubmitSenderIDResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Watch

Response Types:

- <a href="https://pkg.go.dev/github.com/prelude-so/go-sdk">prelude</a>.<a href="https://pkg.go.dev/github.com/prelude-so/go-sdk#WatchPredictResponse">WatchPredictResponse</a>
- <a href="https://pkg.go.dev/github.com/prelude-so/go-sdk">prelude</a>.<a href="https://pkg.go.dev/github.com/prelude-so/go-sdk#WatchSendEventsResponse">WatchSendEventsResponse</a>
- <a href="https://pkg.go.dev/github.com/prelude-so/go-sdk">prelude</a>.<a href="https://pkg.go.dev/github.com/prelude-so/go-sdk#WatchSendFeedbacksResponse">WatchSendFeedbacksResponse</a>

Methods:

- <code title="post /v2/watch/predict">client.Watch.<a href="https://pkg.go.dev/github.com/prelude-so/go-sdk#WatchService.Predict">Predict</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/prelude-so/go-sdk">prelude</a>.<a href="https://pkg.go.dev/github.com/prelude-so/go-sdk#WatchPredictParams">WatchPredictParams</a>) (<a href="https://pkg.go.dev/github.com/prelude-so/go-sdk">prelude</a>.<a href="https://pkg.go.dev/github.com/prelude-so/go-sdk#WatchPredictResponse">WatchPredictResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /v2/watch/event">client.Watch.<a href="https://pkg.go.dev/github.com/prelude-so/go-sdk#WatchService.SendEvents">SendEvents</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/prelude-so/go-sdk">prelude</a>.<a href="https://pkg.go.dev/github.com/prelude-so/go-sdk#WatchSendEventsParams">WatchSendEventsParams</a>) (<a href="https://pkg.go.dev/github.com/prelude-so/go-sdk">prelude</a>.<a href="https://pkg.go.dev/github.com/prelude-so/go-sdk#WatchSendEventsResponse">WatchSendEventsResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /v2/watch/feedback">client.Watch.<a href="https://pkg.go.dev/github.com/prelude-so/go-sdk#WatchService.SendFeedbacks">SendFeedbacks</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/prelude-so/go-sdk">prelude</a>.<a href="https://pkg.go.dev/github.com/prelude-so/go-sdk#WatchSendFeedbacksParams">WatchSendFeedbacksParams</a>) (<a href="https://pkg.go.dev/github.com/prelude-so/go-sdk">prelude</a>.<a href="https://pkg.go.dev/github.com/prelude-so/go-sdk#WatchSendFeedbacksResponse">WatchSendFeedbacksResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
