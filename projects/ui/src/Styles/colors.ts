export const colors = {
  kumquatOrange: 'hsl(22, 100%, 94%)', //#FFEADE
  tangerineOrange: 'hsl(20, 92%, 84%)', //#FCCCB3
  peachOrange: 'hsl(20, 75%, 64%)', //#E98D5E
  grapefruitOrange: 'hsl(20, 76%, 47%)', //#D75B1D
  pumpkinOrange: 'hsl(19, 100%, 27%)', //#8B2E00

  flashlightGold: 'hsl(39, 100%, 95%)', //#FFF6E5
  sunGold: 'hsl(39, 100%, 59%)', //#FFB831
  novaGold: 'hsl(39, 78%, 44%)', //#C68A19

  groveGreen: 'hsl(161, 100%, 96%)', //	#ebfff9
  standGreen: '#39f0b8',
  forestGreen: 'hsl(161, 88%, 42%)', //#0dc98e

  dropBlue: 'hsl(202, 100%, 98%)', //#F7FCFF
  splashBlue: 'hsla(198, 100%, 93%, 1)', //#DEF5FF
  puddleBlue: 'hsl(198, 100%, 84%)', //#B1E7FF
  pondBlue: 'hsl(198, 81%, 67%)', // #6AC7F0
  lakeBlue: 'hsl(198, 71%, 60%)', //#54B7E3
  seaBlue: 'hsl(198, 72%, 46%)', //#2196C9
  oceanBlue: 'hsl(198, 84%, 37%)', //#0F7FB1
  darkOceanBlue: '#305781',
  neptuneBlue: 'hsl(204, 65%, 18%)', //#10344C
  planeOfWaterBlue: 'hsl(204, 59%, 13%)', //#0E2636

  lightJanuaryGrey: 'hsl(0, 0%, 99%)',
  januaryGrey: 'hsl(0, 0%, 97%)', //#F9F9F9',
  februaryGrey: 'hsl(0, 0%, 94%)', //#F2F2F2',
  darkFebruaryGrey: 'hsl(0, 0%, 92%)',
  marchGrey: 'hsl(240, 1%, 87%)', //#DFDFE0
  aprilGrey: 'hsl(216, 13%, 85%)', //#D4D8DE
  mayGrey: 'hsl(214, 12%, 78%)', //#C2C8D0
  juneGrey: 'hsl(216, 10%, 70%)', //#ADB3BC
  septemberGrey: 'hsl(200, 3%, 44%)', //#6E7477
  novemberGrey: 'hsl(200, 5%, 22%)', //#35393B

  scrollbarBorderGrey: '#E5E5E5',
  scrollbarBackgroundGrey: 'hsl(0,0%,98.5%)',

  boxShadow: 'hsla(0, 0%, 0%, 0.1)',
  darkerBoxShadow: 'hsla(0,0%,0%, .15)',

  envoyPink: '#d14a9b',
  errorRed: 'hsl(349 34% 50%)',
};

export function hslToHSLA(hsl: string, alpha: number | string) {
  let hslContent = hsl.split(')')[0].split('(')[1];

  return `hsla(${hslContent}, ${alpha})`;
}