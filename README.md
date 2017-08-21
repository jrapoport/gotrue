# GoTrue - User management for APIs

GoTrue is a small open-source API written in golang, that can act as a self-standing
API service for handling user registration and authentication for JAM projects.

It's based on OAuth2 and JWT and will handle user signup, authentication and custom
user data.

## Configuration

You may configure GoTrue using either a configuration file named `.env`,
environment variables, or a combination of both. Environment variables are prefixed with `GOTRUE_`, and will always have precedence over values provided via file.

### Top-Level

```
GOTRUE_SITE_URL=https://example.netlify.com/
```

`SITE_URL` - `string` **required**

The base URL your site is located at. Currently used in combination with other settings to construct URLs used in emails.

`OPERATOR_TOKEN` - `string` *Multi-instance mode only*

The shared secret with an operator (usually Netlify) for this microservice. Used to verify requests have been proxied through the operator and
the payload values can be trusted.

### API

```
GOTRUE_API_HOST=localhost
PORT=9999
```

`API_HOST` - `string`

Hostname to listen on.

`PORT` (no prefix) / `API_PORT` - `number`

Port number to listen on. Defaults to `8081`.

`API_ENDPOINT` - `string` *Multi-instance mode only*

Controls what endpoint Netlify can access this API on.

### Database

```
GOTRUE_DB_DRIVER=sqlite3
DATABASE_URL=gotrue.db
```

`DB_DRIVER` - `string` **required**

Chooses what dialect of database you want. Choose from `mongo`, `sqlite3`, `mysql`, or `postgres`.

`DATABASE_URL` (no prefix) / `DB_DATABASE_URL` - `string` **required**

Connection string for the database. See the [gorm examples](https://github.com/jinzhu/gorm/blob/gh-pages/documents/database.md) for more details.

`DB_NAMESPACE` - `string`

Adds a prefix to all table names.

`DB_AUTOMIGRATE` - `bool`

If enabled, creates missing tables and columns upon startup.

### Logging

```
LOG_LEVEL=debug
```

`LOG_LEVEL` - `string`

Controls what log levels are output. Choose from `panic`, `fatal`, `error`, `warn`, `info`, or `debug`. Defaults to `info`.

`LOG_FILE` - `string`

If you wish logs to be written to a file, set `log_file` to a valid file path.

### JSON Web Tokens (JWT)

```
GOTRUE_JWT_SECRET=supersecretvalue
GOTRUE_JWT_EXP=3600
GOTRUE_JWT_AUD=netlify
```

`JWT_SECRET` - `string` **required**

The secret used to sign JWT tokens with.

`JWT_EXP` - `number`

How long tokens are valid for, in seconds. Defaults to 3600 (1 hour).

`JWT_AUD` - `string`

The default JWT audience. Use audiences to group users.

`JWT_ADMIN_GROUP_NAME` - `string`

The name of the admin group (if enabled). Defaults to `admin`.

`JWT_ADMIN_GROUP_DISABLED` - `bool`

The first user created will be made an admin.
Set to `true` to turn off this behavior.
Defaults to `false`.

`JWT_DEFAULT_GROUP_NAME` - `string`

The default group to assign all new users to.

### External Authentication Providers

We support `github`, `gitlab`, and `bitbucket` for external authentication.
Use the names as the keys underneath `external` to configure each separately.

```
GOTRUE_EXTERNAL_GITHUB_CLIENT_ID=myappclientid
GOTRUE_EXTERNAL_GITHUB_SECRET=clientsecretvaluessssh
```

No external providers are required, but you must provide the required values if you choose to enable any.

`EXTERNAL_X_CLIENT_ID` - `string` **required**

The OAuth2 Client ID registered with the external provider.

`EXTERNAL_X_SECRET` - `string` **required**

The OAuth2 Client Secret provided by the external provider when you registered.

`EXTERNAL_X_REDIRECT_URI` - `string` **required for gitlab**

The URI a OAuth2 provider will redirect to with the `code` and `state` values.

`EXTERNAL_X_URL` - `string`

The base URL used for constructing the URLs to request authorization and access tokens. Used by `gitlab` only. Defaults to `https://gitlab.com`.

### E-Mail

Sending email is not required, but highly recommended for password recovery.
If enabled, you must provide the required values below.

```
GOTRUE_MAILER_ADMIN_EMAIL=support@example.com
GOTRUE_MAILER_PORT=25
GOTRUE_MAILER_SUBJECTS_CONFIRMATION="Please confirm"
```

`MAILER_ADMIN_EMAIL` - `string` **required**

The `From` email address for all emails sent.

`MAILER_HOST` - `string` **required**

The mail server hostname to send emails through.

`MAILER_PORT` - `number` **required**

The port number to connect to the mail server on.

`MAILER_USER` - `string`

If the mail server requires authentication, the username to use.

`MAILER_PASS` - `string`

If the mail server requires authentication, the password to use.

`MAILER_MAX_FREQUENCY` - `number`

Controls the minimum amount of time that must pass before sending another signup confirmation or password reset email. The value is the number of seconds. Defaults to 900 (15 minutes).

`MAILER_AUTOCONFIRM` - `bool`

If you do not require email confirmation, you may set this to `true`. Defaults to `false`.

`MAILER_MEMBER_FOLDER` - `string`

The folder on `site_url` where the `confirm`, `recover`, and `confirm-email`
pages are located. This is used in combination with `site_url` to generate the 
URLs used in the emails. Defaults to `/member`.

`MAILER_SUBJECTS_INVITE` - `string`

Email subject to use for user invite. Defaults to `You have been invited`.

`MAILER_SUBJECTS_CONFIRMATION` - `string`

Email subject to use for signup confirmation. Defaults to `Confirm Your Signup`.

`MAILER_SUBJECTS_RECOVERY` - `string`

Email subject to use for password reset. Defaults to `Reset Your Password`.

`MAILER_SUBJECTS_EMAIL_CHANGE` - `string`

Email subject to use for email change confirmation. Defaults to `Confirm Email Change`.

`MAILER_TEMPLATES_INVITE` - `string`

URL path to an email template to use when inviting a user. Defaults to `/.netlify/gotrue/templates/invite.html`
`SiteURL`, `Email`, and `ConfirmationURL` variables are available.

Default Content (if template is unavailable):
```html
<h2>You have been invited</h2>

<p>You have been invited to create a user on {{ .SiteURL }}. Follow this link to accept the invite:</p>
<p><a href="{{ .ConfirmationURL }}">Accept the invite</a></p>
```

`MAILER_TEMPLATES_CONFIRMATION` - `string`

URL path to an email template to use when confirming a signup. Defaults to `/.netlify/gotrue/templates/confirm.html`
`SiteURL`, `Email`, and `ConfirmationURL` variables are available.

Default Content (if template is unavailable):
```html
<h2>Confirm your signup</h2>

<p>Follow this link to confirm your account:</p>
<p><a href="{{ .ConfirmationURL }}">Confirm your mail</a></p>
```

`MAILER_TEMPLATES_RECOVERY` - `string`

URL path to an email template to use when resetting a password. Defaults to `/.netlify/gotrue/templates/recover.html`
`SiteURL`, `Email`, and `ConfirmationURL` variables are available.

Default Content (if template is unavailable):
```html
<h2>Reset Password</h2>

<p>Follow this link to reset the password for your account:</p>
<p><a href="{{ .ConfirmationURL }}">Reset Password</a></p>
```

`MAILER_TEMPLATES_EMAIL_CHANGE` - `string`

URL path to an email template to use when confirming the change of an email address. Defaults to `/.netlify/gotrue/templates/email-change.html`
`SiteURL`, `Email`, `NewEmail`, and `ConfirmationURL` variables are available.

Default Content (if template is unavailable):
```html
<h2>Confirm Change of Email</h2>

<p>Follow this link to confirm the update of your email from {{ .Email }} to {{ .NewEmail }}:</p>
<p><a href="{{ .ConfirmationURL }}">Change Email</a></p>
```


## Endpoints

GoTrue exposes the following endpoints:

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
    "id":1,
    "email":"email@example.com",
    "confirmation_sent_at": "2016-05-15T20:49:40.882805774-07:00",
    "created_at": "2016-05-15T19:53:12.368652374-07:00",
    "updated_at": "2016-05-15T19:53:12.368652374-07:00"
  }
  ```

* **POST /verify**

  Verify a registration or a password recovery. Type can be `signup` or `recover`
  and the `token` is a token returned from either `/signup` or `/recover`.

  ```json
  {
    "type": "signup",
    "token": "confirmation-code-delivered-in-email"
  }
  ```

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
  the password and refresh_token grant types

  ```
  grant_type=password&username=email@example.com&password=secret
  ```

  or

  ```
  grant_type=refresh_token&refresh_token=my-refresh-token
  ```

  Once you have an access token, you can access the methods requiring authentication
  by settings the `Authorization: Bearer YOUR_ACCESS_TOKEN_HERE` header.

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
    "id":1,
    "email":"email@example.com",
    "confirmation_sent_at": "2016-05-15T20:49:40.882805774-07:00",
    "created_at": "2016-05-15T19:53:12.368652374-07:00",
    "updated_at": "2016-05-15T19:53:12.368652374-07:00"
  }
  ```

* **PUT /user**

  Update a user (Requires authentication). Apart from changing email/password, this
  method can be used to set custom user data.

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
    "id":1,
    "email":"email@example.com",
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
