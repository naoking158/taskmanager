<template>
  <v-col cols="12" class="mt-4">
    <v-form @submit.prevent="addComment">
      <v-textarea
        v-model="newComment"
        label="Add a comment"
        outlined
      ></v-textarea>
      <v-btn type="submit" color="primary" :disabled="!newComment.trim()">
        Add Comment
      </v-btn>
    </v-form>
  </v-col>
</template>
