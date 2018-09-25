<template>
  <v-layout row wrap text-xs-center>

    <v-flex xs12>
      <v-card>
        <h1>Chart</h1>
      </v-card>
    </v-flex>

    <v-flex xs12 md6>
      <v-card>
        <h1>Top pages</h1>
      </v-card>
    </v-flex>

    <v-flex xs12 md6>
      <v-card>
        <h1>Top referrers</h1>
      </v-card>
    </v-flex>

  </v-layout>
</template>

<script>

export default {
  created() {
    console.log('created');
  },
};
</script>
