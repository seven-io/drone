<p align="center">
  <img src="https://www.seven.io/wp-content/uploads/Logo.svg" width="250" alt="seven logo" />
</p>

<h1 align="center">seven SMS &amp; Voice for Drone CI</h1>

<p align="center">
  Send SMS or place text-to-speech calls from <a href="https://www.drone.io/">Drone CI</a> pipelines via the seven gateway.
</p>

<p align="center">
  <a href="LICENSE"><img src="https://img.shields.io/badge/License-MIT-teal.svg" alt="MIT License" /></a>
  <img src="https://img.shields.io/badge/Drone-CI-fc4a1a" alt="Drone CI" />
  <img src="https://img.shields.io/badge/Docker-image-2496ed" alt="Docker image" />
</p>

---

## Features

- **Drone Plugin** - Drop-in pipeline step using the `seven-io/drone` Docker image
- **SMS or Voice** - Switch between message types via the `type` setting
- **Flash SMS** - Optional flash messages that bypass the inbox

## Prerequisites

- A [Drone CI](https://www.drone.io/) installation
- A [seven account](https://www.seven.io/) with API key ([How to get your API key](https://help.seven.io/en/developer/where-do-i-find-my-api-key))

## Usage

```yaml
- name: seven
  image: seven-io/drone
  settings:
    api_key: my_seven_api_key
    flash:   1
    from:    Drone CI
    text:    Drone CI has finished!
    to:      +491716992343
    type:    sms
```

### Required settings

| Setting | Description |
|---------|-------------|
| `api_key` | Your seven API key |
| `text` | Message body |
| `to` | Recipient phone number |

### Optional settings

| Setting | Description |
|---------|-------------|
| `from` | Sender ID. Up to 11 alphanumeric (SMS) or 16 numeric characters. Must be a verified number or a [shared number](https://www.seven.io/en/docs/glossary/shared-numbers/) |
| `flash` | Send as flash SMS (SMS only) |
| `type` | `sms` (default) or `voice` |

## Support

Need help? Feel free to [contact us](https://www.seven.io/en/company/contact/) or [open an issue](https://github.com/seven-io/drone/issues).

## License

[MIT](LICENSE)
