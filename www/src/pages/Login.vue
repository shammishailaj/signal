<template>
  <v-container grid-list-xl text-xs-center>
    <v-flex xs10 offset-xs1 sm8 offset-sm2 md6 offset-md3>
      <v-card id="login-card" class="elevation-3">

        <v-flex xs12 sm8 offset-sm2 md6 offset-md3>
          <v-text-field
            v-model="username"
            label="Username"
            outline
            ></v-text-field>
        </v-flex>
        <v-flex xs12 sm8 offset-sm2 md6 offset-md3>
          <v-text-field
            v-model="password"
            label="Password"
            outline
            type="password"
            @keyup.enter.native="login"
            ></v-text-field>
        </v-flex>
        <v-btn color="primary" @click="login">Login</v-btn>

      </v-card>
    </v-flex>
  </v-container>
</template>

<script>

export default {
  data: () => ({
    username: '',
    password: '',
  }),
  methods: {
    login() {
      console.log(this.username, this.password);
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
