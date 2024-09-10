const local: App.I18n.Schema = {
  system: {
    title: '防火墙管理系统',
    navTitle: '防火墙管理系统',
    updateTitle: '系统版本更新通知',
    updateContent: '检测到系统有新版本发布，是否立即刷新页面？',
    updateConfirm: '立即刷新',
    updateCancel: '稍后再说',
    loginTitle: '开箱即用',
    loginContent: '基于 nftables 用于替代 firewall-cmd 繁琐操作的解决方案',
    welcome: '欢迎使用'
  },
  common: {
    action: '操作',
    add: '新增',
    error: '错误',
    addSuccess: '添加成功',
    backToHome: '返回首页',
    batchDelete: '批量删除',
    cancel: '取消',
    close: '关闭',
    check: '勾选',
    expandColumn: '展开列',
    columnSetting: '列设置',
    config: '配置',
    confirm: '确认',
    delete: '删除',
    deleteSuccess: '删除成功',
    confirmDelete: '确认删除吗？',
    edit: '编辑',
    index: '序号',
    keywordSearch: '请输入关键词搜索',
    logout: '退出登录',
    logoutConfirm: '确认退出登录吗？',
    lookForward: '敬请期待',
    modify: '修改',
    modifySuccess: '修改成功',
    noData: '无数据',
    operate: '操作',
    pleaseCheckValue: '请检查输入的值是否合法',
    refresh: '刷新',
    reset: '重置',
    search: '搜索',
    switch: '切换',
    tip: '提示',
    trigger: '触发',
    update: '更新',
    updateSuccess: '更新成功',
    userCenter: '个人中心',
    yesOrNo: {
      yes: '是',
      no: '否'
    },
    dialog: {
      info: '信息',
      success: '成功',
      warning: '警告',
      error: '错误',
      failed: '失败'
    },
    dataLoading: '数据加载中...',
    copySuccess: '复制成功',
    copy: '复制'
  },
  request: {
    logout: '请求失败后登出用户',
    logoutMsg: '用户状态失效，请重新登录',
    logoutWithModal: '请求失败后弹出模态框再登出用户',
    logoutWithModalMsg: '用户状态失效，请重新登录',
    refreshToken: '请求的token已过期，刷新token',
    tokenExpired: 'token已过期',
    '500': '服务器500异常，请稍后联系技术人员确认'
  },
  theme: {
    themeSchema: {
      title: '主题模式',
      light: '亮色模式',
      dark: '暗黑模式',
      auto: '跟随系统'
    },
    grayscale: '灰色模式',
    colourWeakness: '色弱模式',
    layoutMode: {
      title: '布局模式',
      vertical: '左侧菜单模式',
      'vertical-mix': '左侧菜单混合模式',
      horizontal: '顶部菜单模式',
      'horizontal-mix': '顶部菜单混合模式',
      reverseHorizontalMix: '一级菜单与子级菜单位置反转'
    },
    recommendColor: '应用推荐算法的颜色',
    recommendColorDesc: '推荐颜色的算法参照',
    themeColor: {
      title: '主题颜色',
      primary: '主色',
      info: '信息色',
      success: '成功色',
      warning: '警告色',
      error: '错误色',
      followPrimary: '跟随主色'
    },
    scrollMode: {
      title: '滚动模式',
      wrapper: '外层滚动',
      content: '主体滚动'
    },
    page: {
      animate: '页面切换动画',
      mode: {
        title: '页面切换动画类型',
        'fade-slide': '滑动',
        fade: '淡入淡出',
        'fade-bottom': '底部消退',
        'fade-scale': '缩放消退',
        'zoom-fade': '渐变',
        'zoom-out': '闪现',
        none: '无'
      }
    },
    fixedHeaderAndTab: '固定头部和标签栏',
    header: {
      height: '头部高度',
      breadcrumb: {
        visible: '显示面包屑',
        showIcon: '显示面包屑图标'
      }
    },
    tab: {
      visible: '显示标签栏',
      cache: '缓存标签页',
      height: '标签栏高度',
      mode: {
        title: '标签栏风格',
        chrome: '谷歌风格',
        button: '按钮风格'
      }
    },
    sider: {
      inverted: '深色侧边栏',
      width: '侧边栏宽度',
      collapsedWidth: '侧边栏折叠宽度',
      mixWidth: '混合布局侧边栏宽度',
      mixCollapsedWidth: '混合布局侧边栏折叠宽度',
      mixChildMenuWidth: '混合布局子菜单宽度'
    },
    footer: {
      visible: '显示底部',
      fixed: '固定底部',
      height: '底部高度',
      right: '底部局右'
    },
    watermark: {
      visible: '显示全屏水印',
      text: '水印文本'
    },
    themeDrawerTitle: '主题配置',
    pageFunTitle: '页面功能',
    configOperation: {
      copyConfig: '复制配置',
      copySuccessMsg: '复制成功，请替换 src/theme/settings.ts 中的变量 themeSettings',
      resetConfig: '重置配置',
      resetSuccessMsg: '重置成功'
    }
  },
  route: {
    login: '登录',
    403: '无权限',
    404: '页面不存在',
    500: '服务器错误',
    406: '无授权',
    'iframe-page': '外链页面',
    home: '首页',

    'user-center': '个人中心',

    system: '系统管理',
    system_shell: 'Shell',
    system_basic: '基本设置',
    audit: '审计日志',
    audit_login: '登录日志',
    audit_settings: '配置日志',
    policy: '本地策略',
    policy_input: '入站策略',
    'policy_input-limit': '入站流控',
    policy_output: '出站策略',
    'policy_output-limit': '出站流控',
    route: '路由策略',
    route_snat: '源地址转换',
    route_dnat: '目的地址转换',
    route_filter: '转发策略',
    route_limit: '转发流控'
  },
  page: {
    login: {
      common: {
        loginOrRegister: '登录 / 注册',
        userNamePlaceholder: '请输入用户名',
        phonePlaceholder: '请输入手机号',
        codePlaceholder: '请输入验证码',
        passwordPlaceholder: '请输入密码',
        confirmPasswordPlaceholder: '请再次输入密码',
        codeLogin: '验证码登录',
        confirm: '确定',
        back: '返回',
        validateSuccess: '验证成功',
        loginSuccess: '登录成功',
        welcomeBack: '欢迎回来，{userName} ！'
      },
      pwdLogin: {
        title: '密码登录',
        rememberMe: '记住我',
        forgetPassword: '忘记密码？',
        register: '注册账号',
        otherAccountLogin: '其他账号登录',
        otherLoginMode: '其他登录方式',
        superAdmin: '超级管理员',
        admin: '管理员',
        user: '普通用户'
      },
      codeLogin: {
        title: '验证码登录',
        getCode: '获取验证码',
        reGetCode: '{time}秒后重新获取',
        sendCodeSuccess: '验证码发送成功',
        imageCodePlaceholder: '请输入图片验证码'
      },
      register: {
        title: '注册账号',
        agreement: '我已经仔细阅读并接受',
        protocol: '《用户协议》',
        policy: '《隐私权政策》'
      },
      resetPwd: {
        title: '重置密码'
      },
      bindWeChat: {
        title: '绑定微信'
      },
      bindCode: {
        scanningBinding: '扫码绑定',
        manualInput: '手动输入',
        tip: '请在客户端中手动输入',
        completeBinding: '完成绑定'
      }
    },
    about: {
      title: '关于',
      introduction: `firewall-web 是一款用于设置Linux防火墙的服务器，基于nftables`,
      projectInfo: {
        title: '项目信息',
        version: '版本',
        latestBuildTime: '最新构建时间',
        githubLink: 'Github 地址',
        previewLink: '预览地址'
      },
      prdDep: '生产依赖',
      devDep: '开发依赖'
    },

    home: {
      document_project: '项目文档',
      document_project_link: '项目文档(外链)',
      greeting: '早安，{userName}, 今天又是充满活力的一天!',
      weatherDesc: '今日多云转晴，20℃ - 25℃!',
      projectCount: '项目数',
      todo: '待办',
      message: '消息',
      downloadCount: '下载量',
      registerCount: '注册量',
      schedule: '作息安排',
      study: '学习',
      work: '工作',
      rest: '休息',
      entertainment: '娱乐',
      visitCount: '访问量',
      turnover: '成交额',
      dealCount: '成交量',
      projectNews: {
        title: '项目动态',
        moreNews: '更多动态',
        desc1: 'Soybean 在2021年5月28日创建了开源项目 soybean-admin!',
        desc2: 'Yanbowe 向 soybean-admin 提交了一个bug，多标签栏不会自适应。',
        desc3: 'Soybean 准备为 soybean-admin 的发布做充分的准备工作!',
        desc4: 'Soybean 正在忙于为soybean-admin写项目说明文档！',
        desc5: 'Soybean 刚才把工作台页面随便写了一些，凑合能看了！'
      },
      creativity: '创意',
      userCount: '用户数',
      assetCount: '资产数',
      accountCount: '账号数',
      online: '在线用户',
      licenseNum: '授权数',
      licenseUsageRate: '授权使用率',
      onlineTrend: '在线趋势',
      today: '今日',
      yesterday: '昨日',
      usedLicense: '已用授权',
      notusedLicense: '未用授权'
    },
    function: {
      tab: {
        tabOperate: {
          title: '标签页操作',
          addTab: '添加标签页',
          addTabDesc: '跳转到关于页面',
          closeTab: '关闭标签页',
          closeCurrentTab: '关闭当前标签页',
          closeAboutTab: '关闭"关于"标签页',
          addMultiTab: '添加多标签页',
          addMultiTabDesc1: '跳转到多标签页页面',
          addMultiTabDesc2: '跳转到多标签页页面(带有查询参数)'
        },
        tabTitle: {
          title: '标签页标题',
          changeTitle: '修改标题',
          change: '修改',
          resetTitle: '重置标题',
          reset: '重置'
        }
      },
      multiTab: {
        routeParam: '路由参数',
        backTab: '返回 function_tab'
      },
      toggleAuth: {
        toggleAccount: '切换账号',
        authHook: '权限钩子函数 `hasAuth`',
        superAdminVisible: '超级管理员可见',
        adminVisible: '管理员可见',
        adminOrUserVisible: '管理员和用户可见'
      },
      request: {
        repeatedErrorOccurOnce: '重复请求错误只出现一次',
        repeatedError: '重复请求错误',
        repeatedErrorMsg1: '自定义请求错误 1',
        repeatedErrorMsg2: '自定义请求错误 2'
      }
    },
    manage: {
      common: {
        status: {
          enable: '启用',
          disable: '禁用'
        }
      },
      role: {
        title: '角色列表',
        roleName: '角色名称',
        roleCode: '角色编码',
        roleStatus: '角色状态',
        roleDesc: '角色描述',
        menuAuth: '菜单权限',
        buttonAuth: '按钮权限',
        form: {
          roleName: '请输入角色名称',
          roleCode: '请输入角色编码',
          roleStatus: '请选择角色状态',
          roleDesc: '请输入角色描述'
        },
        addRole: '新增角色',
        editRole: '编辑角色'
      },
      user: {
        title: '用户列表',
        userName: '用户名',
        userGender: '性别',
        nickName: '昵称',
        userPhone: '手机号',
        userEmail: '邮箱',
        userStatus: '用户状态',
        userRole: '用户角色',
        form: {
          userName: '请输入用户名',
          userGender: '请选择性别',
          nickName: '请输入昵称',
          userPhone: '请输入手机号',
          userEmail: '请输入邮箱',
          userStatus: '请选择用户状态',
          userRole: '请选择用户角色'
        },
        addUser: '新增用户',
        editUser: '编辑用户',
        gender: {
          male: '男',
          female: '女'
        }
      },
      menu: {
        home: '首页',
        title: '菜单列表',
        id: 'ID',
        parentId: '父级菜单ID',
        menuType: '菜单类型',
        menuName: '菜单名称',
        routeName: '路由名称',
        routePath: '路由路径',
        pathParam: '路径参数',
        layout: '布局',
        page: '页面组件',
        i18nKey: '国际化key',
        icon: '图标',
        localIcon: '本地图标',
        iconTypeTitle: '图标类型',
        order: '排序',
        constant: '常量路由',
        keepAlive: '缓存路由',
        href: '外链',
        hideInMenu: '隐藏菜单',
        activeMenu: '高亮的菜单',
        multiTab: '支持多页签',
        fixedIndexInTab: '固定在页签中的序号',
        query: '路由参数',
        button: '按钮',
        buttonCode: '按钮编码',
        buttonDesc: '按钮描述',
        menuStatus: '菜单状态',
        form: {
          home: '请选择首页',
          menuType: '请选择菜单类型',
          menuName: '请输入菜单名称',
          routeName: '请输入路由名称',
          routePath: '请输入路由路径',
          pathParam: '请输入路径参数',
          page: '请选择页面组件',
          layout: '请选择布局组件',
          i18nKey: '请输入国际化key',
          icon: '请输入图标',
          localIcon: '请选择本地图标',
          order: '请输入排序',
          keepAlive: '请选择是否缓存路由',
          href: '请输入外链',
          hideInMenu: '请选择是否隐藏菜单',
          activeMenu: '请选择高亮的菜单的路由名称',
          multiTab: '请选择是否支持多标签',
          fixedInTab: '请选择是否固定在页签中',
          fixedIndexInTab: '请输入固定在页签中的序号',
          queryKey: '请输入路由参数Key',
          queryValue: '请输入路由参数Value',
          button: '请选择是否按钮',
          buttonCode: '请输入按钮编码',
          buttonDesc: '请输入按钮描述',
          menuStatus: '请选择菜单状态',
          port: '请输入端口号'
        },
        addMenu: '新增菜单',
        editMenu: '编辑菜单',
        addChildMenu: '新增子菜单',
        type: {
          directory: '目录',
          menu: '菜单'
        },
        iconType: {
          iconify: 'iconify图标',
          local: '本地图标'
        }
      }
    },
    userSetting: {
      username: '用户名称',
      password: '密码',
      oldPassword: '旧密码',
      newPassword: '新密码',
      confirmPassword: '确认密码',
      email: '邮箱地址',
      phone: '手机号码',
      submit: '提交',
      basicSetting: '基本设置',
      securitySetting: '安全设置',
      userInfoSetting: '个人账户信息设置',
      passwordSetting: '密码设置'
    },

    basic: {
      sessionSettings: '会话配置',
      emailSettings: '邮件配置',
      messageSettings: '短信配置',
      fileServer: '文件服务',
      webTimeout: 'web超时时间(分钟)',
      nativePassword: '本地密码',
      disableLogin: '禁止登录',
      notify: '通知',
      test: '测试',
      sendTest: '发送邮件测试',
      emailServer: '邮件地址',
      emailPort: '邮件端口',
      sendEmail: '发件邮箱',
      protocol: '连接协议',
      emailAccount: '发件账号',
      emailPassword: '发件密码',
      parameterDescription: '参数说明',
      messageDescription: '接收号码:{mobile} 短信内容:{content} 有效期:{validity}  验证码:{code}',
      builtIn: '内置',
      parameterType: '参数类型',
      requestMethod: '请求方式',
      requestUrl: '请求地址',
      requestParameters: '请求参数',
      messageContent: '短信内容',
      fileserver1: '文件服务器一',
      fileserver2: '文件服务器二',
      addr: '连接地址',
      workDir: '工作目录',
      loginSetting: '登录设置',
      totpTip: '注意: 手机令牌功能需确保服务器时间尽量接近国际标准时间。',
      totpDesc: '手机误差时间: 1-30之间,每增加1，则允许前后时间误差增加30秒',
      emailCodeValidity: '邮件验证码有效期(分钟)',
      messageCodeValidity: '短信验证码有效期(分钟)',
      totpIssue: '手机令牌标识',
      totpOffset: '手机误差时间',
      mix: '双因素认证',
      first: '一步认证',
      second: '二步认证',
      native: '本地认证',
      mobile: '手机令牌',
      email: '邮件认证',
      message: '短信认证',
      emailService: '邮件服务',
      emailBackup: '邮件备份',
      noLimit: '无限制'
    },
    audit: {
      loginStatus: '登录状态',
      startTime: '开始时间',
      endTime: '结束时间',
      online: '在线',
      offline: '离线',
      loginResult: '登录结果',
      logOff: '强制下线',
      logOffTip: '确定要强制下线吗？',
      opration: '操作类型',
      method: '请求方法',
      path: '请求路径',
      operateResult: '操作结果',
      requestBody: '请求内容',
      responseBody: '响应内容'
    },

    shell: {
      total: '',
      upload: '已上传',
      download: '已下载',
      progress: '进度',
      speed: '平均速度',
      size: '文件大小',
      usedTime: '已用时间',
      downloadSuccess: '下载成功',
      skip: '跳过',
      wsError: '未知错误，可能是登录失效，请刷新页面重试!'
    },
    firewallPolicy: {
      list: '策略列表',
      policy: '策略',
      comment: '备注',
      action: '动作',
      sourceIp: '来源 IP',
      port: '端口',
      protocol: '协议',
      allIp: '所有IP',
      partialIp: '指定IP',
      portValidationFailure: '端口格式错误，请检查！',
      ipValidationFailure: 'IP格式错误，请检查！',
      portTip: '多个端口英文逗号分隔，如: 80,8080-8090',
      ipTip: '多个IP英文逗号分隔，如:192.168.1.1,192.168.2.0/24,192.168.1.2-192.168.1.10',
      accept: '允许',
      reject: '拒绝',
      drop: '丢弃',
      position: '位置',
      start: '首部',
      end: '尾部',
      beforePosition: '之前',
      afterPosition: '之后',
      option: '选项',
      pingReply: 'Ping 回复',
      pingRequest: 'Ping 请求',
      newTcp: '新连接',
      establishedTcp: '已建立连接',
      relatedTcp: '相关连接',
      untrackedTcp: '未追踪连接',
      invalidTcp: '无法识别连接',
      defaultPolicy: '默认策略',
      all: '全部',
      speed: '速度',
      allEthernet: '所有网口',
      sourceEthernet: '入网口',
      destinationEthernet: '出网口',
      dynamicIp: '动态IP',
      destIp: '目的IP',
      nat: '地址转换',
      destPort: '目的端口',
      sourcePort: '源端口',
      natPort: '转换端口',
      portType: '端口类型'
    }
  },
  form: {
    required: '不能为空',
    userName: {
      required: '请输入用户名',
      invalid: '用户名格式不正确'
    },
    phone: {
      required: '请输入手机号',
      invalid: '手机号格式不正确'
    },
    pwd: {
      required: '请输入密码',
      invalid: '密码格式不正确，6-18位字符，包含字母、数字、下划线'
    },

    confirmPwd: {
      required: '请输入确认密码',
      invalid: '两次输入密码不一致'
    },
    code: {
      required: '请输入验证码',
      invalid: '验证码格式不正确'
    },
    email: {
      required: '请输入邮箱',
      invalid: '邮箱格式不正确'
    },
    userPwd1: {
      required: '请输入密码',
      invalid: '密码长度不小于{length}位'
    },
    userPwd2: {
      required: '密码长度不小于{length}位',
      invalid: '密码至少包含数字、字母 两类'
    },
    userPwd3: {
      required: '密码长度不小于{length}位',
      invalid: '密码至少包含数字、字母、特殊字符 三类'
    },
    userPwd4: {
      required: '密码长度不小于{length}位',
      invalid: '密码必须包含数字、大写字母、小写字母、特殊字符 四类'
    },
    departmentName: {
      required: '请输入部门名称',
      invalid: '部门名称不能为空'
    },
    name: {
      required: '请输入名称',
      invalid: '名称不能为空'
    },
    level: {
      required: '请输入等级',
      invalid: '请输入等级范围 {range}'
    },
    loginname: {
      required: '请输入登录名称',
      invalid: '在登录时使用，不允许包含中文字符、空格、加号(+)、冒号(:)、斜杠(/)，最大长度为200'
    },
    accountValidTime: {
      required: '请输入账号有效期',
      invalid: '结束时间不能小于开始时间'
    },
    authentication: {
      required: '请输入认证方式',
      invalid: '认证方式不能为空'
    },
    webTimeout: {
      required: '不能为空',
      invalid: '取值范围: {range}'
    },
    url: {
      required: '不能为空',
      invalid: '请输入正确的URL地址'
    },
    range: {
      required: '不能为空',
      invalid: '取值范围: {range}'
    },
    key: '变量名',
    value: '变量值'
  },
  dropdown: {
    closeCurrent: '关闭',
    closeOther: '关闭其它',
    closeLeft: '关闭左侧',
    closeRight: '关闭右侧',
    closeAll: '关闭所有'
  },
  icon: {
    themeConfig: '主题配置',
    themeSchema: '主题模式',
    lang: '切换语言',
    fullscreen: '全屏',
    fullscreenExit: '退出全屏',
    reload: '刷新页面',
    collapse: '折叠菜单',
    expand: '展开菜单',
    pin: '固定',
    unpin: '取消固定'
  },
  datatable: {
    itemCount: '共 {total} 条',
    onlineTime: '在线时长',
    page: '页',
    name: '名称',
    username: '用户名称',
    loginname: '登录名称',
    email: '邮件地址',
    phone: '手机号码',
    department: '部门',
    role: '角色',
    roleName: '角色名称',
    status: '状态',
    operation: '操作',
    action: '操作',
    description: '描述',
    createTime: '创建时间',
    updateTime: '更新时间',
    createUser: '创建用户',
    updateUser: '更新用户',
    level: '等级',
    builtIn: '内置',
    account: '账号',
    password: '密码',
    confirmPassword: '确认密码',
    code: '验证码',
    permission: '权限',
    menu: '菜单',
    authentication: '认证',
    lastLoginTime: '最后登录时间',
    locked: '锁定',
    invalidAccount: '无效账号',
    passwordExpired: '密码过期',
    normal: '正常',
    lastChangePasswordTime: '上次改密时间',
    all: '全部',
    valid: '有效',
    passwordValid: '密码有效期',
    accountValid: '账号有效期',
    globalConfig: '全局配置',
    passwordLength: '密码长度',
    passwordComplexity: '密码强度',
    expiredAction: '过期处理',
    differChecked: '相同检查',
    date: {
      year: '年',
      month: '月',
      day: '日',
      week: '周',
      hour: '时',
      minute: '分',
      second: '秒'
    },
    bindTOTP: '重新绑定手机令牌',
    prot: '服务端口',
    loginMethod: '登录方式',
    secretKey: '登录密钥',
    secretCipher: '密钥密码',
    updatePassword: '更新密码',
    loginTime: '登录时间',
    logoutTime: '登出时间'
  }
};

export default local;
