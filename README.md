# 🦇 &nbsp;Gothic

![GitHub Workflow Status](https://img.shields.io/github/workflow/status/jrapoport/gothic/test?style=flat-square) [![Go Report Card](https://goreportcard.com/badge/github.com/jrapoport/gothic?style=flat-square&)](https://goreportcard.com/report/github.com/jrapoport/gothic) ![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/jrapoport/gothic?style=flat-square) [![GitHub](https://img.shields.io/github/license/jrapoport/gothic?style=flat-square)](https://github.com/jrapoport/gothic/blob/master/LICENSE)

[![Buy Me A Coffee](https://img.shields.io/badge/buy%20me%20a%20coffee-☕-6F4E37?style=flat-square)](https://www.buymeacoffee.com/jrapoport)

Gothic is a user registration and authentication microservice written in Go.
It's based on OAuth2 and JWT and will handle user signup, authentication and
custom user data.

## Project History

This project was originally forked from 
[Netlify's GoTrue](https://github.com/netlify/gotrue).

The purpose was to adopt newer, more developer friendly technologies like
[Gorm](https://gorm.io/), [gRPC](https://grpc.io/), and [gRPC
Web](https://github.com/grpc/grpc-web); newer versions of critical libraries
like [JWT v4](https://github.com/dgrijalva/jwt-go); and migrate away from older
libraries that are deprecated with [security
flaws](https://github.com/gobuffalo/uuid).

These changes allow for advances like self-contained database migration,
expanded database driver support (e.g., PostgreSQL), and gRPC support. Broadly
speaking, they are intended to make it easier to modify and use the microservice
outside of Netlify tool chain, and in a more active development environment.

While the Netlify team did a good job with GoTrue, their use in production means
they cannot easily adopt these kinds of significant changes. In many cases, they
will likely never make them given the impacts to their tooling, deployment, and
production systems — which makes perfect sense for their situation.

I'd like to thank Netlify team for their hard work on the original version of
this microservice.

### NOTE: gRPC / gRPC Web support is in its infancy.

Currently only the scaffolding is in place with a few supported methods like
healthcheck and settings. The intention is to refactor the API over time to
support both REST & gRPC.

## Configuration

You may configure Gothic using either a configuration file named `.env`,
environment variables, or a combination of both. Environment variables are
prefixed with `GOTHIC_`, and will always have precedence over values provided
via file.

### Top-Level

```properties
GOTHIC_SITE_URL=https://example.gothic.com/
```

`GOTHIC_SITE_URL` - `string` **required**

The base URL your site is located at. Currently used in combination with other
settings to construct URLs used in emails.

`GOTHIC_DISABLE_SIGNUP` - `bool`

When signup is disabled the only way to create new users is through invites.
Defaults to `false`, all signups enabled.

`GOTHIC_RATE_LIMIT` - `string`

Header on which to rate limit the `/token` endpoint.

### API

```properties
GOTHIC__HOST=localhost
GOTHIC_REST_PORT=9999 # the http REST server port
GOTHIC_RPC_PORT=3001 # the gRPC server port
GOTHIC_RPCWEB_PORT=6001 # the gRPC Web server port
```

`GOTHIC_HOST` - `string`

Hostname to listen on.

`GOTHIC_REST_PORT` - `number`

Port number for the HTTP REST API server to listen on. Defaults to `8081`.

`GOTHIC_RPC_PORT` - `number`

Port number for the gRPC API server to listen on. Defaults to `3001`.

`GOTHIC_RPCWEB_PORT` - `number`

Port number for the gRPC Web API server to listen on. Defaults to `6001`.

`GORTHIC_REQUEST_ID` - `string`

If you wish to inherit a request ID from the incoming request, specify the name
in this value.

### Database

```properties
GOTHIC_DB_DRIVER=mysql
GOTHIC_DB_URL=root@localhost
GOTHIC_DB_NAME=gothic
```

`GOTHIC_DB_DRIVER` - `string` **required**

Chooses what dialect of database you want. Must be `mysql`.

`GOTHIC_DB_URL` `string` **required**

Connection string for the database.

`GOTHIC_DB_NAME` - `string`

The name of the database to create automatically. Defaults to `gothic`.

`GOTHIC_DB_NAMESPACE` - `string`

Adds a prefix to all table names.

**Migrations Note**

Migrations *WILL BE* applied automatically if you start with 
```
GOTHIC_DB_AUTOMIGRATE=true
```

### Logging

```properties
GOTHIC_LOG_LEVEL=debug # available without GOTHIC prefix (exception)
GOTHIC_LOG_FILE=/var/log/go/gothic.log
```

`GOTHIC_LOG_LEVEL` - `string`

Controls what log levels are output. Choose from `panic`, `fatal`, `error`,
`warn`, `info`, or `debug`. Defaults to `info`.

`GOTHIC_LOG_FILE` - `string`

If you wish logs to be written to a file, set `log_file` to a valid file path.

### Opentracing
Currently, only the Datadog tracer is supported.

```properties
GOTHIC_TRACING_ENABLED=true
GOTHIC_TRACING_HOST=127.0.0.1
GOTHIC_TRACING_PORT=8126
GOTHIC_TRACING_TAGS="tag1:value1,tag2:value2"
GOTHIC_SERVICE_NAME="gothic"
```

`GOTHIC_TRACING_ENABLED` - `bool`

Whether tracing is enabled or not. Defaults to `false`.

`GOTHIC_TRACING_HOST` - `bool`

The tracing destination.

`GOTHIC_TRACING_PORT` - `bool`

The port for the tracing host.

`GOTHIC_TRACING_TAGS` - `string`

A comma separated list of key:value pairs. These key value pairs will be added
as tags to all opentracing spans.

`GOTHIC_TRACING_SERVICE_NAME` - `string`

The name to use for the service.

### JSON Web Tokens (JWT)

```properties
GOTHIC_JWT_SECRET=supersecretvalue
GOTHIC_JWT_METHOD=HS256
GOTHIC_JWT_EXP=3600
GOTHIC_JWT_AUD=gothic
```

`GOTHIC_JWT_SECRET` - `string` **required**

The secret used to sign JWT tokens with.

`GOTHIC_JWT_METHOD` - `string`

The method used to sign JWT tokens. Defaults to `HS256` (HMAC256).

`GOTHIC_JWT_EXP` - `number`

How long tokens are valid for, in seconds. Defaults to 3600 (1 hour).

`GOTHIC_JWT_AUD` - `string`

The default JWT audience. Use audiences to group users.

`GOTHIC_JWT_ADMIN_GROUP` - `string`

The name of the admin group (if enabled). Defaults to `admin`.

`GOTHIC_JWT_DEFAULT_GROUP` - `string`

The default group to assign all new users to.

### External Authentication Providers

We support `bitbucket`, `github`, `gitlab`, and `google` for external
authentication. Use the names as the keys underneath `external` to configure
each separately.

```properties
GOTHIC_EXTERNAL_GITHUB_CLIENT_ID=myappclientid
GOTHIC_EXTERNAL_GITHUB_SECRET=clientsecretvaluessssh
```

No external providers are required, but you must provide the required values if
you choose to enable any.

`GOTHIC_EXTERNAL_[PROVIDER]_ENABLED` - `bool`

Whether this external provider is enabled or not

`GOTHIC_EXTERNAL_[PROVIDER]_CLIENT_ID` - `string` **required**

The OAuth2 Client ID registered with the external provider.

`GOTHIC_EXTERNAL_[PROVIDER]_SECRET` - `string` **required**

The OAuth2 Client Secret provided by the external provider when you registered.

`GOTHIC_EXTERNAL_[PROVIDER]_REDIRECT_URI` - `string` **required for gitlab**

The URI a OAuth2 provider will redirect to with the `code` and `state` values.

`GOTHIC_EXTERNAL_[PROVIDER]_URL` - `string`

The base URL used for constructing the URLs to request authorization and access
tokens. Used by `gitlab` only. Defaults to `https://gitlab.com`.

### E-Mail

Sending email is not required, but highly recommended for password recovery.
If enabled, you must provide the required values below.

```properties
GOTHIC_SMTP_HOST=smtp.mandrillapp.com
GOTHIC_SMTP_PORT=587
GOTHIC_SMTP_USER=smtp-delivery@example.com
GOTHIC_SMTP_PASS=correcthorsebatterystaple
GOTHIC_SMTP_ADMIN_EMAIL=support@example.com
GOTHIC_MAILER_SUBJECTS_CONFIRMATION="Please confirm"
```

`GOTHIC_SMTP_ADMIN_EMAIL` - `string` **required**

The `From` email address for all emails sent.

`GOTHIC_SMTP_HOST` - `string` **required**

The mail server hostname to send emails through.

`GOTHIC_SMTP_PORT` - `number` **required**

The port number to connect to the mail server on.

`GOTHIC_SMTP_USER` - `string`

If the mail server requires authentication, the username to use.

`GOTHIC_SMTP_PASS` - `string`

If the mail server requires authentication, the password to use.

`GOTHIC_SMTP_MAX_FREQUENCY` - `number`

Controls the minimum amount of time that must pass before sending another signup
confirmation or password reset email. The value is the number of seconds.
Defaults to 900 (15 minutes).

`GOTHIC_MAILER_AUTOCONFIRM` - `bool`

If you do not require email confirmation, you may set this to `true`. Defaults
to `false`.

`GOTHIC_MAILER_URLPATHS_INVITE` - `string`

URL path to use in the user invite email. Defaults to `/`.

`GOTHIC_MAILER_URLPATHS_CONFIRMATION` - `string`

URL path to use in the signup confirmation email. Defaults to `/`.

`GOTHIC_MAILER_URLPATHS_RECOVERY` - `string`

URL path to use in the password reset email. Defaults to `/`.

`GOTHIC_MAILER_URLPATHS_EMAIL_CHANGE` - `string`

URL path to use in the email change confirmation email. Defaults to `/`.

`GOTHIC_MAILER_SUBJECTS_INVITE` - `string`

Email subject to use for user invite. Defaults to `You have been invited`.

`GOTHIC_MAILER_SUBJECTS_CONFIRMATION` - `string`

Email subject to use for signup confirmation. Defaults to `Confirm Your Signup`.

`GOTHIC_MAILER_SUBJECTS_RECOVERY` - `string`

Email subject to use for password reset. Defaults to `Reset Your Password`.

`GOTHIC_MAILER_SUBJECTS_EMAIL_CHANGE` - `string`

Email subject to use for email change confirmation. Defaults to `Confirm Email
Change`.

`GOTHIC_MAILER_TEMPLATES_INVITE` - `string`

URL path to an email template to use when inviting a user.
`SiteURL`, `Email`, and `ConfirmationURL` variables are available.

Default Content (if template is unavailable):

```html
<h2>You have been invited</h2>

<p>You have been invited to create a user on {{ .SiteURL }}. Follow this link to
accept the invite:</p> <p><a href="{{ .ConfirmationURL }}">Accept the
invite</a></p> ```

`GOTHIC_MAILER_TEMPLATES_CONFIRMATION` - `string`

URL path to an email template to use when confirming a signup.
`SiteURL`, `Email`, and `ConfirmationURL` variables are available.

Default Content (if template is unavailable):

```html
<h2>Confirm your signup</h2>

<p>Follow this link to confirm your user:</p>
<p><a href="{{ .ConfirmationURL }}">Confirm your mail</a></p>
```

`MAILER_TEMPLATES_RECOVERY` - `string`

URL path to an email template to use when resetting a password.
`SiteURL`, `Email`, and `ConfirmationURL` variables are available.

Default Content (if template is unavailable):

```html
<h2>Reset Password</h2>

<p>Follow this link to reset the password for your user:</p>
<p><a href="{{ .ConfirmationURL }}">Reset Password</a></p>
```

`GOTHIC_MAILER_TEMPLATES_EMAIL_CHANGE` - `string`

URL path to an email template to use when confirming the change of an email
address. `SiteURL`, `Email`, `NewEmail`, and `ConfirmationURL` variables are
available.

Default Content (if template is unavailable):

```html
<h2>Confirm Change of Email</h2>

<p>Follow this link to confirm the update of your email from {{ .Email }} to {{
.NewEmail }}:</p> <p><a href="{{ .ConfirmationURL }}">Change Email</a></p> ```

`GOTHIC_WEBHOOK_URL` - `string`

Url of the webhook receiver endpoint. This will be called when events like
`validate`, `signup` or `login` occur.

`GOTHIC_WEBHOOK_SECRET` - `string`

Shared secret to authorize webhook requests. This secret signs the [JSON Web
Signature](https://tools.ietf.org/html/draft-ietf-jose-json-web-signature-41) of
the request. You *should* use this to verify the integrity of the request.
Otherwise others can feed your webhook receiver with fake data.

`GOTHIC_WEBHOOK_RETRIES` - `number`

How often Gothic should try a failed hook.

`GOTHIC_WEBHOOK_TIMEOUT_SEC` - `number`

Time between retries (in seconds).

`GOTHIC_WEBHOOK_EVENTS` - `list`

Which events should trigger a webhook. You can provide a comma separated list.
For example to listen to all events, provide the values `validate,signup,login`.

## Endpoints

Gothic exposes the following endpoints:

* **GET /settings**

  Returns the publicly available settings for this gothic instance.

  ```json
  {
    "external": {
      "bitbucket": true,
      "github": true,
      "gitlab": true,
      "google": true
    },
    "disable_signup": false,
    "autoconfirm": false
  }
  ```

* **POST /signup**

  Register a new user with an email and password.

  ```json
  {
    "email": "email@example.com",
    "password": "secret"
  }
  ```

  Returns:

  ```json
  {
    "id": "11111111-2222-3333-4444-5555555555555",
    "email": "email@example.com",
    "confirmation_sent_at": "2016-05-15T20:49:40.882805774-07:00",
    "created_at": "2016-05-15T19:53:12.368652374-07:00",
    "updated_at": "2016-05-15T19:53:12.368652374-07:00"
  }
  ```

* **POST /invite**

  Invites a new user with an email.

  ```json
  {
    "email": "email@example.com"
  }
  ```

  Returns:

  ```json
  {
    "id": "11111111-2222-3333-4444-5555555555555",
    "email": "email@example.com",
    "confirmation_sent_at": "2016-05-15T20:49:40.882805774-07:00",
    "created_at": "2016-05-15T19:53:12.368652374-07:00",
    "updated_at": "2016-05-15T19:53:12.368652374-07:00",
    "invited_at": "2016-05-15T19:53:12.368652374-07:00"
  }
  ```

* **POST /verify**

  Verify a registration or a password recovery. Type can be `signup` or
  `recovery` and the `token` is a token returned from either `/signup` or
  `/recover`.

  ```json
  {
    "type": "signup",
    "token": "confirmation-code-delivered-in-email",
    "password": "12345abcdef"
  }
  ```

  `password` is required for signup verification if no existing password exists.

  Returns:

  ```json
  {
    "access_token": "jwt-token-representing-the-user",
    "token_type": "bearer",
    "expires_in": 3600,
    "refresh_token": "a-refresh-token"
  }
  ```

* **POST /recover**

  Password recovery. Will deliver a password recovery mail to the user based on
  email address.

  ```json
  {
    "email": "email@example.com"
  }
  ```

  Returns:

  ```json
  {}
  ```

* **POST /token**

  This is an OAuth2 endpoint that currently implements
  the password, refresh_token, and authorization_code grant types

  ```
  grant_type=password&username=email@example.com&password=secret
  ```

  or

  ```
  grant_type=refresh_token&refresh_token=my-refresh-token
  ```

  Once you have an access token, you can access the methods requiring
  authentication by settings the `Authorization: Bearer YOUR_ACCESS_TOKEN_HERE`
  header.

  Returns:

  ```json
  {
    "access_token": "jwt-token-representing-the-user",
    "token_type": "bearer",
    "expires_in": 3600,
    "refresh_token": "a-refresh-token"
  }
  ```

* **GET /user**

  Get the JSON object for the logged in user (requires authentication)

  Returns:

  ```json
  {
    "id": "11111111-2222-3333-4444-5555555555555",
    "email": "email@example.com",
    "confirmation_sent_at": "2016-05-15T20:49:40.882805774-07:00",
    "created_at": "2016-05-15T19:53:12.368652374-07:00",
    "updated_at": "2016-05-15T19:53:12.368652374-07:00"
  }
  ```

* **PUT /user**

  Update a user (Requires authentication). Apart from changing email/password,
  this method can be used to set custom user data.

  ```json
  {
    "email": "new-email@example.com",
    "password": "new-password",
    "data": {
      "key": "value",
      "number": 10,
      "admin": false
    }
  }
  ```

  Returns:

  ```json
  {
    "id": "11111111-2222-3333-4444-5555555555555",
    "email": "email@example.com",
    "confirmation_sent_at": "2016-05-15T20:49:40.882805774-07:00",
    "created_at": "2016-05-15T19:53:12.368652374-07:00",
    "updated_at": "2016-05-15T19:53:12.368652374-07:00"
  }
  ```

* **POST /logout**

  Logout a user (Requires authentication).

  This will revoke all refresh tokens for the user. Remember that the JWT tokens
  will still be valid for stateless auth until they expires.

## TODO

* Schema for custom user data in config file
