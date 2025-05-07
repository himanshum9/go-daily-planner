// AJAX functions for Daily Planner

// Todo operations
function updateTodo(id, completed) {
    fetch(`/planner/todos/${id}`, {
        method: 'PUT',
        headers: {
            'Content-Type': 'application/json',
        },
        body: JSON.stringify({ completed: completed })
    })
    .then(response => response.json())
    .then(data => {
        if (data.error) {
            showAlert('error', data.error);
        } else {
            location.reload();
        }
    });
}

function deleteTodo(id) {
    if (confirm('Are you sure you want to delete this todo?')) {
        fetch(`/planner/todos/${id}`, {
            method: 'DELETE'
        })
        .then(response => response.json())
        .then(data => {
            if (data.error) {
                showAlert('error', data.error);
            } else {
                location.reload();
            }
        });
    }
}

// Priority operations
function updatePriority(id, completed) {
    fetch(`/planner/priorities/${id}`, {
        method: 'PUT',
        headers: {
            'Content-Type': 'application/json',
        },
        body: JSON.stringify({ completed: completed })
    })
    .then(response => response.json())
    .then(data => {
        if (data.error) {
            showAlert('error', data.error);
        } else {
            location.reload();
        }
    });
}

function deletePriority(id) {
    if (confirm('Are you sure you want to delete this priority?')) {
        fetch(`/planner/priorities/${id}`, {
            method: 'DELETE'
        })
        .then(response => response.json())
        .then(data => {
            if (data.error) {
                showAlert('error', data.error);
            } else {
                location.reload();
            }
        });
    }
}

// Contact operations
function updateContact(id, completed) {
    fetch(`/planner/contacts/${id}`, {
        method: 'PUT',
        headers: {
            'Content-Type': 'application/json',
        },
        body: JSON.stringify({ completed: completed })
    })
    .then(response => response.json())
    .then(data => {
        if (data.error) {
            showAlert('error', data.error);
        } else {
            location.reload();
        }
    });
}

function deleteContact(id) {
    if (confirm('Are you sure you want to delete this contact task?')) {
        fetch(`/planner/contacts/${id}`, {
            method: 'DELETE'
        })
        .then(response => response.json())
        .then(data => {
            if (data.error) {
                showAlert('error', data.error);
            } else {
                location.reload();
            }
        });
    }
}

// Water intake operations
function updateWaterIntake(glasses) {
    fetch('/planner/water-intake', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
        },
        body: JSON.stringify({ glasses: glasses })
    })
    .then(response => response.json())
    .then(data => {
        if (data.error) {
            showAlert('error', data.error);
        } else {
            location.reload();
        }
    });
}

// Thought operations
function generateThought() {
    fetch('/planner/thought/generate', {
        method: 'POST'
    })
    .then(response => response.json())
    .then(data => {
        if (data.error) {
            showAlert('error', data.error);
        } else {
            document.getElementById('thoughtContent').textContent = data.content;
        }
    });
}

// Helper functions
function showAlert(type, message) {
    const alertDiv = document.createElement('div');
    alertDiv.className = `alert alert-${type} alert-dismissible fade show`;
    alertDiv.innerHTML = `
        ${message}
        <button type="button" class="btn-close" data-bs-dismiss="alert"></button>
    `;
    
    const container = document.querySelector('.container');
    container.insertBefore(alertDiv, container.firstChild);
    
    setTimeout(() => {
        alertDiv.remove();
    }, 5000);
} 