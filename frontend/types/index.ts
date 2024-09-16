export type errorResponse = {
  message: string;
}

export type Workspace = {
  id: string;
  name: string;
  description: string;
}

type taskStatus = 'TODO' | 'In Progress' | 'DONE' | 'On Hold'

export type Task = {
  id: string;
  workspace_id?: string;
  title: string;
  description: string;
  status: taskStatus;
  due_date: Date;
  created_by?: string;
  assigned_to?: string;
  parent_task_id?: string;
  created_at?: Date;
  updated_at?: Date;
}

export type Comment = {
  id: string;
  task_id: string;
  user_id: string;
  content: string;
  created_at: Date;
  updated_at: Date;
}
