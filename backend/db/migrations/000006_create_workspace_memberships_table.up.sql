CREATE TABLE workspace_memberships (
    id UUID PRIMARY KEY,
    workspace_id UUID NOT NULL,
    user_id UUID NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (workspace_id) REFERENCES workspaces(id),
    FOREIGN KEY (user_id) REFERENCES users(id),
    UNIQUE (workspace_id, user_id)
);

CREATE TRIGGER update_workspace_memberships_modtime 
BEFORE UPDATE ON workspace_memberships 
FOR EACH ROW 
EXECUTE FUNCTION update_modified_column();
