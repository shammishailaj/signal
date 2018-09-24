<p align="center">
  <h2 align="center">Signal</h2>
  <p align="center">Web and Mobile Analytics</p>
</p>

---

Signal lets you automatically capture every user interaction from web, mobile, and cloud services: clicks, submits, transactions, emails, and much more. You can then analyze it all retroactively.

1. [Repo organization](#repo-organization)
2. [Tracking snippet](#tracking-snippet)

---


## Repo organization

This repo is splitteed into the following components:

| Name | Description |
| ---- | ----------- |
| [api](api) | Signal's API which receive events and process requests |
| [docs](docs) | Contains all the high level documentation for the Signal platform |
| [web](web) | Signal's web tracking library |
| [www](www) | Signal's dashboard |



## Tracking snippet
```html
<script type="text/javascript" async src="https://[SIGNAL_API_INSTANCE]/js?id=[PROJECT_TRACKING_ID]"></script>
```
