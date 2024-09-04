import { h } from 'vue';
import type { TreeOption } from 'naive-ui';
// import SvgIcon from '@/components/custom/svg-icon.vue';
import TwemojiCardFileBox from '~icons/twemoji/card-file-box';
import Fa6SolidFolderClosed from '~icons/fa6-solid/folder-closed';
import FluentEmojiFileFolder from '~icons/fluent-emoji/file-folder';

export type Result = {
  id: number;
  name: string;
  children?: Result[];
  parentId: number;
  suffix?: any;
  prefix?: any;
};

export function createTreePrefix(data: any) {
  if (!data.children || data.children.length === 0) {
    return () => h(TwemojiCardFileBox, {});
  }
  return () => h(FluentEmojiFileFolder, {});
}

export function updateTreePrefix(
  _keys: Array<string | number>,
  _option: Array<TreeOption | null>,
  meta: {
    node: TreeOption | null;
    action: 'expand' | 'collapse' | 'filter';
  }
) {
  if (!meta.node) return;
  switch (meta.action) {
    case 'expand':
      meta.node.prefix = () => h(FluentEmojiFileFolder, {});
      break;
    case 'collapse':
      meta.node.prefix = () => h(Fa6SolidFolderClosed, {});
      break;
    default:
  }
}

export function generateTreeData(data: Result) {
  const res: Result = {
    id: data.id,
    parentId: data.parentId,
    name: data.name
  };

  if (data.children && data.children.length > 0) {
    const tmpChildren: Result[] = [];
    data.children.forEach(ch => {
      tmpChildren[tmpChildren.length] = generateTreeData(ch);
    });
    res.children = tmpChildren;
  }

  res.prefix = createTreePrefix(data);
  return res;
}
