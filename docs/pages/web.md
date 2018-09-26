# Web tracker

## Usage

### Tracking snippet

```html
<script type="text/javascript" async src="https://[SIGNAL_API_HOST]/js?id=[PROJECT_TRACKING_ID]"></script>
```

### NPM

```bash
npm i -s @astrocorp/signal
```

```javascript
import signal from '@astrocorp/signal';

signal.init('[PROJECT_TRACKING_ID]');
```


## Content Security Policy

If you use a [Content Security Policy (CSP)](https://developer.mozilla.org/en-US/docs/Web/HTTP/CSP)
to specify security policies for your website, Signal requires the following CSP directives
(replace yoursignal.com with the URL to your Signal instance):
```
script-src: yoursignalapi.com;
img-src: yoursignalapi.com;
```

