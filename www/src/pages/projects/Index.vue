<template>
  <v-container grid-list-xl text-xs-center>
  <v-layout row wrap text-xs-center>

    <v-flex xs12>
      <v-card>
        <h1>Chart</h1>
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
  }),
  created() {
    console.log('created');
  },
};
</script>

<style scoped lang="scss">

.subcard {
  height: 100%;
  border-radius: 4px;
}

</style>
