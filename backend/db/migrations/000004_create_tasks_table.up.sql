CREATE TABLE tasks (
    id UUID PRIMARY KEY,
    workspace_id UUID NOT NULL,
    title VARCHAR(200) NOT NULL,
    description TEXT,
    status VARCHAR(20) NOT NULL CHECK (status IN ('TODO', 'In Progress', 'DONE', 'On Hold')),
    created_by UUID NOT NULL,
    assigned_to UUID,
    parent_task_id UUID,
    due_date DATE,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (workspace_id) REFERENCES workspaces(id),
    FOREIGN KEY (created_by) REFERENCES users(id),
    FOREIGN KEY (assigned_to) REFERENCES users(id),
    FOREIGN KEY (parent_task_id) REFERENCES tasks(id)
);

CREATE TRIGGER update_tasks_modtime 
BEFORE UPDATE ON tasks 
FOR EACH ROW 
EXECUTE FUNCTION update_modified_column();
