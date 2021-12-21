![](https://www.sms77.io/wp-content/uploads/2019/07/sms77-Logo-400x79.png "sms77 Logo")

# sms77 Plugin for Drone.io CI

Send SMS or make text-to-speech calls.

## Prerequisites

- An API Key from sms77 - can be created in your [developer dashboard](https://app.sms77.io/developer)
- [Drone CI](https://www.drone.io/)

## Usage

```yml
- name: sms77
  image: sms77io/drone
  settings:
    api_key: my_sms77_api_key
    debug: 0
    flash: 1
    from: Drone CI
    text: Drone CI has finished!
    to: +491771783130
    type: sms
```

Required settings:

* api_key
* text
* to

Optional settings:

* debug: if enabled, API validates inputs but messages don't get sent
* flash: send as flash SMS (SMS only)
* from: sender identifier up to 11 alphanumeric (SMS only) or 16 numeric characters - must be a verified number or a [shared number](https://www.sms77.io/en/docs/glossary/shared-numbers/)
* type: must be either "voice" or "sms" (defaults to "sms")

## Support

Need help? Feel free to [contact us](https://www.sms77.io/en/company/contact/).

[![MIT](https://img.shields.io/badge/License-MIT-teal.svg)](LICENSE)