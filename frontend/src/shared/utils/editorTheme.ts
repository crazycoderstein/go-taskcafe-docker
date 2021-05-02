import theme from 'App/ThemeStyles';

const colors = {
  almostBlack: 'rgb(38, 44, 73)',
  lightBlack: 'rgb(16, 22, 58)',
  bgPrimary: 'rgb(16, 22, 58)',
  almostWhite: 'rgb(194, 198, 220)',
  white: '#FFF',
  white10: 'rgb(194, 198, 220)',
  black: '#000',
  black10: 'rgba(0, 0, 0, 0.1)',
  primary: 'rgb(115, 103, 240)',
  greyLight: '#F4F7FA',
  grey: '#E8EBED',
  greyMid: '#C5CCD3',
  greyDark: '#DAE1E9',
};

export const base = {
  ...colors,
  fontFamily: 'Open Sans',
  fontFamilyMono: "'SFMono-Regular',Consolas,'Liberation Mono', Menlo, Courier,monospace",
  fontWeight: 400,
  zIndex: 1000000,
  link: colors.primary,
  placeholder: '#B1BECC',
  textSecondary: '#fff',
  textLight: colors.white,
  textHighlight: '#b3e7ff',
  textHighlightForeground: colors.white,
  selected: colors.primary,
  codeComment: '#6a737d',
  codePunctuation: '#5e6687',
  codeNumber: '#d73a49',
  codeProperty: '#c08b30',
  codeTag: '#3d8fd1',
  codeString: '#032f62',
  codeSelector: '#6679cc',
  codeAttr: '#c76b29',
  codeEntity: '#22a2c9',
  codeKeyword: '#d73a49',
  codeFunction: '#6f42c1',
  codeStatement: '#22a2c9',
  codePlaceholder: '#3d8fd1',
  codeInserted: '#202746',
  codeImportant: '#c94922',

  blockToolbarBackground: colors.bgPrimary,
  blockToolbarTrigger: colors.white,
  blockToolbarTriggerIcon: colors.white,
  blockToolbarItem: colors.white,
  blockToolbarText: colors.white,
  blockToolbarHoverBackground: colors.primary,
  blockToolbarDivider: colors.almostWhite,

  noticeInfoBackground: '#F5BE31',
  noticeInfoText: colors.almostBlack,
  noticeTipBackground: '#9E5CF7',
  noticeTipText: colors.white,
  noticeWarningBackground: '#FF5C80',
  noticeWarningText: colors.white,
};

export const dark = {
  ...base,
  background: colors.almostBlack,
  text: colors.almostWhite,
  code: colors.almostWhite,
  cursor: colors.white,
  divider: '#4E5C6E',
  placeholder: '#52657A',

  toolbarBackground: colors.bgPrimary,
  toolbarHoverBackground: colors.primary,
  toolbarInput: colors.almostWhite,
  toolbarItem: colors.white,

  tableDivider: colors.lightBlack,
  tableSelected: colors.primary,
  tableSelectedBackground: '#002333',

  quote: colors.greyDark,
  codeBackground: colors.black,
  codeBorder: colors.lightBlack,
  codeString: '#3d8fd1',
  horizontalRule: colors.lightBlack,
  imageErrorBackground: 'rgba(0, 0, 0, 0.5)',

  scrollbarBackground: colors.black,
  scrollbarThumb: colors.lightBlack,
};

export default dark;
