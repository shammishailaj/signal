<template>
  <v-container grid-list-xl text-xs-center>
  <v-layout row wrap text-xs-center>

    <v-flex xs12 sm4 text-xs-center
      class="card"
      id="create-project"
      @click="create_project_dialog = true">
      <v-card class="project-card">
        <h1 id="create-project-plus">+</h1><br />
        <h4 class="headline mb-0">Create project</h4>
      </v-card>
    </v-flex>

    <v-flex xs12 sm4 text-xs-center class="card" v-for="project in projects" :key="project.id">
      <router-link :to="{path: `/projects/${project.id}`}">
        <v-card class="project-card">
          <v-card-title primary-title>
            <h4 class="headline mb-0">{{ project.name }}</h4>
          </v-card-title>
        </v-card>
      </router-link>
    </v-flex>

    <v-dialog
      v-model="create_project_dialog"
      width="500"
      @keydown.esc="create_project_dialog = false"
      >
      <v-card class="elevation-3">
        <v-container grid-list-xl text-xs-center>

          <v-flex xs12 sm8 offset-sm2 md6 offset-md3>
            <v-text-field
              v-model="new_projec_name"
              label="Name"
              @keyup.enter.native="create_project"
              ></v-text-field>
          </v-flex>
      <v-btn
        color="primary"
        @click="create_project"
        :loading="is_loading">Create</v-btn>

        </v-container>
      </v-card>
    </v-dialog>

  </v-layout>
  </v-container>
</template>


<script>
import api from '@/services/api';

export default {
  data: () => ({
    is_loading: false,
    create_project_dialog: false,
    new_projec_name: '',
    projects: [],
  }),
  async created() {
    this.fetch_data();
  },
  methods: {
    async fetch_data() {
      this.is_loading = true;

      try {
        const res = await api.get('/api/v1/projects');
        this.projects = res.data.data;
      } catch (err) {
        console.error(err);
      } finally {
        this.is_loading = false;
      }
    },
    async create_project() {
      this.is_loading = true;

      try {
        const res = await api.post('/api/v1/projects', { name: this.new_projec_name });
        this.projects.push(res.data.data);
        this.new_projec_name = '';
        this.create_project_dialog = false;
      } catch (err) {
        console.error(err);
      } finally {
        this.is_loading = false;
      }
    },
  },
};
</script>

<style scoped lang="scss">

.card {

  a {
    text-decoration: none;
  }
}

.project-card {
  border-radius: 4px;
  height: 25vh;
  color: #2196F3;
}

#create-project-plus {
  padding-top: 4vh;
}

#create-project {
  cursor: pointer;
}

</style>
