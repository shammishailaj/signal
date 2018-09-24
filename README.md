<p align="center">
  <h2 align="center">Signal</h2>
  <p align="center">Simple and beautiful Analytics</p>
</p>

---

Signal lets you automatically capture every user interaction from web, mobile, and cloud services: clicks, submits, transactions, emails, and much more. You can then analyze it all retroactively.

1. [Repo organization](#repo-organization)
2. [Usage](#usage)

---


## Repo organization

This repo is splitteed into the following components:

| Name | Description |
| ---- | ----------- |
| [api](api) | Signal's API which receive events and process requests |
| [docs](docs) | Contains all the high level documentation for the Signal platform |
| [web](web) | Signal's web tracking library |
| [www](www) | Signal's dashboard |





## Usage


### Tracking snippet

```html
<script type="text/javascript" async src="https://[SIGNAL_API_INSTANCE]/js?id=[PROJECT_TRACKING_ID]"></script>
```


### Content Security Policy

If you use a [Content Security Policy (CSP)](https://developer.mozilla.org/en-US/docs/Web/HTTP/CSP)
to specify security policies for your website, Signal requires the following CSP directives
(replace yoursignal.com with the URL to your Signal instance):
```
script-src: yoursignal.com;
connect-src: yoursignal.com;
img-src: yoursignal.com;
```
