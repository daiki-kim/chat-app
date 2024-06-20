let token = null;
let socket = null;

function login() {
    const email = document.getElementById('email').value;
    const password = document.getElementById('password').value;

    fetch('/api/v1/login', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify({ email, password })
    })
    .then(response => response.json())
    .then(data => {
        token = data.token;
        document.getElementById('login').style.display = 'none';
        document.getElementById('chat').style.display = 'block';
        fetchRooms();
    })
    .catch(error => console.error('Error:', error));
}

function fetchRooms() {
    fetch('/api/v1/rooms', {
        method: 'GET',
        headers: {
            'Authorization': `Bearer ${token}`
        }
    })
    .then(response => response.json())
    .then(data => {
        const roomsDiv = document.getElementById('rooms');
        roomsDiv.innerHTML = '';
        data.rooms.forEach(room => {
            const roomDiv = document.createElement('div');
            roomDiv.innerText = room.name;
            roomDiv.onclick = () => joinRoom(room.id);
            roomsDiv.appendChild(roomDiv);
        });
    })
    .catch(error => console.error('Error:', error));
}

function joinRoom(roomId) {
    if (socket) {
        socket.close();
    }

    socket = new WebSocket(`ws://localhost:8080/ws/rooms/${roomId}`);

    socket.onmessage = event => {
        const message = JSON.parse(event.data);
        displayMessage(message);
    };

    fetchMessages(roomId);
}

function fetchMessages(roomId) {
    fetch(`/api/v1/rooms/${roomId}/messages`, {
        method: 'GET',
        headers: {
            'Authorization': `Bearer ${token}`
        }
    })
    .then(response => response.json())
    .then(data => {
        const messagesDiv = document.getElementById('messages');
        messagesDiv.innerHTML = '';
        data.messages.forEach(message => {
            displayMessage(message);
        });
    })
    .catch(error => console.error('Error:', error));
}

function displayMessage(message) {
    const messagesDiv = document.getElementById('messages');
    const messageDiv = document.createElement('div');
    messageDiv.innerText = `[${message.timestamp}] ${message.sender_id}: ${message.content}`;
    messagesDiv.appendChild(messageDiv);
}

function sendMessage() {
    const messageInput = document.getElementById('messageInput');
    const message = messageInput.value;
    messageInput.value = '';

    socket.send(JSON.stringify({ content: message }));
}
