<template>
  <v-container>
    <h1 class="text-h4 mb-6">User page</h1>
    <v-row justify="center">
      <v-col cols="12" sm="8" md="6">
        <v-card class="mx-auto">
          <v-card-text>
            <v-form @submit.prevent="updateSetting">
              <v-text-field
                v-model="user.name"
                label="Username"
                prepend-icon="mdi-account"
                :rules="[v => !!v || 'Username is required']"
                required></v-text-field>
              <v-text-field
                v-model="user.display_name"
                label="Display Name"
                prepend-icon="mdi-account-card"
                :rules="[v => !!v || 'Display name is required']"
                required></v-text-field>
              <v-btn
                type="submit"
                color="primary"
                block
                class="mt-4"
              >
                Update Settings
              </v-btn>
            </v-form>
          </v-card-text>
        </v-card>
      </v-col>
    </v-row>
  </v-container>
</template>

<script setup lang="ts">
import { ref } from 'vue'

const user = ref({
  name: 'Taro',
  display_name: 'aaaaaaaaaaa',
})

const updateSetting = () => {
  alert('update setting is clicked')
}
</script>

<style scoped>
.container {
  display: flex;
  flex-direction: column;
}
</style>
