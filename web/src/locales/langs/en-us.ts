const local: App.I18n.Schema = {
  system: {
    title: 'Firewalld',
    navTitle: 'Firewalld',
    updateTitle: 'System Version Update Notification',
    updateContent: 'A new version of the system has been detected. Do you want to refresh the page immediately?',
    updateConfirm: 'Refresh immediately',
    updateCancel: 'Later',
    loginTitle: 'Out of the box',
    loginContent: 'Nftable-based solutions used to replace the tedious operations of firewall-cmd',
    welcome: 'Welcome to '
  },
  common: {
    error: 'Error',
    action: 'Action',
    add: 'Add',
    addSuccess: 'Add Success',
    backToHome: 'Back to home',
    batchDelete: 'Batch Delete',
    cancel: 'Cancel',
    close: 'Close',
    check: 'Check',
    expandColumn: 'Expand Column',
    columnSetting: 'Column Setting',
    config: 'Config',
    confirm: 'Confirm',
    delete: 'Delete',
    deleteSuccess: 'Delete Success',
    confirmDelete: 'Are you sure you want to delete?',
    edit: 'Edit',
    index: 'Index',
    keywordSearch: 'Please enter keyword',
    logout: 'Logout',
    logoutConfirm: 'Are you sure you want to log out?',
    lookForward: 'Coming soon',
    modify: 'Modify',
    modifySuccess: 'Modify Success',
    noData: 'No Data',
    operate: 'Operate',
    pleaseCheckValue: 'Please check whether the value is valid',
    refresh: 'Refresh',
    reset: 'Reset',
    search: 'Search',
    switch: 'Switch',
    tip: 'Tip',
    trigger: 'Trigger',
    update: 'Update',
    updateSuccess: 'Update Success',
    userCenter: 'User Center',
    yesOrNo: {
      yes: 'Yes',
      no: 'No'
    },
    dialog: {
      info: 'Info',
      success: 'Success',
      warning: 'Warning',
      error: 'Error',
      failed: 'Failed'
    },
    dataLoading: 'Data Loading...',
    copySuccess: 'Copy Success',
    copy: 'Copy'
  },
  request: {
    logout: 'Logout user after request failed',
    logoutMsg: 'User status is invalid, please log in again',
    logoutWithModal: 'Pop up modal after request failed and then log out user',
    logoutWithModalMsg: 'User status is invalid, please log in again',
    refreshToken: 'The requested token has expired, refresh the token',
    tokenExpired: 'The requested token has expired',
    '500': 'The server is abnormal. Contact technical support later'
  },
  theme: {
    themeSchema: {
      title: 'Theme Schema',
      light: 'Light',
      dark: 'Dark',
      auto: 'Follow System'
    },
    grayscale: 'Grayscale',
    colourWeakness: 'Colour Weakness',
    layoutMode: {
      title: 'Layout Mode',
      vertical: 'Vertical Menu Mode',
      horizontal: 'Horizontal Menu Mode',
      'vertical-mix': 'Vertical Mix Menu Mode',
      'horizontal-mix': 'Horizontal Mix menu Mode',
      reverseHorizontalMix: 'Reverse first level menus and child level menus position'
    },
    recommendColor: 'Apply Recommended Color Algorithm',
    recommendColorDesc: 'The recommended color algorithm refers to',
    themeColor: {
      title: 'Theme Color',
      primary: 'Primary',
      info: 'Info',
      success: 'Success',
      warning: 'Warning',
      error: 'Error',
      followPrimary: 'Follow Primary'
    },
    scrollMode: {
      title: 'Scroll Mode',
      wrapper: 'Wrapper',
      content: 'Content'
    },
    page: {
      animate: 'Page Animate',
      mode: {
        title: 'Page Animate Mode',
        fade: 'Fade',
        'fade-slide': 'Slide',
        'fade-bottom': 'Fade Zoom',
        'fade-scale': 'Fade Scale',
        'zoom-fade': 'Zoom Fade',
        'zoom-out': 'Zoom Out',
        none: 'None'
      }
    },
    fixedHeaderAndTab: 'Fixed Header And Tab',
    header: {
      height: 'Header Height',
      breadcrumb: {
        visible: 'Breadcrumb Visible',
        showIcon: 'Breadcrumb Icon Visible'
      }
    },
    tab: {
      visible: 'Tab Visible',
      cache: 'Tab Cache',
      height: 'Tab Height',
      mode: {
        title: 'Tab Mode',
        chrome: 'Chrome',
        button: 'Button'
      }
    },
    sider: {
      inverted: 'Dark Sider',
      width: 'Sider Width',
      collapsedWidth: 'Sider Collapsed Width',
      mixWidth: 'Mix Sider Width',
      mixCollapsedWidth: 'Mix Sider Collapse Width',
      mixChildMenuWidth: 'Mix Child Menu Width'
    },
    footer: {
      visible: 'Footer Visible',
      fixed: 'Fixed Footer',
      height: 'Footer Height',
      right: 'Right Footer'
    },
    watermark: {
      visible: 'Watermark Full Screen Visible',
      text: 'Watermark Text'
    },
    themeDrawerTitle: 'Theme Configuration',
    pageFunTitle: 'Page Function',
    configOperation: {
      copyConfig: 'Copy Config',
      copySuccessMsg: 'Copy Success, Please replace the variable "themeSettings" in "src/theme/settings.ts"',
      resetConfig: 'Reset Config',
      resetSuccessMsg: 'Reset Success'
    }
  },
  route: {
    login: 'Login',
    403: 'No Permission',
    404: 'Page Not Found',
    500: 'Server Error',
    'iframe-page': 'Iframe',
    home: 'Home',
    'user-center': 'User Center',
    system: 'System Manage',
    system_basic: 'Basic Settings',
    system_shell: 'Shell',
    audit: 'Audit',
    audit_login: 'Login',
    audit_settings: 'Settings',
    406: 'No License',
    policy: 'Native Policy',
    policy_input: 'Input Policy',
    'policy_input-limit': 'Input Limit',
    policy_output: 'Output Policy',
    'policy_output-limit': 'Output Limit',
    route: 'Route Policy',
    route_snat: 'SNAT',
    route_dnat: 'DNAT',
    route_filter: 'Forward Policy'
  },
  page: {
    login: {
      common: {
        loginOrRegister: 'Login / Register',
        userNamePlaceholder: 'Please enter user name',
        phonePlaceholder: 'Please enter phone number',
        codePlaceholder: 'Please enter verification code',
        passwordPlaceholder: 'Please enter password',
        confirmPasswordPlaceholder: 'Please enter password again',
        codeLogin: 'Verification code login',
        confirm: 'Confirm',
        back: 'Back',
        validateSuccess: 'Verification passed',
        loginSuccess: 'Login successfully',
        welcomeBack: 'Welcome back, {userName} !'
      },
      pwdLogin: {
        title: 'Password Login',
        rememberMe: 'Remember me',
        forgetPassword: 'Forget password?',
        register: 'Register',
        otherAccountLogin: 'Other Account Login',
        otherLoginMode: 'Other Login Mode',
        superAdmin: 'Super Admin',
        admin: 'Admin',
        user: 'User'
      },
      codeLogin: {
        title: 'Verification Code Login',
        getCode: 'Get verification code',
        reGetCode: 'Reacquire after {time}s',
        sendCodeSuccess: 'Verification code sent successfully',
        imageCodePlaceholder: 'Please enter image verification code'
      },
      register: {
        title: 'Register',
        agreement: 'I have read and agree to',
        protocol: '《User Agreement》',
        policy: '《Privacy Policy》'
      },
      resetPwd: {
        title: 'Reset Password'
      },
      bindWeChat: {
        title: 'Bind WeChat'
      },
      bindCode: {
        scanningBinding: 'Scanning Binding',
        manualInput: 'Manual Input',
        tip: 'Enter a value manually on the client',
        completeBinding: 'Complete Binding'
      }
    },
    about: {
      title: 'About',
      introduction: `firewall-web is a server for setting up Linux firewalls, based on nftables`,
      projectInfo: {
        title: 'Project Info',
        version: 'Version',
        latestBuildTime: 'Latest Build Time',
        githubLink: 'Github Link',
        previewLink: 'Preview Link'
      },
      prdDep: 'Production Dependency',
      devDep: 'Development Dependency'
    },
    home: {
      document_project: 'Project Document',
      document_project_link: 'Project Document(External Link)',
      greeting: 'Good morning, {userName}, today is another day full of vitality!',
      weatherDesc: 'Today is cloudy to clear, 20℃ - 25℃!',
      projectCount: 'Project Count',
      todo: 'Todo',
      message: 'Message',
      downloadCount: 'Download Count',
      registerCount: 'Register Count',
      schedule: 'Work and rest Schedule',
      study: 'Study',
      work: 'Work',
      rest: 'Rest',
      entertainment: 'Entertainment',
      visitCount: 'Visit Count',
      turnover: 'Turnover',
      dealCount: 'Deal Count',
      projectNews: {
        title: 'Project News',
        moreNews: 'More News',
        desc1: 'Soybean created the open source project soybean-admin on May 28, 2021!',
        desc2: 'Yanbowe submitted a bug to soybean-admin, the multi-tab bar will not adapt.',
        desc3: 'Soybean is ready to do sufficient preparation for the release of soybean-admin!',
        desc4: 'Soybean is busy writing project documentation for soybean-admin!',
        desc5: 'Soybean just wrote some of the workbench pages casually, and it was enough to see!'
      },
      creativity: 'Creativity',
      userCount: 'Users',
      assetCount: 'Assets',
      accountCount: 'Accounts',
      online: 'Online',
      licenseNum: 'License',
      onlineTrend: 'Online Trend',
      licenseUsageRate: 'License Usage Rate',
      today: 'Today',
      yesterday: 'Yesterday',
      usedLicense: 'Used License',
      notusedLicense: 'Not Used License'
    },
    function: {
      tab: {
        tabOperate: {
          title: 'Tab Operation',
          addTab: 'Add Tab',
          addTabDesc: 'To about page',
          closeTab: 'Close Tab',
          closeCurrentTab: 'Close Current Tab',
          closeAboutTab: 'Close "About" Tab',
          addMultiTab: 'Add Multi Tab',
          addMultiTabDesc1: 'To MultiTab page',
          addMultiTabDesc2: 'To MultiTab page(with query params)'
        },
        tabTitle: {
          title: 'Tab Title',
          changeTitle: 'Change Title',
          change: 'Change',
          resetTitle: 'Reset Title',
          reset: 'Reset'
        }
      },
      multiTab: {
        routeParam: 'Route Param',
        backTab: 'Back function_tab'
      },
      toggleAuth: {
        toggleAccount: 'Toggle Account',
        authHook: 'Auth Hook Function `hasAuth`',
        superAdminVisible: 'Super Admin Visible',
        adminVisible: 'Admin Visible',
        adminOrUserVisible: 'Admin and User Visible'
      },
      request: {
        repeatedErrorOccurOnce: 'Repeated Request Error Occurs Once',
        repeatedError: 'Repeated Request Error',
        repeatedErrorMsg1: 'Custom Request Error 1',
        repeatedErrorMsg2: 'Custom Request Error 2'
      }
    },
    manage: {
      common: {
        status: {
          enable: 'Enable',
          disable: 'Disable'
        }
      },
      role: {
        title: 'Role List',
        roleName: 'Role Name',
        roleCode: 'Role Code',
        roleStatus: 'Role Status',
        roleDesc: 'Role Description',
        menuAuth: 'Menu Auth',
        buttonAuth: 'Button Auth',
        form: {
          roleName: 'Please enter role name',
          roleCode: 'Please enter role code',
          roleStatus: 'Please select role status',
          roleDesc: 'Please enter role description'
        },
        addRole: 'Add Role',
        editRole: 'Edit Role'
      },
      user: {
        title: 'User List',
        userName: 'User Name',
        userGender: 'Gender',
        nickName: 'Nick Name',
        userPhone: 'Phone Number',
        userEmail: 'Email',
        userStatus: 'User Status',
        userRole: 'User Role',
        form: {
          userName: 'Please enter user name',
          userGender: 'Please select gender',
          nickName: 'Please enter nick name',
          userPhone: 'Please enter phone number',
          userEmail: 'Please enter email',
          userStatus: 'Please select user status',
          userRole: 'Please select user role'
        },
        addUser: 'Add User',
        editUser: 'Edit User',
        gender: {
          male: 'Male',
          female: 'Female'
        }
      },
      menu: {
        home: 'Home',
        title: 'Menu List',
        id: 'ID',
        parentId: 'Parent ID',
        menuType: 'Menu Type',
        menuName: 'Menu Name',
        routeName: 'Route Name',
        routePath: 'Route Path',
        pathParam: 'Path Param',
        layout: 'Layout Component',
        page: 'Page Component',
        i18nKey: 'I18n Key',
        icon: 'Icon',
        localIcon: 'Local Icon',
        iconTypeTitle: 'Icon Type',
        order: 'Order',
        constant: 'Constant',
        keepAlive: 'Keep Alive',
        href: 'Href',
        hideInMenu: 'Hide In Menu',
        activeMenu: 'Active Menu',
        multiTab: 'Multi Tab',
        fixedIndexInTab: 'Fixed Index In Tab',
        query: 'Query Params',
        button: 'Button',
        buttonCode: 'Button Code',
        buttonDesc: 'Button Desc',
        menuStatus: 'Menu Status',
        form: {
          home: 'Please select home',
          menuType: 'Please select menu type',
          menuName: 'Please enter menu name',
          routeName: 'Please enter route name',
          routePath: 'Please enter route path',
          pathParam: 'Please enter path param',
          page: 'Please select page component',
          layout: 'Please select layout component',
          i18nKey: 'Please enter i18n key',
          icon: 'Please enter iconify name',
          localIcon: 'Please enter local icon name',
          order: 'Please enter order',
          keepAlive: 'Please select whether to cache route',
          href: 'Please enter href',
          hideInMenu: 'Please select whether to hide menu',
          activeMenu: 'Please select route name of the highlighted menu',
          multiTab: 'Please select whether to support multiple tabs',
          fixedInTab: 'Please select whether to fix in the tab',
          fixedIndexInTab: 'Please enter the index fixed in the tab',
          queryKey: 'Please enter route parameter Key',
          queryValue: 'Please enter route parameter Value',
          button: 'Please select whether it is a button',
          buttonCode: 'Please enter button code',
          buttonDesc: 'Please enter button description',
          menuStatus: 'Please select menu status',
          port: 'Please enter port'
        },
        addMenu: 'Add Menu',
        editMenu: 'Edit Menu',
        addChildMenu: 'Add Child Menu',
        type: {
          directory: 'Directory',
          menu: 'Menu'
        },
        iconType: {
          iconify: 'Iconify Icon',
          local: 'Local Icon'
        }
      }
    },
    userSetting: {
      username: 'username',
      password: 'password',
      oldPassword: 'Old password',
      newPassword: 'New password',
      confirmPassword: 'Confirm password',
      email: 'Email',
      phone: 'Phone',
      submit: 'submit',
      basicSetting: 'Basic Setting',
      securitySetting: 'Security Setting',
      userInfoSetting: 'Account information Settings',
      passwordSetting: 'Password Setting'
    },

    basic: {
      loginSetting: 'Login Setting',
      sessionSettings: 'Session',
      emailSettings: 'Email',
      messageSettings: 'Message',
      fileServer: 'File Server',
      webTimeout: 'Web Timeout(min)',
      nativePassword: 'Native Password',
      disableLogin: 'Disable login',
      notify: 'Notify',
      test: 'Test',
      emailServer: 'Server',
      emailPort: 'Port',
      sendEmail: 'Email',
      protocol: 'Protocol',
      emailAccount: 'Account',
      emailPassword: 'Password',
      sendTest: 'Send Test',
      parameterDescription: 'Parameter Description',
      messageDescription:
        'Receiving phone number :{mobile} Text message content :{content} validity period :{validity} Verification code :{code}',
      builtIn: 'Built-in',
      parameterType: 'Parameter Type',
      requestMethod: 'Request Method',
      requestUrl: 'Request URL',
      requestParameters: 'Request Parameters',
      messageContent: 'Message Content',
      fileserver1: 'File Server 1',
      fileserver2: 'File Server 2',
      addr: 'Address',
      workDir: 'Work Directory',
      totpTip:
        'Note: The mobile token function needs to ensure that the server time is as close to international standard time as possible.',
      totpDesc:
        'Mobile phone error time: between 1 and 30, each increase of 1, the time error is allowed to increase by 30 seconds',
      emailCodeValidity: 'Email Code Validity(min)',
      messageCodeValidity: 'Message Code Validity(min)',
      totpIssue: 'TOTP Issue',
      totpOffset: 'TOTP Offset',
      mix: 'Mix',
      first: 'First',
      second: 'Second',
      native: 'Native',
      mobile: 'Mobile',
      email: 'Email',
      message: 'Message',
      emailService: 'Email Service',
      emailBackup: 'Email Backup',
      noLimit: 'No Limit'
    },
    audit: {
      loginStatus: 'Login Status',
      startTime: 'Start Time',
      endTime: 'End Time',
      online: 'Online',
      offline: 'Offline',
      loginResult: 'Login Result',
      logOff: 'Log Off',
      logOffTip: 'Are you sure you want to force this user offline',
      opration: 'Opration',
      method: 'Method',
      path: 'Path',
      operateResult: 'Operate Result',
      requestBody: 'Request Body',
      responseBody: 'Response Body'
    },

    shell: {
      total: '',
      upload: 'Upload',
      download: 'Download',
      progress: 'Progress',
      speed: 'Speed',
      size: 'File Size',
      usedTime: 'Used Time',
      downloadSuccess: 'Download Success',
      skip: 'Skip',
      wsError: 'Unknown error, may be invalid login, please refresh the page and try again!'
    },
    firewallPolicy: {
      list: 'Policy List',
      policy: 'Policy',
      comment: 'Comment',
      action: 'Action',
      sourceIp: 'Source IP',
      port: 'Port',
      protocol: 'Protocol',
      allIp: 'All IP',
      partialIp: 'Partial IP',
      portValidationFailure: 'Port format error, please check',
      ipValidationFailure: 'IP format error, please check',
      portTip: 'Multiple ports separated by commas, for example: 80,8080-8090',
      ipTip: 'Multiple ips separated by commas, for example: 192.168.1.1,192.168.2.0/24,192.168.1.2-192.168.1.10',
      accept: 'Accept',
      reject: 'Reject',
      drop: 'Drop',
      position: 'Position',
      start: 'Start',
      end: 'End',
      beforePosition: 'Before Position',
      afterPosition: 'After Position',
      option: 'Option',
      pingReply: 'Ping Reply',
      pingRequest: 'Ping Request',
      newTcp: 'New Connection',
      establishedTcp: 'Established Connection',
      relatedTcp: 'Related Connection',
      untrackedTcp: 'Untracked Connection',
      invalidTcp: 'Invalid Connection',
      defaultPolicy: 'Default Policy',
      all: 'All',
      speed: 'Speed',
      allEthernet: 'All Ethernet',
      sourceEthernet: 'Source Ethernet',
      destinationEthernet: 'Destination Ethernet',
      dynamicIp: 'Dynamic IP',
      destIp: 'Destination IP',
      nat: 'NAT',
      destPort: 'Destination Port',
      sourcePort: 'Source Port',
      natPort: 'NAT Port'
    }
  },
  form: {
    required: 'Cannot be empty',
    userName: {
      required: 'Please enter user name',
      invalid: 'User name format is incorrect'
    },
    phone: {
      required: 'Please enter phone number',
      invalid: 'Phone number format is incorrect'
    },
    pwd: {
      required: 'Please enter password',
      invalid: '6-18 characters, including letters, numbers, and underscores'
    },
    confirmPwd: {
      required: 'Please enter password again',
      invalid: 'The two passwords are inconsistent'
    },
    code: {
      required: 'Please enter verification code',
      invalid: 'Verification code format is incorrect'
    },
    email: {
      required: 'Please enter email',
      invalid: 'Email format is incorrect'
    },
    userPwd1: {
      required: 'Please enter password',
      invalid: 'Must contain at least {length} characters'
    },
    userPwd2: {
      required: 'Must contain at least {length} characters',
      invalid: 'Must contain at least letters and numbers'
    },
    userPwd3: {
      required: 'Must contain at least {length} characters',
      invalid: 'Must contains at least digits, letters, and special characters'
    },
    userPwd4: {
      required: 'Must contain at least {length} characters',
      invalid: 'Must contain digits, lowercase letters, uppercase letters, and special characters'
    },
    departmentName: {
      required: 'Please enter department name',
      invalid: 'department name cannot be empty'
    },
    name: {
      required: 'Please enter name',
      invalid: 'Name cannot be empty'
    },
    level: {
      required: 'Please enter level',
      invalid: 'Please enter a range of grades {range}'
    },
    loginname: {
      required: 'Please enter login name',
      invalid:
        'Account must not contain Chinese characters, space, plus "+", colon ":", or slash "/", and it cannot exceed 200 characters. '
    },
    accountValidTime: {
      required: 'Please enter the account validity period',
      invalid: 'The end time cannot be later than the start time'
    },
    authentication: {
      required: 'Please enter the authentication',
      invalid: 'The authentication cannot be empty'
    },
    webTimeout: {
      required: 'cannot be empty',
      invalid: 'Value range: {range}'
    },
    url: {
      required: 'cannot be empty',
      invalid: 'Please enter the correct URL address'
    },
    range: {
      required: 'cannot be empty',
      invalid: 'Value range: {range}'
    },
    key: 'Key',
    value: 'Value'
  },
  dropdown: {
    closeCurrent: 'Close Current',
    closeOther: 'Close Other',
    closeLeft: 'Close Left',
    closeRight: 'Close Right',
    closeAll: 'Close All'
  },
  icon: {
    themeConfig: 'Theme Configuration',
    themeSchema: 'Theme Schema',
    lang: 'Switch Language',
    fullscreen: 'Fullscreen',
    fullscreenExit: 'Exit Fullscreen',
    reload: 'Reload Page',
    collapse: 'Collapse Menu',
    expand: 'Expand Menu',
    pin: 'Pin',
    unpin: 'Unpin'
  },
  datatable: {
    itemCount: 'Total {total} items',
    onlineTime: 'Online Time',
    page: 'Page',
    name: 'Name',
    username: 'User Name',
    loginname: 'Login Name',
    email: 'Email',
    phone: 'Phone',
    department: 'Department',
    role: 'Role',
    roleName: 'Role Name',
    status: 'Status',
    operation: 'Operation',
    action: 'Action',
    description: 'Description',
    createTime: 'Create Time',
    updateTime: 'Update Time',
    createUser: 'Create User',
    updateUser: 'Update User',
    level: 'Level',
    builtIn: 'Built-in',
    account: 'Account',
    password: 'Password',
    confirmPassword: 'Confirm Password',
    code: 'Code',
    permission: 'Permission',
    menu: 'Menu',
    authentication: 'Authentication',
    lastLoginTime: 'Last Login Time',
    locked: 'Locked',
    invalidAccount: 'Invalid Account',
    passwordExpired: 'Password Expired',
    normal: 'Normal',
    lastChangePasswordTime: 'Change Password At',
    all: 'All',
    valid: 'Valid',
    passwordValid: 'Password validity',
    accountValid: 'Account validity',
    globalConfig: 'Global Configuration',
    passwordLength: 'Password length',
    passwordComplexity: 'Password complexity',
    expiredAction: 'Expired action',
    differChecked: 'Different check',
    date: {
      year: 'Year',
      month: 'Month',
      day: 'Day',
      week: 'Week',
      hour: 'Hour',
      minute: 'Minute',
      second: 'Second'
    },
    bindTOTP: 'Rebind the TOTP',
    prot: 'Port',
    loginMethod: 'Login Method',
    secretKey: 'Secret Key',
    secretCipher: 'Secret Cipher',
    updatePassword: 'Update Password',
    loginTime: 'Login Time',
    logoutTime: 'Logout Time'
  }
};

export default local;
