import './style.css';

import {Greet} from '../wailsjs/go/main/App';

document.querySelector('#app').innerHTML = `
    <div id="result">Please enter your name below 👇</div>
    <div id="input">
        <input id="name" type="text" autocomplete="off" />
    <button onclick="greet()">Greet</button>
    </div>
`;

let nameElement = document.getElementById("name");
nameElement.focus();
let resultElement = document.getElementById("result");

// Setup the greet function
window.greet = function () {
    // Get name
    let name = nameElement.value;

    // Check if the input is empty
    if (name === "") return;

    // Call App.Greet(name)
    try {
        Greet(name)
            .then((result) => {
                // Update result with data back from App.Greet()
                resultElement.innerText = result;
            })
            .catch((err) => {
                console.error(err);
            });
    } catch (err) {
        console.error(err);
    }
};
