import store from 'store';

import api from '@/services/api';
import router from '@/router';

const TOKEN_KEY = 'signal_token';

class Auth {
  constructor() {
    this._token = store.get(TOKEN_KEY);
    api.token = this._token;
  }

  async login(username, password) {
    const data = {
      username,
      password,
    };
    let res = null;

    try {
      res = await api.post('/api/v1/login', data);
      this._token = res.data.data.token;
      store.set(TOKEN_KEY, this._token);
      api.token = this._token;
    } catch (err) {
      if (err.response && err.response.data && err.response.data.error) {
        throw new Error(err.response.data.error);
      }
      throw err;
    }
  }

  logout() {
    store.remove(TOKEN_KEY);
    this._token = null;
    api.token = null;
    router.push('/login');
  }

  is_authenticated() {
    if (this._token) {
      return true;
    }
    return false;
  }
}

export default new Auth();
