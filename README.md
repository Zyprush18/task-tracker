# Task Tracker CLI

**Task Tracker CLI** is a simple command-line interface (CLI) application to manage your tasks. This tool allows you to add, update, delete, and track tasks efficiently based on their status.

## Features

- ✅ **Add Tasks**: Add new tasks with a description.
- ✏️ **Update Tasks**: Modify the description of an existing task.
- ❌ **Delete Tasks**: Remove tasks from the list.
- 🔄 **Mark Tasks**: Change the status of a task to `in-progress` or `done`.
- 📋 **List Tasks**: View all tasks or filter them by status (`todo`, `in-progress`, `done`).

## Task Structure

Each task contains the following properties:

- `id`: A unique identifier for the task.
- `description`: A short description of the task.
- `status`: The current status (`todo`, `in-progress`, or `done`).
- `createdAt`: Timestamp of when the task was created.
- `updatedAt`: Timestamp of the last time the task was updated.

---

## Getting Started

### 1. Clone the Repository
```bash
git clone https://github.com/Zyprush18/task-tracker.git
cd task-tracker
```

### 2. Build the Application
```bash
go build -o task-cli main.go
```

---

## Commands and Examples

### ➕ Add a New Task
```bash
task-cli add "Buy groceries"
```
**Example output:**
```bash
Task added successfully (ID: 1)
```

---

### ✏️ Update an Existing Task
```bash
task-cli update <task_id> "New description"
```
**Example output:**
```bash
Task updated successfully
```

---

### ❌ Delete a Task
```bash
task-cli delete <task_id>
```
**Example output:**
```bash
Task deleted successfully
```

---

### 🔄 Mark Task as In-Progress
```bash
task-cli mark-in-progress <task_id>
```
**Example output:**
```bash
Task updated successfully
```

---

### ✅ Mark Task as Done
```bash
task-cli mark-done <task_id>
```
**Example output:**
```bash
Task updated successfully
```

---

### 📋 List All Tasks
```bash
task-cli list
```
**Example output:**
```json
[
  {
    "id": 1,
    "description": "Buy groceries",
    "status": "todo",
    "created_at": "2025-05-14T18:57:04.767335827+08:00",
    "updated_at": "0001-01-01T00:00:00Z"
  }
]
```

---

### 📂 List Tasks by Status
```bash
task-cli list <status>
```
Available statuses:
- `todo`
- `in-progress`
- `done`

**Example output:**
```json
[
  {
    "id": 2,
    "description": "Write documentation",
    "status": "in-progress",
    "created_at": "2025-05-14T19:10:00.123456789+08:00",
    "updated_at": "2025-05-14T19:15:00.987654321+08:00"
  }
]
```

---

## Project Idea

The CLI Task Manager is a lightweight tool for managing tasks directly from your terminal. It enhances productivity by allowing you to focus on tasks without leaving the command line.

For more inspiration and ideas, check out the [Task Tracker Project on Roadmap.sh](https://roadmap.sh/projects/task-tracker).