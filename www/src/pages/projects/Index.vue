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
              <td class="text-xs-left">{{ props.item.page }}</td>
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
              <td class="text-xs-left">{{ props.item.referrer }}</td>
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

import api from '@/services/api';

export default {
  data: () => ({
    views: [],
    top_pages: [],
    top_referrers: [],
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
    this.fetch_data();
  },
  methods: {
    async fetch_data() {
      try {
        const results = await Promise.all([
          api.get(`/api/v1/projects/${this.$route.params.project_id}/views`),
          api.get(`/api/v1/projects/${this.$route.params.project_id}/pages`),
          api.get(`/api/v1/projects/${this.$route.params.project_id}/referrers`),
        ]);
        this.views = results[0].data.data;
        this.top_pages = results[1].data.data;
        this.top_referrers = results[2].data.data;
        Vue.nextTick(() => {
          this.render_chart();
        });
      } catch (error) {
        console.error(error);
      }
    },
    render_chart() {
      const series = this.views.map(point => point.views);
      const dates = this.views.map((point) => {
        const d = new Date(point.date);
        return [d.getFullYear(), d.getMonth() + 1, d.getDate()].join('/');
      });
      // console.log(series, dates);

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
          boundaryGap: false,
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
