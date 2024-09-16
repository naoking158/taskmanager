<template>
  <v-dialog v-model="dialog" max-width="500px">
    <v-card>
      <v-card-title>Edit Task</v-card-title>
      <v-card-text>
        <v-text-field
          v-model="task.title"
          label="Title"
          required
        ></v-text-field>
        <v-textarea
          v-model="task.description"
          label="Description"
        ></v-textarea>
        <v-select
          label="タスク状態"
          :items="['TODO', 'In Progress', 'DONE', 'On Hold']"
        ></v-select>
        期限
        <v-row>
          <v-col class="d-flex align-center" cols="4" min-width="2em">
            <!-- 日付編集ダイアログ -->
            <v-btn
              @click="dateDialog = !dateDialog"
            >
              {{ formatDate(task.due_date) }}
              <v-dialog v-model="dateDialog">
                <v-date-picker
                  v-model="taskDueDate"
                  @update:modelValue="updateDate()"
                  @click:save="dateDialog=false"
                  @click:cancel="dateDialog=false"
                >
                </v-date-picker>
              </v-dialog>
            </v-btn>
          </v-col>
          <v-col class="d-flex align-center" cols="4" min-width="2em">
            <!-- 時間編集ダイアログ -->
            <v-btn
              @click="dateDialog = !dateDialog"
            >
              {{ formatTime(task.due_date) }}
              <v-dialog v-model="dateDialog">
                <v-date-picker
                  v-model="taskDueDate"
                  @update:modelValue="updateDate()"
                  @click:save="dateDialog=false"
                  @click:cancel="dateDialog=false"
                >
                </v-date-picker>
              </v-dialog>
            </v-btn>
          </v-col>
        </v-row>
      </v-card-text>
      <v-card-actions>
        <v-spacer></v-spacer>
        <v-btn color="blue darken-1" text @click="dialog = false">Cancel</v-btn>
        <v-btn color="blue darken-1" text @click="createTask">Create</v-btn>
      </v-card-actions>
    </v-card>
  </v-dialog>
</template>
