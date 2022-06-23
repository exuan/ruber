import {UploadRequestOption as RcCustomRequestOptions} from "rc-upload/lib/interface";
import {uploadUrl, upload} from "@/services/api";
import {message} from 'antd';

export function tree(data: any, pid = 0, key = 'parent_id') {
  const result = [];
  // eslint-disable-next-line no-restricted-syntax
  for (const i in data) {
    if (data[i][key] === pid) {
      const temp = data[i];
      const children = tree(data, data[i].id, key);
      if (children.length) {
        temp.children = children;
      }
      result.push(temp);
    }
  }

  return result;
}

export const imageRequest  = async (options: RcCustomRequestOptions) => {
  const {file, onSuccess} = options;
  if (typeof file !== 'object') {
    return "";
  }
  let ext: string | undefined = '';
  if ("name" in file) {
    ext = file?.name?.split('.').pop();
  }
  if (!ext) {
    message.error('图片格式错误');
    return "";
  }

  const res = await uploadUrl(ext)
  if (!res || res.code !== 0) {
    message.error(res?.msg || '请求错误');
    return "";
  }

  try {
    const result = await upload(res.data.presigned_url, file);
    if (onSuccess) {
      onSuccess(res.data.url, result);
    }
  } catch (e) {
    message.error(e ? e.message : '请求错误')
    return "";
  }

  return res.data.url
}


