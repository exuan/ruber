import type { Settings as LayoutSettings } from '@ant-design/pro-layout';
import { PageLoading } from '@ant-design/pro-layout';
import type { RequestConfig, RunTimeLayoutConfig } from 'umi';
import { history } from 'umi';
import { notification } from 'antd';
import RightContent from '@/components/RightContent';
import Footer from '@/components/Footer';
import { currentUser as queryCurrentUser } from '@/services/user';
import type { ResponseError ,RequestOptionsInit} from 'umi-request';

const isDev = process.env.NODE_ENV === 'development';
const loginPath = '/user/login';

/** 获取用户信息比较慢的时候会展示一个 loading */
export const initialStateConfig = {
  loading: <PageLoading />,
};

/**
 * @see  https://umijs.org/zh-CN/plugins/plugin-initial-state
 * */
export async function getInitialState(): Promise<{
  settings?: Partial<LayoutSettings>;
  currentUser?: API.CurrentUser;
  fetchUserInfo?: () => Promise<API.CurrentUser | undefined>;
}> {
  const fetchUserInfo = async () => {
    try {
      const res = await queryCurrentUser();
      if (!res|| res.code !== 0) {
        throw new Error(res?.msg || "请登录");
      }
      return res.data;
    } catch (error) {
      history.push(loginPath);
    }
    return undefined;
  };
  // 如果是登录页面，不执行
  if (history.location.pathname !== loginPath) {
    const currentUser = await fetchUserInfo();
    return {
      fetchUserInfo,
      currentUser,
      settings: {},
    };
  }
  return {
    fetchUserInfo,
    settings: {},
  };
}



const authHeaderInterceptor = (url: string, options: RequestOptionsInit) => {
const token = localStorage?.getItem('x-ruber-backend-token');

  let authHeader = {}
  if ( url.indexOf('v1') > -1) {
    authHeader = {'x-ruber-backend-token':token}
  }

  return {
    url: `${url}`,
    options: { ...options, interceptors: true, headers: authHeader },
  };
};

export const request: RequestConfig = {
  prefix: isDev ? '' : 'your-domain',
  requestInterceptors: [authHeaderInterceptor],
  credentials: 'include',
  errorHandler: (error: ResponseError) => {
    const { response } = error;
    //
    // if (response && response.status) {
    //   const { status, statusText, url } = response;
    //   const requestErrorMessage = messages['app.request.error'];
    //   const errorMessage = `${requestErrorMessage} ${status}: ${url}`;
    //   const errorDescription = messages[`app.request.${status}`] || statusText;
    //   notification.error({
    //     message: errorMessage,
    //     description: errorDescription,
    //   });
    // }

    if (!response) {
      notification.error({
        description: '您的网络发生异常，无法连接服务器',
        message: '网络异常',
      });
    }
    throw error;
  },
};


// ProLayout 支持的api https://procomponents.ant.design/components/layout
export const layout: RunTimeLayoutConfig = ({ initialState }) => {
  return {
    rightContentRender: () => <RightContent />,
    disableContentMargin: false,
    waterMarkProps: {
      content: initialState?.currentUser?.username,
    },
    footerRender: () => <Footer />,
    onPageChange: () => {
      const { location } = history;
      // 如果没有登录，重定向到 login
      if (!initialState?.currentUser && location.pathname !== loginPath) {
        history.push(loginPath);
      }
    },
    menuHeaderRender: undefined,
    // 自定义 403 页面
    // unAccessible: <div>unAccessible</div>,
    ...initialState?.settings,
  };
};
