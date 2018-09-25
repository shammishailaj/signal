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
        <v-btn color="primary" @click="login">Login</v-btn>

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

export default {
  data: () => ({
    username: '',
    password: '',
    error_msg: '',
    success_msg: '',
  }),
  methods: {
    async login() {
      console.log(this.username, this.password);
      try {
        await auth.login(this.username, this.password);
        this.error_msg = '';
        this.success_msg = 'Success';
        // change page
      } catch (err) {
        this.error_msg = err.toString();
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
