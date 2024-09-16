<template>
  <v-container v-if="comments">
    <h2 class="text-h5 mb-2">Comments</h2>

    <v-container class="comments">
      <v-card v-for="comment in comments" :key="comment.id" :border=true>
        <div class="my-2 text-pre-wrap">
          <v-list-item-subtitle class="subtitle">
            {{ new Date(comment.created_at).toLocaleString() }} Posted by: {{ comment.user_id }} 
          </v-list-item-subtitle>
          <v-card-text>
            {{ comment.content }}
          </v-card-text>
        </div>
      </v-card>
    </v-container>
  </v-container>
</template>

<script setup lang="ts">
import type { Comment } from '~/types';

defineProps<{
  comments: Comment[] | undefined;
}>()
</script>

<style scoped>
.subtitle {
  padding-left: 1em;
}

.comments {
  display: flex;
  flex-direction: column;
  gap: 0.5em;
}
</style>
