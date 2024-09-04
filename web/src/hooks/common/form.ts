import { ref, toValue } from 'vue';
import type { ComputedRef, Ref } from 'vue';
import type { FormInst } from 'naive-ui';
import {
  REG_CODE_SIX,
  REG_DIGIT,
  REG_EMAIL,
  REG_LOGIN_NAME,
  REG_LOWERCASE,
  REG_OTHER_CHAR,
  REG_PHONE,
  REG_PWD,
  REG_UPPERCASE,
  REG_USER_NAME
} from '@/constants/reg';
import { $t } from '@/locales';

export interface PasswordComplexity {
  digit: boolean;
  lower: boolean;
  upper: boolean;
  other: boolean;
}

export const validatePasswordComplexity = (password: string) => {
  const complexity: PasswordComplexity = {
    digit: false,
    lower: false,
    upper: false,
    other: false
  };

  if (REG_DIGIT.test(password)) complexity.digit = true;
  if (REG_UPPERCASE.test(password)) complexity.upper = true;
  if (REG_LOWERCASE.test(password)) complexity.lower = true;
  if (REG_OTHER_CHAR.test(password)) complexity.other = true;

  return complexity;
};

export function useFormRules() {
  const patternRules = {
    userName: {
      pattern: REG_USER_NAME,
      message: $t('form.userName.invalid'),
      trigger: 'change'
    },
    loginName: {
      pattern: REG_LOGIN_NAME,
      message: $t('form.loginname.invalid'),
      trigger: 'change'
    },
    phone: {
      pattern: REG_PHONE,
      message: $t('form.phone.invalid'),
      trigger: 'change'
    },
    pwd: {
      pattern: REG_PWD,
      message: $t('form.pwd.invalid'),
      trigger: 'change'
    },
    code: {
      pattern: REG_CODE_SIX,
      message: $t('form.code.invalid'),
      trigger: 'change'
    },
    email: {
      pattern: REG_EMAIL,
      message: $t('form.email.invalid'),
      trigger: 'change'
    },
    level: {
      type: 'number',
      min: 11,
      max: 9999,
      message: $t('form.level.invalid', { range: '11-9999' }),
      trigger: 'change'
    },
    accountValidTime: {
      trigger: 'change',
      validator(_rule, value) {
        if (value && value[1] < new Date().getTime()) {
          return new Error($t('form.accountValidTime.invalid'));
        }

        return true;
      }
    }
  } satisfies Record<string, App.Global.FormRule>;

  const formRules = {
    userName: [createRequiredRule($t('form.userName.required')), patternRules.userName],
    phone: [patternRules.phone],
    pwd: [createRequiredRule($t('form.pwd.required')), patternRules.pwd],
    code: [createRequiredRule($t('form.code.required')), patternRules.code],
    email: [patternRules.email],
    loginPwd: [createRequiredRule($t('form.pwd.required'))],
    name: [createRequiredRule($t('form.name.required'))],
    level: [createRequiredNumberRule($t('form.level.required')), patternRules.level],
    loginname: [createRequiredRule($t('form.loginname.required')), patternRules.loginName],
    accountValidTime: [patternRules.accountValidTime]
  } satisfies Record<string, App.Global.FormRule[]>;

  /** the default required rule */
  const defaultRequiredRule = createRequiredRule($t('form.required'));

  function createRequiredRule(message: string): App.Global.FormRule {
    return {
      required: true,
      message,
      trigger: ['change', 'input']
    };
  }

  function createRequiredNumberRule(message: string, typeStr: any = 'number'): App.Global.FormRule {
    return {
      type: typeStr,
      required: true,
      message,
      trigger: 'change'
    };
  }

  /** create a rule for confirming the password */
  function createConfirmPwdRule(pwd: string | Ref<string> | ComputedRef<string>) {
    const confirmPwdRule: App.Global.FormRule[] = [
      { required: true, message: $t('form.confirmPwd.required') },
      {
        asyncValidator: (rule, value) => {
          if (value.trim() !== '' && value !== toValue(pwd)) {
            return Promise.reject(rule.message);
          }
          return Promise.resolve();
        },
        message: $t('form.confirmPwd.invalid'),
        trigger: 'input'
      }
    ];
    return confirmPwdRule;
  }

  function createUserPassword(length: number, complex: number, required = true) {
    const res: any[] = [
      {
        // createRequiredRule($t('form.pwd.required'))
        trigger: 'change',
        validator(_rule: any, value: string) {
          if (!value) return true;
          if (value.length < length) return new Error($t('form.userPwd1.invalid', { length }));
          // 密码必须符合密码强度
          const complexity = validatePasswordComplexity(value);

          switch (complex) {
            case 2:
              if (!complexity.digit || !(complexity.lower || complexity.upper))
                return new Error($t('form.userPwd2.invalid'));
              break;

            case 3:
              if (!complexity.digit || !(complexity.lower || complexity.upper) || !complexity.other)
                return new Error($t('form.userPwd3.invalid'));
              break;
            case 4:
              if (!complexity.digit || !complexity.lower || !complexity.upper || !complexity.other)
                return new Error($t('form.userPwd4.invalid'));
              break;

            default:
              return true;
          }
          return true;
        }
      }
    ];

    if (required) {
      res.push(createRequiredRule($t('form.pwd.required')));
    }
    return res;
  }

  function createUserConfirmPassword(value1: string, required = true) {
    const res: any[] = [
      {
        // createRequiredRule($t('form.pwd.required'))
        trigger: 'change',
        validator(_rule: any, value: string) {
          if (!value && required) return true;
          if (!value1 && !value) return true;
          if (value !== value1) {
            return new Error($t('form.confirmPwd.invalid'));
          }
          return true;
        }
      }
    ];
    if (required) {
      res.push(createRequiredRule($t('form.confirmPwd.required')));
    }

    return res;
  }

  return {
    patternRules,
    formRules,
    defaultRequiredRule,
    createRequiredRule,
    createConfirmPwdRule,
    createUserPassword,
    createUserConfirmPassword,
    createRequiredNumberRule
  };
}

export function useNaiveForm() {
  const formRef = ref<FormInst | null>(null);

  async function validate() {
    await formRef.value?.validate();
  }

  async function restoreValidation() {
    formRef.value?.restoreValidation();
  }

  return {
    formRef,
    validate,
    restoreValidation
  };
}
