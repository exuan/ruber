import {request} from 'umi';

export async function uploadUrl(ext: string| undefined, options?: Record<string, any>) {
  return request<API.Default<API.Upload>>('/v1/uploadUrl', {
    method: 'POST',
    data: {ext},
    ...(options || {}),
  });
}

export async function upload(url: string, body: any) {
  return request(url, {
    method: 'PUT',
    body,
  });
}
