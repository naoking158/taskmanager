CREATE TABLE comments (
    id UUID PRIMARY KEY,
    task_id UUID NOT NULL,
    user_id UUID NOT NULL,
    content TEXT NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (task_id) REFERENCES tasks(id),
    FOREIGN KEY (user_id) REFERENCES users(id)
);

CREATE TRIGGER update_comments_modtime 
BEFORE UPDATE ON comments 
FOR EACH ROW 
EXECUTE FUNCTION update_modified_column();
