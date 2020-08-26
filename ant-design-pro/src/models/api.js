import createModel from "@/lib/utils/createModel";
import request from '@/utils/request';
import { func } from "prop-types";

async function createApi(params = {}) {
  return request(`${params.url}`, {
    method: 'POST',
    body: params.body,
  });
}

async function readApi(params = {}) {
  return request(`${params.url}`, {
    method: 'GET',
    params: {
      ...params.query
    },
  });
}

async function updateApi(params = {}) {
  return request(`${params.url}`, {
    method: 'PUT',
    body: params.body,
  });
}

async function destroyApi(params = {}) {
  return request(`${params.url}?${params.query}`, {
    method: 'DELETE',
  });
}

function requestEffect(apiMethod) {
  return function* effect({ payload = {} }, { call }) {
    return yield call(apiMethod, payload);
  };
}

export default {
  namespace: 'api',
  state: {},

  effects: {
    _get: requestEffect(readApi),
    _post: requestEffect(createApi),
    _put: requestEffect(updateApi),
    _delete: requestEffect(destroyApi),
  },
}
