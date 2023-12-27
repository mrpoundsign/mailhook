*mailhook* is a simple email to webhook server.

Only tested with Discord.

## Usage

Defaults to looking for `mailhook.yml` in `/etc/mailhook.yml`

```bash
mailhook
```

You can specify a `mailhook.yml`` in the command line.
```bash
mailhook -c mailhook.yml
```

## Configuration

Only a single user/password is supported.

Sample `mailhook.yml`:

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