// @ts-ignore
/* eslint-disable */
import {request} from 'umi';

export async function login(body: API.LoginParams, options?: { [key: string]: any }) {
  return request<API.Default<API.LoginResult>>('/v1/login', {
    method : 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    data   : body,
    ...(options || {}),
  });
}

export async function logout(options?: { [key: string]: any }) {
  return request<Record<string, any>>('/v1/logout', {
    method: 'POST',
    ...(options || {}),
  });
}

export async function currentUser() {
  return request<API.Default<API.CurrentUser>>('/v1/user', {
    method : 'GET',
    headers: {
      'Content-Type': 'application/json',
    },
  });
}
