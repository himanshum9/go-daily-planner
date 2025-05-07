// Add Todo
function addTodo() {
    const title = document.getElementById('todoTitle').value;
    const description = document.getElementById('todoDescription').value;
    const dueDate = document.getElementById('todoDueDate').value;

    fetch('/planner/todos', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
        },
        body: JSON.stringify({
            title: title,
            description: description,
            dueDate: dueDate,
        }),
    })
    .then(response => response.json())
    .then(data => {
        if (data.error) {
            alert(data.error);
        } else {
            location.reload();
        }
    })
    .catch(error => {
        console.error('Error:', error);
        alert('Failed to add todo');
    });
}

// Update Todo
function updateTodo(id, completed) {
    fetch(`/planner/todos/${id}`, {
        method: 'PUT',
        headers: {
            'Content-Type': 'application/json',
        },
        body: JSON.stringify({
            completed: completed,
        }),
    })
    .then(response => response.json())
    .then(data => {
        if (data.error) {
            alert(data.error);
        }
    })
    .catch(error => {
        console.error('Error:', error);
        alert('Failed to update todo');
    });
}

// Delete Todo
function deleteTodo(id) {
    if (confirm('Are you sure you want to delete this todo?')) {
        fetch(`/planner/todos/${id}`, {
            method: 'DELETE',
        })
        .then(response => response.json())
        .then(data => {
            if (data.error) {
                alert(data.error);
            } else {
                location.reload();
            }
        })
        .catch(error => {
            console.error('Error:', error);
            alert('Failed to delete todo');
        });
    }
}

// Add Priority
function addPriority() {
    const title = document.getElementById('priorityTitle').value;
    const description = document.getElementById('priorityDescription').value;

    fetch('/planner/priorities', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
        },
        body: JSON.stringify({
            title: title,
            description: description,
        }),
    })
    .then(response => response.json())
    .then(data => {
        if (data.error) {
            alert(data.error);
        } else {
            location.reload();
        }
    })
    .catch(error => {
        console.error('Error:', error);
        alert('Failed to add priority');
    });
}

// Update Priority
function updatePriority(id, completed) {
    fetch(`/planner/priorities/${id}`, {
        method: 'PUT',
        headers: {
            'Content-Type': 'application/json',
        },
        body: JSON.stringify({
            completed: completed,
        }),
    })
    .then(response => response.json())
    .then(data => {
        if (data.error) {
            alert(data.error);
        }
    })
    .catch(error => {
        console.error('Error:', error);
        alert('Failed to update priority');
    });
}

// Delete Priority
function deletePriority(id) {
    if (confirm('Are you sure you want to delete this priority?')) {
        fetch(`/planner/priorities/${id}`, {
            method: 'DELETE',
        })
        .then(response => response.json())
        .then(data => {
            if (data.error) {
                alert(data.error);
            } else {
                location.reload();
            }
        })
        .catch(error => {
            console.error('Error:', error);
            alert('Failed to delete priority');
        });
    }
}

// Add Contact
function addContact() {
    const name = document.getElementById('contactName').value;
    const type = document.getElementById('contactType').value;
    const description = document.getElementById('contactDescription').value;

    fetch('/planner/contacts', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
        },
        body: JSON.stringify({
            name: name,
            type: type,
            description: description,
        }),
    })
    .then(response => response.json())
    .then(data => {
        if (data.error) {
            alert(data.error);
        } else {
            location.reload();
        }
    })
    .catch(error => {
        console.error('Error:', error);
        alert('Failed to add contact');
    });
}

// Update Contact
function updateContact(id, completed) {
    fetch(`/planner/contacts/${id}`, {
        method: 'PUT',
        headers: {
            'Content-Type': 'application/json',
        },
        body: JSON.stringify({
            completed: completed,
        }),
    })
    .then(response => response.json())
    .then(data => {
        if (data.error) {
            alert(data.error);
        }
    })
    .catch(error => {
        console.error('Error:', error);
        alert('Failed to update contact');
    });
}

// Delete Contact
function deleteContact(id) {
    if (confirm('Are you sure you want to delete this contact?')) {
        fetch(`/planner/contacts/${id}`, {
            method: 'DELETE',
        })
        .then(response => response.json())
        .then(data => {
            if (data.error) {
                alert(data.error);
            } else {
                location.reload();
            }
        })
        .catch(error => {
            console.error('Error:', error);
            alert('Failed to delete contact');
        });
    }
}

// Update Water Intake
function updateWaterIntake(glasses) {
    fetch('/planner/water-intake', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
        },
        body: JSON.stringify({
            glasses: glasses,
        }),
    })
    .then(response => response.json())
    .then(data => {
        if (data.error) {
            alert(data.error);
        } else {
            location.reload();
        }
    })
    .catch(error => {
        console.error('Error:', error);
        alert('Failed to update water intake');
    });
}

// Add Thought
function addThought() {
    const content = document.getElementById('thoughtContent').value;

    fetch('/planner/thought', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
        },
        body: JSON.stringify({
            content: content,
        }),
    })
    .then(response => response.json())
    .then(data => {
        if (data.error) {
            alert(data.error);
        } else {
            location.reload();
        }
    })
    .catch(error => {
        console.error('Error:', error);
        alert('Failed to add thought');
    });
}

// Generate Thought
function generateThought() {
    fetch('/planner/thought/generate', {
        method: 'POST',
    })
    .then(response => response.json())
    .then(data => {
        if (data.error) {
            alert(data.error);
        } else {
            document.getElementById('thoughtContent').textContent = data.content;
        }
    })
    .catch(error => {
        console.error('Error:', error);
        alert('Failed to generate thought');
    });
} 