import './style.css';

import {Greet} from '../wailsjs/go/main/App';

document.querySelector('#app').innerHTML = `
    <div id="result">Please enter your name below 👇</div>
    <div id="input">
        <input id="name" type="text" autocomplete="off" />
    <button onclick="greet()">Greet</button>
    </div>
`;

usernameElement = document.getElementById('username');

// Setup the greet function
window.start = function() { 
    username = nameElement.value;
