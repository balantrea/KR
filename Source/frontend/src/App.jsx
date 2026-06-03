import React, { useState, useEffect } from 'react';
import { api } from './api';
import {
  LogOut,
  Plus,
  Trash2,
  Edit2,
  Shield,
  Activity,
  User
} from 'lucide-react';

export default function App() {
  const [token, setToken] = useState(() => localStorage.getItem('token'));
  const [isRegister, setIsRegister] = useState(false);
  const [username, setUsername] = useState('');
  const [password, setPassword] = useState('');
  const [error, setError] = useState('');
  const [sports, setSports] = useState([]);
  const [name, setName] = useState('');
  const [type, setType] = useState('');
  const [editingId, setEditingId] = useState(null);
  const [newUsername, setNewUsername] = useState('');

  useEffect(() => {
    if (token) loadSports();
  }, [token]);

  const loadSports = async () => {
    try {
      const data = await api.getSports();
      setSports(data || []);
    } catch (err) {
      console.error(err);
    }
  };

  const handleAuth = async (e) => {
    e.preventDefault();
    setError('');
    try {
      if (isRegister) {
        await api.register(username, password);
        setIsRegister(false);
        alert('Регистрация успешна! Теперь войдите.');
      } else {
        const userToken = await api.login(username, password);
        setToken(userToken);
      }
      setUsername('');
      setPassword('');
    } catch (err) {
      setError(err.message);
    }
  };

  const handleLogout = () => {
    api.logout();
    setToken(null);
    setSports([]);
  };

  const handleSubmitSport = async (e) => {
    e.preventDefault();
    if (!name || !type) return;
    try {
      if (editingId) {
        await api.updateSport(editingId, name, type);
        setEditingId(null);
      } else {
        await api.createSport(name, type);
      }
      setName('');
      setType('');
      loadSports();
    } catch (err) {
      alert(err.message);
    }
  };

  const handleEdit = (sport) => {
    setEditingId(sport.id);
    setName(sport.name);
    setType(sport.type);
  };

  const handleDelete = async (id) => {
    if (window.confirm('Удалить этот вид спорта?')) {
      try {
        await api.deleteSport(id);
        loadSports();
      } catch (err) {
        alert(err.message);
      }
    }
  };

  const handleChangeUsername = async () => {
    if (!newUsername) return;
    try {
      await api.updateUsername(newUsername);
      alert('Имя пользователя изменено');
      setNewUsername('');
    } catch (err) {
      alert(err.message);
    }
  };

  const handleDeleteAccount = async () => {
    const confirmed = window.confirm('Вы действительно хотите удалить аккаунт?');
    if (!confirmed) return;
    try {
      await api.deleteAccount();
      api.logout();
      setToken(null);
      alert('Аккаунт удален');
    } catch (err) {
      alert(err.message);
    }
  };

  if (!token) {
    return (
        <div style={styles.authContainer}>
          <div style={styles.authCard}>
            <div style={{ textAlign: 'center', marginBottom: 20 }}>
              <Shield size={48} color="#4f46e5" />
              <h2 style={{ marginTop: 10, color: '#111827' }}>
                {isRegister ? 'Регистрация' : 'Вход в систему'}
              </h2>
            </div>
            {error && (
                <div style={styles.error}>{error}</div>
            )}
            <form onSubmit={handleAuth} style={styles.form}>
              <input
                  style={styles.input}
                  type="text"
                  placeholder="Имя пользователя"
                  value={username}
                  onChange={(e) => setUsername(e.target.value)}
                  required
              />
              <input
                  style={styles.input}
                  type="password"
                  placeholder="Пароль"
                  value={password}
                  onChange={(e) => setPassword(e.target.value)}
                  required
              />
              <button style={styles.btnPrimary} type="submit">
                {isRegister ? 'Создать аккаунт' : 'Войти'}
              </button>
            </form>
            <p style={{ textAlign: 'center', marginTop: 15, fontSize: 14 }}>
            <span style={{ color: '#666' }}>
              {isRegister ? 'Уже есть аккаунт?' : 'Нет аккаунта?'}
            </span>{' '}
              <button
                  style={styles.btnLink}
                  onClick={() => {
                    setIsRegister(!isRegister);
                    setError('');
                  }}
              >
                {isRegister ? 'Войти' : 'Зарегистрироваться'}
              </button>
            </p>
          </div>
        </div>
    );
  }

  return (
      <div style={styles.dashboard}>
        <header style={styles.header}>
          <div style={{ display: 'flex', alignItems: 'center', gap: 10 }}>
            <Activity size={28} color="#4f46e5" />
            <h1 style={{ fontSize: 20, margin: 0, color: '#111827' }}>
              Спортивный Менеджер
            </h1>
          </div>
          <button style={styles.btnDanger} onClick={handleLogout}>
            <LogOut size={16} />
            Выйти
          </button>
        </header>
        <main style={styles.mainContent}>
          <div style={styles.card}>
            <div style={{ display: 'flex', alignItems: 'center', gap: 10, marginBottom: 20 }}>
              <User size={22} />
              <h3 style={{ margin: 0, color: '#111827' }}>Профиль</h3>
            </div>
            <div style={{ display: 'flex', gap: 10 }}>
              <input
                  style={styles.input}
                  type="text"
                  placeholder="Новое имя пользователя"
                  value={newUsername}
                  onChange={(e) => setNewUsername(e.target.value)}
              />
              <button style={styles.btnPrimary} onClick={handleChangeUsername}>
                Изменить имя
              </button>
            </div>
            <button style={{ ...styles.btnDanger, marginTop: 20 }} onClick={handleDeleteAccount}>
              Удалить аккаунт
            </button>
          </div>
          <div style={{ ...styles.card, marginTop: 20 }}>
            <h3 style={{ margin: '0 0 10px 0', color: '#111827' }}>
              {editingId ? 'Редактировать запись' : 'Добавить вид спорта'}
            </h3>
            <form onSubmit={handleSubmitSport} style={{ display: 'flex', gap: 10 }}>
              <input
                  style={styles.input}
                  type="text"
                  placeholder="Название"
                  value={name}
                  onChange={(e) => setName(e.target.value)}
                  required
              />
              <input
                  style={styles.input}
                  type="text"
                  placeholder="Тип"
                  value={type}
                  onChange={(e) => setType(e.target.value)}
                  required
              />
              <button style={styles.btnPrimary} type="submit">
                {editingId ? <Edit2 size={16} /> : <Plus size={16} />}
                {editingId ? 'Сохранить' : 'Добавить'}
              </button>
            </form>
          </div>
          <div style={{ ...styles.card, marginTop: 20 }}>
            <h3 style={{ margin: '0 0 5px 0', color: '#111827' }}>
              Ваши спортивные дисциплины
            </h3>
            <p style={{ color: '#6b7280', fontSize: 13, margin: '0 0 15px 0' }}>
              Другие пользователи не видят эти данные.
            </p>
            {sports.length === 0 ? (
                <p style={{ textAlign: 'center', padding: 20, color: '#9ca3af', margin: 0 }}>
                  У вас пока нет записей.
                </p>
            ) : (
                <table style={styles.table}>
                  <thead>
                  <tr style={{ background: '#f3f4f6' }}>
                    <th style={styles.th}>ID</th>
                    <th style={styles.th}>Название</th>
                    <th style={styles.th}>Тип</th>
                    <th style={styles.th}>Действия</th>
                  </tr>
                  </thead>
                  <tbody>
                  {sports.map((sport) => (
                      <tr key={sport.id} style={{ borderBottom: '1px solid #e5e7eb' }}>
                        <td style={styles.td}>{sport.id}</td>
                        <td style={styles.td}>
                          <strong style={{ color: '#111827' }}>{sport.name}</strong>
                        </td>
                        <td style={styles.td}>{sport.type}</td>
                        <td style={{ ...styles.td, display: 'flex', gap: 10 }}>
                          <button style={styles.btnIconEdit} onClick={() => handleEdit(sport)}>
                            <Edit2 size={14} />
                          </button>
                          <button style={styles.btnIconDelete} onClick={() => handleDelete(sport.id)}>
                            <Trash2 size={14} />
                          </button>
                        </td>
                      </tr>
                  ))}
                  </tbody>
                </table>
            )}
          </div>
        </main>
      </div>
  );
}

const styles = {
  authContainer: { display: 'flex', justifyContent: 'center', alignItems: 'center', height: '100vh', background: '#f3f4f6', fontFamily: 'sans-serif' },
  authCard: { background: '#fff', padding: 30, borderRadius: 8, boxShadow: '0 4px 6px -1px rgba(0,0,0,0.1)', width: 350 },
  form: { display: 'flex', flexDirection: 'column', gap: 15 },
  input: { padding: '10px', borderRadius: 6, border: '1px solid #d1d5db', fontSize: 14, width: '100%', boxSizing: 'border-box', background: '#fff', color: '#000' },
  btnPrimary: { display: 'flex', alignItems: 'center', justifyContent: 'center', gap: 5, padding: '10px', background: '#4f46e5', color: '#fff', border: 'none', borderRadius: 6, cursor: 'pointer', fontWeight: 'bold' },
  btnLink: { background: 'none', border: 'none', color: '#4f46e5', cursor: 'pointer', fontWeight: 'bold', padding: 0 },
  error: { color: '#ef4444', background: '#fee2e2', padding: 10, borderRadius: 6, marginBottom: 15, fontSize: 14, textAlign: 'center' },
  dashboard: { minHeight: '100vh', background: '#f9fafb', fontFamily: 'sans-serif' },
  header: { display: 'flex', justifyContent: 'space-between', alignItems: 'center', padding: '15px 30px', background: '#fff', borderBottom: '1px solid #e5e7eb' },
  btnDanger: { display: 'flex', alignItems: 'center', gap: 5, padding: '8px 12px', background: '#ef4444', color: '#fff', border: 'none', borderRadius: 6, cursor: 'pointer' },
  mainContent: { padding: '30px', maxWidth: 1000, margin: '0 auto' },
  card: { background: '#fff', padding: 20, borderRadius: 8, boxShadow: '0 1px 3px rgba(0,0,0,0.05)' },
  table: { width: '100%', borderCollapse: 'collapse' },
  th: { padding: 12, textAlign: 'left', borderBottom: '2px solid #e5e7eb', fontSize: 14, color: '#374151' },
  td: { padding: 12, fontSize: 14, color: '#4b5563' },
  btnIconEdit: { background: '#f3f4f6', border: 'none', padding: 6, borderRadius: 4, cursor: 'pointer', color: '#4b5563', display: 'flex', alignItems: 'center' },
  btnIconDelete: { background: '#fee2e2', border: 'none', padding: 6, borderRadius: 4, cursor: 'pointer', color: '#ef4444', display: 'flex', alignItems: 'center' }
};