<template>
  <v-container grid-list-xl text-xs-center>
    <v-layout row wrap text-xs-center>

      <v-flex xs12>
        <v-card>
          <div id="chart"></div>
        </v-card>
      </v-flex>

      <v-flex xs12 md6>
        <v-card class="subcard">
          <v-data-table
            :headers="page_headers"
            :items="top_pages"
            hide-actions
            class="elevation-1"
            >
            <template slot="items" slot-scope="props">
              <td class="text-xs-left">{{ props.item.path }}</td>
              <td class="text-xs-left">{{ props.item.views }}</td>
            </template>
          </v-data-table>
        </v-card>
      </v-flex>

      <v-flex xs12 md6>
        <v-card class="subcard">
          <v-data-table
            :headers="referrers_headers"
            :items="top_referrers"
            hide-actions
            class="elevation-1"
            >
            <template slot="items" slot-scope="props">
              <td class="text-xs-left">{{ props.item.path }}</td>
              <td class="text-xs-left">{{ props.item.views }}</td>
            </template>
          </v-data-table>
        </v-card>
      </v-flex>

    </v-layout>
  </v-container>
</template>

<script>
import echarts from 'echarts';
import Vue from 'vue';

export default {
  data: () => ({
    top_pages: [
      { path: '/san', views: '2,802' },
      { path: '/san/implementations', views: '517' },
      { path: '/san/versions/v1.0.0', views: '365' },
      { path: '/san/san-vs', views: '347' },
      { path: '/rocket', views: '309' },
      { path: '/', views: '250' },
      { path: '/rocket/docker', views: '156' },
    ],
    top_referrers: [
      { path: '(direct)', views: '3,241' },
      { path: 'news.ycombinator.com', views: '625' },
      { path: 'github.com', views: '330' },
      { path: 'old.reddit.com', views: '237' },
      { path: 't.co', views: '176' },
      { path: 'kerkour.com', views: '42' },
      { path: 'm.me', views: '17' },
    ],
    page_headers: [
      {
        text: 'Top Pages',
        align: 'left',
        sortable: false,
        value: 'path',
      },
      {
        text: 'Views',
        align: 'left',
        sortable: false,
        value: 'views',
      },
    ],
    referrers_headers: [
      {
        text: 'Top Referrers',
        align: 'left',
        sortable: false,
        value: 'path',
      },
      {
        text: 'Views',
        align: 'left',
        sortable: false,
        value: 'views',
      },
    ],
    chart: null,
  }),
  created() {
    Vue.nextTick(() => {
      this.render_chart();
    });
  },
  methods: {
    render_chart() {
      const raw_data =[{ x: '2018-08-24T22:00:00.000Z', y: 28 }, { x: '2018-08-25T22:00:00.000Z', y: 12 }, { x: '2018-08-26T22:00:00.000Z', y: 121 }, { x: '2018-08-27T22:00:00.000Z', y: 48 }, { x: '2018-08-28T22:00:00.000Z', y: 383 }, { x: '2018-08-29T22:00:00.000Z', y: 106 }, { x: '2018-08-30T22:00:00.000Z', y: 50 }, { x: '2018-08-31T22:00:00.000Z', y: 8 }, { x: '2018-09-01T22:00:00.000Z', y: 15 }, { x: '2018-09-02T22:00:00.000Z', y: 38 }, { x: '2018-09-03T22:00:00.000Z', y: 55 }, { x: '2018-09-04T22:00:00.000Z', y: 59 }, { x: '2018-09-05T22:00:00.000Z', y: 18 }, { x: '2018-09-06T22:00:00.000Z', y: 74 }, { x: '2018-09-07T22:00:00.000Z', y: 22 }, { x: '2018-09-08T22:00:00.000Z', y: 36 }, { x: '2018-09-09T22:00:00.000Z', y: 89 }, { x: '2018-09-10T22:00:00.000Z', y: 48 }, { x: '2018-09-11T22:00:00.000Z', y: 77 }, { x: '2018-09-12T22:00:00.000Z', y: 47 }, { x: '2018-09-13T22:00:00.000Z', y: 49 }, { x: '2018-09-14T22:00:00.000Z', y: 5 }, { x: '2018-09-15T22:00:00.000Z', y: 7 }, { x: '2018-09-16T22:00:00.000Z', y: 10 }, { x: '2018-09-17T22:00:00.000Z', y: 2303 }, { x: '2018-09-18T22:00:00.000Z', y: 2998 }, { x: '2018-09-19T22:00:00.000Z', y: 1902 }, { x: '2018-09-20T22:00:00.000Z', y: 1506 }, { x: '2018-09-21T22:00:00.000Z', y: 832 }, { x: '2018-09-22T22:00:00.000Z', y: 729 }, { x: '2018-09-23T22:00:00.000Z', y: 652 }, { x: '2018-09-24T22:00:00.000Z', y: 438 }]; // eslint-disable-line
      const series = raw_data.map(point => point.y);
      const dates = raw_data.map((point) => {
        const d = new Date(point.x);
        return [d.getFullYear(), d.getMonth() + 1, d.getDate()].join('/');
      });

      this.chart = echarts.init(document.getElementById('chart'));

      // specify chart configuration item and data
      const option = {
        tooltip: {
          trigger: 'axis',
          position(pt) {
            return [pt[0], '10%'];
          },
        },
        xAxis: {
          type: 'category',
          data: dates,
        },
        yAxis: {
          type: 'value',
        },
        series: [{
          name: 'views',
          type: 'line',
          smooth: true,
          areaStyle: {
            color: new echarts.graphic.LinearGradient(0, 0, 0, 1, [{
              offset: 0,
              color: 'rgb(255, 158, 68)',
            }, {
              offset: 1,
              color: 'rgb(255, 70, 131)',
            }]),
          },
          data: series,
        }],
      };

      // use configuration item and data specified to show chart
      this.chart.setOption(option);
    },
  },
};
</script>

<style scoped lang="scss">

.subcard {
  height: 100%;
  border-radius: 4px;
}

#chart {
  width: 100%;
  height: 50vh;
}

</style>
