<template>
  <div>
    <v-data-table
      class="elevation-1"
      hide-actions
      :items="accounts"
      :headers="headers"
      :loading="is_loading"
      >
      <template slot="items" slot-scope="props">
        <td class="text-xs-left">{{ props.item }}</td>
      </template>
    </v-data-table>
  </div>
</template>

<script>
import api from '@/services/api';

export default {
  data: () => ({
    is_loading: false,
    headers: [
      { text: 'Username', align: 'left', sortable: false },
    ],
    accounts: [],
  }),
  async created() {
    this.fetch_data();
  },
  methods: {
    async fetch_data() {
      this.is_loading = true;

      try {
        const res = await api.get('/api/v1/accounts');
        this.accounts = res.data.data;
      } catch (err) {
        console.error(err);
      } finally {
        this.is_loading = false;
      }
    },
  },
};
</script>
