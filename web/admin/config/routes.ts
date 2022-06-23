export default [
  { path: '/', redirect: '/welcome' },
  {
    path: '/user',
    layout: false,
    routes: [
      { path: '/user', routes: [{ name: '登录', path: '/user/login', component: './user/Login' }] },
      { component: './404' },
    ],
  },
  { path: '/welcome', name: '欢迎', icon: 'smile', component: './Welcome' },
  {
    name: '菜谱管理',
    icon: 'OrderedListOutlined',
    routes: [
      { path: '/recipe/index', name: '首页', icon: 'smile', component: './recipe/Index' },
      { path: '/recipe/category', name: '分类', icon: 'smile', component: './recipe/Category' },
      { path: '/recipe/list', name: '菜谱', icon: 'smile', component: './recipe/List' },
      ]
  },

  { component: './404' },
];
