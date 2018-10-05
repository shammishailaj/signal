# Web tracker

## Usage

### Tracking snippet

```html
<script type="text/javascript" async src="https://[SIGNAL_API_HOST]/js?id=[PROJECT_TRACKING_ID]"></script>
```

### NPM

```bash
npm i -s @bloom42/signal
```

```javascript
import signal from '@bloom42/signal';

signal.init('[PROJECT_TRACKING_ID]', { endpoint: '[SIGNAL_API_HOST]' });
```


## Content Security Policy

If you use a [Content Security Policy (CSP)](https://developer.mozilla.org/en-US/docs/Web/HTTP/CSP)
to specify security policies for your website, Signal requires the following CSP directives
(replace yoursignal.com with the URL to your Signal instance):
```
script-src: yoursignalapi.com;
img-src: yoursignalapi.com;
```

