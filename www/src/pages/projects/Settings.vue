<template>
  <v-flex xs12 sm6 offset-sm3 md4 offset-md4 v-if="project">
    <v-text-field
      label="Tracking ID"
      outline
      readonly
      v-model="project.tracking_id"
      ></v-text-field>
      <v-btn color="error" @click="delete_project">Delete project</v-btn>
  </v-flex>
</template>

<script>
import api from '@/services/api';
import router from '@/router';

export default {
  data: () => ({
    is_loading: false,
    project: null,
  }),
  async created() {
    this.fetch_data();
  },
  methods: {
    async fetch_data() {
      this.is_loading = true;

      try {
        const res = await api.get(`/api/v1/projects/${this.$route.params.project_id}`);
        this.project = res.data.data;
      } catch (err) {
        console.error(err);
      } finally {
        this.is_loading = false;
      }
    },
    async delete_project() {
      this.is_loading = true;

      try {
        await api.delete(`/api/v1/projects/${this.$route.params.project_id}`);
        router.push('/');
      } catch (err) {
        console.error(err);
      } finally {
        this.is_loading = false;
      }
    },
  },
};
</script>
