// Code with Kids - JavaScript Learning Examples
// This file contains additional JavaScript examples for learning

// 1. Object Literals
const student = {
    name: "Alex",
    age: 12,
    grade: "6th",
    subjects: ["Math", "Science", "English"],
    greet: function() {
        return `Hello! I'm ${this.name}, a ${this.age}-year-old ${this.grade} grader.`;
    }
};

// 2. Classes (ES6)
class Animal {
    constructor(name, type) {
        this.name = name;
        this.type = type;
    }
    
    makeSound() {
        return `${this.name} the ${this.type} makes a sound!`;
    }
    
    getInfo() {
        return `${this.name} is a ${this.type}`;
    }
}

// 3. Arrow Functions
const multiply = (a, b) => a * b;
const isEven = num => num % 2 === 0;
const doubleArray = arr => arr.map(num => num * 2);

// 4. Template Literals
const createMessage = (name, activity) => {
    return `${name} is learning ${activity} today! 🎉`;
};

// 5. Destructuring
const person = {
    firstName: "Sam",
    lastName: "Johnson",
    age: 10,
    hobbies: ["coding", "reading", "sports"]
};

const { firstName, lastName, hobbies } = person;

// 6. Spread Operator
const numbers = [1, 2, 3];
const moreNumbers = [...numbers, 4, 5, 6];

// 7. Array Methods
const fruits = ["apple", "banana", "orange", "grape"];
const longFruits = fruits.filter(fruit => fruit.length > 5);
const upperFruits = fruits.map(fruit => fruit.toUpperCase());

// 8. Async/Await Example
async function fetchData(url) {
    try {
        const response = await fetch(url);
        const data = await response.json();
        return data;
    } catch (error) {
        console.error('Error fetching data:', error);
        return null;
    }
}

// 9. Local Storage Example
function saveToStorage(key, value) {
    localStorage.setItem(key, JSON.stringify(value));
}

function loadFromStorage(key) {
    const item = localStorage.getItem(key);
    return item ? JSON.parse(item) : null;
}

// 10. Event Handling
function setupEventListeners() {
    // Example of adding event listeners
    const buttons = document.querySelectorAll('button');
    buttons.forEach(button => {
        button.addEventListener('click', function() {
            console.log('Button clicked:', this.textContent);
        });
    });
}

// Export functions for use in HTML
window.JSLearning = {
    student,
    Animal,
    multiply,
    isEven,
    doubleArray,
    createMessage,
    person,
    numbers,
    moreNumbers,
    fruits,
    longFruits,
    upperFruits,
    fetchData,
    saveToStorage,
    loadFromStorage,
    setupEventListeners
};

// Console examples for learning
console.log('=== JavaScript Learning Examples ===');
console.log('Student object:', student);
console.log('Student greeting:', student.greet());

const dog = new Animal('Buddy', 'dog');
console.log('Animal example:', dog.makeSound());

console.log('Arrow function example:', multiply(4, 5));
console.log('Template literal:', createMessage('Emma', 'JavaScript'));

console.log('Destructuring:', firstName, lastName);
console.log('Spread operator:', moreNumbers);
console.log('Array methods - long fruits:', longFruits);
console.log('Array methods - upper fruits:', upperFruits);
