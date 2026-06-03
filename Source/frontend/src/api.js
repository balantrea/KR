const API_URL = 'http://localhost:8080';

export const api = {
  async request(endpoint, method = 'GET', body = null) {
    const token = localStorage.getItem('token');
    const options = {
      method,
      headers: {
        'Content-Type': 'application/json',
      },
    };
    if (token) {
      options.headers.Authorization = `Bearer ${token}`;
    }
    if (body) {
      options.body = JSON.stringify(body);
    }
    const response = await fetch(`${API_URL}${endpoint}`, options);
    if (!response.ok) {
      const text = await response.text();
      throw new Error(text);
    }
    if (response.status === 204) return null;
    return response.json().catch(() => null);
  },

  async register(username, password) {
    return this.request('/register', 'POST', {
      username,
      password,
    });
  },

  login: async (username, password) => {
    const res = await fetch(`${API_URL}/login`, {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json'
      },
      body: JSON.stringify({
        username,
        password
      })
    });
    if (!res.ok) {
      throw new Error('Неверный логин или пароль');
    }
    const data = await res.json();
    localStorage.setItem('token', data.token);
    return data.token;
  },

  async getSports() {
    return this.request('/sports', 'GET');
  },

  async createSport(name, type) {
    return this.request('/sports', 'POST', { name, type });
  },

  async updateSport(id, name, type) {
    return this.request(`/sports/${id}`, 'PUT', { name, type });
  },

  async deleteSport(id) {
    return this.request(`/sports/${id}`, 'DELETE');
  },

  async updateUsername(username) {
    return this.request('/profile/username', 'PUT', {
      username,
    });
  },

  async deleteAccount() {
    return this.request('/profile', 'DELETE');
  },

  logout() {
    localStorage.removeItem('token');
  },
};