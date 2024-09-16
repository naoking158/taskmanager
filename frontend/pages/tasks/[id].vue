<template>
  <v-container>
    <TaskDetail :task="task" />
  </v-container>
  <v-container>
    <Comments :comments="comments" />
  </v-container>
</template>

<script setup lang="ts">
import { useRoute } from 'vue-router'
import { useTaskByID } from '~/composables/useTask'
import { useComments } from '~/composables/useComments';

import TaskDetail from '~/components/Task/TaskDetail.vue'
import Comments from '~/components/Comment/Comments.vue'

const route = useRoute()
const taskId = route.params.id as string

const { data: task } = useTaskByID(taskId)
const { data: comments } = useComments(taskId)
</script>
