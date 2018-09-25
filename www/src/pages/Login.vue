<template>
  <v-container grid-list-xl text-xs-center>
    <v-flex xs10 offset-xs1 sm8 offset-sm2 md6 offset-md3>
      <v-card id="login-card" class="elevation-3">

        <v-flex xs12 sm8 offset-sm2 md6 offset-md3>
          <v-text-field
            v-model="username"
            label="Username"
            ></v-text-field>
        </v-flex>
        <v-flex xs12 sm8 offset-sm2 md6 offset-md3>
          <v-text-field
            v-model="password"
            label="Password"
            type="password"
            @keyup.enter.native="login"
            ></v-text-field>
        </v-flex>
        <v-btn
          color="primary"
          @click="login"
          :loading="is_loading">Login</v-btn>

        <v-alert
        :value="error_msg !== ''"
        type="error">
          {{ error_msg }}
        </v-alert>
        <v-alert
        :value="success_msg !== ''"
        type="success">
          {{ success_msg }}
        </v-alert>

      </v-card>
    </v-flex>
  </v-container>
</template>

<script>
import auth from '@/services/auth';
import router from '@/router';

export default {
  data: () => ({
    is_loading: false,
    username: '',
    password: '',
    error_msg: '',
    success_msg: '',
  }),
  methods: {
    async login() {
      this.is_loading = true;
      try {
        await auth.login(this.username, this.password);
        this.error_msg = '';
        this.success_msg = 'Success';
        router.push('/');
      } catch (err) {
        this.error_msg = err.toString();
      } finally {
        this.is_loading = false;
      }
    },
  },
};

</script>

<style scoped lang="scss">

#login-card {
  padding: 20px;
  border-radius: 4px;
  margin-top: 50px;
}

</style>
