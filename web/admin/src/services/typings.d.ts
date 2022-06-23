// @ts-ignore
/* eslint-disable */

declare namespace API {
  type Default<T> = {
    msg: string;
    code: number;
    data: T;
  };

  type DefaultList<T> = {
    msg: string;
    code: number;
    data: {
      total: number,
      items: T[]
    };
  };

  type Upload = {
    url: string;
    presigned_url: string;
  };

  type CurrentUser = {
    id?: string;
    username?: string;
    avatar?: string;
  };

  type LoginResult = {
    token?: sting;
  };
  type SelectItem = {
    label?: string;
    value?: any;
  }

  type RecipeIndex = {
    carousel?: string[];
  };

  type RecipeCategoryItem = {
    id?: number;
    parent_id?: number;
    name?: string;
    sort?: number;
    status?: number;
    creator?: string;
    modifier?: string;
    create_time?: number;
    update_time?: number;
  };

  type RecipeItem = {
    id?: number;
    category_id?: number;
    title?: string;
    image?: string;
    content?: string;
    weight?: number;
    favorites?: number;
    sort?: number;
    status?: number;
    creator?: string;
    modifier?: string;
    create_time?: number;
    update_time?: number;
  };

  type PageParams = {
    page?: number;
    size?: number;
  };


  type RuleListItem = {
    key?: number;
    disabled?: boolean;
    href?: string;
    avatar?: string;
    name?: string;
    owner?: string;
    desc?: string;
    callNo?: number;
    status?: number;
    updatedAt?: string;
    createdAt?: string;
    progress?: number;
  };

  type RuleList = {
    data?: RuleListItem[];
    /** 列表的内容总数 */
    total?: number;
    success?: boolean;
  };

  type FakeCaptcha = {
    code?: number;
    status?: string;
  };

  type LoginParams = {
    username?: string;
    password?: string;
    autoLogin?: boolean;
    type?: string;
  };

  type ErrorResponse = {
    /** 业务约定的错误码 */
    errorCode: string;
    /** 业务上的错误信息 */
    errorMessage?: string;
    /** 业务上的请求是否成功 */
    success?: boolean;
  };

  type NoticeIconList = {
    data?: NoticeIconItem[];
    /** 列表的内容总数 */
    total?: number;
    success?: boolean;
  };

  type NoticeIconItemType = 'notification' | 'message' | 'event';

  type NoticeIconItem = {
    id?: string;
    extra?: string;
    key?: string;
    read?: boolean;
    avatar?: string;
    title?: string;
    status?: string;
    datetime?: string;
    description?: string;
    type?: NoticeIconItemType;
  };
}
