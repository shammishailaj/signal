<template>
  <v-app>

    <v-navigation-drawer
      v-model="drawer"
      fixed
      clipped
      :mini-variant="minidrawer"
      app v-if="is_authenticated()">
      <!--
        <v-list class="pa-1">
        <v-list-tile v-if="minidrawer" @click.stop="minidrawer = !minidrawer">
        <v-list-tile-action>
        <v-icon>chevron_right</v-icon>
        </v-list-tile-action>
        </v-list-tile>
        <v-list-tile-action v-if="!minidrawer">
        <v-btn icon @click.stop="minidrawer = !minidrawer">
        <v-icon>chevron_left</v-icon>
        </v-btn>
        </v-list-tile-action>
        </v-list>
      -->
      <v-list v-if="$route.path.indexOf('/projects/') === -1" dense>
        <v-list-tile
          v-for="item in items"
          :key="item.text"
          :to="{path: `/${item.path}`}">
          <v-list-tile-action>
            <v-icon>{{ item.icon }}</v-icon>
          </v-list-tile-action>
          <v-list-tile-content>
            <v-list-tile-title>
              {{ item.text }}
            </v-list-tile-title>
          </v-list-tile-content>
        </v-list-tile>
      </v-list>

      <v-list v-else dense>
        <v-list-tile
          v-for="item in project_items"
          :key="item.text"
          :to="{path: `/projects/${$route.params.project_id}/${item.path}`}">
          <v-list-tile-action>
            <v-icon>{{ item.icon }}</v-icon>
          </v-list-tile-action>
          <v-list-tile-content>
            <v-list-tile-title>
              {{ item.text }}
            </v-list-tile-title>
          </v-list-tile-content>
        </v-list-tile>
      </v-list>

      <v-list id="drawer-other-links">
        <v-list-tile to="/about">
          <v-list-tile-title class="grey--text text--lighten-1">about</v-list-tile-title>
        </v-list-tile>
      </v-list>
    </v-navigation-drawer>

    <v-toolbar color="light-blue" dark dense fixed clipped-left app v-if="is_authenticated()">
      <v-toolbar-side-icon @click.stop="drawer = !drawer"></v-toolbar-side-icon>
      <router-link to="/">
        <v-toolbar-title class="mr-5 align-center">
          <span class="title">Signal</span>
        </v-toolbar-title>
      </router-link>

      <v-spacer></v-spacer>

      <v-btn
        color="white"
        outline
        small
        >
        <v-icon left>restore</v-icon>
        Last 30 days
      </v-btn>

      <v-btn icon @click="logout">
        <v-icon>power_settings_new</v-icon>
      </v-btn>
    </v-toolbar>

    <v-content>
      <router-view/>
    </v-content>

  </v-app>
</template>

<script>
import auth from '@/services/auth';

export default {
  data: () => ({
    drawer: true,
    minidrawer: false,
    items: [
      { icon: 'dashboard', text: 'Projects', path: 'projects' },
      { icon: 'group', text: 'Accounts', path: 'accounts' },
    ],
    project_items: [
      { icon: 'home', text: 'Home', path: 'home' },
      { icon: 'settings', text: 'Settings', path: 'settings' },
    ],
  }),
  created() {
    if (this.$vuetify.breakpoint.mdAndDown) {
      this.drawer = false;
    }
  },
  methods: {
    is_authenticated() {
      return auth.is_authenticated();
    },
    logout() {
      auth.logout();
    },
  },
};
</script>

<style lang="scss">
#app {
  font-family: 'Avenir', Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  background-color: #fafbfc;
}
#nav {
  padding: 30px;
  a {
    font-weight: bold;
    color: #2c3e50;
    &.router-link-exact-active {
      color: #42b983;
    }
  }
}

.v-toolbar {
  a {
    color: white;
    text-decoration: none;
  }
}

.v-btn {
  border-radius: 4px;
}

#drawer-other-links {
  position: absolute;
  bottom: 0px;
  width: 100%;
}
</style>
