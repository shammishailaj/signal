import axios from 'axios';
import store from 'store';

import router from '@/router';

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

  async get(url, options) {
    let res = null;

    try {
      res = await this._client.get(url, options);
    } catch (err) {
      if (this._token && err.response.status === 401) {
        this.logout();
      } else {
        throw err;
      }
    }
    return res;
  }

  async post(url, data, options) {
    let res = null;

    try {
      res = await this._client.post(url, data, options);
    } catch (err) {
      if (this._token && err.response.status === 401) {
        this.logout();
      } else {
        throw err;
      }
    }
    return res;
  }

  async delete(url, options) {
    let res = null;

    try {
      res = await this._client.delete(url, options);
    } catch (err) {
      if (this._token && err.response.status === 401) {
        this.logout();
      } else {
        throw err;
      }
    }
    return res;
  }

  logout() {
    store.remove('signal_token');
    this._token = null;
    router.push('/login');
  }
}

export default new Api();
