import { h } from 'vue';
import { NButton, NSpace } from 'naive-ui';

import { useAuth } from '@/hooks/business/auth';
import {
  CHANGE_FIREWALL_POLICY_POSITION,
  DELETE_FIREWALL_POLICY,
  UPDATE_FIREWALL_POLICY
} from '@/utils/permissions_consts';
import { $t } from '@/locales';

export function createAction(row: any, Fn: Function) {
  const { hasAuth } = useAuth();

  const opActions = [
    {
      permission: UPDATE_FIREWALL_POLICY,
      name: $t('common.edit'),
      disabled: false,
      type: 'success',
      onClick: (e: any) => {
        e.stopPropagation();
        Fn(row, 'update');
      }
    },
    {
      permission: CHANGE_FIREWALL_POLICY_POSITION,
      name: $t('page.firewallPolicy.position'),
      disabled: false,
      type: 'info',
      onClick: (e: any) => {
        e.stopPropagation();
        Fn(row, 'position');
      }
    },
    {
      permission: DELETE_FIREWALL_POLICY,
      name: $t('common.delete'),
      disabled: false,
      type: 'error',
      onClick: (e: any) => {
        e.stopPropagation();
        Fn(row, 'delete');
      }
    }
  ];

  return h(NSpace, { wrap: false, justify: 'center' }, () =>
    opActions.map((op: any) => {
      if (hasAuth(op.permission)) {
        return h(
          NButton,
          {
            type: op.type,
            quaternary: true,
            onClick: op.onClick
          },
          {
            default: () => op.name
            // icon: () => h(NIcon, null, { default: () => h(op.icon) })
          }
        );
      }

      return null;
    })
  );
}
