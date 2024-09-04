import { $t } from '@/locales';

export function userOnlineChartOptions(): any {
  const color = ['#8e9dff', '#26deca'];

  return {
    tooltip: {
      trigger: 'axis',
      axisPointer: {
        type: 'cross',
        label: {
          backgroundColor: '#6a7985'
        }
      }
    },
    legend: {
      data: [$t('page.home.yesterday'), $t('page.home.today')]
    },
    grid: {
      left: '3%',
      right: '4%',
      bottom: '3%',
      containLabel: true
    },
    xAxis: [
      {
        type: 'category',
        boundaryGap: false,
        data: [] as string[]
      }
    ],
    yAxis: [
      {
        type: 'value'
      }
    ],
    series: [
      {
        color: color[0],
        name: $t('page.home.yesterday'),
        type: 'line',
        smooth: true,
        // showSymbol: false,/

        zlevel: 3,

        areaStyle: {
          color: {
            type: 'linear',
            x: 0,
            y: 0,
            x2: 0,
            y2: 1,
            colorStops: [
              {
                offset: 0.25,
                color: color[0]
              },
              {
                offset: 1,
                color: '#ffffff00'
              }
            ]
          }
        },
        emphasis: {
          focus: 'series'
        },
        data: [] as number[]
      },
      {
        color: color[1],
        name: $t('page.home.today'),
        type: 'line',
        smooth: true,
        // showSymbol: false,

        zlevel: 3,

        areaStyle: {
          color: {
            type: 'linear',
            x: 0,
            y: 0,
            x2: 0,
            y2: 1,
            colorStops: [
              {
                offset: 0.25,
                color: color[1]
              },
              {
                offset: 1,
                color: '#ffffff00'
              }
            ]
          }
        },
        emphasis: {
          focus: 'series'
        },
        data: [] as number[]
      }
    ]
  };
}

export function userLineChartOpts(): any {
  const color = ['#8e9dff', '#26deca'];
  return {
    tooltip: {
      trigger: 'axis',
      axisPointer: {
        type: 'cross',
        label: {
          backgroundColor: '#6a7985'
        }
      }
    },
    legend: {
      data: [$t('page.home.yesterday'), $t('page.home.today')]
    },
    grid: {
      left: '3%',
      right: '4%',
      bottom: '3%',
      containLabel: true
    },
    xAxis: {
      type: 'category',
      boundaryGap: false,
      data: [] as string[]
    },
    yAxis: {
      type: 'value'
    },
    series: [
      {
        color: color[0],
        name: $t('page.home.yesterday'),
        type: 'line',
        smooth: true,
        areaStyle: {
          color: {
            type: 'linear',
            x: 0,
            y: 0,
            x2: 0,
            y2: 1,
            colorStops: [
              {
                offset: 0.25,
                color: color[0]
              },
              {
                offset: 1,
                color: '#ffffff00'
              }
            ]
          }
        },
        emphasis: {
          focus: 'series'
        },
        data: [] as number[]
      },
      {
        color: color[1],
        name: $t('page.home.today'),
        type: 'line',
        smooth: true,
        areaStyle: {
          color: {
            type: 'linear',
            x: 0,
            y: 0,
            x2: 0,
            y2: 1,
            colorStops: [
              {
                offset: 0.25,
                color: color[1]
              },
              {
                offset: 1,
                color: '#ffffff00'
              }
            ]
          }
        },
        emphasis: {
          focus: 'series'
        },
        data: []
      }
    ]
  };
}

interface HomeData {
  assets: number;
  users: number;
  online: number;
  accounts: number;
  yesterday: number[];
  today: number[];
  xAxis: string[];
  license: number;
}

export function parseData(data: any): HomeData {
  const xAxis = [
    '02:00',
    '04:00',
    '06:00',
    '08:00',
    '10:00',
    '12:00',
    '14:00',
    '16:00',
    '18:00',
    '20:00',
    '22:00',
    '00:00'
  ];
  const { assets, users, online, accounts, yesterday, today, license } = data;

  const res = xAxis.slice(0, today.length);

  const yester = yesterday.slice(0, today.length);
  return {
    assets,
    users,
    online,
    accounts,
    yesterday: yester,
    today,
    xAxis: res,
    license
  };
}

export function licenseOptions(): any {
  return {
    title: {
      // text: `{value|${chartData}} {unit|%}`,
      text: '',
      top: 'center',
      left: 'center',
      textStyle: {
        rich: {
          value: {
            fontSize: 50,

            fontWeight: 700
          },
          unit: {
            fontSize: 20,

            padding: [0, 0, 6, 1],
            fontWeight: 700
          }
        }
      }
    },
    legend: {
      data: [$t('page.home.usedLicense'), $t('page.home.notusedLicense')]
    },
    series: [
      {
        type: 'pie',
        clockwise: false,
        radius: ['58%', '70%'],
        zlevel: 3,
        data: [
          {
            value: 0,
            label: {
              formatter: '',
              fontSize: 20
            },
            name: $t('page.home.notusedLicense'),
            itemStyle: {
              color: '#00000000'
            }
          },
          {
            value: 0,
            name: $t('page.home.usedLicense'),
            label: {
              formatter: '',
              fontSize: 20
            },
            itemStyle: {
              color: '#00000000'
            }
          }
        ],
        startAngle: -90
      }
    ]
  };
}
