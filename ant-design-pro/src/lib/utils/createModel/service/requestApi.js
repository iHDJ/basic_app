import { stringify } from 'qs';
import superRequest from '@/utils/request';

function getMetaContent(name) {
  const meta = document.head.querySelector(`meta[name='${name}']`);

  if (!meta) {
    return null;
  }

  return meta.getAttribute('content');
}

function request(url, options = {}) {
  const newOptions = {
    ...options,
    credentials: 'same-origin',
    headers: {}
  };

  const { method } = newOptions;
  if (method === 'POST' || method === 'PUT' || method === 'DELETE') {
    newOptions.headers = {
      'X-CSRF-Token': getMetaContent('csrf-token'),

      ...newOptions.headers,
    };
  }

  const accessToken = localStorage.getItem("access_token")

  if (accessToken) {
    newOptions.headers["AccessToken"] = accessToken
  }

  return superRequest(url, newOptions);
}

export async function createApi(params = {}) {
  return request(`${params.url}`, {
    method: 'POST',
    body: params.body,
  });
}

export async function readApi(params = {}) {
  return request(`${params.url}?${stringify(params.query, { arrayFormat: 'brackets' })}`);
}

export async function updateApi(params = {}) {
  return request(`${params.url}`, {
    method: 'PUT',
    body: params.body,
  });
}

export async function destroyApi(params = {}) {
  return request(`${params.url}?${stringify(params.query, { arrayFormat: 'brackets' })}`, {
    method: 'DELETE',
  });
}
