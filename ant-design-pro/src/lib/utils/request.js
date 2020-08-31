import superRequest from '@/utils/request';
import { csrfToken, requestHeaders } from './setting';

export default function request(url, options = {}) {
  let { headers } = options;

  const { method } = options;
  if (method === 'POST' || method === 'PUT' || method === 'DELETE') {
    headers = {
      ...headers,
      'X-CSRF-Token': csrfToken,
    };
  }

  return superRequest(url, {
    ...options,
    credentials: 'same-origin',
    headers: {
      ...headers,
      ...requestHeaders,
    },
  });
}
