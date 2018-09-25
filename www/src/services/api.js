import axios from 'axios';

class Api {
  constructor() {
    this._token = null;
    this._client = axios.create({ baseURL: process.env.VUE_APP_API_BASE_URL });
  }

  set token(token) {
    this._token = token;
    if (token) {
      this._client.defaults.headers.common.Authorization = this.getAuthorizationHeader();
    } else {
      this._client.defaults.headers.common.Authorization = undefined;
    }
  }

  get token() {
    return this._token;
  }

  getAuthorizationHeader() {
    const token = this._token;
    let header = null;

    if (token) {
      header = `Bearer ${token}`;
    }
    return header;
  }

  get(url, options) {
    return this._client.get(url, options);
  }

  post(url, data, options) {
    return this._client.post(url, data, options);
  }

  delete(url, options) {
    return this._client.delete(url, options);
  }
}

export default new Api();
