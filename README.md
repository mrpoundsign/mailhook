*mailhook* is a simple email to webhook server.

Only tested with Discord.

## Usage

Defaults to looking for `mailhook.yaml` in `/etc/mailhook.yaml`

```bash
mailhook
```

You can specify a `mailhook.yaml`` in the command line.
```bash
mailhook -c mailhook.yaml
```

## Configuration

Only a single user/password is supported.

Sample `mailhook.yaml`:

```yaml
port: 1025
host: 0.0.0.0
auth:
  username: username
  password: password
hooks:
  - name: One
    address: text@mailhook
    url: https://discord.com/api/webhooks/...
    html_markdown: false
  - name: two
    address: html@mailhook
    url: https://discord.com/api/webhooks/...
    html_markdown: true
```

If you need to use html emails, set `html_markdown: true`. Otherwise it will use the text portion of the email.

### Adding Headers to HTTP Requests

Headers can be added to the request using the `headers` key.
```yaml
hoooks:
  - name: One
      address: text@mailhook
      url: https://discord.com/api/webhooks/...
      html_markdown: false
      headers:
        - cache-control: none
          content-type: text/plain
```

## Building

You need to have `go` installed. If you don't have it, go to https://golang.org/doc/install

Now, get the `mailhook` repo.

```bash
git clone https://github.com/mrpoundsign/mailhook.git
cd mailhook
```

Do the go things.

```bash
go get
```

Now build it.

```bash
go build -o mailhook cmd/mailhook/main.go
```